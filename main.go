package main

import (
	"fmt"
	"os"

	"github.com/KatsuyaAkasaka/pit/cmd"
)

func main() {
	c := cmd.Cmd()
	if err := c.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
