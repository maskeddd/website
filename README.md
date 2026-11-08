# Website

## Project structure

```text
main.go                 Site generation and local preview server
pages/                  Page templates
components/             Reusable templates and shared layout
assets/css/input.css    Tailwind source CSS
public/                 Static files copied unchanged into dist/
  js/                   Plain JavaScript
  images/               Images
dist/                   Generated HTML, CSS, and copied static files
build.sh                Full site build
```

Edit `.templ` files, not the generated `*_templ.go` files beside them. Keep source
files out of `dist/`; it is ignored by Git and can be regenerated.

## Development

From the project root, run:

```sh
nix develop --command air
```

If direnv has already loaded the development shell, run `air` directly. Stop any
existing development process before starting it.

Open <http://127.0.0.1:8080>. Air runs `build.sh` when source files change, then
starts the Go program with `-serve`. The preview server serves `dist/` on
`127.0.0.1:8081`; Air's proxy on port `8080` automatically refreshes the browser
after each rebuild. Use the proxy URL, not port `8081` or a local HTML file.
Preview responses disable caching so CSS changes show up immediately. Changes
inside `dist/` do not trigger rebuilds.

Press Ctrl+C to stop Air and the preview server.

## Static generation

Build the complete site from the project root:

```sh
nix develop --command sh ./build.sh
```

Inside the development shell, run `sh ./build.sh` directly. The build:

1. Generates Go code from the templates.
2. Copies the contents of `public/` into `dist/` (excluding `.gitkeep` placeholders).
3. Compiles `assets/css/input.css` directly into `dist/css/style.css`.
4. Builds the Go program and renders `dist/index.html`.

`go run .` only renders HTML; it does not compile CSS or copy static files.

Only `dist/` needs to be deployed; neither Air nor the Go preview server is needed
in production.
