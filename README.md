![Image](https://raw.githubusercontent.com/mj-azimi/code_attack/refs/heads/main/DALL%C2%B7E%202025-01-04%2000.48.34%20-%20A%20futuristic%20office%20workspace%20with%20a%20computer%20screen%20displaying%20code.%20The%20screen%20shows%20a%20programming%20environment%20with%20lines%20of%20code%20being%20analyzed%20by%20.webp)


# Repository Code Analyzer with Ollama

This project is written in Go and aims to analyze the code in a repository. It uses the Ollama model to assess the code and provides a report on issues and weaknesses in the code.

## Features
- Connect to a Git repository.
- Use the Ollama model to analyze the code.
- Report issues and weaknesses in the code.
- Display the percentage of code issues.

## Prerequisites
To run this project, you need the following:
- Go (version 1.18 or higher)
- An internet connection to use the Ollama model
- Access to a Git repository

### ollama
Read the ollama documentation and install a suitable model first.
!["Ollama WebSite"]("https://ollama.com")

## Initial configuration
1. Open the config.json file
2. Enter your ollama server url
3. Enter your model name

## How to Run

1. To compile the project, use the following command:
    ```shell
    go build .
    ```

2. To run the application, use the following command:
    ```shell
    ./code_attack
    ```

3. After running the application, it will automatically connect to the Git repository and analyze the code for issues.

### Configuration Settings (Optional)
If you want to change some model or repository connection settings, you can edit the `config.json` configuration file.

## How to Use
After running the application, it will analyze the code in your repository. A report showing the issues and weaknesses in the code will be displayed, including the percentage of identified problems.

## Suggestions for Future Improvements:
- Support for analyzing code in different programming languages.
- Improve the issue reporting with more details.
- Add the ability to save reports as files.
