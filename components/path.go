package components

import "context"

type pathKey struct{}

func WithPath(ctx context.Context, path string) context.Context {
	return context.WithValue(ctx, pathKey{}, path)
}

func CurrentPath(ctx context.Context) string {
	p, _ := ctx.Value(pathKey{}).(string)
	return p
}

type navKey struct{}

func WithNavLinks(ctx context.Context, links []NavLink) context.Context {
	return context.WithValue(ctx, navKey{}, links)
}

func NavLinks(ctx context.Context) []NavLink {
	links, _ := ctx.Value(navKey{}).([]NavLink)
	return links
}
