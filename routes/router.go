package routes

import (
	"github.com/gorilla/sessions"
	"github.com/iwmh/go-echo-auth/model"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	"net/http"
)

func Router(app *model.Application, e *echo.Echo) {

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Static("/", "/static")

	// Create User
	e.POST("/api/users", func(c echo.Context) error {
		u := new(model.User)
		err := c.Bind(u)
		if err != nil {
			return c.JSON(http.StatusBadRequest, u)
		}
		app.DB.Create(&u)
		return c.JSON(http.StatusOK, u)
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
}
