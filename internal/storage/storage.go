package storage

import "errors"

type Operation string

const (
	OperationCreate Operation = "create"
	OperationTake   Operation = "take"
	OperationAdd    Operation = "add"
	OperationShow   Operation = "show"
)

type BankStorage struct {
	list map[int]float64
}

func New() *BankStorage {
	list := make(map[int]float64)
	return &BankStorage{
		list,
	}
}

func (b *BankStorage) Create(id int, balance float64) error {
	b.list[id] = balance
	return nil
}

func (b *BankStorage) List() (map[int]float64, error) {
	return b.list, nil
}

func (b *BankStorage) Show(id int) (float64, error) {
	return b.list[id], nil
}

func (b *BankStorage) UpdateBalance(id int, changingInBalance float64, operation Operation) error {
	if _, ok := b.list[id]; !ok {
		switch {
		case operation == OperationTake && changingInBalance <= b.list[id]:
			b.list[id] = b.list[id] - changingInBalance
		case operation == OperationAdd:
			b.list[id] = b.list[id] + changingInBalance
		default:
			return errors.New("invalid operation")
		}
	}
	return nil
}
