package memory_cache

import "errors"

type operation string

var (
	NotFoundError          = errors.New("not found")
	NotEnoughtBalanceError = errors.New("not enough balance")
	InvalidOperationError  = errors.New("invalid operation")
	WrongAccountIdError    = errors.New("wrong account id")
)

const (
	OperationCreate operation = "create"
	OperationTake   operation = "take"
	OperationAdd    operation = "add"
	OperationShow   operation = "show"
)

type BankStorage struct {
	// todo: mutex
	list map[int]float64
}

func New() BankStorage {
	list := make(map[int]float64)
	return BankStorage{
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

func (b *BankStorage) UpdateBalance(id int, changingInBalance float64, operationType operation) error {
	if _, ok := b.list[id]; !ok {
		return WrongAccountIdError
	}
	if operationType == OperationTake && changingInBalance > b.list[id] {
		return NotEnoughtBalanceError
	}
	switch {
	case operationType == OperationTake && changingInBalance <= b.list[id]:
		b.list[id] = b.list[id] - changingInBalance
	case operationType == OperationAdd:
		b.list[id] = b.list[id] + changingInBalance
	default:
		return InvalidOperationError
	}
	return nil
}
