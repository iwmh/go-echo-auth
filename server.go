package main

import (
	"github.com/gorilla/sessions"
	"github.com/iwmh/go-echo-auth/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	"net/http"
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

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/login", func(c echo.Context) error {
		sess, _ := session.Get("ses_echo", c)
		sess.Options = &sessions.Options{
			Path:     "/",
			MaxAge:   86400 * 7,
			HttpOnly: true,
		}
		sess.Values["foo"] = "bar"
		sess.Save(c.Request(), c.Response())
		return c.NoContent(http.StatusOK)
	})
	type Cookie struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	}
	e.GET("/cookie", func(c echo.Context) error {
		cookie, err := c.Cookie("ses_echo")
		if err != nil {
			return err
		}
		co := &Cookie{
			Name:  cookie.Name,
			Value: cookie.Value,
		}
		return c.JSON(http.StatusOK, co)
	})
	e.Logger.Fatal(e.Start(":1323"))
}
