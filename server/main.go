package main

import (
	"context"
	"embed"
	"flag"
	"html/template"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"strings"
	"sync/atomic"
	"syscall"
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
	var path string
	flag.StringVar(&path, "path", "", "path")
	flag.Parse()

	if path == "" {
		path, _ = os.Executable()
		if strings.Contains(path, "go-build") || strings.Contains(path, "___go_build_") {
			path, _ = os.Getwd()
		}
		path = strings.ReplaceAll(path, `\`, `/`)
		path = path[:strings.LastIndex(path, "/server")]
	}

	var err error
	if homeTpl, err = template.ParseFS(tpl, "tpl/home.html"); err != nil {
		panic(err)
	}
	if viewTpl, err = template.ParseFS(tpl, "tpl/view.html"); err != nil {
		panic(err)
	}

	notePath = path + "/notefile"
	prefixLen = len(notePath) + 1
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
	http.HandleFunc("/{$}", home)
	http.HandleFunc("/view/{path}", view)
	server := &http.Server{Addr: ":1024"}
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("ListenAndServe.err:", err)
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-quit
	if err := server.Shutdown(context.Background()); err != nil {
		log.Fatal("Shutdown.err:", err)
	}
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
			w.Write([]byte(`<a class="file" href="/view/`))
			w.Write([]byte(url.PathEscape(p[prefixLen:])))
			w.Write([]byte(`"><small>📄</small> `))
			w.Write([]byte(d.Name()))
			w.Write([]byte("</a>"))
		}
		w.Write([]byte("</li>\n"))
	}
}

func home(w http.ResponseWriter, _ *http.Request) {
	homeTpl.Execute(w, homeText.Load())
}

type ViewData struct {
	Title   string
	Nav     string
	Content string
}

func view(w http.ResponseWriter, r *http.Request) {
	p := r.PathValue("path")
	if len(p) < 2 {
		http.NotFound(w, r)
		return
	}
	b, err := os.ReadFile(notePath + "/" + p)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	viewTpl.Execute(w, &ViewData{
		Title:   p,
		Nav:     strings.ReplaceAll(p, "/", "📌"),
		Content: string(b),
	})
}
