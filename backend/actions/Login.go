package actions

import (
	"log"
	"strconv"
	"time"

	"github.com/Mamush-Meshesha/backend/client"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
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

	claims["exp"] = time.Now().Add(time.Hour * 7).Unix()
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
