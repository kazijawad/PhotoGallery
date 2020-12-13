package controllers

import (
	"net/http"

	"github.com/gorilla/schema"
)

// parseForm uses gorilla/schema to decode form data.
func parseForm(r *http.Request, dst interface{}) error {
	if err := r.ParseForm(); err != nil {
		return err
	}
	dec := schema.NewDecoder()
	dec.IgnoreUnknownKeys(true)
	if err := dec.Decode(dst, r.PostForm); err != nil {
		return err
	}
	return nil
}
