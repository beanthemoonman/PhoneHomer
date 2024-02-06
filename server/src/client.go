package main

import (
	"sync"
	"time"
)

var clients = make(map[string]*Client)

var clientsSync = make(map[string]*sync.Mutex)

var syncMapLock = sync.Mutex{}

type Client struct {
	Id         string    `json:"id"`
	IpAddress  string    `json:"ipAddress"`
	Error      string    `json:"error"`
	PhoneHome  bool      `json:"phoneHome"`
	LastUpdate time.Time `json:"lastUpdate"`
}

func GetClient(id string) *Client {
	if _, ok := clients[id]; !ok {
		return nil
	}
	return clients[id]
}

func UpdateClient(id string, ipAddress string, error string) *Client {
	lock := getLock(id)
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
		delete(clients, id)
		clients[id] = &Client{
			Id:         client.Id,
			IpAddress:  ipAddress,
			Error:      error,
			PhoneHome:  client.PhoneHome,
			LastUpdate: time.Now(),
		}
	}
	lock.Unlock()
	return clients[id]
}

func UpdateClientStatus(id string, phoneHome bool) bool {
	lock := getLock(id)
	lock.Lock()
	client := GetClient(id)
	if client == nil {
		lock.Unlock()
		return false
	}
	delete(clients, id)
	clients[id] = &Client{
		Id:         client.Id,
		IpAddress:  client.IpAddress,
		Error:      client.Error,
		PhoneHome:  phoneHome,
		LastUpdate: client.LastUpdate,
	}
	lock.Unlock()
	return true
}

func getLock(id string) *sync.Mutex {
	syncMapLock.Lock()
	if _, ok := clientsSync[id]; !ok {
		clientsSync[id] = &sync.Mutex{}
	}
	syncMapLock.Unlock()
	return clientsSync[id]
}
