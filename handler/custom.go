package handler

func customRoute(ctx *Context, a *App) bool {
	for _, rt := range a.Routes {
		matches := rt.Pattern.FindStringSubmatch(ctx.URL.Path)
		if len(matches) <= 0 {
			continue
		}
		if len(matches) > 1 {
			ctx.Params = matches[1:]
		}
		rt.Handler(ctx)
		return true
	}
	return false
}
