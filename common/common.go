package common

import "log"

type JsonResult struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func WriteLogErr(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

const (
	Success = 0
	Failure = 1
)
