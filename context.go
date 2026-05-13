package rakhsh

import (
	"encoding/json"
	"net/http"
)

type Context struct {
	Writer  http.ResponseWriter
	Request *http.Request
	Params  map[string]string
}

func (c *Context) JSON(status int, v any) {
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(status)
	json.NewEncoder(c.Writer).Encode(v)
}

func (c *Context) String(status int, msg string) {
	c.Writer.WriteHeader(status)
	c.Writer.Write([]byte(msg))
}
