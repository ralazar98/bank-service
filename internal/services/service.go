package services

type Operation string

type Cache interface {
	Create(id int, balance float64) error
	Update(id int, changingInBalance float64, operation Operation) error
	Get(id int) (float64, error)
}

type bank struct {
	c Cache
}

func New(c Cache) *bank {
	return &bank{
		c: c,
	}
}

func (b *bank) GetBalance(userID int) (float64, error) {
	balance, err := b.c.Get(userID)
	if err != nil {
		return 0, err
	}
	return balance, nil
}
