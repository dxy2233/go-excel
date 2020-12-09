package main

import (
	"context"
	"fmt"

	"github.com/tencentyun/scf-go-lib/cloudfunction"
)

// DefineEvent test
type DefineEvent struct {
	Key1 string `json:"key1"`
	Key2 string `json:"key2"`
}

func hello(ctx context.Context, event DefineEvent) (string, error) {
	fmt.Println("key1:", event.Key1)
	fmt.Println("key2:", event.Key2)
	return "hello", nil
}

func main() {
	cloudfunction.Start(hello)

	// r := router.Lent()
	// http.ListenAndServe(":80", r)
}
