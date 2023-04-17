package models

import (
	"log"
	"net/http"
	"todo-server/initializers"

	"github.com/gin-gonic/gin"
)

func PostTodo(c *gin.Context) {
	var Body struct {
		Title string
		Text  string
	}
	b := Body
	c.Bind(&b)
	_, err := initializers.DB.Query("INSERT INTO todos(title, text) VALUES($1, $2)", b.Title, b.Text)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, err)
	// println(stmt)
	// println(b.Title)
	// println(b.Text)

}

func GetTodo(c *gin.Context) {
	type Body struct {
		Title string `json:"Title"`
		Text  string `json:"Text"`
	}

	res := []Body{}

	rows, err := initializers.DB.Query("SELECT * FROM todos")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var title, text string
		var id int
		err = rows.Scan(&title, &text, &id)
		if err != nil {
			log.Fatal(err)
		}
		res = append(res, Body{title, text})
		// println(title, text)
	}
	c.IndentedJSON(200, res)
}

func DelTodo(c *gin.Context) {
	var Body struct {
		ID int
	}
	b := Body
	c.Bind(&b)
	println(b.ID)
	rows, err := initializers.DB.Query("DELETE FROM todos WHERE id=($1)", b.ID)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(200, err)
	defer rows.Close()
}
