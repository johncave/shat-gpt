package main

import "log"

type ErrorResponse struct {
	Code    int    `json:"status"`
	Message string `json:"message"`
}

func logError(desc string, err error) {
	if err != nil {
		log.Println("Error encountered in", desc, err)
	}
}
