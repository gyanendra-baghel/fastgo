package fastgo

import (
	"log"
	"net/http"
)

type HandlerFunc func(*Ctx)

type App struct {
	middlewares []HandlerFunc
	server      *http.Server
}

func New() *App {
	return &App{}
}

// Add middleware (router or global handler)
func (a *App) Use(handler HandlerFunc) {
	a.middlewares = append(a.middlewares, handler)
}

func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := NewCtx(w, r)

	var index int
	var run func(error)

	run = func(err error) {
		if err != nil {
			ctx.LastErr = err
			return
		}

		if index < len(a.middlewares) {
			mw := a.middlewares[index]
			index++
			defer func() {
				if rec := recover(); rec != nil {
					run(HTTPErrorf(http.StatusInternalServerError, "panic: %v", rec))
				}
			}()
			mw(ctx)
		}
	}

	ctx.Next = func(err error) {
		ctx.LastErr = err
		run(err)
	}

	run(nil)
}

func (a *App) Listen(addr string) error {
	a.server = &http.Server{
		Addr:    addr,
		Handler: a,
	}
	log.Printf("🚀 Listening on %s", addr)
	return a.server.ListenAndServe()
}
