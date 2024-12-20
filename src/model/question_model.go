package model

import "github.com/mouday/go-pass/src/utils"

const (
	QUESTION_TYPE_SINGLE   = 0
	QUESTION_TYPE_MULTIPLE = 1
)

type Option struct {
	Label  string `json:"label"`
	Status bool   `json:"status"`
}

type QuestionModel struct {
	Id            int64           `json:"id" gorm:"primaryKey"`
	Title         string          `json:"title"`
	OptionOkCount int             `json:"optionOkCount"`          // 正确选项个数
	Score         int             `json:"score" gorm:"default:0"` // 得分
	Tags          []string        `json:"tags" gorm:"serializer:json"`
	Options       []Option        `json:"options" gorm:"serializer:json"`
	Status        bool            `json:"status"`
	CreateTime    utils.LocalTime `json:"createTime" gorm:"type:datetime;autoCreateTime"`
	UpdateTime    utils.LocalTime `json:"updateTime" gorm:"type:datetime;autoUpdateTime"`
}

// 自定义表名
func (QuestionModel) TableName() string {
	return "tb_question"
}
