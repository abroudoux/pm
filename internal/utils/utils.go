package utils

import "os"

func PrintErrorAndExit(err error) {
	println(err)
	os.Exit(1)
} 