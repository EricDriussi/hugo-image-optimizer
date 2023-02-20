package main

import (
	"github.com/EricDriussi/hugo-image-optimizer/cmd"
	"github.com/EricDriussi/hugo-image-optimizer/internal/config"
)

func main() {
	config.Load()
	cmd.Execute()
}

// TODO.Allow for compression config based on cwebp parameters
// TODO.Make conf file optional, default to curr values and ask for file if no matching dirs
// TODO.Also compress webp if file size excedes X threshold
