package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "GitHub CLI",
		Usage: "Run Docker commands with GitHub repository",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "url",
				Usage:    "GitHub repository URL",
				Required: true,
			},
			&cli.StringFlag{
				Name:  "username, u",
				Usage: "GitHub username",
			},
			&cli.StringFlag{
				Name:  "password, p",
				Usage: "GitHub password",
			},
		},
		Action: func(c *cli.Context) error {
			repoURL := c.String("url")
			username := c.String("username")
			password := c.String("password")

			if username != "" && password != "" {
				repoURL = strings.Replace(repoURL, "https://", fmt.Sprintf("https://%s:%s@", username, password), 1)
			}

			repoName := extractRepoName(repoURL)

			dockerCommand := fmt.Sprintf(`docker run --rm srcd/hercules hercules --burndown --burndown-files --burndown-people --devs --pb %s | docker run --rm -i -v "$(pwd):/io" srcd/hercules labours -f pb -m all -o /io/%s.png && git ls-files | xargs wc -l`, repoURL, repoName)
			
			fmt.Println("Running command:", dockerCommand)
			
			cmd := exec.Command("bash", "-c", dockerCommand)
			output, err := cmd.CombinedOutput()
			if err != nil {
				return fmt.Errorf("error running command: %v\nOutput: %s", err, output)
			}

			fmt.Println(string(output))
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func extractRepoName(url string) string {
	splitURL := strings.Split(url, "/")
	return splitURL[len(splitURL)-1]
}