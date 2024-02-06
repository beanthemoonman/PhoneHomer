package main

// StartTunnelRequest is the request object for the StartTunnel and StopTunnel methods
type StartTunnelRequest struct {
	Id string `json:"id"`
}

// UpdateRequest is the request object for the Update method
type UpdateRequest struct {
	Id    string `json:"id"`
	Error string `json:"error"`
}

// Response is the response object to be used when no other specific object is wanted
type Response struct {
	Error string `json:"error"`
}
