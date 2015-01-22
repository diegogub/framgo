package framgo

import (
	"github.com/gorilla/schema"
)

var decoder = schema.NewDecoder()

func MapPost(i interface{}, post map[string][]string) error {
	return decoder.Decode(&i, post)
}
