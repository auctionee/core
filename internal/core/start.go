package core

import (
	"time"

	"github.com/auctionee/core/internal/core/data"
	"github.com/auctionee/core/internal/core/sync"
	"github.com/auctionee/core/internal/helper"
	"github.com/auctionee/core/internal/links/goods"
	"github.com/auctionee/core/pkg/models"
)

func start(user models.UserInfo, info models.StartInfo) error {
	//надо проверить, принадлежит ли оборудование челику
	//чекаем есть ли GUID
	if err := goods.ChekGoodOwnership(user, info.Good.Name); err != nil {
		return err
	}
	//дампаем все что есть
	sync.LocalDB.Dump()

	AUID := helper.GetAUID()
	TTL, err := time.ParseDuration(info.DurationMinutes)
	if err != nil {
		return err
	}
	//добавляем
	sync.LocalDB.Add(data.InternalInfo{
		TTL:    TTL,
		AUID:   AUID,
		Winner: "",
	})
	//по таймеру убиваем локально и вливаем новую базу
	go time.AfterFunc(TTL, func() {
		sync.LocalDB.Kill(AUID)
		sync.LocalDB.Poll()
	})
	return nil
}
