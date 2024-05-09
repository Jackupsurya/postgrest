package models

type Todo struct {
	Id   int    `json:"id,omitempty"`
	Task string `json:"task,omitempty"`
	Done bool   `json:"done,omitempty"`
}

type Response struct {
	Success bool
	Data    interface{}
	Error   error
}

type PostgrestErrorResponse struct {
	Code    string      `json:"code,omitempty"`
	Details interface{} `json:"details,omitempty"`
	Hint    interface{} `json:"hint,omitempty"`
	Message string      `json:"message,omitempty"`
}
