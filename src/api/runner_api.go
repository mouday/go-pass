package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mouday/go-pass/src/config"
	"github.com/mouday/go-pass/src/form"
	"github.com/mouday/go-pass/src/model"
	"github.com/mouday/go-pass/src/utils"
	"github.com/mouday/go-pass/src/vo"
)

type RunnerForm struct {
	RunnerId    string `json:"runnerId"`
	Title       string `json:"title"`
	Url         string `json:"url" `
	AccessToken string `json:"accessToken" `
	Status      bool   `json:"status" `
}

func AddRunner(ctx *gin.Context) {
	form := &RunnerForm{}
	ctx.BindJSON(&form)

	row := &model.RunnerModel{
		Title:       form.Title,
		Url:         form.Url,
		AccessToken: form.AccessToken,
		Status:      form.Status,
		RunnerId:    utils.GetUuidV4(),
	}

	db := config.GetDB()
	db.Model(&model.RunnerModel{}).Create(&row)

	vo.Success(ctx, nil)
}

func UpdateRunner(ctx *gin.Context) {
	form := &RunnerForm{}
	ctx.BindJSON(&form)

	data := &model.RunnerModel{
		Title:       form.Title,
		Url:         form.Url,
		AccessToken: form.AccessToken,
		Status:      form.Status,
	}

	db := config.GetDB()
	db.Model(&model.RunnerModel{}).Where("runner_id = ?", form.RunnerId).Updates(&data)

	vo.Success(ctx, nil)
}

func UpdateRunnerStatus(ctx *gin.Context) {
	form := &RunnerForm{}
	ctx.BindJSON(&form)

	db := config.GetDB()

	db.Model(&model.RunnerModel{}).Where("runner_id = ?", form.RunnerId).Update("status", form.Status)

	vo.Success(ctx, nil)
}

func RemoveRunner(ctx *gin.Context) {
	form := &RunnerForm{}
	ctx.BindJSON(&form)

	db := config.GetDB()

	db.Where("runner_id = ?", form.RunnerId).Delete(&model.RunnerModel{})

	vo.Success(ctx, nil)
}

func GetRunner(ctx *gin.Context) {
	form := &RunnerForm{}
	ctx.BindJSON(&form)

	db := config.GetDB()
	row := &model.RunnerModel{}
	db.Model(&model.RunnerModel{}).Where("runner_id = ?", form.RunnerId).Find(&row)

	vo.Success(ctx, row)
}

func GetRunnerList(ctx *gin.Context) {

	form := &form.PageForm{
		Page:   1,
		Size:   10,
		Status: 0,
	}

	ctx.BindJSON(&form)

	db := config.GetDB()

	taskList := &[]model.RunnerModel{}
	var count int64

	db.Model(&model.RunnerModel{}).Count(&count)

	db.Model(&model.RunnerModel{}).Order("id desc").Find(&taskList)

	vo.Success(ctx, gin.H{
		"list":  taskList,
		"total": count,
	})
}
