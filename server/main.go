package main

import (
	"embed"
	"flag"
	"html/template"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync/atomic"
	"time"
)

//go:embed tpl
var tpl embed.FS

//go:embed static
var static embed.FS

var (
	notePath         string
	prefixLen        int
	homeTpl, viewTpl *template.Template
	homeText         = &atomic.Value{}
)

func init() {
	var dirPath string
	flag.StringVar(&dirPath, "path", "E:/github/note", "project dir path")
	flag.Parse()

	var err error
	if homeTpl, err = template.ParseFS(tpl, "tpl/index.html"); err != nil {
		panic(err)
	}
	if viewTpl, err = template.ParseFS(tpl, "tpl/view.html"); err != nil {
		panic(err)
	}

	notePath = dirPath + "/notefile"
	prefixLen = len(notePath)
	load()
	go func() {
		t := time.Tick(time.Minute)
		for range t {
			load()
		}
	}()
}

func main() {
	http.Handle("/static/", http.FileServer(http.FS(static)))
	http.HandleFunc("/", index)
	http.HandleFunc("/view", view)
	http.ListenAndServe(":1024", nil)
}

func load() {
	buf := &strings.Builder{}
	read(notePath, buf)
	homeText.Store(template.HTML(buf.String()))
}

func read(path string, w io.Writer) {
	dirs, err := os.ReadDir(path)
	if err != nil {
		return
	}
	for _, d := range dirs {
		w.Write([]byte("<li>"))
		p := path + "/" + d.Name()
		if d.IsDir() {
			w.Write([]byte(`<span class="dir"><span>📘</span> `))
			w.Write([]byte(d.Name()))
			w.Write([]byte(`</span><ul class="sub-ul">`))
			read(p, w)
			w.Write([]byte("</ul>"))
		} else {
			w.Write([]byte(`<a class="file" href="/view?path=`))
			w.Write([]byte(url.QueryEscape(p[prefixLen:])))
			w.Write([]byte(`"><small>📄</small> `))
			w.Write([]byte(d.Name()))
			w.Write([]byte("</a>"))
		}
		w.Write([]byte("</li>\n"))
	}
}

func index(w http.ResponseWriter, _ *http.Request) {
	homeTpl.Execute(w, homeText.Load())
}

type ViewData struct {
	Title   string
	Nav     string
	Content string
}

func view(w http.ResponseWriter, r *http.Request) {
	p := r.URL.Query().Get("path")
	if len(p) < 2 {
		http.NotFoundHandler()
		return
	}
	b, err := os.ReadFile(notePath + p)
	if err != nil {
		http.NotFoundHandler()
		return
	}
	viewTpl.Execute(w, &ViewData{
		Title:   p,
		Nav:     strings.ReplaceAll(p, "/", "📌"),
		Content: string(b),
	})
}
