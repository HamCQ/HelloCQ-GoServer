package main

import (
	"hellocq/bootstrap"
	"hellocq/common/conf"
	"hellocq/common/log"

	"hellocq/app/server/routers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	bootstrap.Init()

	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())
	e.Logger.SetOutput(log.LogIO)

	r := routers.Router{Log: log.NewLog()}

	r.SetRouter(e)

	log.NewLog().Infoln("starting server at:", conf.AppConf.App.HostPort)
	e.Start(conf.AppConf.App.HostPort)
}
