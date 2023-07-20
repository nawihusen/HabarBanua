package mysql

import (
	"be-service-saksi-management/domain"
	"database/sql"
)

type mysqlAccountRepository struct {
	Conn *sql.DB
}

// NewMySQLAccountRepository is constructor of MySQL repository
func NewMySQLAccountRepository(Conn *sql.DB) domain.AccountMySQLRepository {
	return &mysqlAccountRepository{Conn}
}
