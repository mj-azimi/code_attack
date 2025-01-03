package main

import (
	aiService "code_attack/Services"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	var repository string

	fmt.Println("Enter repository:")
	fmt.Scanln(&repository)

	fmt.Println("start clone project ...go")

	cmd := exec.Command("git" ,"clone",repository )
	cmd.Output();
	fmt.Println("clone project successfully")

	

	parts := strings.Split(repository, "/")
	repoName := parts[len(parts)-1]
	repoName = strings.TrimSuffix(repoName, ".git")

	fileContents, err := getFile(repoName)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for filename, content := range fileContents {
		aiRes , errorRes := aiService.Generate(content)
		if errorRes != nil{
		fmt.Printf("Error: %v", errorRes)
		}
		outputFilename := fmt.Sprintf("output_%s.txt", sanitizeFilename(filename))

		err := saveData(outputFilename, filename, aiRes)
		if err != nil {
			fmt.Printf("Error saving data to file %s: %v\n", outputFilename, err)
		} else {
			fmt.Printf("Data saved to file: %s\n", outputFilename)
		}
	}

	os.RemoveAll(repoName)
	fmt.Println("All files were successfully analyzed.")
}

func getFile(dirName string) (map[string]string, error) {
	fileContents := make(map[string]string)

	err := filepath.WalkDir(dirName, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			fmt.Printf("Error processing path %s: %v\n", path, err)
			return err
		}

		if d.IsDir() && d.Name() == ".git" {
			return filepath.SkipDir
		}

		if !d.IsDir() {
			fmt.Printf("Reading file: %s\n", path)

			file, err := os.Open(path)
			if err != nil {
				fmt.Printf("Error opening file %s: %v\n", path, err)
				return err
			}
			defer file.Close()

			content := new(strings.Builder)
			_, err = io.Copy(content, file)
			if err != nil {
				fmt.Printf("Error reading file content %s: %v\n", path, err)
				return err
			}

			fileContents[path] = content.String()
		}
		return nil
	})

	if err != nil {
		fmt.Printf("Error traversing files: %v", err)
	}

	return fileContents, nil
}

func saveData(outputFilename, originalFilename, aiOutput string) error {
	file, err := os.Create(outputFilename)
	if err != nil {
		fmt.Printf("Error creating file: %v", err)
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("Original File: %s\n\nAI Output:\n%s", originalFilename, aiOutput))
	if err != nil {
		fmt.Printf("Error writing to file: %v", err)
	}

	return nil
}

func sanitizeFilename(filename string) string {
	return strings.ReplaceAll(strings.ReplaceAll(filename, "/", "_"), "\\", "_")
}

