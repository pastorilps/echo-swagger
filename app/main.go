package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
	_ "github.com/pastorilps/echo-swagger/app/docs"
	middleware_cors "github.com/pastorilps/echo-swagger/middleware"
	_userHttpDelivery "github.com/pastorilps/echo-swagger/users/delivery/http"
	_userRepo "github.com/pastorilps/echo-swagger/users/repository"
	_userUseCase "github.com/pastorilps/echo-swagger/users/usecase"
	"github.com/subosito/gotenv"
)

type Env struct {
	db *sql.DB
}

func init() {
	gotenv.Load()
}

// @title Echo Swagger Example API
// @version 1.0
// @description This is a sample server server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3001
// @BasePath /v1
// @schemes http

func main() {
	var (
		dbHost = os.Getenv("HOST_DATABASE")
		dbPort = os.Getenv("PORT_DATABASE")
		dbUser = os.Getenv("USER_DATABASE")
		dbPass = os.Getenv("PASS_DATABASE")
		dbName = os.Getenv("DB_DATABASE")
	)

	conn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable  ", dbHost, dbPort, dbUser, dbPass, dbName)
	dbConn, err := sql.Open("postgres", conn)
	if err != nil {
		log.Fatalln(err)
	}

	err = dbConn.Ping()
	if err != nil {
		log.Fatalln(err)

	}

	defer func() {
		err := dbConn.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}()

	defer dbConn.Close()

	// Echo instance
	e := echo.New()

	// Middleware
	mid := middleware_cors.InitMiddleware()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(mid.CORS)

	userRepo := _userRepo.NewUserRepo(dbConn)
	userUseCase := _userUseCase.NewUserUseCase(userRepo)
	_userHttpDelivery.NewUserHandler(e, userUseCase)

	// Start server
	e.Logger.Fatal(e.Start(":3001"))
}
