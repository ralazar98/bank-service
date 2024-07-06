package accounts

import "errors"

var (
	NotFoundErr = errors.New("not found")
)

type MemStore struct {
	list map[string]float64
}

func NewMemStore() *MemStore {
	list := make(map[string]float64)
	return &MemStore{
		list,
	}
}

func (m MemStore) Add(name string, balance float64) error {
	m.list[name] = balance
	return nil
}
func (m MemStore) List() (map[string]float64, error) {
	return m.list, nil
}

func (m MemStore) Show(name string) (float64, error) {
	return m.list[name], nil
}

func (m MemStore) AddMoney(name string, addedBalance float64) error {
	m.list[name] += addedBalance
	return nil
}
