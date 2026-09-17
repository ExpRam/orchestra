package main

import (
	"os"

	"github.com/expram/orchestra/internal/cli"
)

func main() {
	os.Exit(cli.Execute())
}
