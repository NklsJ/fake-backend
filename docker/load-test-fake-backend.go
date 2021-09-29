package main

import (
	"fmt"
	"time"
	"github.com/aerogo/aero"
)

type TestObject struct {
	Subfield1 string `json:"subfield1"`
	Subfield2 string `json:"subfield2"`
}

type Response struct {
	Name string `json:"name"`
	Description string `json:"description"`
	SomeArray []TestObject `json:"someArray"`
}

type FileUploadResponse struct {
	Message string `json:"message"`
}

func main() {
	app := aero.New()

	var testObjectsArray []TestObject
	testObjectsArray = append(testObjectsArray, TestObject{
			Subfield1: "foo",
			Subfield2: "bar",
		}, TestObject{
			Subfield1: "123",
			Subfield2: "456",
		}, TestObject{
			Subfield1: "test",
			Subfield2: "test",
		},
	)

	app.Get("/json-testing-endpoint", func(ctx aero.Context) error {
		time.Sleep(80 * time.Millisecond)

		return ctx.JSON(Response{
			Name: "Load Test",
			Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
			SomeArray: testObjectsArray,
		})
	})

	app.Post("/file-upload-testing-endpoint", func(ctx aero.Context) error {
		time.Sleep(1500 * time.Millisecond)

		if len(ctx.Request().Header("X-Log-Content-Length")) > 0 {
			fmt.Println("Content-Length: ", ctx.Request().Header("Content-Length"))
		}

		return ctx.JSON(FileUploadResponse{
			Message: "ok",
		})
	})

	app.Run()
}
