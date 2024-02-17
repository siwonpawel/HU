package ui

import (
	"embed"
	"io/fs"
	"os"
)

//go:embed html static
var efs embed.FS

func getFS(embeded bool) fs.FS {
	if embeded {
		return efs
	} else {
		return os.DirFS("ui")
	}
}
