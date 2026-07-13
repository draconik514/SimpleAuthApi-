package main

import(
	"backend/config"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	"backend/routes"
)

func main(){
	log.Println("Starting Simple Auth Backend")

	config.ConnectDB()

	r := gin.Default()

	r.GET("/health", func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{
			"status" : "OK",
			"database" : "connected",
			"service" : "Simple Auth Api",
		})
	})

	routes.SetupRoutes(r)

	port := ":8080"
	log.Printf("Server running on http://localhost%s", port)
	log.Println("API ENDPOINT :")

	log.Printf("GET : http://localhost%s/health", port)

	log.Printf("POST : http://localhost%s/api/register", port)

	if err := r.Run(port); err != nil {
		log.Fatal("Failed to start server", err)
	}

}