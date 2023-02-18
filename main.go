package main

import (
	"github.com/EricDriussi/hugo-image-optimizer/cmd"
	"github.com/EricDriussi/hugo-image-optimizer/internal/config"
)

// TODO.refactor - modelling
func main() {
	config.Load()
	cmd.Execute()
}
