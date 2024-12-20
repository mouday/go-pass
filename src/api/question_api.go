package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mouday/go-pass/src/config"
	"github.com/mouday/go-pass/src/form"
	"github.com/mouday/go-pass/src/model"
	"github.com/mouday/go-pass/src/vo"
	"gorm.io/gorm"
)

func getQuestionOptionOkCount(options []model.Option) int {

	total := 0
	for _, option := range options {
		if option.Status == true {
			total++
		}
	}

	return total

}

func AddQuestion(ctx *gin.Context) {
	form := model.QuestionModel{
		Tags: []string{},
	}
	ctx.BindJSON(&form)
	form.OptionOkCount = getQuestionOptionOkCount(form.Options)

	db := config.GetDB()
	db.Model(&model.QuestionModel{}).Create(&form)

	vo.Success(ctx, form)
}

func UpdateQuestion(ctx *gin.Context) {
	form := model.QuestionModel{}

	ctx.BindJSON(&form)
	form.OptionOkCount = getQuestionOptionOkCount(form.Options)

	db := config.GetDB()
	db.Model(&model.QuestionModel{}).Where("id = ?", form.Id).Updates(&form)

	vo.Success(ctx, nil)
}

type QuestionAnswerResultForm struct {
	Id     uint `json:"id"`
	Result bool `json:"result"`
}

func UpdateQuestionAnswerResult(ctx *gin.Context) {
	form := QuestionAnswerResultForm{}

	ctx.BindJSON(&form)

	db := config.GetDB()

	if form.Result {
		db.Model(&model.QuestionModel{}).Where("id = ?", form.Id).Update("score", gorm.Expr("`score` + ?", 1))
	} else {
		db.Model(&model.QuestionModel{}).Where("id = ?", form.Id).Update("score", gorm.Expr("`score` - ?", 3))
	}

	vo.Success(ctx, nil)
}

func UpdateQuestionStatus(ctx *gin.Context) {
	params := &model.QuestionModel{}
	ctx.BindJSON(&params)

	db := config.GetDB()

	db.Model(&model.QuestionModel{}).Where("id = ?", params.Id).Update("status", params.Status)

	// err := service.ChangeQuestionStatus(params.QuestionId, params.Status)

	// if err != nil {
	// 	vo.Error(ctx, -1, err.Error())
	// } else {

	// }

	vo.Success(ctx, nil)

}

func RemoveQuestion(ctx *gin.Context) {
	row := &model.QuestionModel{}
	ctx.BindJSON(&row)

	db := config.GetDB()

	db.Where("id = ?", row.Id).Delete(&model.QuestionModel{})

	vo.Success(ctx, nil)
}

func GetQuestion(ctx *gin.Context) {

	from := model.QuestionModel{}
	ctx.BindJSON(&from)

	db := config.GetDB()
	row := model.QuestionModel{}

	db.Model(&model.QuestionModel{}).Where("id = ?", from.Id).Find(&row)

	if row.Id == 0 {
		vo.Error(ctx, 404, "Not Found")
	} else {
		vo.Success(ctx, row)
	}
}

type QuestionDetailResponse struct {
	Data         model.QuestionModel `json:"data"`
	Total        int64               `json:"total"`
	NextId       int64               `json:"nextId"`
	SuccessCount int64               `json:"successCount"`
	ErrerCount   int64               `json:"errerCount"`
}

func GetQuestionDetail(ctx *gin.Context) {

	from := model.QuestionModel{}
	ctx.BindJSON(&from)

	db := config.GetDB()
	row := model.QuestionModel{}

	resposne := QuestionDetailResponse{}

	db.Model(&model.QuestionModel{}).Where("id = ?", from.Id).Find(&row)
	resposne.Data = row

	if row.Id != 0 {

		var total int64
		var successCount int64
		var errerCount int64

		db.Model(&model.QuestionModel{}).Count(&total)
		db.Model(&model.QuestionModel{}).Where("score > 0").Count(&successCount)
		db.Model(&model.QuestionModel{}).Where("score < 0").Count(&errerCount)
		resposne.Total = total
		resposne.SuccessCount = successCount
		resposne.ErrerCount = errerCount

		nextRow := model.QuestionModel{}
		db.Model(&model.QuestionModel{}).Where("id != ?", row.Id).Order("score asc, update_time asc, id desc").First(&nextRow)

		resposne.NextId = nextRow.Id

		vo.Success(ctx, resposne)
	} else {
		vo.Error(ctx, 404, "Not Found")

	}
}

func GetQuestionList(ctx *gin.Context) {

	params := &form.PageForm{}

	ctx.BindJSON(&params)

	db := config.GetDB()

	QuestionList := []model.QuestionModel{}
	var count int64

	db.Model(&model.QuestionModel{}).Count(&count)

	db.Model(&model.QuestionModel{}).Order("score asc, update_time desc, id desc").Limit(params.GetSize()).Offset(params.PageOffset()).Find(&QuestionList)

	vo.Success(ctx, gin.H{
		"list":  QuestionList,
		"total": count,
	})
}
