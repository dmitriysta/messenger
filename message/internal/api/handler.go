package api

import (
	"net/http"
	"strconv"

	"github.com/dmitriysta/messenger/message/internal/interfaces"
	"github.com/dmitriysta/messenger/message/internal/pkg/errors"

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
		}).Error(errors.InvalidRequestBody)

		c.JSON(http.StatusBadRequest, gin.H{"error": errors.InvalidRequestBody})
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
		}).Error(errors.ErrorCreatingMessage)

		c.JSON(http.StatusInternalServerError, gin.H{"error": errors.ErrorCreatingMessage})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":        message.Id,
		"userId":    message.UserID,
		"channelId": message.ChannelID,
		"content":   message.Content,
	})
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
		}).Error(errors.InvalidChannelId)

		c.JSON(http.StatusBadRequest, gin.H{"error": errors.InvalidChannelId})
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
		}).Error(errors.ErrorChannelIdNotInt)

		c.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrorChannelIdNotInt})
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
		}).Error(errors.ErrorGettingMessages)

		c.JSON(http.StatusInternalServerError, gin.H{"error": errors.ErrorGettingMessages})
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
		}).Error(errors.ErrorMessageIdNotInt)

		c.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrorMessageIdNotInt})
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
		}).Error(errors.ErrorBindingRequestBody)

		c.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrorBindingRequestBody})
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
		}).Error(errors.ErrorGettingMessages)

		c.JSON(http.StatusInternalServerError, gin.H{"error": errors.ErrorGettingMessages})
		return
	}

	message.Content = updateReq.Content
	if err := h.messageService.UpdateMessage(ctx, message); err != nil {
		traceID := span.Context().(jaeger.SpanContext).TraceID().String()
		h.logger.WithFields(logrus.Fields{
			"module":  "message",
			"handler": "UpdateMessageHandler",
			"content": message.Content + " -> " + updateReq.Content,
			"traceId": traceID,
			"error":   err.Error(),
		}).Error(errors.ErrorUpdatingMessage)

		c.JSON(http.StatusInternalServerError, gin.H{"error": errors.ErrorUpdatingMessage})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":        message.Id,
		"userId":    message.UserID,
		"channelId": message.ChannelID,
		"content":   message.Content,
	})
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
		}).Error(errors.ErrorMessageIdNotInt)

		c.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrorMessageIdNotInt})
		return
	}

	if err := h.messageService.DeleteMessage(ctx, messageId); err != nil {
		traceID := span.Context().(jaeger.SpanContext).TraceID().String()
		h.logger.WithFields(logrus.Fields{
			"module":  "message",
			"handler": "DeleteMessageHandler",
			"traceId": traceID,
			"error":   err.Error(),
		}).Error(errors.ErrorDeletingMessage)

		c.JSON(http.StatusInternalServerError, gin.H{"error": errors.ErrorDeletingMessage})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Message deleted successfully",
	})
}
