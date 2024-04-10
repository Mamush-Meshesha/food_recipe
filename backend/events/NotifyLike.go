package events

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Mamush-Meshesha/backend/client"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gopkg.in/gomail.v2"
)

func Notify(c *gin.Context) {
	var eventPayload Payload
	if err := c.BindJSON(&eventPayload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON payload"})
		return
	}
	log.Println(eventPayload.Event.Data.New)

	recipeID := eventPayload.Event.Data.New.RecipeID

	user, err := getOwnerEmailFromDatabase(recipeID)
	if err != nil {
		log.Println("Error retrieving owner's email:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve owner's email"})
		return
	}

	var userQuery struct {
		User []User `graphql:"users(where: {id: {_eq: $id}})"`
	}
	queryVars := map[string]interface{}{
		"id": eventPayload.Event.Data.New.UserID,
	}
	err = client.Query(&userQuery, queryVars)
	if err != nil {
		log.Println("Error querying user:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query user"})
		return
	}
	if len(userQuery.User) == 0 {
		log.Println("User not found:", user.ID)
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	err = sendEmailNotification(user.Email, fmt.Sprintf("%s %s", userQuery.User[0].FirstName, userQuery.User[0].LastName))
	if err != nil {
		log.Println("Error sending email notification:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send email notification"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Email notification sent to recipe owner"})
}

func getOwnerEmailFromDatabase(recipeID int) (*User, error) {
	variables := map[string]interface{}{
		"id": recipeID,
	}
	var queryResponse RecipeUser
	err := client.Query(&queryResponse, variables)
	if err != nil {
		return nil, err
	}

	if len(queryResponse.Recipe) > 0 {

		return &queryResponse.Recipe[0].User, nil
	}
	return nil, errors.New("Recipe owner email not found")
}

func sendEmailNotification(email string, fullName string) error {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	application_password := os.Getenv("APPLICATION_PASSWORD")

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", "mam1620she@gmail.com")
	mailer.SetHeader("To", email)
	mailer.SetHeader("Subject", "New Like Notification")
	mailer.SetBody("text/plain", fmt.Sprintf("Your recipe has been liked by a user %s!", fullName))

	dialer := gomail.NewDialer("smtp.gmail.com", 587, "mam1620she@gmail.com", application_password)

	err = dialer.DialAndSend(mailer)
	if err != nil {
		log.Println("Error sending email notification:", err)
		return err
	}
	return nil
}

type Payload struct {
	Event Event `json:"event"`
}
type Event struct {
	Data             Data             `json:"data"`
	Op               string           `json:"op"`
	SessionVariables SessionVariables `json:"session_variables"`
}

type Data struct {
	New NewData     `json:"new"`
	Old interface{} `json:"old"`
}

type NewData struct {
	CreatedAt time.Time `json:"created_at"`
	ID        int       `json:"id"`
	RecipeID  int       `json:"recipe_id"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    int       `json:"user_id"`
}

type SessionVariables struct {
	XHasuraRole   string `json:"x-hasura-role"`
	XHasuraUserID string `json:"x-hasura-user-id"`
}

type RecipeUser struct {
	Recipe []struct {
		User  User   `graphql:"user"`
		ID    int    `json:"id"`
		Title string `json:"title"`
	} `graphql:" recipe(where: {id: {_eq: $id}})"`
}
type User struct {
	LastName  string `graphql:"last_name"`
	FirstName string `graphql:"first_name"`
	Email     string `json:"email"`
	ID        int    `json:"id"`
}
