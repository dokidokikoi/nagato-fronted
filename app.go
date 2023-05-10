package main

import (
	"context"
	"fmt"
	initTask "nagato/internal/init"
	"nagato/internal/model"
	"nagato/internal/service"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// App struct
type App struct {
	ctx     context.Context
	service service.Service
}

// NewApp creates a new App application struct
func NewApp() *App {
	initTask.Init("./application.yml")
	return &App{service: service.NewService()}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	_, err := a.service.Conf().GetConfig(ctx)
	if err != nil {
		panic(err)
	}
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) GetFolderChildren(uid string) gin.H {
	fileSrv := service.NewService().File()
	data, code, err := fileSrv.GetFolderByUid(uid)
	if err != nil {
		return gin.H{
			"code": code,
			"msg":  err.Error(),
		}
	}
	return gin.H{
		"code": code,
		"data": data,
	}
}

func (a *App) Upload(filePath, puuid string) gin.H {
	if filePath == "" {
		return gin.H{
			"code": -1,
			"msg":  "文件不存在",
		}
	}
	fileSrv := service.NewService().File()
	slice := strings.Split(filePath, "/")
	code, err := fileSrv.Upload(slice[len(slice)-1], filePath, puuid)
	if err != nil {
		return gin.H{
			"code": code,
			"msg":  err.Error(),
		}
	}
	return gin.H{
		"code": code,
		"data": nil,
	}
}

func (a *App) CreateDir(dirName, puuid string) gin.H {
	fileSrv := service.NewService().File()
	code, err := fileSrv.CreateDir(dirName, puuid)
	if err != nil {
		return gin.H{
			"code": code,
			"msg":  err.Error(),
		}
	}
	return gin.H{
		"code": code,
		"data": nil,
	}
}

func (a *App) CreateBlank(example model.CreateBlankTemp) gin.H {
	fileSrv := service.NewService().Blank()
	p, err := example.CreateBlank()
	if err != nil {
		if err != nil {
			return gin.H{
				"code": -1,
				"msg":  err.Error(),
			}
		}
	}
	code, err := fileSrv.CreateBlank(*p)
	if err != nil {
		return gin.H{
			"code": code,
			"msg":  err.Error(),
		}
	}
	return gin.H{
		"code": code,
		"data": nil,
	}
}

func (a *App) BlankList() gin.H {
	data, code, err := service.NewService().Blank().List()
	if err != nil {
		return gin.H{
			"code": code,
			"msg":  err.Error(),
		}
	}
	return gin.H{
		"code": code,
		"data": data,
	}
}

func (a *App) GetBlank(idStr string) gin.H {
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return gin.H{
			"code": -1,
			"msg":  "blank不存在",
		}
	}
	data, code, err := service.NewService().Blank().Get(uint(id))
	if err != nil {
		return gin.H{
			"code": code,
			"msg":  err.Error(),
		}
	}
	return gin.H{
		"code": code,
		"data": data,
	}
}
