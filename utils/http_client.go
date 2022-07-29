package utils

import (
	"github.com/gofiber/fiber/v2"
)

func HttpClient(method string, URI string) (err error, status int, body []byte) {
	status = fiber.StatusOK

	a := fiber.AcquireAgent()
	req := a.Request()

	req.Header.SetMethod(method)
	req.Header.Set("accept", "application/json")
	req.SetRequestURI(URI)

	err = a.Parse()
	if err != nil {
		status = fiber.StatusInternalServerError
		return
	}

	code, body, errs := a.Bytes()
	if errs != nil {
		err = errs[0]
		status = code
		return
	}

	return
}
