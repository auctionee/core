package main

import (
	"context"
	"log"

	"github.com/auctionee/core/internal"
	"github.com/auctionee/core/pkg/server"
	"golang.org/x/sync/errgroup"
)

func main() {
	cfg := internal.Config{}
	s := server.NewServer(&cfg)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		return s.ListenAndServe()
	})
	log.Fatal(g.Wait())
}
