package main

import (
	"log"
	"net/http"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

func main() {
	// Initialize the PocketBase app
	app := pocketbase.New()

	// Register routes during the server's initialization
	app.OnServe().BindFunc(func(se *core.ServeEvent) error {
		// Register a PUT route for "/api/like/{postId}"
		se.Router.PUT("/api/like/{postId}", func(e *core.RequestEvent) error {
			// Get the postId from the path parameter
			postId := e.Request.PathValue("postId")

			// Increment likes for the specified postId
			_, err := e.App.DB().
				NewQuery("UPDATE posts SET likes = likes + 1 WHERE id = {id}").
				Bind(dbx.Params{"id": postId}).
				Execute()
			if err != nil {
				return e.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to increment likes"})
			}

			// Retrieve the updated record by ID
			record, err := e.App.FindRecordById("posts", postId)
			if err != nil {
				return e.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve record"})
			}

			// Return the updated record as JSON
			return e.JSON(http.StatusOK, record)
		})

		return se.Next()
	})

	// Start the PocketBase app
	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
