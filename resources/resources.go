package resources

import (
	"bytes"
	"embed"
	_ "embed"
	"fmt"
)

//go:embed html/*.gohtml
var html embed.FS

//go:embed css/*.css
var css embed.FS

//go:embed js/*.js
var js embed.FS

func Html() embed.FS {
	return html
}

func Css() ([]byte, error) {
	return build(css, "css")
}

func Js() ([]byte, error) {
	return build(js, "js")
}

func build(source embed.FS, dirName string) ([]byte, error) {
	var filePaths []string

	entries, err := source.ReadDir(dirName)
	if err != nil {
		//return nil, err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		filePaths = append(filePaths, fmt.Sprintf("%s/%s", dirName, entry.Name()))
	}

	return mergeFiles(source, filePaths...)
}

func mergeFiles(source embed.FS, filePaths ...string) ([]byte, error) {
	buffer := bytes.Buffer{}

	for _, fp := range filePaths {
		fc, err := source.ReadFile(fp)

		if err != nil {
			return nil, err
		}
		buffer.Write(fc)
	}

	return buffer.Bytes(), nil
}
