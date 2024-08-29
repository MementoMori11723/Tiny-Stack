package handler

import (
	"context"
	"net/http"

	"github.com/MementoMori11723/Tiny-Stack/client/pages"
)

func AboutHandler(w http.ResponseWriter, r *http.Request) {
  about := pages.About()
  about.Render(context.Background(), w)
}
