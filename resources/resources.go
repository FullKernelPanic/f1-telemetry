package resources

import (
	_ "embed"
)

//go:embed html/index.html
var indexHtml string

func IndexHtml() string {
	return indexHtml
}
