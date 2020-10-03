package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/afief/mockidi/handler"
	"github.com/afief/mockidi/store"
)

func main() {
	ctx := context.Background()
	store := store.NewStore()
	handler := handler.NewHandlers(ctx, store)

	http.HandleFunc("/", handler.Init)
	fmt.Println("Listening :3000")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		panic(err.Error())
	}
}
