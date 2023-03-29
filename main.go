package main

import (
	// "net/http"

	"github.com/dilyara4949/BookStore/pkg/db"
	"github.com/dilyara4949/BookStore/pkg/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	DB := db.Init()
	h := handlers.New(DB)
	router := gin.Default()
	router.GET("/book", h.GetBooks)
	router.GET("book/:id", h.GetBook)
	router.GET("/", h.Home)
	

	router.Run("localhost:8000")
}

// lsof -i :8000
// kill -9  <PID>

// % docker build .  
// docker run -d -p 8000:8080 288ca3be3974   (image id)  
// docker run --name assigment2 -p 8000:8000 -d 1feb2c724dab


// module github.com/dilyara4949/BookStore



// # Use an official Golang runtime as a parent image
// FROM golang:1.17-alpine AS builder

// # Set the working directory to /app
// WORKDIR /app

// # Copy the source code and config file into the container at /app
// COPY . .

// # Build the Go app
// RUN go build -o main .

// # Use an official Alpine Linux image as a parent image
// FROM alpine:latest

// # Copy the built executable and config file from the previous stage
// COPY --from=builder /app/main /app/main
// #COPY --from=builder /app/config.yaml /app/config.yaml

// # Set the working directory to /app
// WORKDIR /app

// # Expose port 8080
// EXPOSE 8080

// # Run the command to start the app with the config file
// CMD ["/app/main"]


// # version: "3.8"

// # services:
// #   database:
// #     container_name: database
// #     image: postgres:12.8
// #     restart: always
// #     environment:
// #       - POSTGRES_USER=postgres
// #       - POSTGRES_PASSWORD=12345
// #       - POSTGRES_DB=bookStore
// #     ports:
// #       - 5432:5432
// #     volumes:
// #       - db:/var/lib/postgresql/data 

// # volumes:
// #   db:

// # docker-compose.yml