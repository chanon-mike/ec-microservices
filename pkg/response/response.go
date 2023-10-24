package response

import "github.com/labstack/echo/v4"

type (
	MsgResponse struct {
		Message string `json:"message"`
	}
)

// Return json of error response
func ErrResponse(c echo.Context, statusCode int, message string) error {
	return c.JSON(statusCode, &MsgResponse{Message: message})
}

// Return json of success response
func SuccessResponse(c echo.Context, statusCode int, data any) error {
	return c.JSON(statusCode, data)
}
