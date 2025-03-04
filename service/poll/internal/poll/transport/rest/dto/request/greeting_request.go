package request

type GreetingRequest struct {
	Name string `json:"name" validate:"required"`
}
