package graph

import "bank-service/internal/repository/postgresql"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	DB *postgresql.BankStorage
}
