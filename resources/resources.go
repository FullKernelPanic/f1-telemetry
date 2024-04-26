package resources

import (
	"embed"
	_ "embed"
)

//go:embed html/*.gohtml
var htmls embed.FS

func Htmls() embed.FS {
	return htmls
}
