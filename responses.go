package main

// JSONResponse is the JSON formatter
type JSONResponse struct {
	Meta interface{} `json:"meta"`
	Data interface{} `json:"data"`
}

// JSONErrorResponse is the JSON formatter for error responses
type JSONErrorResponse struct {
	Error *APIError `json:"error"`
}

// APIError is part of the JSONErrorResponse
type APIError struct {
	Status int16  `json:"status"`
	Title  string `json:"title"`
}
