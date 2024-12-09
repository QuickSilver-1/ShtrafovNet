package postgres

import (
	"auction/internal/domain/interfaces"
	derr "auction/internal/infrastructure/errors"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type PostgresDb struct {
    ip     string
    port   int
    nameDB string
    user   string
    pass   string
    conn   *sql.DB
}

// NewDb создает новое соединение с базой данных PostgreSQL
func NewDb(ip, name, user, pass string, port int, logger interfaces.LoggerRepo) *PostgresDb {
    db := &PostgresDb{
        ip:     ip,
        port:   port,
        nameDB: name,
        user:   user,
        pass:   pass,
    }

    logger.Info("connection start")
    err := db.Connect()

    if err != nil {
        logger.Error("database connection error:" + err.Error())
        panic(err)
    }

    logger.Debug("connection successful")
    return db
}

// Connect устанавливает соединение с базой данных
func (db *PostgresDb) Connect() error {

    sqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
        db.ip, db.port, db.user, db.pass, db.nameDB)
    conn, err := sql.Open("postgres", sqlInfo)

    if err != nil {
        return &derr.ErrDatabaseConnection{
            Err:  fmt.Sprintf("Failed to connect to the postgres database with data - host: %s, port: %d, database: %s, user: %s, pass: %s", db.ip, db.port, db.nameDB, db.user, db.pass),
            Code: 400,
        }
    }

    db.conn = conn

    return nil
}

// Query выполняет SQL-запрос к базе данных
func (db *PostgresDb) Query(sql string, args ...any) (*sql.Rows, error) {
    rows, err := db.conn.Query(sql, args...)

    if err != nil {
        return nil, &derr.QueryErr{
            Err:  fmt.Sprintf("Database query error: %v", err),
            Code: 500,
        }
    }

    return rows, nil
}

// Close закрывает соединение с базой данных
func (db *PostgresDb) Close() {
    db.conn.Close()
}