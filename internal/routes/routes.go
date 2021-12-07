package routes

import (
	"github.com/creasty/defaults"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
	"users/internal/users"
)

func Routes() {
	router := gin.Default()
	router.POST("/person/add", addUser)
	router.PUT("/person/edit", editUser)
	router.DELETE("/person/delete", deleteUser)

	router.GET("/persons/list", listUsers)

	router.GET("/health", healthCheck)

	if err := router.Run(":8080"); err != nil {
		log.Fatal().Err(err).Msg("Unable to run routes")
	}
}

// Add user
func addUser(c *gin.Context) {
	var p users.User
	if err := c.BindJSON(&p); err != nil {
		log.Error().Err(err).Msg("Unable to addPerson")
		return
	}
	if err := defaults.Set(p); err != nil {
		log.Error().Err(err).Msg("Failed to set defaults")
		return
	}
	// TODO: Figure out how to manage lists
	//xlist.XList.AddPerson(p)
	c.IndentedJSON(http.StatusCreated, p)
}

// Edit user
func editUser(c *gin.Context) {
	var users users.User
	if err := c.BindJSON(&users); err != nil {
		log.Error().Err(err).Msg("Unable to addPerson")
		return
	}
}

// Remove user
func deleteUser(c *gin.Context) {

}

// list users
func listUsers(c *gin.Context) {

}

func healthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
