package util

import (
	"assignment2/model"
)

func PrintResponse(success bool, data any, err string)model.Response {
	return model.Response{
		Success: success,
		Data:    data,
		Error:   err,
	}
}