package model

import "github.com/mouday/go-pass/src/utils"

type AnswerLogModel struct {
	Id        uint   `json:"id"`
	QuestionId string `json:"questionId"`
	Answer    string `json:"answer"`
	Status int              `json:"status"`
	CreateTime utils.LocalTime `gorm:"type:datetime;autoCreateTime" json:"createTime"`
	UpdateTime utils.LocalTime `gorm:"type:datetime;autoUpdateTime" json:"updateTime"`
}

// 自定义表名
func (AnswerLogModel) TableName() string {
	return "tb_answer_log"
}
