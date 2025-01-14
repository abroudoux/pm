package repository

import (
	"fmt"
	"os"

	"github.com/abroudoux/pm/internal/actions"
)

func FlagMode() error {
	flag := os.Args[1]

	if flag == "--file" || flag == "-f" {
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
	} else if flag == "--help" || flag == "-h" {
		printHelpMenu()
	} else if flag == "--root" || flag == "-r" {
		err := actions.GoToFileReference()
		if err != nil {
			return err
		}
	} else if flag == "--version" || flag == "-v" {
		println("2.0.0")
	} else {
		err := actions.RunCommandInReferenceFileDirectory()
		if err != nil {
			return err
		}
	}

	return nil
}

func printHelpMenu() {
	println("Usage: pm [options] [command]")
	fmt.Printf("  %-20s %s\n", "pm [command]", "Go to the directory of the project and run the command")
	fmt.Printf("  %-20s %s\n", "pm [--file | -f] [file]", "Set the target file")
	fmt.Printf("  %-20s %s\n", "pm [--help | -h]", "Show this help menu")
}