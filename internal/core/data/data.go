package data

import (
	"time"
)

type AUID string

type AuctionsDump map[AUID]InternalInfo

type InternalInfo struct {
	TTL    time.Duration `json:"TTL"`
	AUID   AUID          `json:"AUID"`
	Winner string        `json:"Winner"`
	Bet    int           `json:"Bet"`
}
