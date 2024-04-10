package actions

import (
	"log"

	"github.com/Mamush-Meshesha/backend/client"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Reset(c *gin.Context) {
	body := ResetInputBody{}
	err := c.ShouldBindJSON(&body)

	if err != nil {
		log.Println("error", err)
		c.JSON(404, gin.H{
			"message": "error resetting password",
		})
		return
	}

	log.Println("Reset request received for UserID:", body.Id)

	user := GetUserByIdQuery{}
	err = client.Query(&user, map[string]interface{}{
		"id": body.Id,
	})

	if err != nil {
		log.Println("error", err)
		c.JSON(401, gin.H{
			"message": err.Error(),
		})
		return
	}

	if len(user.GetUserById) == 0 {
		log.Println("User not found for UserID:", body.Id)
		c.JSON(404, gin.H{
			"message": "user not found",
		})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), 8)
	if err != nil {
		log.Println("error", err)
		c.JSON(500, gin.H{
			"message": "error hashing the new password",
		})
		return
	}

	updateUserPassword := UpdateUserPasswordMutation{}
	updateVariables := map[string]interface{}{
		"id":       body.Id,
		"password": string(hashedPassword),
	}

	err = client.Mutation(&updateUserPassword, updateVariables)

	if err != nil {
		log.Println("error", err.Error())
		c.JSON(403, gin.H{
			"message": "unable to update user password",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "password reset successful",
	})
}

type UpdateUserPasswordMutation struct {
	UpdateUserByPk struct {
		Id int64 `graphql:"id"`
	} `graphql:"update_users_by_pk(pk_columns: {id: $id},_set: {password: $password})"`
}

type GetUserByIdQuery struct {
	GetUserById []struct {
		Id int64 `json:"id"`
	} `graphql:"users(where: {id: {_eq: $id}})"`
}

type ResetInputBody struct {
	Password string `json:"password"`
	Id       int64  `json:"id"`
}
