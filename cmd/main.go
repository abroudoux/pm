package main

import (
	_ "embed"
	"os"

	"github.com/abroudoux/pm/internal/actions"
	"github.com/abroudoux/pm/internal/repository"
	"github.com/abroudoux/pm/internal/utils"
)

func main() {
	err := actions.CheckIfReferenceFileExists()
	if err != nil {
		utils.PrintErrorAndExit(err)
	}

	if len(os.Args) > 1 {
		err := repository.FlagMode()
		if err != nil {
			utils.PrintErrorAndExit(err)
		}
	}

	println("No command provided. Use 'pm --help' to see the available options.")
}
