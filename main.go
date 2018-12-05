package main

import (
	"fmt"
	"github.com/microsoft/go-winio"
	"os"
)

func main() {
	var err error
	l, err := winio.ListenPipe("\\\\.\\pipe\\pipe-test", nil)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	defer l.Close()
	fmt.Print("done\n")
}
