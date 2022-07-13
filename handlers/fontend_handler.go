package handlers

import (
	"bytes"
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

func FrontendAddService(c *gin.Context) {
	c.HTML(http.StatusOK, "add.gohtml", nil)
}

func FrontendAddProcess(c *gin.Context) {
	model := model.ServiceProcess{
		Brand:   c.PostForm("brand"),
		Service: c.PostForm("services"),
		Source:  c.PostForm("source"),
		Date:    c.PostForm("date"),
		Model:   c.PostForm("model"),
		Serial:  c.PostForm("serial"),
		Price:   c.PostForm("price"),
	}

	data, err := json.MarshalIndent(model, "", " ")
	if err != nil {
		log.Println(err)
	}

	http.Post("http://localhost:8080/api/service", "application/json", bytes.NewBuffer(data))
	c.Redirect(http.StatusMovedPermanently, "/")

}

func FrontendEditService(c *gin.Context) {
	id := c.Param("id")
	req, err := http.Get("http://localhost:8080/api/service/" + id)
	if err != nil {
		log.Println(err)
	}

	defer req.Body.Close()
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println(err)
	}

	var service model.Service
	json.Unmarshal(body, &service)

	c.HTML(http.StatusOK, "edit.gohtml", service)
}
