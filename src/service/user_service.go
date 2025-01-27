package service

import (
	"fmt"

	"github.com/mouday/go-pass/src/config"
	"github.com/mouday/go-pass/src/model"
	"github.com/mouday/go-pass/src/utils"
)

func AddUser() {
	db := config.GetDB()
	userRow := model.UserModel{
		UserId:   utils.GetUuidV4(),
		Username: "admin",
		Password: utils.EncodePassword("123456"),
		Status:   true,
	}

	db.Model(&model.UserModel{}).Create(userRow)
	fmt.Println("userRow:", userRow)
}
