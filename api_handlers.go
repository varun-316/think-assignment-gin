package main

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getDetails(c *gin.Context, db *sql.DB) {
	username := c.Param("username")
	c.IndentedJSON(http.StatusOK, gin.H{"data": getDetailsByUsername(username, db)})
}

// func getAuthentication(c *gin.Context, db *sql.DB) {
// 	username := c.Param("username")
// 	password := c.Param("password")
//
// 	user := getUserAuth(username, db)
// 	if password != user.Password {
// 		return
// 	}
// 	return
// }
