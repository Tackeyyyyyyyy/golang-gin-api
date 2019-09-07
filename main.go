package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"bytes"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// db connection
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/name_list")
	if err != nil {
		fmt.Print(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Print(err.Error())
	}
	type Name struct {
		Id         int
		First_Name string
		Last_Name  string
	}
	router := gin.Default()

	v1 := router.Group("/v1")

	// GET
	v1.GET("/names", func(c *gin.Context) {
		var (
			name  Name
			names []Name
		)
		rows, err := db.Query("select id, first_name, last_name from names;")
		if err != nil {
			fmt.Print(err.Error())
		}
		for rows.Next() {
			err = rows.Scan(&name.Id, &name.First_Name, &name.Last_Name)
			names = append(names, name)
			if err != nil {
				fmt.Print(err.Error())
			}
		}
		defer rows.Close()
		c.JSON(http.StatusOK, gin.H{
			"result": names,
			"count":  len(names),
		})
	})

	// POST
	v1.POST("/name", func(c *gin.Context) {
		var buffer bytes.Buffer
		firstName := c.PostForm("first_name")
		lastName := c.PostForm("last_name")
		stmt, err := db.Prepare("insert into names (first_name, last_name) values(?,?);")
		if err != nil {
			fmt.Print(err.Error())
		}
		_, err = stmt.Exec(firstName, lastName)

		if err != nil {
			fmt.Print(err.Error())
		}

		buffer.WriteString(firstName)
		buffer.WriteString(" ")
		buffer.WriteString(lastName)
		defer stmt.Close()
		name := buffer.String()
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf(" %s successfully created", name),
		})
	})

	router.Run(":3000")
}
