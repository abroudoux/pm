package repository

import (
	"fmt"
	"net/http"
	"os"

	"github.com/abroudoux/pm/internal/actions"
)

func FlagMode() error {
	flag := os.Args[1]

	switch flag {
	case "--file", "-f":
		if len(os.Args) == 2 {
			err := actions.PrintReferenceFile()
			if err != nil {
				return err
			}
		} else {
			referenceFile := os.Args[2]
			err := actions.SetReferenceFile(referenceFile)
			if err != nil {
				return err
			}
		}
	case "--root", "-r":
		err := actions.GoToFileReference()
		if err != nil {
			return err
		}
	case "--version", "-v":
		latestVersion, err := getLatestRelease()
		if err != nil {
			return err
		}

		fmt.Printf("Latest version: %s\n", latestVersion)
	case "--help", "-h":
		printHelpMenu()
	default:
		err := actions.RunCommandInReferenceFileDirectory()
		if err != nil {
			return err
		}
	}

	return nil
}

func getLatestRelease() (string, error) {
	url := "https://api.github.com/repos/abroudoux/pm/releases/latest"
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}

	latestVersion := res.Header.Get("tag_name")
	return latestVersion, nil
}

func printHelpMenu() {
	println("Usage: pm [options] [command]")
	fmt.Printf("  %-20s %s\n", "pm [command]", "Go to the directory of the project and run the command")
	fmt.Printf("  %-20s %s\n", "pm [--file | -f] [file]", "Set the target file")
	fmt.Printf("  %-20s %s\n", "pm [--help | -h]", "Show this help menu")
}