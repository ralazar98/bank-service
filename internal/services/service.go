package services

type Cache interface {
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

func (b *bank) GetBalance(user TakeBalance) (float64, error) {
	balance, err := b.c.Get(user.ID)
	if err != nil {
		return 0, err
	}
	return balance, nil
}
