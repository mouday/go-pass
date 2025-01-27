package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mouday/go-pass/src/api"
)

/* 注册路由 */
func RegistRouter(app *gin.Engine) {
	// question
	app.POST("/api/addQuestion", api.AddQuestion)
	app.POST("/api/updateQuestion", api.UpdateQuestion)
	// app.POST("/api/updateQuestionStatus", api.UpdateQuestionStatus)
	app.POST("/api/removeQuestion", api.RemoveQuestion)
	app.POST("/api/getQuestion", api.GetQuestion)
	app.POST("/api/getQuestionDetail", api.GetQuestionDetail)
	app.POST("/api/getQuestionList", api.GetQuestionList)
	app.POST("/api/updateQuestionAnswerResult", api.UpdateQuestionAnswerResult)

	// auth
	// app.POST("/api/login", api.Login)

	// runner
	// app.POST("/api/addRunner", api.AddRunner)
	// app.POST("/api/updateRunner", api.UpdateRunner)
	// app.POST("/api/updateRunnerStatus", api.UpdateRunnerStatus)
	// app.POST("/api/removeRunner", api.RemoveRunner)
	// app.POST("/api/getRunner", api.GetRunner)
	// app.POST("/api/getRunnerList", api.GetRunnerList)

	// task
	// app.POST("/api/addTask", api.AddTask)
	// app.POST("/api/updateTask", api.UpdateTask)
	// app.POST("/api/updateTaskStatus", api.UpdateTaskStatus)
	// app.POST("/api/removeTask", api.RemoveTask)
	// app.POST("/api/getTask", api.GetTask)
	// app.POST("/api/getTaskList", api.GetTaskList)
	// app.POST("/api/runTask", api.RunTask)
	// app.POST("/api/startTask", api.StartTask)
	// app.POST("/api/stopTask", api.StopTask)

	// log
	// app.POST("/api/getTaskLogList", api.GetTaskLogList)
	// app.POST("/api/reportTaskStatus", api.ReportTaskStatus)
	// app.POST("/api/getTaskLogDetail", api.GetTaskLogDetail)
}
