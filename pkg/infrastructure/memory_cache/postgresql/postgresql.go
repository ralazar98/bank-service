package postgresql

import (
	"bank-service/internal/config"
	"bank-service/internal/entity"
	"bank-service/internal/services"
	"github.com/jackc/pgx"
	"time"
)

type BankStorage struct {
	conn *pgx.ConnPool
}

func New(conn *pgx.ConnPool) *BankStorage {
	return &BankStorage{
		conn: conn,
	}
}

func NewConnConfig(cfg *config.Config) pgx.ConnConfig {
	return pgx.ConnConfig{
		Host:     cfg.Host,
		Port:     cfg.Port,
		User:     cfg.Username,
		Password: cfg.Password,
		Database: cfg.Database,
	}
}

func NewConnect(conConf pgx.ConnConfig) (*pgx.ConnPool, error) {
	pool, err := pgx.NewConnPool(pgx.ConnPoolConfig{
		ConnConfig:     conConf,
		MaxConnections: 20,
	})
	return pool, err
}

func (s *BankStorage) CreateAccount(user *services.CreateAccount) (*entity.User, error) {

	date := time.Now().UTC().Format("2006-01-02")
	q := `
		INSERT INTO public.bank_storage (user_id, balance,date_created,date_updated)
		VALUES ($1,$2,$3,$4)
		returning user_id,balance;
	`
	var entityUser entity.User
	err := s.conn.QueryRow(q, user.UserID, user.Balance, date, date).Scan(&entityUser.ID, &entityUser.Balance.Sum)
	if err != nil {
		return nil, err
	}
	return &entityUser, err
}

func (s *BankStorage) GetBalance(user *services.GetBalance) (*entity.User, error) {
	q := `
			SELECT user_id, balance FROM public.bank_storage where user_id = $1
	`
	var entityUser entity.User
	err := s.conn.QueryRow(q, user.UserID).Scan(&entityUser.ID, &entityUser.Balance.Sum)
	return &entityUser, err
}
func (s *BankStorage) UpdateBalance(user *services.UpdateBalance) (*entity.User, error) {
	q := `
			UPDATE public.bank_storage SET balance = balance+$2,date_updated=$3 WHERE user_id = $1
			RETURNING user_id, balance
	`
	dateUpdate := time.Now().UTC().Format("2006-01-02")
	var entityUser entity.User

	err := s.conn.QueryRow(q, user.UserID, user.ChangingInBalance, dateUpdate).
		Scan(&entityUser.ID, &entityUser.Balance.Sum)
	return &entityUser, err

}
