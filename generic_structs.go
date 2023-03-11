package main

type ErrorResponse struct {
	Code    int    `json:"status"`
	Message string `json:"message"`
}
