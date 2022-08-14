package controllers

import "net/http"

type PingPongCtrl struct{}

type pingOut struct {
	Message string `json:"message"`
}

func (c *PingPongCtrl) Ping(*http.Request) (*pingOut, error) {
	// return nil, handlers.ErrUnauthorized()

	return &pingOut{Message: "pong"}, nil
}
