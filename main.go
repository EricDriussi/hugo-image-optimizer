package main

import (
	"hugo-images/cmd"
	"hugo-images/internal/config"
)

func main() {
	config.Load()
	cmd.Execute()
}
