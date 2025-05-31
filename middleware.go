package fastgo

import (
	"context"
	"net/http"
	"time"
)

// Timeout middleware cancels requests taking longer than d and returns 504
func Timeout(d time.Duration) HandlerFunc {
	return func(c *Ctx) {
		ctx, cancel := context.WithTimeout(c.Req.Context(), d)
		defer cancel()

		done := make(chan struct{})
		panicChan := make(chan interface{})
		go func() {
			defer func() {
				if r := recover(); r != nil {
					panicChan <- r
				}
			}()
			c.Req = c.Req.WithContext(ctx)
			c.Next(nil)
			close(done)
		}()

		select {
		case <-ctx.Done():
			c.Next(HTTPErrorf(http.StatusGatewayTimeout, "Request timed out"))
		case p := <-panicChan:
			c.Next(HTTPErrorf(http.StatusInternalServerError, "panic: %v", p))
		case <-done:
		}
	}
}

// Static serves static files from dirPath at / prefix
func Static(dirPath string) HandlerFunc {
	fs := http.FileServer(http.Dir(dirPath))
	return func(c *Ctx) {
		fs.ServeHTTP(c.Res, c.Req)
	}
}
