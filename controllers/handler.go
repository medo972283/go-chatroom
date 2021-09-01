package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/medo972283/go-chatroom/controllers/api"
)

func webRouters(e *echo.Echo) {

	// Login page
	e.GET("/login", Login)

	// Homepage
	e.GET("/", Homepage)
	e.GET("/homepage", Homepage)

	// Chatroom page
	e.POST("/chatroom", Chatroom)
}

func apiEntrypoints(e *echo.Echo) {
	// Create user
	e.POST("/users", api.CreateUser)

	// Get list of chatrooms
	e.GET("/chatrooms", api.IndexChatrooms)

	// Get a specific chatroom
	e.GET("/chatrooms/:id", api.ViewChatroom)

	// Create chatroom
	e.POST("/chatrooms", api.CreateChatroom)

	// Create message
	e.POST("/messages", api.CreateMessage)

	// Login the system
	e.POST("/login", api.Login)
}

func webSocketConnection(e *echo.Echo) {
	// Build message channel in the chatroom
	e.GET("/chatroom/ws", api.MessageChannel)
}

func AttachHandler(e *echo.Echo) {

	// Attach Web Controller router
	webRouters(e)

	// Attach API enrty point router
	apiEntrypoints(e)

	// Attach web socket connection entry point
	webSocketConnection(e)
}
