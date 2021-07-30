package service

import (
	"Problem5/models"
	util "Problem5/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetArticles(c *gin.Context) ([]models.Article, error) {
	fmt.Println("Inside Service")
	articles, err := util.ReadJsonFile(c)
	if err != nil {
		return articles, err
	}
	return articles, nil
}

func GetArticleById(c *gin.Context) (models.Article, error) {
	var article models.Article
	articles, err := util.ReadJsonFile(c)
	if err != nil {
		return article, err
	}

	id, _ := strconv.Atoi(c.Param("id"))
	for _, artVal := range articles {
		if artVal.Author_id == id {
			article = artVal
			fmt.Println("ID is", article)
			break
		}
	}
	if article.Author_id != 0 {
		return article, nil
	} else {
		return article, fmt.Errorf("no Matching article present")
	}
}

func AddArticle(c *gin.Context, article models.Article) (string, error) {
	articles, err := util.ReadJsonFile(c)
	if err != nil {
		return "", err
	}
	articles = append(articles, article)
	articleToAdd, _ := json.Marshal(articles)
	err = ioutil.WriteFile("data.json", articleToAdd, 0644)
	if err != nil {
		fmt.Println(err)
		return "", fmt.Errorf("error while writing on a Json File")
	}

	return "Article Added Successfully", nil
}

func UpdateArticle(c *gin.Context, id int, article models.Article) (string, error) {
	articles, err := util.ReadJsonFile(c)
	if err != nil {
		return "", err
	}
	key := util.CheckIfIdExists(articles, id)
	fmt.Println("test", key, article)
	if key != -1 {
		if article.Content != "" {
			articles[key].Content = article.Content
		}
		if article.Intro != "" {
			articles[key].Intro = article.Intro
		}
		if article.Title != "" {
			articles[key].Title = article.Title
		}

		articleToAdd, err := json.Marshal(articles)
		if err != nil {
			return "", fmt.Errorf("error marshalling Article data")
		}
		err = ioutil.WriteFile("data.json", articleToAdd, 0644)
		if err != nil {
			fmt.Println(err)
			return "", fmt.Errorf("error while writing on json file")
		}
		return "Article Updated Successfully", nil
	} else {
		return "Article ID not found", nil
	}
}
