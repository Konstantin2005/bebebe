package converter

import "github.com/microcosm-cc/bluemonday"

// policy описывает, какие теги/атрибуты мы оставляем.
var policy = func() *bluemonday.Policy {
	p := bluemonday.StrictPolicy() //

	p.AllowElements("p", "br", "hr", "blockquote", "pre", "code")
	p.AllowElements("h1", "h2", "h3", "h4", "h5", "h6")
	p.AllowElements("strong", "em", "del")
	p.AllowElements("ul", "ol", "li")
	p.AllowElements("a", "img")
	p.AllowElements("table", "thead", "tbody", "tr", "th", "td")

	p.AllowAttrs("src", "alt").OnElements("img")
	p.AllowAttrs("href").OnElements("a")

	p.AllowAttrs("class").Matching(bluemonday.SpaceSeparatedTokens).
		OnElements("code")

	return p
}()

func sanitize(raw string) string {
	return policy.Sanitize(raw)
}
