package handler

import (
	"context"
	"net/http"

	"github.com/MementoMori11723/Tiny-Stack/client/pages"
)

func ExampleHandler(w http.ResponseWriter, r *http.Request) { 
  example := pages.Example()
  example.Render(context.Background(), w)
}
