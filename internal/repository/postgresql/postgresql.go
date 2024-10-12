package postgresql

import (
	"bank-service/internal/entity"
	"bank-service/internal/services"
	"github.com/jackc/pgx"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
	"time"
)

type BankStorage struct {
	conn *pgx.ConnPool
}

func New() *BankStorage {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbHost := os.Getenv("DB_HOST")
	dbPortStr := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort, err := strconv.Atoi(dbPortStr)
	if err != nil {
		log.Fatal("Error converting DB_PORT to integer")
	}

	connConf := pgx.ConnConfig{
		Host:     dbHost,
		Port:     uint16(dbPort),
		User:     dbUser,
		Password: dbPassword,
		Database: dbName,
	}
	pool, err := pgx.NewConnPool(pgx.ConnPoolConfig{
		ConnConfig:     connConf,
		MaxConnections: 20,
	})
	if err != nil {
		log.Println(err.Error())
	}
	log.Print("Подключено к БД")
	return &BankStorage{
		conn: pool,
	}
}

func (s *BankStorage) CreateAccount(user *services.CreateAccount) (*entity.User, error) {

	date := time.Now().UTC().Format("2006-01-02")
	request := `
		INSERT INTO public.bank_storage (user_id, balance,date_created,date_updated)
		VALUES ($1,$2,$3,$4)
		returning user_id,balance;
	`
	var entityUser entity.User
	err := s.conn.QueryRow(request, user.UserID, user.Balance, date, date).Scan(&entityUser.ID, &entityUser.Balance.Sum)
	if err != nil {
		return nil, err
	}
	return &entityUser, err
}

func (s *BankStorage) GetBalance(user *services.GetBalance) (*entity.User, error) {
	request := `
			SELECT user_id, balance FROM public.bank_storage where user_id = $1
	`
	var entityUser entity.User
	err := s.conn.QueryRow(request, user.UserID).Scan(&entityUser.ID, &entityUser.Balance.Sum)
	return &entityUser, err
}
func (s *BankStorage) UpdateBalance(user *services.UpdateBalance) (*entity.User, error) {
	request := `
			UPDATE public.bank_storage SET balance = balance+$2,date_updated=$3 WHERE user_id = $1
			RETURNING user_id, balance
	`
	dateUpdate := time.Now().UTC().Format("2006-01-02")
	var entityUser entity.User

	err := s.conn.QueryRow(request, user.UserID, user.ChangingInBalance, dateUpdate).
		Scan(&entityUser.ID, &entityUser.Balance.Sum)
	return &entityUser, err

}
