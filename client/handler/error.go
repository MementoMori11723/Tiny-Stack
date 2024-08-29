package handler

import (
	"context"
	"net/http"

	"github.com/MementoMori11723/Tiny-Stack/client/pages"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request) {
  error := pages.Error()
  error.Render(context.Background(), w)
}
