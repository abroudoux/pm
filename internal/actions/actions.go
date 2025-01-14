package actions

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var referenceFilePath string = "/tmp/pm_reference_file"

func CheckIfReferenceFileExists() error {
	referenceFileContent, err := os.ReadFile(referenceFilePath)
	if err != nil {
		err := createReferenceFile()
		if err != nil {
			return err
		}
	}

	if string(referenceFileContent) == "" {
		return fmt.Errorf("reference file not found")
	}

	return nil
}

func createReferenceFile() error {
	file, err := os.Create(referenceFilePath)
	if err != nil {
		return err
	}

	defer file.Close()

	err = SetReferenceFile("package.json")
	if err != nil {
		return err
	}

	println("Reference file successfully created!")
	return nil
}

func SetReferenceFile(referenceFile string) error {
	err := os.WriteFile(referenceFilePath, []byte(referenceFile), 0644)
	if err != nil {
		return err
	}

	println("Target file set to", referenceFile)
	return nil
}

func PrintReferenceFile() error {
	referenceFileContent, err := os.ReadFile(referenceFilePath)
	if err != nil {
		return err
	}

	println("Current reference file: " + strings.TrimSpace(string(referenceFileContent)))
	return nil
}

func getReferenceFile() (string, error) {
	referenceFileContent, err := os.ReadFile(referenceFilePath)
	if err != nil {
		return "", err
	}

	if string(referenceFileContent) == "" {
		return "", fmt.Errorf("reference file not found")
	}

	return string(referenceFileContent), nil
}

func GoToFileReference() error {
	for {
		if checkIfReferenceFileInCurrentDirectory() {
			currentDir, err := getCurrentWorkingDirectory()
			if err != nil {
				return err
			}

			println("Current directory: " + currentDir)
			return nil
		}

		err := moveBack()
		if err != nil {
			return err
		}

		if isRootDirectory() {
			return fmt.Errorf("reference file not found")
		}
	}
}

func getCurrentWorkingDirectory() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	return dir, nil
}

func isRootDirectory() bool {
	currentDir, _ := os.Getwd()
	return currentDir == "/"
}

func moveBack() error {
	err := os.Chdir("..")
	if err != nil {
		return err
	}

	return nil
}

func checkIfReferenceFileInCurrentDirectory() bool {
	referenceFile, err := getReferenceFile()
	if err != nil {
		return false
	}

	_, err = os.Stat(referenceFile)
	return err == nil
}

func RunCommandInReferenceFileDirectory() error {
	err := GoToFileReference()
	if err != nil {
		return err
	}

	command := ""
	for _, arg := range os.Args[1:] {
		command += arg + " "
	}

	command = command[:len(command)-1]
	partsCommand := strings.Fields(command)
	if len(partsCommand) < 1 {
		return fmt.Errorf("command not found")
	}

	cmd := exec.Command(partsCommand[0], partsCommand[1:]...)
	err = cmd.Run()
	if err != nil {
		return err
	}

	println("Command executed successfully!")

	return nil
}