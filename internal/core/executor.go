package core

import (
	"fmt"

	"github.com/auctionee/core/internal/core/data"
	"github.com/auctionee/core/internal/links/auth"
	"github.com/auctionee/core/pkg/models"
)

func Execute(req models.Request) error {
	if req.Bet == true && req.Start == true {
		return fmt.Errorf("error: more than one action")
	}
	if err := auth.Permissions(req.UserInfo); err != nil {
		return err
	}
	if req.Bet {
		if err := doBet(req.UserInfo, req.BetInfo.Amount, data.AUID(req.BetInfo.AUID)); err != nil {
			return err
		}
	}
	if req.Start {
		if err := start(req.UserInfo, req.StartInfo); err != nil {
			return err
		}
	}
	return nil
}
