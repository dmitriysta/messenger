package api

import (
	"github.com/dmitriysta/messenger/user/internal/interfaces"
	"github.com/dmitriysta/messenger/user/internal/pkg/errors"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"github.com/uber/jaeger-client-go"
	"net/http"
	"strconv"
)

type UserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserHandler struct {
	userService interfaces.UserService
	logger      *logrus.Logger
	tracer      opentracing.Tracer
}

func NewUserHandler(userService interfaces.UserService, logger *logrus.Logger, tracer opentracing.Tracer) *UserHandler {
	return &UserHandler{
		userService: userService,
		logger:      logger,
		tracer:      tracer,
	}
}

func (h *UserHandler) CreateUserHandler(c *gin.Context) {
	span, ctx := opentracing.StartSpanFromContext(c.Request.Context(), "CreateUserHandler")
	defer span.Finish()

	var userRequest UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		traceID := span.Context().(jaeger.SpanContext).TraceID().String()
		h.logger.WithFields(logrus.Fields{
			"module":  "user",
			"handler": "CreateUserHandler",
			"traceId": traceID,
			"error":   err.Error(),
		}).Error(errors.InvalidRequestBody)

		c.JSON(http.StatusBadRequest, gin.H{"error": errors.InvalidRequestBody})
		return
	}

	user, err := h.userService.CreateUser(ctx, userRequest.Username, userRequest.Email, userRequest.Password)
	if err != nil {
		traceID := span.Context().(jaeger.SpanContext).TraceID().String()
		h.logger.WithFields(logrus.Fields{
			"module":   "user",
			"handler":  "CreateUserHandler",
			"traceId":  traceID,
			"username": userRequest.Username,
			"email":    userRequest.Email,
			"password": userRequest.Password,
			"error":    err.Error(),
		}).Error(errors.ErrorCreatingUser)

		c.JSON(http.StatusInternalServerError, gin.H{"error": errors.ErrorCreatingUser})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":       user.Id,
		"username": user.Name,
		"email":    user.Email,
	})
}

func (h *UserHandler) GetUserHandler(c *gin.Context) {
	span, ctx := opentracing.StartSpanFromContext(c.Request.Context(), "GetUserHandler")
	defer span.Finish()

	userIDStr := c.Param("id")
	if userIDStr == "" {
		traceID := span.Context().(jaeger.SpanContext).TraceID().String()
		h.logger.WithFields(logrus.Fields{
			"module":  "user",
			"handler": "GetUserHandler",
			"traceId": traceID,
		}).Error(errors.ErrorEmptyUserId)

		c.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrorEmptyUserId})
		return
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		traceID := span.Context().(jaeger.SpanContext).TraceID().String()
		h.logger.WithFields(logrus.Fields{
			"module":  "user",
			"handler": "GetUserHandler",
			"traceId": traceID,
			"error":   err.Error(),
		}).Error(errors.InvalidUserId)

		c.JSON(http.StatusBadRequest, gin.H{"error": errors.InvalidUserId})
		return
	}

	user, err := h.userService.GetUserById(ctx, userID)
	if err != nil {
		traceID := span.Context().(jaeger.SpanContext).TraceID().String()
		h.logger.WithFields(logrus.Fields{
			"module":  "user",
			"handler": "GetUserHandler",
			"traceId": traceID,
			"userId":  userID,
			"error":   err.Error(),
		}).Error(errors.ErrorGettingUser)

		c.JSON(http.StatusInternalServerError, gin.H{"error": errors.ErrorGettingUser})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":       user.Id,
		"username": user.Name,
		"email":    user.Email,
	})
}

func (h *UserHandler) UpdateUserHandler(c *gin.Context) {
	span, ctx := opentracing.StartSpanFromContext(c.Request.Context(), "UpdateUserHandler")
	defer span.Finish()

	userIDStr := c.Param("id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		traceID := span.Context().(jaeger.SpanContext).TraceID().String()
		h.logger.WithFields(logrus.Fields{
			"module":  "user",
			"handler": "UpdateUserHandler",
			"traceId": traceID,
			"error":   err.Error(),
		}).Error(errors.InvalidUserId)

		c.JSON(http.StatusBadRequest, gin.H{"error": errors.InvalidUserId})
		return
	}

	var userRequest UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		traceID := span.Context().(jaeger.SpanContext).TraceID().String()
		h.logger.WithFields(logrus.Fields{
			"module":  "user",
			"handler": "UpdateUserHandler",
			"traceId": traceID,
			"error":   err.Error(),
		}).Error(errors.InvalidRequestBody)

		c.JSON(http.StatusBadRequest, gin.H{"error": errors.InvalidRequestBody})
		return
	}

	user, err := h.userService.GetUserById(ctx, userID)
	if err != nil {
		traceID := span.Context().(jaeger.SpanContext).TraceID().String()
		h.logger.WithFields(logrus.Fields{
			"module":  "user",
			"handler": "UpdateUserHandler",
			"traceId": traceID,
			"userId":  userID,
			"error":   err.Error(),
		}).Error(errors.ErrorGettingUser)

		c.JSON(http.StatusInternalServerError, gin.H{"error": errors.ErrorGettingUser})
		return
	}

	user.Name = userRequest.Username
	user.Email = userRequest.Email
	user.Password = userRequest.Password

	if err := h.userService.UpdateUser(ctx, user); err != nil {
		traceID := span.Context().(jaeger.SpanContext).TraceID().String()
		h.logger.WithFields(logrus.Fields{
			"module":   "user",
			"handler":  "UpdateUserHandler",
			"traceId":  traceID,
			"username": user.Name + " -> " + userRequest.Username,
			"email":    user.Email + " -> " + userRequest.Email,
			"password": user.Password + " -> " + userRequest.Password,
			"error":    err.Error(),
		}).Error(errors.ErrorUpdatingUser)

		c.JSON(http.StatusInternalServerError, gin.H{"error": errors.ErrorUpdatingUser})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":       user.Id,
		"username": user.Name,
		"email":    user.Email,
	})
}

func (h *UserHandler) DeleteUserHandler(c *gin.Context) {
	span, ctx := opentracing.StartSpanFromContext(c.Request.Context(), "DeleteUserHandler")
	defer span.Finish()

	userIDStr := c.Param("id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		traceID := span.Context().(jaeger.SpanContext).TraceID().String()
		h.logger.WithFields(logrus.Fields{
			"module":  "user",
			"handler": "DeleteUserHandler",
			"traceId": traceID,
			"error":   err.Error(),
		}).Error(errors.InvalidUserId)

		c.JSON(http.StatusBadRequest, gin.H{"error": errors.InvalidUserId})
		return
	}

	if err := h.userService.DeleteUser(ctx, userID); err != nil {
		traceID := span.Context().(jaeger.SpanContext).TraceID().String()
		h.logger.WithFields(logrus.Fields{
			"module":  "user",
			"handler": "DeleteUserHandler",
			"traceId": traceID,
			"userId":  userID,
			"error":   err.Error(),
		}).Error(errors.ErrorDeletingUser)

		c.JSON(http.StatusInternalServerError, gin.H{"error": errors.ErrorDeletingUser})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User deleted successfully",
	})
}
