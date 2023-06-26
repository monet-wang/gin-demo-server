// pkg/server/server.go

package server

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"mark-server/internal/app/user"
	"mark-server/internal/infrastructure/database"
)

func Start() {

	// Database connection parameters
	dbHost := "localhost"
	dbPort := 3306
	dbUser := "root"
	dbPassword := "root"
	dbName := "yidan-test"

	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbName)

	// Open a connection to the MySQL database
	db, err := sql.Open("mysql", dbURI)
	if err != nil {
		fmt.Printf("Failed to connect to the database: %v\n", err)
		return
	}

	// Ping the database to check the connection
	err = db.Ping()
	if err != nil {
		fmt.Printf("Failed to ping the database: %v\n", err)
		return
	}

	defer db.Close()

	userRepo := database.NewMySQLUserRepository(db)
	userService := user.NewUserService(userRepo)
	userHandler := user.NewUserHandler(userService)

	router := gin.Default()

	// Register the user handler routes
	router.GET("/users", userHandler.GetUsers)
	router.PUT("/users/:id", userHandler.UpdateUser)
	router.DELETE("/users/:id", userHandler.DeleteUser)
	router.POST("/users/create", userHandler.CreateUser)

	router.Run(":8080")
}
