package services

import (
	"bebebe/internal/utils"
	"html/template"
)

type ContentService struct {
	htmlUtils *utils.HTMLUtils
}

func NewContentService() *ContentService {
	return &ContentService{
		htmlUtils: &utils.HTMLUtils{},
	}
}

func (cs *ContentService) ProcessUserInput(input string) template.HTML {
	// Обрабатываем пользовательский ввод
	return cs.htmlUtils.FormatContent(input)
}

func (cs *ContentService) SanitizeText(text string) string {
	// Просто экранируем для безопасного отображения
	return cs.htmlUtils.EscapeHTML(text)
}

func (cs *ContentService) DecodeStoredData(encoded string) string {
	// Декодируем сохраненные данные
	return cs.htmlUtils.UnescapeHTML(encoded)
}
