package data

import (
	"time"
)

type AUID string
type InternalInfo struct {
	TTL    time.Duration
	AUID   AUID
	Winner string
	Bet    int
}
type AuctionsDump map[AUID]InternalInfo
