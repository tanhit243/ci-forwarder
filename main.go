package main

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"

	"log"

	"os"
)

type build struct {
	Event  string `json:"event"`
	Branch string `json:"branch"`
	Status string `json:"status"`
	Link   string `json:"link"`
}

type release struct {
	Event string `json:"event"`
	Build build  `json:"build"`
}

func main() {
	router := gin.Default()
	var releasePayload release
	router.POST("/ci-forwarder", func(c *gin.Context) {

		if err := c.BindJSON(&releasePayload); err != nil {
			return
		}
		logCiRequestBody(releasePayload)
		c.JSON(http.StatusOK, releasePayload)
	})

	router.Run("localhost:2424")
}

func logCiRequestBody(releasePayload release) {
	// log to custom file
	LOG_FILE, _ := filepath.Abs("./tmp/development.log")
	// open log file
	logFile, err := os.OpenFile(LOG_FILE, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Panic(err)
	}
	defer logFile.Close()

	// Set log out put and enjoy :)
	log.SetOutput(logFile)

	// optional: log date-time, filename, and line number
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	log.Printf("%+v\n", releasePayload)
}
