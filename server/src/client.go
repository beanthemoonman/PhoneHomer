package main

import (
	"sync"
	"time"
)

// Clients is a map of Client structs that gets used for the API
var clients = make(map[string]*Client)

// ClientsSync is a map of Mutexes that allow us to make sure that only one thread is touching a client object at one time
var clientsSync = make(map[string]*sync.Mutex)

// syncMapLock gets used in GetLock to prevent to threads from creating a lock for a Client at the same time
var syncMapLock = sync.Mutex{}

type Client struct {
	Id         string    `json:"id"`
	IpAddress  string    `json:"ipAddress"`
	Error      string    `json:"error"`
	PhoneHome  bool      `json:"phoneHome"`
	LastUpdate time.Time `json:"lastUpdate"`
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
	lock := GetLock(id)
	lock.Lock()
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
	lock.Unlock()
	return clients[id]
}

// UpdateClientStatus returns false if it can't find a Client for id, otherwise it updates the PhoneHome property of the Client
func UpdateClientStatus(id string, phoneHome bool) bool {
	lock := GetLock(id)
	lock.Lock()
	client := GetClient(id)
	if client == nil {
		lock.Unlock()
		return false
	}
	clients[id].PhoneHome = phoneHome
	lock.Unlock()
	return true
}

// GetLock implements a double-checking checking method for lock creation. Returns the lock.
func GetLock(id string) *sync.Mutex {
	if _, ok := clientsSync[id]; !ok {
		syncMapLock.Lock()
		if _, ok2 := clientsSync[id]; !ok2 {
			clientsSync[id] = &sync.Mutex{}
		}
		syncMapLock.Unlock()
	}
	return clientsSync[id]
}
