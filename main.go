package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Message struct {
	Message string `json:"message"`
}

/*func main() {
	hello := Message{
		Message: "hello world",
	}
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(hello)
	})

	http.ListenAndServe(fmt.Sprintf("0.0.0.0:%s", os.Getenv("PORT")), mux)
}*/

func main() {
	hello := Message{
		Message: "hello world",
	}
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	r.Use(cors.New(config))

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, hello)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(fmt.Sprintf("0.0.0.0:%s", port))
}
