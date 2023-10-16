package controller

import (
	"fmt"
	"log"

	"examplo.com/crud-project/src/configuration/rest_err/validation"
	"examplo.com/crud-project/src/controller/model/request"
	"examplo.com/crud-project/src/controller/model/response"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {

		log.Printf("There are some incorrect filds, error=%s\n", err.Error())
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	fmt.Println(userRequest)
	response := response.UserResponse{
		ID:    "test",
		Email: userRequest.Email,
		Name:  userRequest.Name,
		Age:   userRequest.Age,
	}
	return
}
