package main

import (
    "log"
    "net/http"

    "github.com/labstack/echo/v5"
	"github.com/pocketbase/dbx"
    "github.com/pocketbase/pocketbase"
    "github.com/pocketbase/pocketbase/apis"
    "github.com/pocketbase/pocketbase/core"
)

func main() {
    app := pocketbase.New()

    app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
        e.Router.AddRoute(echo.Route{
            Method: http.MethodPut,
            Path:   "/api/like/:postId",
            Handler: func(c echo.Context) error {
                id := c.PathParam("postId")

				_, err := app.DB().Update("posts", dbx.Params{
					"likes": dbx.NewExp("likes + 1"),
				}, dbx.HashExp{"id": id}).Execute()

                if err != nil {
                    return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to increment likes"})
                }

                record, err := app.Dao().FindRecordById("posts", id)
                if err != nil {
                    return apis.NewNotFoundError("The record does not exist.", err)
                }

                return c.JSON(http.StatusOK, record)
            },
            Middlewares: []echo.MiddlewareFunc{
                apis.ActivityLogger(app),
                // Add authentication middleware if necessary
            },
        })

        return nil
    })

    if err := app.Start(); err != nil {
        log.Fatal(err)
    }
}
