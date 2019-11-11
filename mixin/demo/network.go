package main

import (
	"context"
	"log"

	"github.com/lilasxie/mixin-sdk/mixin"
)

func doReadNetworkInfo(ctx context.Context, user *mixin.User) {
	network, err := user.ReadNetworkInfo(ctx)
	if err != nil {
		log.Panicln(err)
	}
	printJSON("read network info", network)
}
