package model

import "fmt"

type ErrorModel struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (me ErrorModel) Error() string {
	return fmt.Sprintf("Status: %d", me.Status)
}
