package core

import (
	"fmt"

	"github.com/auctionee/core/internal/core/data"
	"github.com/auctionee/core/pkg/models"
)

func Execute(req models.Request) error {
	if req.Bet == true && req.Start == true {
		return fmt.Errorf("error: more than one action")
	}
	if req.Bet {
		doBet(req.UserInfo.Login, req.BetInfo.Amount, data.AUID(req.BetInfo.AUID))
	}
	if req.Start {
		if err := start(req.StartInfo); err != nil {
			return err
		}
	}
	return nil
}
