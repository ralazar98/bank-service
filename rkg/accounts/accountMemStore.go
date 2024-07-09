package accounts

import (
	"errors"
	"sync"
)

var (
	NotFoundErr         = errors.New("not found")
	NotEnoughBalanceErr = errors.New("not enough balance")
)

type operation string

const (
	OperationTake operation = "take"
)

type MemStore struct {
	list map[string]float64
	mu   *sync.Mutex
}

func NewMemStore() *MemStore {
	list := make(map[string]float64)
	return &MemStore{
		list,
	}
}

func (m *MemStore) Add(id string, balance float64) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.list[id] = balance
	return nil
}
func (m *MemStore) List() (map[string]float64, error) {
	return m.list, nil
}

func (m *MemStore) Show(id string) (float64, error) {
	return m.list[id], nil
}

func (m *MemStore) ChangeBalance(name string, operation operation, changingInBalance float64) error {
	if operation == OperationTake {
		if changingInBalance < m.list[name] {
			m.list[name] -= changingInBalance
		} else {
			return NotEnoughBalanceErr
		}
	} else if operation == "add" {
		m.list[name] += changingInBalance
	} else {
		return NotFoundErr
	}
	return nil
}
