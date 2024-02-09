package main

import (
	"sync"
	"time"
)

// clients is a map of Client structs that gets used for the API
var clients = make(map[string]*Client)

// syncMapLock gets used in GetLock to prevent to threads from creating a lock for a Client at the same time
var syncMapLock = &sync.RWMutex{}

// Client represents the client servers connecting to the PhoneHomer API
type Client struct {
	Id         string    `json:"id"`         //Unique ID for the client
	IpAddress  string    `json:"ipAddress"`  //Client IP
	Error      string    `json:"error"`      //Error if the client wants to report trouble
	PhoneHome  bool      `json:"phoneHome"`  //Should the client phone home
	LastUpdate time.Time `json:"lastUpdate"` //When did the client last reach out
}

// GetClient returns the Client pointer if it exists in clients otherwise it returns nil
func GetClient(id string) *Client {
	if _, ok := clients[id]; !ok {
		return nil
	}
	return clients[id]
}

// UpdateClient creates a Client in clients if no client exists, otherwise it updates the existing object
func UpdateClient(id string, ipAddress string, error string) *Client {
	syncMapLock.Lock()
	client := GetClient(id)
	if client == nil {
		clients[id] = &Client{
			Id:         id,
			IpAddress:  ipAddress,
			Error:      error,
			PhoneHome:  false,
			LastUpdate: time.Now(),
		}
	} else {
		clients[id].IpAddress = ipAddress
		clients[id].Error = error
		clients[id].LastUpdate = time.Now()
	}
	client = clients[id]
	syncMapLock.Unlock()
	return client
}

// UpdateClientStatus returns false if it can't find a Client for id, otherwise it updates the PhoneHome property of the Client
func UpdateClientStatus(id string, phoneHome bool) bool {
	syncMapLock.Lock()
	client := GetClient(id)
	if client == nil {
		syncMapLock.Unlock()
		return false
	}
	clients[id].PhoneHome = phoneHome
	syncMapLock.Unlock()
	return true
}

func getClientsAsList() []*Client {
	syncMapLock.RLock()
	var ret = make([]*Client, 0)
	for _, client := range clients {
		ret = append(ret, client)
	}
	syncMapLock.RUnlock()
	return ret
}
