package controller

// Controller
type Controller struct {
}

// NewController
func NewController() *Controller {
	return &Controller{}
}

// Message
type Message struct {
	Message string `json:"message" example:"message"`
}
