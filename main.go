package main

import (
	"optimize/cmd"
	"optimize/internal/config"
)

// TODO.refactor - modelling
func main() {
	config.Load()
	cmd.Execute()
}
