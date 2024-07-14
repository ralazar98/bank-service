package services

type Cache interface {
	Create(id int, balance float64) error
	Update(id int, changingInBalance float64, operation Operation) error
}

type bank struct {
	c Cache
}

func New(c Cache) *bank {
	return &bank{
		c: c,
	}
}

func (b *bank) Create(user CreateAcc) error {
	err := b.c.Create(user.ID, user.Balance)
	if err != nil {
		return err
	}
	return nil
}
