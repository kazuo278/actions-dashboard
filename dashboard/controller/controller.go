package controller

import (
	"github/dashboard/controller/rest"
	"github/dashboard/controller/websocket"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Static("/dashboard", "./static-content")
	e.GET("/ws", websocket.Socket)
	e.GET("/actions/history", rest.GetHistory)
	e.POST("/actions/history", rest.PostHistory)
	e.PUT("/actions/history", rest.PutHistory)
	e.GET("/actions/count", rest.GetHistoryCount)
	e.GET("/actions/time", rest.GetHistoryTime)
	e.POST("/actions/runner/registration-token", rest.PostRunnerRegistrationToken)
	e.Logger.Fatal(e.Start(":8080"))
}
