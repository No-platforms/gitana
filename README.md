

# Gitana CLI

Gitana CLI is a powerful command-line tool designed for generating Dockerized Git analytics and exporting results as PNG images. This tool provides various statistics from Git repositories, making it easier for developers to analyze their code contributions and project history.

## Features

- **First Commit Info**: Retrieve the date and author of the first commit in the repository.
- **Contributors List**: List all contributors to the repository along with their commit counts.
- **Branch Management**: Count and list branches, including merged branches.
- **Commit Count**: Get the total number of commits in the repository.
- **Line Count**: Count lines of code in the repository.
- **Docker Integration**: Generate Docker commands for GitHub repository analytics.

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

## Usage

### General Command Structure

```bash
gitana [command] [subcommand]
```

### Available Commands

- **Git Commands**:
  - `git first_commit`: Get information about the first commit.
  - `git contributors`: List all contributors.
  - `git branch_count`: Count all branches created and merged.
  - `git branches`: List all merge commits in the repository.
  - `git branches_count`: Count all merge commits.
  - `git commit_count`: Count total commits in the repository.

- **Other Commands**:
  - `stat [repository_url]`: Generate Docker commands for analytics.
  - `lines`: Count lines of code in the repository.

### Examples

1. Get information about the first commit:
   ```bash
   ./gitana git first_commit
   ```

2. List all contributors:
   ```bash
   ./gitana git contributors
   ```

3. Count all branches:
   ```bash
   ./gitana git branch_count
   ```

4. List all merge commits:
   ```bash
   ./gitana git branches
   ```

5. Count total commits:
   ```bash
   ./gitana git commit_count
   ```

6. Generate Docker commands for a specific repository:
   ```bash
   ./gitana stat https://github.com/user/repo.git
   ```

7. Count lines of code in the repository:
   ```bash
   ./gitana lines
   ```

## Contribution

Contributions are welcome! If you would like to contribute to Gitana CLI, please follow these steps:

1. Fork the repository.
2. Create a new branch (`git checkout -b feature/YourFeature`).
3. Make your changes and commit them (`git commit -m 'Add some feature'`).
4. Push to the branch (`git push origin feature/YourFeature`).
5. Open a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Thanks to [urfave/cli](https://github.com/urfave/cli) for providing an easy-to-use CLI framework.
- Thanks to [Docker](https://www.docker.com/) for making containerization accessible.

---

For more information or questions, please open an issue or contact me directly!
```
