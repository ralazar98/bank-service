package postgresql

import (
	"bank-service/configs"
	"bank-service/internal/entity"
	"github.com/jackc/pgx"
	"time"
)

type BankStorage struct {
	conn *pgx.ConnPool
}

func New(cfg configs.DataBaseConfig) (*BankStorage, error) {

	connConf := pgx.ConnConfig{
		Host:     cfg.Host,
		Port:     uint16(cfg.Port),
		User:     cfg.User,
		Password: cfg.Password,
		Database: cfg.DBName,
	}
	pool, err := pgx.NewConnPool(pgx.ConnPoolConfig{
		ConnConfig:     connConf,
		MaxConnections: cfg.MaxConnections,
	})
	if err != nil {

		return nil, err
	}
	return &BankStorage{
		conn: pool,
	}, nil
}

func (s *BankStorage) CreateAccount(user *entity.CreateAccount) (*entity.User, error) {

	request := `
		INSERT INTO public.bank_storage (user_id, balance,date_created,date_updated)
		VALUES ($1,$2,$3,$4)
		returning user_id,balance;
	`
	var entityUser entity.User
	err := s.conn.QueryRow(request, user.UserID, user.Balance, time.Now(), time.Now()).Scan(&entityUser.ID, &entityUser.Balance.Sum)
	if err != nil {
		return nil, err
	}
	return &entityUser, nil
}

func (s *BankStorage) GetBalance(user *entity.GetBalance) (*entity.User, error) {
	request := `
			SELECT user_id, balance FROM public.bank_storage where user_id = $1
	`
	var entityUser entity.User
	err := s.conn.QueryRow(request, user.UserID).Scan(&entityUser.ID, &entityUser.Balance.Sum)
	if err != nil {
		return nil, err
	}
	return &entityUser, err
}

func (s *BankStorage) Update(user *entity.UpdateBalance) (*entity.User, error) {
	return nil, nil
}
