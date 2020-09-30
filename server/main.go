package main

import (
	"context"
	"net/http"

	"github.com/afief/mockidi/handler"
	"github.com/afief/mockidi/store"
)

func main() {
	ctx := context.Background()
	store := store.NewStore()
	handler := handler.NewHandlers(ctx, store)

	http.HandleFunc("/", handler.Init)
	http.ListenAndServe(":3000", nil)
}
