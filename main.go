package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"website/components"
	"website/pages"

	"github.com/a-h/templ"
)

type route struct {
	Path  string
	Label string
	Page  templ.Component
}

var routes = []route{
	{"/", "About", pages.Index()},
	{"/programming", "Programming", pages.Programming()},
	{"/contact", "Contact", pages.Contact()},
}

func main() {
	serve := flag.Bool("serve", false, "serve the generated site for local development")
	addr := flag.String("addr", "127.0.0.1:8081", "preview server listen address")
	flag.Parse()

	nav := make([]components.NavLink, len(routes))
	for i, r := range routes {
		nav[i] = components.NavLink{Label: r.Label, Href: r.Path}
	}

	const rootPath = "dist"
	for _, r := range routes {
		if err := writePage(rootPath, r, nav); err != nil {
			log.Fatalf("failed to write page %s: %v", r.Path, err)
		}
	}

	if *serve {
		files := http.FileServer(http.Dir(rootPath))
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Cache-Control", "no-store")
			files.ServeHTTP(w, r)
		})
		log.Printf("Preview server listening on http://%s", *addr)
		log.Fatal(http.ListenAndServe(*addr, handler))
	}
}

func writePage(root string, r route, nav []components.NavLink) error {
	dir := filepath.Join(root, r.Path)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}

	f, err := os.Create(filepath.Join(dir, "index.html"))
	if err != nil {
		return err
	}

	ctx := components.WithPath(context.Background(), r.Path)
	ctx = components.WithNavLinks(ctx, nav)
	if err := r.Page.Render(ctx, f); err != nil {
		f.Close()
		return err
	}
	return f.Close()
}
