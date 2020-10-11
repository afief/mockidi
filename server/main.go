package main

import (
	"fmt"
	"net/http"

	"github.com/afief/mockidi/handler"
	"github.com/afief/mockidi/store"
)

func main() {
	store := store.NewStore()

	http.HandleFunc("/", handler.NewHandlers(store))
	fmt.Println("Listening :3000")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		panic(err.Error())
	}
}
