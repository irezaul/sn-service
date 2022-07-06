package handlers

import (
	"encoding/json"
	"ginorm/model"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FrontendService(c *gin.Context) {
	req, err := http.Get("http://localhost:8080/api/service")
	if err != nil {
		log.Println(err)
	}

	defer req.Body.Close()
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println(err)
	}

	var service []model.Service
	json.Unmarshal(body, &service)

	c.HTML(http.StatusOK, "index.gohtml", service)
}
