package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/holefillingco-ltd/Shupple-api/src/api/service"
)

type UserController struct{}

func (usercontroller UserController) Index(c *gin.Context) {
	var userService service.UserService
	p, err := userService.GetAll()

	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

func (usercontroller UserController) Create(c *gin.Context) {
	var userService service.UserService
	p, err := userService.CreateUser(c)

	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}