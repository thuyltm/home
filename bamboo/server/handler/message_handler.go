package handler

import (
	"fmt"
	"home/bamboo/server/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type MessageHandler struct {
	MessageService service.MessageService
}

func NewMessageHandler(messageService service.MessageService) *MessageHandler {
	return &MessageHandler{
		MessageService: messageService,
	}
}

func (h *MessageHandler) CountMessages(c echo.Context) error {
	r, err := h.MessageService.CountMessages()
	if err != nil {
		return c.HTML(http.StatusInternalServerError, err.Error())
	}
	return c.HTML(http.StatusOK, fmt.Sprintf("Hello, Docker! (%d)\n", r))
}

func (h *MessageHandler) CreateMessage(c echo.Context) error {
	m := &service.Message{}
	if err := c.Bind(m); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	err := h.MessageService.CreateMessage(m.Value)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, m)
}
