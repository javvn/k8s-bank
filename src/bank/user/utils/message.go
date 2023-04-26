package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"jawncorp.com/bank/user/domain/model"
)

var messageConfig map[string]interface{}

func init() {
	err := json.Unmarshal([]byte(MessagesConfig), &messageConfig)

	if err != nil {
		panic(err)
	}
}
func GetMessageStatusCode(messageCode string) int {
	return int((messageConfig[messageCode].(map[string]interface{}))["statusCode"].(float64))
}

func GetMessageMessageCode(messageCode string) string {
	return (messageConfig[messageCode].(map[string]interface{}))["messageCode"].(string)
}

func GetMessageMessage(lang string, messageCode string, args ...interface{}) string {
	llang := lang

	if !(llang == "en" || llang == "ko") {
		llang = "ko"
	}

	return fmt.Sprintf(((messageConfig[messageCode].(map[string]interface{}))["message"].(map[string]interface{}))[llang].(string), args...)
}

func GetErrorMessage(context interface{}, lang string, err1 error) (err error) {
	c := context.(echo.Context)

	messageCode := err1.Error()

	errorMessages := &model.ErrorMessages{
		Code:    GetMessageMessageCode(messageCode),
		Message: GetMessageMessage(lang, messageCode),
	}

	return c.JSON(GetMessageStatusCode(messageCode), errorMessages)
}

func SetErrorMessage(messageCode string) (err error) {
	err = errors.New(messageCode)

	return
}
