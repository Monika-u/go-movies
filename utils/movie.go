package utils

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ParseBody(r *http.Request, x interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}

func SetContext(c *gin.Context, ctx context.Context) context.Context {
	for k, v := range c.Keys {
		ctx = context.WithValue(ctx, k, v)
	}
	return ctx
}
