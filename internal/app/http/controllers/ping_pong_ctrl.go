package controllers

type PingPongCtrl struct{}

type pingOut struct {
	Message string `json:"message"`
}

func (c *PingPongCtrl) Ping() (*pingOut, error) {
	return &pingOut{Message: "pong"}, nil
}
