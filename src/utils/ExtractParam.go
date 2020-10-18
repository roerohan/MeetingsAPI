package utils

import (
	"net/http"
	"strings"
)

// ExtractParam the last parameter from `/something/<id>`
func ExtractParam(r *http.Request) string {
	p := strings.Split(r.URL.Path, "/")

	return p[len(p)-1]
}
