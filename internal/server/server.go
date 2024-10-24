package server

import (
	"context"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lahnasti/test/internal/config"
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

func (s *Server) HandleGetAPIData(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	config, err := config.SetupConfig()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to setup config", "error": err.Error()})
		return
	}
	request, err := http.NewRequest("GET", config.API, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create request", "error": err.Error()})
		return
	}
	request.Header.Set(config.Header, config.Token)
	request.Header.Set("Content-Type", "application/json")
	request = request.WithContext(ctx)

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to send request", "error": err.Error()})
		return
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to read response body", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Success", "data": body})
}
