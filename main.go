package main

import (
	"context"
	"feed-management/config"
	"net/http"
)

func main() {
	config.NewEnv()
	logger := config.NewLogger()
	validator := config.NewValidator()
	cc, db := config.NewDatabase()
	defer func() {
		if err := cc.Disconnect(context.Background()); err != nil {
			panic(err)
		}
	}()
	cache := config.NewCache()
	r := config.NewRouter()
	conn, ch := config.NewAmqp()
	defer conn.Close()
	defer ch.Close()
	config.Initialize(&config.Bootstrap{
		DB:        db,
		Cache:     cache,
		Logger:    logger,
		Router:    r,
		Validator: validator,
		Ch:        ch,
	})
	err := http.ListenAndServe("0.0.0.0:3003", r)
	if err != nil {
		panic(err)
	}
}
