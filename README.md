# FastGo 🚀

A minimal, fast, and expressive web framework for Go – inspired by Express.js.


## ✨ Features

- ⚡ Ultra-light and minimal
- ✅ Routing with route params (`/users/:id`)
- 🔄 Middleware support
- 🗂 Static file serving
- 📦 JSON & Form body parsing
- ❓ Query param helpers + default values
- 🔥 Custom error handling
- 📁 Easily composable routers
- 📜 Auto JSON error responses
<!-- - 🧪 Easy to test -->


## 📦 Installation

```bash
go get github.com/gyanendra-baghel/fastgo
````


## 🧩 Example

```go
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


```


## 📚 Usage

### ➤ Route Params

```go
router.Get("/user/:id", func(ctx *fastgo.Ctx) {
	id := ctx.Params["id"]
})
```

### ➤ Query Helpers

```go
ctx.Query("q")                // string
ctx.QueryOrDefault("q", "x")  // default fallback
ctx.QueryInt("page")          // convert to int
ctx.QueryIntOrDefault("p", 1) // int with fallback
```

### ➤ JSON Body Parsing

```go
type Login struct {
	Username string
	Password string
}
var body Login
ctx.BodyJSON(&body)
```

### ➤ Custom Error Handling

```go
app.Use(func(ctx *fastgo.Ctx, next func(error)) {
	err := recover()
	if err != nil {
		ctx.JSON(500, map[string]string{"error": "Internal Server Error"})
		return
	}
	next(nil)
})
```

## 📜 License

MIT © 2025 \[Gyanendra Singh]

