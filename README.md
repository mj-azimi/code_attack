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
<a href="https://ollama.com">Ollama WebSite</a>

![Image](https://private-user-images.githubusercontent.com/3325447/254932576-0d0b44e2-8f4a-4e99-9b52-a5c1c741c8f7.png?jwt=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJnaXRodWIuY29tIiwiYXVkIjoicmF3LmdpdGh1YnVzZXJjb250ZW50LmNvbSIsImtleSI6ImtleTUiLCJleHAiOjE3MzU5NzM5MTIsIm5iZiI6MTczNTk3MzYxMiwicGF0aCI6Ii8zMzI1NDQ3LzI1NDkzMjU3Ni0wZDBiNDRlMi04ZjRhLTRlOTktOWI1Mi1hNWMxYzc0MWM4ZjcucG5nP1gtQW16LUFsZ29yaXRobT1BV1M0LUhNQUMtU0hBMjU2JlgtQW16LUNyZWRlbnRpYWw9QUtJQVZDT0RZTFNBNTNQUUs0WkElMkYyMDI1MDEwNCUyRnVzLWVhc3QtMSUyRnMzJTJGYXdzNF9yZXF1ZXN0JlgtQW16LURhdGU9MjAyNTAxMDRUMDY1MzMyWiZYLUFtei1FeHBpcmVzPTMwMCZYLUFtei1TaWduYXR1cmU9ZTQ1ODlmZmQyNDZiMWY0NGViYjExZDkyMjRjZmVkM2M4MTNmYzUzYTNlMDlmY2Y2OTM1ODI0YzhmZWMxNGUyZiZYLUFtei1TaWduZWRIZWFkZXJzPWhvc3QifQ.ZSoL1QjXTSEQKOrqxqcJEQOEb2byiwURZNrpDBr1ahM)

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
