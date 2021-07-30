package main

import (
	"Problem5/models"
	"Problem5/service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getArticles(c *gin.Context) {
	articles, err := service.GetArticles(c)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.IndentedJSON(http.StatusOK, articles)
	}
}

func getArticleById(c *gin.Context) {
	article, err := service.GetArticleById(c)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
	} else {
		c.IndentedJSON(http.StatusOK, article)
	}
}

func addArticle(c *gin.Context) {
	var article models.Article
	c.BindJSON(&article)
	if article.Author_id >= 0 {
		message, err := service.AddArticle(c, article)
		if err != nil {
			fmt.Println(err)
			c.String(http.StatusNotImplemented, err.Error())
		} else {
			c.String(http.StatusOK, message)
		}
	} else {
		c.String(http.StatusBadRequest, "Error in request data, The author_id passed in Url should be greater than 0")
	}
}

func updateArticle(c *gin.Context) {
	var article models.Article
	c.BindJSON(&article)
	id, _ := strconv.Atoi(c.Param("id"))
	if article.Author_id <= 0 || article.Author_id != id || id <= 0 {
		c.String(http.StatusBadRequest, "Error in request data, The id passed in Url should match the Author_id in the request body and both the id's should be greater than 0")
	} else {
		message, err := service.UpdateArticle(c, id, article)
		if err != nil {
			fmt.Println(err)
			c.String(http.StatusNotImplemented, err.Error())
		} else {
			c.String(http.StatusOK, message)
		}
	}
}

func main() {
	router := gin.Default()
	router.GET("/api/v1/articles", getArticles)
	router.GET("/api/v1/articles/:id", getArticleById)
	router.POST("/api/v1/articles", addArticle)
	router.POST("/api/v1/articles/:id/update", updateArticle)

	router.Run("localhost:5000")
}
