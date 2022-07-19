package core

import (
	"github.com/gangcheng1030/game_script/cmd/superrobot/controller/config"
	"github.com/gangcheng1030/game_script/cmd/superrobot/model"
)

type AccountManager struct {
	accounts map[string]*model.AccountWithStatus
}

func NewAccountManager() *AccountManager {
	// TODO: 使用配置或从数据库加载
	accounts := make(map[string]*model.AccountWithStatus)
	for accountName := range config.AccountCache {
		accounts[accountName] = &model.AccountWithStatus{
			AccountName: accountName,
			Status:      model.ACCOUNT_STATUS_IDLE,
		}
	}
	return &AccountManager{
		accounts: accounts,
	}
}
