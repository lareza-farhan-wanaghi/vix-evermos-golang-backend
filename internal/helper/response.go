package helper

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// @TODO : make helper response

type JSONResp struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Errors  []string    `json:"errors"`
	Data    interface{} `json:"data"`
}

type JSONRespArgs struct {
	Ctx        *fiber.Ctx
	StatusCode int
	Errors     []string
	Data       interface{}
}

// ResponseWithJSON responses to the request with json format data
func ResponseWithJSON(args *JSONRespArgs) error {
	hasAnError := args.Errors != nil
	messagePrefix := "Succeed"
	if hasAnError {
		messagePrefix = "Failed"
	}
	message := fmt.Sprintf("%s to %s data", messagePrefix, strings.ToUpper(args.Ctx.Method()))

	return args.Ctx.Status(args.StatusCode).JSON(&JSONResp{
		Status:  !hasAnError,
		Message: message,
		Errors:  args.Errors,
		Data:    args.Data,
	})
}
