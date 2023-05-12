package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type personalInfo struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Age       string `json:"age"`
	Education string `json:"education"`
}

var Infos = []personalInfo{
	{ID: "1", Name: "AAAA", Age: "23", Education: "B.Tech"},
	{ID: "2", Name: "BBBB", Age: "22", Education: "Degree"},
	{ID: "3", Name: "CCCCC ", Age: "24", Education: "M.Pharmacy"},
}

func main() {
	router := gin.Default()
	router.GET("/", getInfos)
	router.GET("/info/:id", getInfosbyId)
	router.POST("/info", postInfos)
	router.DELETE("/info/:id", deletebyId)

	router.Run("localhost:9000")
}

func getInfos(c *gin.Context) {
	c.JSON(200, Infos)
}

func getInfosbyId(c *gin.Context) {
	id := c.Param("id")

	for _, a := range Infos {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Info not found"})
}

func postInfos(c *gin.Context) {
	var newPersonalInfo personalInfo

	if err := c.BindJSON(&newPersonalInfo); err != nil {
		return
	}

	Infos = append(Infos, newPersonalInfo)
	c.IndentedJSON(http.StatusCreated, newPersonalInfo)
}

func deletebyId(c *gin.Context) {
	id := c.Param("id")
	for i, p := range Infos {
		if p.ID == id {
			Infos = append(Infos[:i], Infos[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "deleted"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "not found"})

}
