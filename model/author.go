package model

import (
	"context"
	"time"

	"github.com/UmaruCMS/article-system/config"
	"github.com/UmaruCMS/article-system/rpc/client/user"
)

type Author struct {
	UserID uint
	Name   string
	Email  string
}

func (author *Author) GetByUserID(userID uint) (*Author, error) {
	rpc := config.RPC
	userInfo := &user.UserInfo{
		Id: uint32(userID),
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	userInfo, err := rpc.UserClient.GetUserInfoByID(ctx, userInfo)
	if err != nil {
		return nil, err
	}
	author.UserID = userID
	author.Name = userInfo.Name
	author.Email = userInfo.Email
	return author, nil
}
