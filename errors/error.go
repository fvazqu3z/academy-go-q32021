package errors

import (
	"log"
	"strconv"
	"time"
)

//ResponseError Struct to manage errors
type ResponseError struct {
	Code int			`json:"code"`
	Date time.Time		`json:"date"`
	Message  string		`json:"message"`
}

//NewResponseError method to create new response error
func NewResponseError(code int,message string, err error) error {
	log.Println(err)
	return ResponseError{
		Message: message,
		Code:    code,
		Date:    time.Now(),
	}
}

//Error default error method
func (err ResponseError) Error() string {
	return "{Code: " + strconv.Itoa(err.Code) + ", Message: " + err.Message + ", Date: " + time.Now().String() + "}"
}