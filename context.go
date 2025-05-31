package fastgo

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

type Ctx struct {
	Req     *http.Request
	Res     http.ResponseWriter
	Params  map[string]string
	LastErr error
	Next    func(err error)
}

func NewCtx(w http.ResponseWriter, r *http.Request) *Ctx {
	return &Ctx{
		Res:    w,
		Req:    r,
		Params: make(map[string]string),
	}
}

// Response helpers
func (c *Ctx) JSON(status int, data any) {
	c.Res.Header().Set("Content-Type", "application/json")
	c.Res.WriteHeader(status)
	json.NewEncoder(c.Res).Encode(data)
}

func (c *Ctx) Text(status int, str string) {
	c.Res.Header().Set("Content-Type", "text/plain")
	c.Res.WriteHeader(status)
	c.Res.Write([]byte(str))
}

// Request helpers
func (c *Ctx) BodyJSON(v any) error {
	defer c.Req.Body.Close()
	return json.NewDecoder(c.Req.Body).Decode(v)
}

func (c *Ctx) Form() (url.Values, error) {
	if err := c.Req.ParseForm(); err != nil {
		return nil, err
	}
	return c.Req.Form, nil
}

func (c *Ctx) Body() ([]byte, error) {
	defer c.Req.Body.Close()
	return io.ReadAll(c.Req.Body)
}

// Query Parameter Helpers

func (c *Ctx) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

func (c *Ctx) Queries() map[string][]string {
	return c.Req.URL.Query()
}

func (c *Ctx) QueryInt(key string) (int, error) {
	return strconv.Atoi(c.Query(key))
}

func (c *Ctx) QueryOrDefault(key, def string) string {
	val := c.Query(key)
	if val == "" {
		return def
	}
	return val
}

func (c *Ctx) QueryIntOrDefault(key string, def int) int {
	valStr := c.Query(key)
	if valStr == "" {
		return def
	}
	if val, err := strconv.Atoi(valStr); err == nil {
		return val
	}
	return def
}
