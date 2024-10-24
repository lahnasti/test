package server

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
}

func NewServer() *Server {
	r := gin.Default()
	return &Server{
		router: r,
	}
}

func (s *Server) Run(address string) {
	s.router.GET("/api", s.HandleGetAPIData)
	log.Printf("Server running on %s", address)
	s.router.Run(address)
}

func (s *Server) HandleGetAPIData(ctx *gin.Context) {
	api := os.Getenv("API_URL")
	header := os.Getenv("HEADER")
	token := fmt.Sprintf("Bearer %s", os.Getenv("TOKEN"))

	request, err := http.NewRequest("GET", api, nil)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create request", "error": err.Error()})
		return
	}
	request.Header.Set(header, token)
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to send request", "error": err.Error()})
		return
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to read response body", "error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Success", "data": body})
}
