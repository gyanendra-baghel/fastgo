package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gyanendra-baghel/fastgo"
)

func main() {
	app := fastgo.New()
  // request timeout
	app.Use(fastgo.Timeout(5 * time.Second))

	router := fastgo.NewRouter()

	router.Get("/", func(c *fastgo.Ctx) {
		c.Text(200, "Welcome to FastGo!")
	})

	router.Get("/search", func(c *fastgo.Ctx) {
		q := c.QueryOrDefault("q", "default")
		c.JSON(200, map[string]string{"query": q})
	})

	app.Use(router.ServeHTTP)

	fmt.Println("Listening on http://localhost:3000")
	http.ListenAndServe(":3000", app)
}
