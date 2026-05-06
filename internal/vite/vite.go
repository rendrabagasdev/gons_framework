package vite

import (
	"fmt"
	"html/template"
	"os"
	"strings"
)

// ViteHelper generates the necessary HTML tags to include Vite assets.
func ViteHelper(entries ...string) template.HTML {
	hotFile := "public/hot"
	var htmlTags []string

	if _, err := os.Stat(hotFile); err == nil {
		content, _ := os.ReadFile(hotFile)
		viteHost := strings.TrimSpace(string(content))
		if viteHost == "" {
			viteHost = "http://localhost:5173"
		}

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
