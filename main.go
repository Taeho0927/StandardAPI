package main

import (
	"fmt"
	"github.com/Taeho0927/goLogger/logger"
	"github/Taeho0927/Standard_API/conf"
	"github/Taeho0927/Standard_API/middlewares"
	"github/Taeho0927/Standard_API/routers"
	"os"
)

func main() {
	if _, err := os.Stat("logs"); os.IsNotExist(err) {
		if err := os.MkdirAll("logs", 0755); err != nil {
			logger.Error(middlewares.SetLogger("main_Mkdir(logs)_Error", err.Error()))
			return
		}
	}
	if _, err := os.Stat("files"); os.IsNotExist(err) {
		if err := os.MkdirAll("files", 0755); err != nil {
			logger.Error(middlewares.SetLogger("main_Mkdir(files)_Error", err.Error()))
			return
		}
	}
	if r, err := routers.RunAPI(); err != nil {
		logger.Error(middlewares.SetLogger("main_RunAPI()_Error", err.Error()))
		return
	} else {
		r.Run(conf.RunBackDomain())
		logger.Debug(middlewares.SetLogger("main_r.Run Excuted", fmt.Sprintf("Serving at %s", conf.RunBackDomain())))
	}
}
