package ai

import (
	"fmt"
	"net/http"
)

type App struct {
	id     int
	idStr  string
	key    string
	client *http.Client
}

var (
	app = NewApp(0, "")
)

func NewApp(appid int, key string) *App {
	app := &App{
		id:     appid,
		key:    key,
		idStr:  fmt.Sprintf("%d", appid),
		client: &http.Client{},
	}
	return app
}

func ResetDefaultApp(appid int, key string) {
	app.id = appid
	app.idStr = fmt.Sprintf("%d", appid)
	app.key = key
}
