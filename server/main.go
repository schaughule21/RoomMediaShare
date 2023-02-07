package main

import (
	"RoomMediaShare/server/api"
	"fmt"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	api.Init(e)

	fmt.Errorf("Server Started on 3030")
	err := e.Start(":3030")

	if err != nil {
		fmt.Errorf("Start error : ", err)
	}
}
