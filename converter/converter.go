package converter

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func ConvertMDToHTML(mdContent string) (string, error) {
	raw, err := toHTML(mdContent)
	if err != nil {
		return "", err
	}
	return sanitize(raw), nil
}

func ConvertFile(inputPath, outputPath string) error {

	in, err := os.ReadFile(inputPath)
	if err != nil {
		return err
	}
	html, err := ConvertMDToHTML(string(in))
	if err != nil {
		return err
	}

	var out io.Writer
	f, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	tplData, err := os.ReadFile("C:\\Users\\kisel\\GolandProjects\\bebebe\\resurse\\pattern\\Ru_langveg.html")
	if err != nil {
		return err
	}
	template := string(tplData)

	const marker = "{{content}}"
	if !strings.Contains(template, marker) {
		fmt.Printf("в шаблоне не найден маркер %s", marker)
		os.Exit(1)
	}
	result := strings.Replace(template, marker, html, 1)

	defer f.Close()
	out = f

	_, err = out.Write([]byte(result))
	return err
}
