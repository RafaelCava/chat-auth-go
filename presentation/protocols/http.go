package protocols

type HttpResponse struct {
	StatusCode int `json:"statusCode" example:"200"`
	Body       any `json:"body"`
}
