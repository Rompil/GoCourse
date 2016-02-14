package main

import (
	"net/http"

	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
)

func main() {
	http.HandleFunc("/markdown", GenerateMarkdown)
	http.Handle("/", http.FileServer(http.Dir("public")))
	http.ListenAndServe(":8080", nil)
}

func GenerateMarkdown(rw http.ResponseWriter, r *http.Request) {
	unsafe_markdown := blackfriday.MarkdownCommon([]byte(r.FormValue("body"))) //Blackfriday itself does nothing to protect against malicious content
	safe_markdown := bluemonday.UGCPolicy().SanitizeBytes(unsafe_markdown)     // So, I use bluemonday as a sanitizer
	rw.Write(safe_markdown)
}
