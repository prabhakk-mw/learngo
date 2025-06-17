package microservices

// This file hosts the code required to cache a microservice.

/*
	When a request comes in, the server should check if the microservice is already running.
	If it is not running, it should start the microservice and cache its address for future requests.
	Care must be taken to ensure that the microservice is not started multiple times,
	and that it stays alive for the duration of the server's lifetime.
*/

import (
	"errors"
	"sync"
)

// Microservice represents a running microservice instance.
type Microservice struct {
	Address string
	// Add other fields as needed (e.g., process handle, etc.)
}

// Cache manages running microservices.
type Cache struct {
	mu           sync.Mutex
	services     map[string]*Microservice
	startService func(name string) (*Microservice, error)
}

// NewCache creates a new Cache with the provided startService function.
func NewCache(startService func(name string) (*Microservice, error)) *Cache {
	return &Cache{
		services:     make(map[string]*Microservice),
		startService: startService,
	}
}

// GetOrStart returns the address of the microservice, starting it if necessary.
func (c *Cache) GetOrStart(name string) (string, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if ms, ok := c.services[name]; ok {
		return ms.Address, nil
	}

	ms, err := c.startService(name)
	if err != nil {
		return "", err
	}
	if ms == nil {
		return "", errors.New("failed to start microservice")
	}
	c.services[name] = ms
	return ms.Address, nil
}
