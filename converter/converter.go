package converter

import (
	"io"
	"os"
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
	if outputPath == "" {
		out = os.Stdout
	} else {
		f, err := os.Create(outputPath)
		if err != nil {
			return err
		}
		defer f.Close()
		out = f
	}

	_, err = out.Write([]byte(html))
	return err
}
