package lib

import (
	"os"
	"strings"
	"testing"
	"text/template"
)

const giganticTemplate = `package testdata

var (
	ySchema = []byte{
{{- range $index, $line := .Lines}}
		{{$line}}
{{- end}}
	}
)
`

func TestRun_giganticSlice(t *testing.T) {
	tmpl, err := template.New("testdata").Parse(giganticTemplate)
	if err != nil {
		t.Fatal(err)
	}

	var lines []string
	for range 200_000 {
		line := strings.Repeat("0x00,", 20)
		lines = append(lines, line)
	}

	file, err := os.CreateTemp(t.TempDir(), "gigantic_*.go")
	if err != nil {
		t.Fatal(err)
	}

	err = tmpl.Execute(file, map[string]any{"Lines": lines})
	if err != nil {
		t.Fatal(err)
	}

	_, err = Run([]string{file.Name()}, 150)
	if err != nil {
		t.Fatal(err)
	}
}
