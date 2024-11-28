package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "Gitana CLI",
		Usage: "Generate dockerized git analytics and export as PNG images",
		Commands: []*cli.Command{
			{
				Name:  "git",
				Usage: "Get various statistics from the git repository",
				Subcommands: []*cli.Command{
					{
						Name:   "first_commit",
						Usage:  "Get the date of the first commit along with the first developer's name in Gregorian format",
						Action: getFirstCommit,
					},
					{
						Name:   "contributors",
						Usage:  "List all contributors to the repository without duplicates",
						Action: listContributors,
					},
					{
						Name:   "merged_branches_count",
						Usage:  "Count all branches created and merged until now",
						Action: countMergedBranches,
					},
					{
						Name:   "branches",
						Usage:  "List all merge commits in the repository",
						Action: listBranches,
					},
					{
						Name:   "branches_count",
						Usage:  "Count all merge commits in the repository",
						Action: countMergeBranches,
					},
					{
						Name:   "commit_count",
						Usage:  "Count total commits in the repository",
						Action: countCommits,
					},
				},
			},
			{
				Name:   "stat",
				Usage:  "Generate hercules git stats from http git url like : 'https://github.com/No-platforms/gitana.git' and export all stats as png images",
				Action: runStat,
			},
			{
				Name:   "lines",
				Usage:  "Count lines of code in the repository",
				Action: countLines,
			},
		},
		EnableBashCompletion: true,
		Action: func(c *cli.Context) error {
			return cli.ShowAppHelp(c)
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func getFirstCommit(c *cli.Context) error {
	cmd := exec.Command("git", "rev-list", "--max-parents=0", "HEAD")
	output, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("error getting first commit hash: %v", err)
	}
	firstCommitHash := strings.TrimSpace(string(output))

	if firstCommitHash == "" {
		return fmt.Errorf("no commits found in this repository")
	}

	cmd = exec.Command("git", "log", "-1", "--pretty=format:%cd %cn", "--date=iso", firstCommitHash)
	output, err = cmd.Output()
	if err != nil {
		return fmt.Errorf("error getting first commit details: %v", err)
	}

	firstCommitInfo := strings.TrimSpace(string(output))
	if firstCommitInfo == "" {
		return fmt.Errorf("no details found for the first commit")
	}

	fmt.Println("First Commit Info:")
	fmt.Println(firstCommitInfo)

	return nil
}

func listContributors(c *cli.Context) error {
	cmd := exec.Command("git", "shortlog", "-sn", "--no-merges") // Use --no-merges to avoid merge commits
	output, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("error listing contributors: %v", err)
	}

	contributors := make(map[string]int)
	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		parts := strings.Fields(line)
		if len(parts) < 2 {
			continue
		}
		count := parts[0]
		name := strings.Join(parts[1:], " ")
		contributors[name] += atoi(count) // Convert string count to int and accumulate
	}

	if len(contributors) == 0 {
		fmt.Println("No contributors found.")
		return nil
	}

	fmt.Println("Contributors:")
	fmt.Printf("%-30s %s\n", "Name", "Commits")
	fmt.Println(strings.Repeat("-", 40))
	for name, count := range contributors {
		fmt.Printf("%-30s %d\n", name, count)
	}
	return nil
}

func countMergedBranches(c *cli.Context) error {
	cmd := exec.Command("git", "branch")
	output, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("error counting branches: %v", err)
	}

	count := len(strings.Split(strings.TrimSpace(string(output)), "\n"))
	fmt.Printf("Total branches created and merged: %d\n", count)
	return nil
}

func listBranches(c *cli.Context) error {
	cmd := exec.Command("git", "log", "--oneline", "--merges")
	output, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("error listing merge branches: %v", err)
	}

	fmt.Println("Merge Commits:")
	fmt.Println(string(output))
	return nil
}

func countMergeBranches(c *cli.Context) error {
	cmd := exec.Command("git", "log", "--oneline", "--merges")
	output, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("error counting merge branches: %v", err)
	}

	count := len(strings.Split(strings.TrimSpace(string(output)), "\n"))
	fmt.Printf("Total merge commits: %d\n", count)
	return nil
}

func countCommits(c *cli.Context) error {
	cmd := exec.Command("git", "rev-list", "--count", "--all", "--remotes")
	output, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("error counting commits: %v", err)
	}

	totalCommits := strings.TrimSpace(string(output))
	fmt.Printf("Total commits in the repository: %s\n", totalCommits)
	return nil
}

func runStat(c *cli.Context) error {
	if c.NArg() < 1 {
		return fmt.Errorf("please provide a GitHub repository URL")
	}

	repoURL := c.Args().Get(0)
	//Don't change Order of below arguments
	modes := []string{
		"burndown-project",
		"burndown-file",
		"burndown-person",
		"languages",
		"old-vs-new",
		"ownership",
		"run-times",
		"sentiment",
	}
	modStr := strings.Join(modes, " -m ")
	dockerCommand := fmt.Sprintf(`docker run --rm srcd/hercules hercules --burndown --burndown-files --devs --couples --burndown-people  --pb %s | docker run --rm -i -v "$(pwd):/io" srcd/hercules labours -f pb --background=white -m %s --disable-projector -o /io/%s.png`, repoURL, modStr, extractRepoName(repoURL))

	fmt.Println("Running command:", dockerCommand)

	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()

	cmd := exec.CommandContext(ctx, "bash", "-c", dockerCommand)

	// Create pipes for stdout and stderr
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return fmt.Errorf("failed to create stdout pipe: %v", err)
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return fmt.Errorf("failed to create stderr pipe: %v", err)
	}

	// Start the command
	fmt.Println("Starting command execution...")
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("failed to start command: %v", err)
	}

	// Create a channel to signal when we're done reading output
	done := make(chan bool)

	// Read stdout and stderr concurrently
	go func() {
		io.Copy(os.Stdout, stdout)
		done <- true
	}()
	go func() {
		io.Copy(os.Stderr, stderr)
		done <- true
	}()

	// Wait for both stdout and stderr to be fully read
	<-done
	<-done

	// Wait for the command to finish
	fmt.Println("Waiting for command to finish...")
	if err := cmd.Wait(); err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			return fmt.Errorf("command failed with exit code %d", exitErr.ExitCode())
		}
		return fmt.Errorf("failed to run command: %v", err)
	}

	fmt.Println("Command execution completed successfully.")
	return nil
}

func countLines(c *cli.Context) error {
	cmd := exec.Command("bash", "-c", "find . -type f -exec wc -l {} + | awk '{ total += $1 } END { print total }'")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error counting lines: %v\nOutput: %s", err, output)
	}

	fmt.Println(string(output))
	return nil
}

// extract report name for create report directory
func extractRepoName(url string) string {
	parts := strings.Split(url, "/")
	return fmt.Sprintf("%s_%s", strings.TrimSuffix(parts[len(parts)-1], ".git"), "gitana_stats")
}

// Helper function to convert string to int safely
func atoi(s string) int {
	val, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return val
}
