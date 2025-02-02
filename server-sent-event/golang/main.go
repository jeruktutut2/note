package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var messageChan chan string

type SSEClient struct {
	Id      int64
	Context echo.Context
}

var SSEClients []SSEClient

func main() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000", "*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	e.GET("/sse/response", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"response": "response sse",
		})
	})
	e.GET("/sse/handle-sse", func(c echo.Context) error {
		// set timeout to 0
		c.SetRequest(c.Request().WithContext(context.Background()))
		c.Response().Header().Set("Content-Type", "text/event-stream")
		c.Response().Header().Set("Cache-Control", "no-cache")
		c.Response().Header().Set("Connection", "keep-alive")
		c.Response().Header().Set("Access-Control-Allow-Origin", "*")

		messageChan = make(chan string)

		defer func() {
			close(messageChan)
			messageChan = nil
			log.Printf("client connection close")
		}()

		flusher, _ := c.Response().Writer.(http.Flusher)

		for {

			select {
			case message := <-messageChan:
				// fmt.Println("message:", message)
				json.NewEncoder(c.Response()).Encode(message)
				// c.Response().Flush()
				flusher.Flush()
			case <-c.Request().Context().Done():
				fmt.Println("connection close")
				return nil
			}
		}
		//return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/sse/send-message/:message", func(c echo.Context) error {
		message := c.Param("message")
		messageChan <- message
		return nil
	})

	e.GET("/sse/handle-sse-without-channel", func(c echo.Context) error {
		// c.SetRequest(c.Request().WithContext(context.Background()))
		c.Response().Header().Set("Content-Type", "text/event-stream")
		c.Response().Header().Set("Cache-Control", "no-cache")
		c.Response().Header().Set("Connection", "keep-alive")
		c.Response().Header().Set("Access-Control-Allow-Origin", "*")

		sseClient := SSEClient{}
		sseClient.Id = time.Now().UnixMilli()
		sseClient.Context = c
		SSEClients = append(SSEClients, sseClient)
		fmt.Println("SSEClients:", SSEClients)

		<-c.Request().Context().Done()

	sseClientsLoop:
		for i := 0; i < len(SSEClients); i++ {
			if SSEClients[i].Id == sseClient.Id {
				SSEClients = append(SSEClients[:i], SSEClients[i+1:]...)
				break sseClientsLoop
			}
		}
		return nil
	})

	e.GET("/sse/send-message-without-channel/:message", func(c echo.Context) error {
		// c.Response().Header().Set("Content-Type", "text/event-stream")
		// c.Response().Header().Set("Cache-Control", "no-cache")
		message := c.Param("message")
		for i := 0; i < len(SSEClients); i++ {
			fmt.Println("send message to:", SSEClients[i], " dengan pesan:", message)
			ctx := SSEClients[i].Context
			// fmt.Println("SSEClients[i]:", SSEClients[i].Id, SSEClients[i].Context)

			headers := ctx.Response().Header()
			// Menyusun output sebagai string
			var output string
			for key, values := range headers {
				output += fmt.Sprintf("%s: %s\n", key, values)
			}
			fmt.Println("headers:", output)

			var messageEvent map[string]interface{}
			messageEvent = map[string]interface{}{
				"message": message,
			}
			// json.NewEncoder(ctx.Response()).Encode(messageEvent)
			jsonData, _ := json.Marshal(messageEvent)
			fmt.Fprintf(ctx.Response(), "data: %s\n\n", jsonData)
			ctx.Response().Flush()
		}
		// return c.JSON(http.StatusOK, "ok")
		return nil
	})

	e.Logger.Fatal(e.Start(":8080"))
}
