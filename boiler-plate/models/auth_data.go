package models

type AppCtxKey string

type AuthData struct {
	Authorization string `header:"Authorization"`
	ISsandBox     bool   `header:"IsSandBox"`
	UserName      string `header:"UserName"`
	UsersId       int64  `header:"UsersId"`
}
