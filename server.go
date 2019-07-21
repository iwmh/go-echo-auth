package main

import (
	"github.com/gorilla/sessions"
	"github.com/iwmh/go-echo-auth/model"
	"github.com/iwmh/go-echo-auth/routes"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
)

func main() {
	e := echo.New()
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))

	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=goechoauth sslmode=disable")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	db.AutoMigrate(&model.User{}, &model.Session{})

	app := &model.Application{DB: db}

	routes.Router(app, e)

	e.Logger.Fatal(e.Start(":1323"))
}
