package api

import (
	"net/http"
	"strconv"

	"github.com/dmitriysta/messenger/message/internal/interfaces"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"github.com/uber/jaeger-client-go"
)

type MessageRequest struct {
	UserId    int    `json:"userId"`
	ChannelID int    `json:"channelId"`
	Content   string `json:"content"`
}

type MessageHandler struct {
	messageService interfaces.MessageService
	logger         *logrus.Logger
	tracer         opentracing.Tracer
}

func NewMessageHandler(messageService interfaces.MessageService, logger *logrus.Logger, tracer opentracing.Tracer) *MessageHandler {
	return &MessageHandler{
		messageService: messageService,
		logger:         logger,
		tracer:         tracer,
	}
}

func (h *MessageHandler) CreateMessageHandler(c *gin.Context) {
	span, ctx := opentracing.StartSpanFromContext(c.Request.Context(), "CreateMessageHandler")
	defer span.Finish()

	var messageRequest MessageRequest

	if err := c.ShouldBindJSON(&messageRequest); err != nil {
		traceID := span.Context().(jaeger.SpanContext).TraceID().String()
		h.logger.WithFields(logrus.Fields{
			"module":  "message",
			"handler": "CreateMessageHandler",
			"traceId": traceID,
			"error":   err.Error(),
		}).Error("Error binding request body")

		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	message, err := h.messageService.CreateMessage(ctx, messageRequest.UserId, messageRequest.ChannelID, messageRequest.Content)
	if err != nil {
		traceID := span.Context().(jaeger.SpanContext).TraceID().String()
		h.logger.WithFields(logrus.Fields{
			"module":    "message",
			"handler":   "CreateMessageHandler",
			"userId":    messageRequest.UserId,
			"channelId": messageRequest.ChannelID,
			"content":   messageRequest.Content,
			"traceId":   traceID,
			"error":     err.Error(),
		}).Error("Error creating message")

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating message"})
		return
	}

	c.JSON(http.StatusCreated, message)
}

func (h *MessageHandler) GetMessagesByChannelIdHandler(c *gin.Context) {
	span, ctx := opentracing.StartSpanFromContext(c.Request.Context(), "GetMessagesByChannelIdHandler")
	defer span.Finish()

	channelIdStr := c.Query("channelId")
	if channelIdStr == "" {
		traceID := span.Context().(jaeger.SpanContext).TraceID().String()
		h.logger.WithFields(logrus.Fields{
			"module":  "message",
			"handler": "GetMessagesByChannelIdHandler",
			"traceId": traceID,
		}).Error("channelId is required")

		c.JSON(http.StatusBadRequest, gin.H{"error": "channelId is required"})
		return
	}

	channelId, err := strconv.Atoi(channelIdStr)
	if err != nil {
		traceID := span.Context().(jaeger.SpanContext).TraceID().String()
		h.logger.WithFields(logrus.Fields{
			"module":  "message",
			"handler": "GetMessagesByChannelIdHandler",
			"traceId": traceID,
			"error":   err.Error(),
		}).Error("channelId must be an integer")

		c.JSON(http.StatusBadRequest, gin.H{"error": "channelId must be an integer"})
		return
	}

	messages, err := h.messageService.GetMessagesByChannelId(ctx, channelId)
	if err != nil {
		traceID := span.Context().(jaeger.SpanContext).TraceID().String()
		h.logger.WithFields(logrus.Fields{
			"module":    "message",
			"handler":   "GetMessagesByChannelIdHandler",
			"channelId": channelId,
			"traceId":   traceID,
			"error":     err.Error(),
		}).Error("Error getting messages")

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting messages"})
		return
	}

	c.JSON(http.StatusOK, messages)
}

func (h *MessageHandler) UpdateMessageHandler(c *gin.Context) {
	span, ctx := opentracing.StartSpanFromContext(c.Request.Context(), "UpdateMessageHandler")
	defer span.Finish()

	messageIdStr := c.Param("id")
	messageId, err := strconv.Atoi(messageIdStr)
	if err != nil {
		traceID := span.Context().(jaeger.SpanContext).TraceID().String()
		h.logger.WithFields(logrus.Fields{
			"module":  "message",
			"handler": "UpdateMessageHandler",
			"traceId": traceID,
			"error":   err.Error(),
		}).Error("Message id is not integer")

		c.JSON(http.StatusBadRequest, gin.H{"error": "Message id must be an integer"})
		return
	}

	var updateReq MessageRequest
	if err := c.ShouldBindJSON(&updateReq); err != nil {
		traceID := span.Context().(jaeger.SpanContext).TraceID().String()
		h.logger.WithFields(logrus.Fields{
			"module":  "message",
			"handler": "UpdateMessageHandler",
			"traceId": traceID,
			"error":   err.Error(),
		}).Error("Error binding request body")

		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	message, err := h.messageService.GetMessageById(ctx, messageId)
	if err != nil {
		traceID := span.Context().(jaeger.SpanContext).TraceID().String()
		h.logger.WithFields(logrus.Fields{
			"module":  "message",
			"handler": "UpdateMessageHandler",
			"traceId": traceID,
			"error":   err.Error(),
		}).Error("Error getting message")

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Message not found"})
		return
	}

	message.Content = updateReq.Content
	if err := h.messageService.UpdateMessage(ctx, message); err != nil {
		traceID := span.Context().(jaeger.SpanContext).TraceID().String()
		h.logger.WithFields(logrus.Fields{
			"module":  "message",
			"handler": "UpdateMessageHandler",
			"traceId": traceID,
			"error":   err.Error(),
		}).Error("Error updating message")

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating message"})
		return
	}

	c.JSON(http.StatusOK, message)
}

func (h *MessageHandler) DeleteMessageHandler(c *gin.Context) {
	span, ctx := opentracing.StartSpanFromContext(c.Request.Context(), "DeleteMessageHandler")
	defer span.Finish()

	messageIdStr := c.Param("id")
	messageId, err := strconv.Atoi(messageIdStr)
	if err != nil {
		traceID := span.Context().(jaeger.SpanContext).TraceID().String()
		h.logger.WithFields(logrus.Fields{
			"module":  "message",
			"handler": "DeleteMessageHandler",
			"traceId": traceID,
			"error":   err.Error(),
		}).Error("Error: message id isn't an integer")

		c.JSON(http.StatusBadRequest, gin.H{"error": "Message id must be an integer"})
		return
	}

	if err := h.messageService.DeleteMessage(ctx, messageId); err != nil {
		traceID := span.Context().(jaeger.SpanContext).TraceID().String()
		h.logger.WithFields(logrus.Fields{
			"module":  "message",
			"handler": "DeleteMessageHandler",
			"traceId": traceID,
			"error":   err.Error(),
		}).Error("Error deleting message")

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting message"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Message deleted successfully"})
}
