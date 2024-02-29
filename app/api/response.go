package api

import (
	"time"
)

type ResponseJson struct {
	Time  time.Time `json:"time"`
	Code  int       `json:"code"`
	Msg   string    `json:"msg"`
	Error string    `json:"error"`
	Data  any       `json:"data"`
}
