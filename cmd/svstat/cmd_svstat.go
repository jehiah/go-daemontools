package main

import (
	"fmt"
	"os"

	"github.com/jehiah/go-daemontools"
)

func main() {
	exitCode := 0
	for _, arg := range os.Args[1:] {
		stat, err := daemontools.Svstat(arg)
		if err != nil {
			exitCode += 1
			fmt.Println(fmt.Sprintf("%s:", arg), err)
		} else {
			fmt.Println(stat)
		}
	}
	os.Exit(exitCode)
}
