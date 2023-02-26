package main

import (
	"github.com/EricDriussi/hugo-image-optimizer/cmd"
	"github.com/EricDriussi/hugo-image-optimizer/internal/config"
)

// TODO.Also compress webp if file size excedes X threshold
func main() {
	config.Load()
	cmd.Execute()
}

// TODO.Allow for compression config based on cwebp parameters
