package accounts

import "errors"

var (
	NotFoundErr         = errors.New("not found")
	NotEnoughBalanceErr = errors.New("not enough balance")
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

func (m MemStore) Add(id string, balance float64) error {
	m.list[id] = balance
	return nil
}
func (m MemStore) List() (map[string]float64, error) {
	return m.list, nil
}

func (m MemStore) Show(id string) (float64, error) {
	return m.list[id], nil
}

func (m MemStore) ChangeBalance(name string, operation string, changingInBalance float64) error {
	if operation == "take" {
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
