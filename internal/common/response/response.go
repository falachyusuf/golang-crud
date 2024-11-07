package response

import (
	"github.com/gofiber/fiber/v2"
	"google.golang.org/protobuf/proto"
)

type response struct {
	Data 		interface{} `json:"data,omitempty"`
	Error 		*int32 		`json:"error,omitempty"`
	ErrorMsg 	*string 		`json:"error_msg,omitempty"`
}

// Create response success
func Success(c *fiber.Ctx, req, data interface{}) error {
	res := &response{
        Data: data,
    }
    return c.Status(fiber.StatusOK).JSON(res)
}

// Create response error
func Error(c *fiber.Ctx, req interface{}, errCode int32, errMsg string) error {
	res := response{
		Error:    proto.Int32(errCode),
		ErrorMsg: proto.String(errMsg),
	}
	return c.Status(fiber.StatusOK).JSON(res)
}