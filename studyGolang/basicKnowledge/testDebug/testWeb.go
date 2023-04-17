package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

type todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

func main() {
	// Connect to SQLite database
	db, err := sql.Open("sqlite3", "./todos.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create "todos" table if it doesn't exist
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS todos (id INTEGER PRIMARY KEY AUTOINCREMENT, title TEXT)")
	if err != nil {
		log.Fatal(err)
	}

	// Create Gin router
	router := gin.Default()

	// Define routes
	router.GET("/todos", func(c *gin.Context) {
		// Get all todos from database
		rows, err := db.Query("SELECT id, title FROM todos")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer rows.Close()

		// Create slice of todos
		todos := []todo{}

		// Loop through rows and add each todo to slice
		for rows.Next() {
			var t todo
			err := rows.Scan(&t.ID, &t.Title)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			todos = append(todos, t)
		}

		// Return slice of todos as JSON
		c.JSON(http.StatusOK, todos)
	})

	router.POST("/todos", func(c *gin.Context) {
		// Parse JSON request body
		var t todo
		err := c.BindJSON(&t)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Insert new todo into database
		result, err := db.Exec("INSERT INTO todos (title) VALUES (?)", t.Title)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Get ID of new todo
		id, err := result.LastInsertId()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Set ID of new todo and return it as JSON
		t.ID = int(id)
		c.JSON(http.StatusOK, t)
	})

	router.DELETE("/todos/:id", func(c *gin.Context) {
		// Get ID parameter from URL
		id := c.Param("id")

		// Delete todo from database
		_, err := db.Exec("DELETE FROM todos WHERE id = ?", id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Return success message as JSON
		c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Todo with ID %s deleted", id)})
	})

	// Start server
	err = router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}