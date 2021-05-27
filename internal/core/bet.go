package core

import (
	"fmt"

	"github.com/auctionee/core/internal/core/data"
	"github.com/auctionee/core/internal/core/sync"
	"github.com/auctionee/core/internal/links/balance"
	"github.com/auctionee/core/pkg/models"
)

func doBet(user models.UserInfo, amount int, AUID data.AUID) error {
	//Идем в базу по ключу auid, получаем ставку, если она ниже - обновляем, в память пишем юзера
	userBalance, err := balance.GetBalance(user)
	if err != nil {
		return err
	}
	if userBalance < amount {
		return fmt.Errorf("low balance!")
	}
	sync.LocalDB.Dump()
	currentLogin, currentBet := sync.LocalDB.GetBet(AUID)
	if currentBet >= amount {
		return fmt.Errorf("your bet is lower than current")
	}
	// последовательность -
	// списываем деньги, возвращаем предыдущему, ставим.

	// если деньги не списались, не делаем ничего
	if err := balance.Charge(user.Login, amount); err != nil {
		return err
	}

	//если не смогли вернуть деньги предыдущему владельцу, то надо вернуть списанные текущему и выйти
	if err := balance.Refund(currentLogin, currentBet); err != nil {
		if err = balance.Refund(user.Login, amount); err != nil {
			return fmt.Errorf("Oops, your money are gone. We will fix it soon :(")
		}
		return err
	}
	sync.LocalDB.SetBet(AUID, amount, user.Login)

	sync.LocalDB.Poll()
	return nil
}
