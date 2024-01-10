package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/dmitriysta/messenger/message/internal/interfaces"
	"github.com/dmitriysta/messenger/message/internal/pkg/errors"

	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"github.com/uber/jaeger-client-go"
)

const (
	HeaderContentType   = "Content-Type"
	MIMEApplicationJSON = "application/json"
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

func (h *MessageHandler) CreateMessageHandler(w http.ResponseWriter, r *http.Request) {
	span, ctx := opentracing.StartSpanFromContext(r.Context(), "CreateMessageHandler")
	defer span.Finish()

	var messageRequest MessageRequest

	if err := json.NewDecoder(r.Body).Decode(&messageRequest); err != nil {
		traceID := span.Context().(jaeger.SpanContext).TraceID().String()
		h.logger.WithFields(logrus.Fields{
			"module":  "message",
			"handler": "CreateMessageHandler",
			"traceId": traceID,
			"error":   err.Error(),
		}).Error(errors.InvalidRequestBody)

		http.Error(w, errors.InvalidRequestBody, http.StatusBadRequest)
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

		http.Error(w, errors.ErrorCreatingMessage, http.StatusInternalServerError)
		return
	}

	w.Header().Set(HeaderContentType, MIMEApplicationJSON)
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(map[string]interface{}{
		"id":        message.Id,
		"userId":    message.UserID,
		"channelId": message.ChannelID,
		"content":   message.Content,
	})
	if err != nil {
		h.logger.Errorf(errors.ErrorEncodingResponse, err)
		http.Error(w, errors.ErrorInternalServer, http.StatusInternalServerError)
		return
	}
}

func (h *MessageHandler) GetMessagesByChannelIdHandler(w http.ResponseWriter, r *http.Request) {
	span, ctx := opentracing.StartSpanFromContext(r.Context(), "GetMessagesByChannelIdHandler")
	defer span.Finish()

	channelIdStr := r.URL.Query().Get("channelId")
	if channelIdStr == "" {
		traceID := span.Context().(jaeger.SpanContext).TraceID().String()
		h.logger.WithFields(logrus.Fields{
			"module":  "message",
			"handler": "GetMessagesByChannelIdHandler",
			"traceId": traceID,
		}).Error(errors.InvalidChannelId)

		http.Error(w, errors.InvalidChannelId, http.StatusBadRequest)
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

		http.Error(w, errors.ErrorChannelIdNotInt, http.StatusBadRequest)
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

		http.Error(w, errors.ErrorGettingMessages, http.StatusInternalServerError)
		return
	}

	w.Header().Set(HeaderContentType, MIMEApplicationJSON)
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(messages)
	if err != nil {
		h.logger.Errorf(errors.ErrorEncodingResponse, err)
		http.Error(w, errors.ErrorInternalServer, http.StatusInternalServerError)
		return
	}
}

func (h *MessageHandler) UpdateMessageHandler(w http.ResponseWriter, r *http.Request) {
	span, ctx := opentracing.StartSpanFromContext(r.Context(), "UpdateMessageHandler")
	defer span.Finish()

	ctxId := r.Context().Value("id")
	messageId, ok := ctxId.(int)
	if !ok {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var updateReq MessageRequest
	if err := json.NewDecoder(r.Body).Decode(&updateReq); err != nil {
		traceID := span.Context().(jaeger.SpanContext).TraceID().String()
		h.logger.WithFields(logrus.Fields{
			"module":  "message",
			"handler": "UpdateMessageHandler",
			"traceId": traceID,
			"error":   err.Error(),
		}).Error(errors.ErrorBindingRequestBody)

		http.Error(w, errors.ErrorBindingRequestBody, http.StatusBadRequest)
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

		http.Error(w, errors.ErrorGettingMessages, http.StatusInternalServerError)
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

		http.Error(w, errors.ErrorUpdatingMessage, http.StatusInternalServerError)
		return
	}

	w.Header().Set(HeaderContentType, MIMEApplicationJSON)
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(map[string]interface{}{
		"id":        message.Id,
		"userId":    message.UserID,
		"channelId": message.ChannelID,
		"content":   message.Content,
	})
	if err != nil {
		h.logger.Errorf(errors.ErrorEncodingResponse, err)
		http.Error(w, errors.ErrorInternalServer, http.StatusInternalServerError)
		return
	}
}

func (h *MessageHandler) DeleteMessageHandler(w http.ResponseWriter, r *http.Request) {
	span, ctx := opentracing.StartSpanFromContext(r.Context(), "DeleteMessageHandler")
	defer span.Finish()

	ctxId := r.Context().Value("id")
	messageId, ok := ctxId.(int)
	if !ok {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
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

		http.Error(w, errors.ErrorDeletingMessage, http.StatusInternalServerError)
		return
	}

	w.Header().Set(HeaderContentType, MIMEApplicationJSON)
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(map[string]interface{}{
		"id": messageId,
	})
	if err != nil {
		h.logger.Errorf(errors.ErrorEncodingResponse, err)
		http.Error(w, errors.ErrorInternalServer, http.StatusInternalServerError)
		return
	}
}
