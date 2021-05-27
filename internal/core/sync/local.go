package sync

import (
	"sync"

	"github.com/auctionee/core/internal/core/data"
)

type LocalDump struct {
	mu   *sync.Mutex
	dump data.AuctionsDump
}
type AuctionsList struct {
	List LocalDump `json:"List"`
}

var LocalDB = LocalDump{}

func (l *LocalDump) Add(data data.InternalInfo) {
	l.mu.Lock()
	l.dump[data.AUID] = data
	l.mu.Unlock()
}
func (l *LocalDump) Kill(AUID data.AUID) {
	l.mu.Lock()
	l.dump[AUID] = data.InternalInfo{}
	l.mu.Unlock()
}
func (l *LocalDump) Dump() {
	l.mu.Lock()
	//собираем с сервера
	l.mu.Unlock()
}
func (l *LocalDump) Poll() {
	l.mu.Lock()
	//вкачиваем на сервер
	l.mu.Unlock()
}
func (l *LocalDump) GetBet(AUID data.AUID) (string, int) {
	return l.dump[AUID].Winner, l.dump[AUID].Bet
}
func (l *LocalDump) SetBet(AUID data.AUID, amount int, who string) {
	l.mu.Lock()
	tmp := l.dump[AUID]
	tmp.Bet = amount
	tmp.Winner = who
	l.dump[AUID] = tmp
	l.mu.Unlock()
}
