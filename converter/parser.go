package converter

import (
	"bytes"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

var md = goldmark.New(
	goldmark.WithExtensions(
		extension.GFM, // таблицы, чек-листы, strikethrough
		extension.Strikethrough,
		extension.Table,
		extension.TaskList,
	),
	goldmark.WithRendererOptions(
		html.WithUnsafe(), // оставляем HTML, но позже всё режем bluemonday
	),
	goldmark.WithParserOptions(
		parser.WithAutoHeadingID(), // для TOC, если понадобится
	),
)

// toHTML парсит Markdown и возвращает НЕсанитизированный HTML.
func toHTML(src string) (string, error) {
	var buf bytes.Buffer
	if err := md.Convert([]byte(src), &buf); err != nil {
		return "", err
	}
	return buf.String(), nil
}
