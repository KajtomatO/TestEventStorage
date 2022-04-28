package storage_memory

import (
	"errors"
	"sync"
)

func NewInMemoryDataStore() *InMemoryDataStore {
	return &InMemoryDataStore{
		map[string]string{},
		sync.RWMutex{},
	}
}

type InMemoryDataStore struct {
	store map[string]string
	lock  sync.RWMutex
}

func (i *InMemoryDataStore) SetDataRecord(id string, data string) {
	i.lock.Lock()
	defer i.lock.Unlock()
	i.store[id] = data
}

func (i *InMemoryDataStore) GetDataRecord(id string) (string, error) {
	i.lock.RLock()
	defer i.lock.RUnlock()

	data := i.store[id]
	if len(data) == 0 {
		//return "", ErrRecordNotFound
		return "", errors.New("test error")
	}

	return data, nil
}
