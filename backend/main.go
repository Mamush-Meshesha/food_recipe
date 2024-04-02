package main

import (
	"log"
	"strconv"
	"time"

	"github.com/Mamush-Meshesha/backend/client"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(201, gin.H{
			"message": " you did it!",
		})
	})

	router.POST("/register", func(c *gin.Context) {
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
	})

	router.POST("/login", func(c *gin.Context) {
		body := LoginInputBody{}
		err := c.ShouldBindJSON(&body)

		if err != nil {
			log.Println("error", err.Error())
			c.JSON(405, gin.H{
				"error": err.Error(),
			})
		}
		user := GetUserByEmailQuery{}
		err = client.Query(&user, map[string]interface{}{
			"email": body.Email,
		})
		if err != nil {
			log.Println("error", err.Error())
		}

		if len(user.GetUserByEmail) < 0 {
			c.JSON(401, gin.H{
				"message": "user by this email not found",
			})
			return
		}

		isPassMatch := bcrypt.CompareHashAndPassword([]byte(user.GetUserByEmail[0].Password), []byte(body.Password)) == nil
		if !isPassMatch {
			c.JSON(401, gin.H{
				"message": "invalid password",
			})
			return
		}
		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)

		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
		claims["iat"] = time.Now().Unix()

		hasuraClaims := map[string]interface{}{
			"x-hasura-allowed-roles": []string{"user", "adminstrator"},
			"x-hasura-default-role":  "user",
			"x-hasura-user-id":       strconv.Itoa(int(user.GetUserByEmail[0].Id)),
		}
		claims["https://hasura.io/jwt/claims"] = hasuraClaims
		tokenString, err := token.SignedString([]byte("0B5G1MO4DK21UihMu8NM5g30Ml5NARjGmiDxve15aHovMqr1XWEHALxlMLQbt3CVENpPBzxVthSc3hqu82Xf"))

		if err != nil {
			c.JSON(201, gin.H{
				"status": "error occured",
			})
		}
		c.JSON(200, gin.H{
			"accessToken": tokenString,
		})
	})

	router.POST("/reset", func(c *gin.Context) {
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
	})

	router.Run(":8480")
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

type UpdateUserPasswordMutation struct {
	UpdateUserByPk struct {
		Id int64 `graphql:"id"`
	} `graphql:"update_users_by_pk(pk_columns: {id: $id},_set: {password: $password})"`
}

type GetUserByEmailQuery struct {
	GetUserByEmail []struct {
		Id        int64  `json:"id"`
		Email     string `json:"email"`
		Password  string `graphql:"password"`
		FirstName string `graphql:"first_name"`
		LastName  string `graphql:"last_name"`
	} `graphql:"users(where: {email: {_eq: $email}})"`
}

type LoginInputBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GetUserByIdQuery struct {
	GetUserById []struct {
		Id int64 `json:"id"`
	}`graphql:"users(where: {id: {_eq: $id}})"`
}

type ResetInputBody struct {
	Password string `json:"password"`
	Id       int64  `json:"id"`
}
