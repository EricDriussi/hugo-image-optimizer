package main

import (
	"optimize/cmd"
	"optimize/internal/config"
)

func main() {
	config.Load()
	cmd.Execute()
}
