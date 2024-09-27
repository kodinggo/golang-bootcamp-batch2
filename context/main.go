package main

import (
	"context"
	"fmt"
)

type UserCtxKey string

var UserIDCtxKey UserCtxKey = "user_id"
var UsernameCtxKey UserCtxKey = "username"

type PersonCtxKey string

var PersonIDCtxKey PersonCtxKey = "user_id"

func main() {
	ctx := context.Background()

	ctx = context.WithValue(ctx, UserIDCtxKey, 123)
	ctx = context.WithValue(ctx, UsernameCtxKey, "joko")
	ctx = context.WithValue(ctx, "data", map[string]string{
		"foo": "bar",
	})

	doSomething(ctx)
}

func doSomething(ctx context.Context) {
	fmt.Println(ctx.Value(PersonIDCtxKey))
	fmt.Println(ctx.Value(UsernameCtxKey))
	data := ctx.Value("data")
	dataMap, ok := data.(map[string]string)
	if ok {
		fmt.Println(dataMap["foo"])
	}
}
