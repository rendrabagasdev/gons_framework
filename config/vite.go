package config

import (
	"fmt"
	"html/template"
	"strings"
)

func Vitehelper(entries ...string) template.HTML {
	env := GetEnv("APP_ENV", "development")
	viteHost := "http://localhost:5173"

	var htmlTags []string

	if env == "development" {

		htmlTags = append(htmlTags, fmt.Sprintf(`<script type="module" src="%s/@vite/client"></script>`, viteHost))

		for _, file := range entries {
			if strings.HasSuffix(file, ".css") {
				htmlTags = append(htmlTags, fmt.Sprintf(`<link rel="stylesheet" href="%s/%s">`, viteHost, file))
			} else if strings.HasSuffix(file, ".js") {
				htmlTags = append(htmlTags, fmt.Sprintf(`<script type="module" src="%s/%s"></script>`, viteHost, file))
			}
		}
	} else {
		for _, file := range entries {
			if strings.HasSuffix(file, ".css") {
				htmlTags = append(htmlTags, `<link rel="stylesheet" href="/css/style.css">`)
			} else if strings.HasSuffix(file, ".js") {
				htmlTags = append(htmlTags, `<script type="module" src="/js/app.js"></script>`)
			}
		}
	}

	return template.HTML(strings.Join(htmlTags, "\n"))
}
