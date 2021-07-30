package utils

import (
	"Problem5/models"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func ReadJsonFile(c *gin.Context) ([]models.Article, error) {
	var articles []models.Article
	articleData, err := ioutil.ReadFile("data.json")
	if err != nil {
		//c.String(http.StatusInternalServerError, "Failed to Read the Json file")
		return articles, fmt.Errorf("failed to Read the Json file")
	}
	err = json.Unmarshal([]byte(articleData), &articles)
	if err != nil {
		return articles, fmt.Errorf("failed to unmarshal the Json file")
	}

	return articles, nil
}

func CheckIfIdExists(articles []models.Article, id int) int {
	for i, val := range articles {
		if val.Author_id == id {
			return i
		}
	}
	return -1
}
