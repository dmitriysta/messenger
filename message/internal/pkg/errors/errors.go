package errors

const (
	InvalidRequestBody      = "Invalid request body"
	ErrorCreatingMessage    = "Error creating message"
	InvalidChannelId        = "Invalid channel id"
	ErrorChannelIdNotInt    = "Channel id is not an integer"
	ErrorGettingMessages    = "Error getting messages"
	ErrorBindingRequestBody = "Error binding request body"
	ErrorUpdatingMessage    = "Error updating message"
	ErrorDeletingMessage    = "Error deleting message"
	ErrorInternalServer     = "Internal server error"
	ErrorEncodingResponse   = "Failed to encode response: %v"
)
