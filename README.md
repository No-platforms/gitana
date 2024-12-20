

# Gitana CLI

Gitana CLI is a powerful command-line tool designed for generating Git analytics. This tool provides various statistics from Git repositories, making it easier for developers to analyze their code contributions and project history.

You can read more on [MEDIUM](https://medium.com/@yiiman/gitana-revolutionizing-git-repository-analysis-cc0d3a545647)
## Features

- **First Commit Info**: Retrieve the date and author of the first commit in the repository.
- **Contributors List**: List all contributors to the repository along with their commit counts.
- **Branch count**: Count and list branches, including merged branches.
- **Commit Count**: Get the total number of commits in the repository.
- **Line Count**: Count lines of code in the repository.
- **Docker Integration**: Generate Docker commands for GitHub repository analytics.
## Prerequisites
  Docker installed on your system

  Git repository URL
## Installation

To install Gitana CLI, follow these steps:

1. Clone the repository:
   ```bash
   git clone https://github.com/No-platforms/gitana.git
   ```

2. Navigate to the project directory:
   ```bash
   cd gitana
   ```

3. Build the application:
   ```bash
   go build -o gitana
   ```

4. (Optional) Move the binary to a directory in your PATH for easier access:
   ```bash
   mv gitana /usr/local/bin/
   ```
5. (Optional) Config permission:
   ```bash
   sudo chmod +x /usr/local/bin/gitana
   ```
   

Another way to build and install easily:
```bash
    make
```

## Usage

### General Command Structure

```bash
gitana [command] [subcommand]
```

### Available Commands

- **Git Commands**:
  - `git first_commit`: Get information about the first commit.
  - `git contributors`: List all contributors.
  - `git merged_branches_count`: Count all branches created and merged.
  - `git branches`: List all merge commits in the repository.
  - `git branches_count`: Count all merge commits.
  - `git commit_count`: Count total commits in the repository.

- **Other Commands**:
  - `stat [repository_url]`: Generate Docker commands for analytics.
  - `lines`: Count lines of code in the repository.

### Examples

1. Get information about the first commit:
   ```bash
   gitana git first_commit
   ```

2. List all contributors:
   ```bash
   gitana git contributors
   ```

3. Count all branches:
   ```bash
   gitana git merged_branches_count
   ```

4. List all merge commits:
   ```bash
   gitana git branches
   ```

5. Count total commits:
   ```bash
   gitana git commit_count
   ```

6. Generate Git Stats based on https://github.com/src-d/hercules package for specific repository

Important Considerations:
* First argument is git repository URL
* Report generation time depends on repository size
* Process may take several minutes for large repositories
* Requires pulling the Hercules Docker image before analysis
* Command Syntax:
   ```bash
   gitana stat https://github.com/user/repo.git
   ```

7. Count lines of code in the repository:
   ```bash
   gitana lines
   ```

## Contribution

Contributions are welcome! If you would like to contribute to Gitana CLI, please follow these steps:

1. Fork the repository.
2. Create a new branch (`git checkout -b feature/YourFeature`).
3. Make your changes and commit them (`git commit -m 'Add some feature'`).
4. Push to the branch (`git push origin feature/YourFeature`).
5. Open a Pull Request.

## License

This project is licensed under the Apache License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Thanks to [urfave/cli](https://github.com/urfave/cli) for providing an easy-to-use CLI framework.
- Thanks to [Docker](https://www.docker.com/) for making containerization accessible.
- Thanks to [src-d/hercules](https://github.com/src-d/hercules) for making containerization git stats and export images.



For more information or questions, please open an issue or contact me directly!


