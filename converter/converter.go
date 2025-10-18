package converter

import (
	"fmt"
	"io"
	"os"
	"strings"
)

const marker = "{{content}}"

func ConvertMDToHTML(mdContent string) (string, error) {
	raw, err := toHTML(mdContent)
	if err != nil {
		return "", err
	}
	return sanitize(raw), nil
}

func Paterns(str string) string {
	if !strings.Contains(str, marker) {
		fmt.Printf("в шаблоне не найден маркер %s", marker)
		os.Exit(1)
	}
	return str
}

func ConvertFile(inputPath, outputPath, PaterPath string) error {

	in, err := os.ReadFile(inputPath)
	if err != nil {
		return err
	}
	html, err := ConvertMDToHTML(string(in))
	if err != nil {
		return err
	}

	var out io.Writer

	if outputPath == "" {

		tplData, err := os.ReadFile(PaterPath)
		if err != nil {
			return err
		}
		template := Paterns(string(tplData))
		result := strings.Replace(template, marker, html, 1)
		fmt.Println(result)
		return nil
	} else {
		f, err := os.Create(outputPath)
		if err != nil {
			return err
		}
		tplData, err := os.ReadFile(PaterPath)
		if err != nil {
			return err
		}

		template := Paterns(string(tplData))
		result := strings.Replace(template, marker, html, 1)

		defer f.Close()
		out = f

		_, err = out.Write([]byte(result))
		return err
	}
}
