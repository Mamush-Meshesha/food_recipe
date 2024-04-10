package actions

import (
	"log"

	"github.com/Mamush-Meshesha/backend/client"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	body := SignupInputBody{}
	err := c.ShouldBindJSON(&body)
	if err != nil {
		log.Println("One", err.Error())
		c.JSON(401, gin.H{
			"message": "Invalid input request",
		})
	}

	user := GetUserByEmailQuery{}
	err = client.Query(&user, map[string]interface{}{
		"email": body.Email,
	})
	if err != nil {
		log.Println("Two", err.Error())

		c.JSON(401, gin.H{
			"error": err.Error(),
		})
	}

	if len(user.GetUserByEmail) > 0 {
		log.Println("three", err.Error())

		c.JSON(401, gin.H{
			"message": "user already exists",
		})
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), 8)
	if err != nil {
		log.Println("four", err.Error())
		c.JSON(501, gin.H{
			"error": err.Error(),
		})
	}
	body.Password = string(hashedPassword)
	insertUser := InsertUserOneMutation{}
	variables := map[string]interface{}{
		"email":      body.Email,
		"first_name": body.FirstName,
		"last_name":  body.LastName,
		"password":   string(hashedPassword),
	}
	err = client.Mutation(&insertUser, variables)

	if err != nil {
		log.Println("five", err.Error())
		c.JSON(501, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(201, gin.H{
		"success": "success",
	})
}

type SignupInputBody struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}
type InsertUserOneMutation struct {
	InsertUserOne struct {
		Id    int64  `graphql:"id"`
		Email string `graphql:"email"`
	} `graphql:"insert_users_one(object: {email: $email, first_name: $first_name, last_name: $last_name, password: $password})"`
}
