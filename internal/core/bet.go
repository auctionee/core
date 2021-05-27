package core

import (
	"github.com/auctionee/core/internal/core/data"
	"github.com/auctionee/core/internal/core/sync"
)

func doBet(who string, amount int, AUID data.AUID) {
	//Идем в базу по ключу auid, получаем ставку, если она ниже - обновляем, в память пишем юзера
	sync.LocalDB.Dump()
	best := sync.LocalDB.GetBet(AUID)
	if best >= amount {
		return
	}

	sync.LocalDB.SetBet(AUID, amount, who)

	sync.LocalDB.Poll()

	//здесь идем в сервис денег и вычитаем масленку копеечку
	//не забыть вернуть, если перебили
}
