package main

type StartTunnelRequest struct {
	Id string `json:"id"`
}

type UpdateRequest struct {
	Id    string `json:"id"`
	Error string `json:"error"`
}

type Response struct {
	Error string `json:"error"`
}
