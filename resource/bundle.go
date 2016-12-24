package resource

import (
	"github.com/bearchit/kiss/log"
	"github.com/bearchit/kiss/sql"
)

type Bundle struct {
	Config *Config
	DB     *sql.DB
	Logger *log.Logger
}

func (b *Bundle) Close() {
	if b.DB != nil {
		b.DB.Close()
	}
}

func NewBundle() (*Bundle, error) {
	c, err := LoadConfig()
	if err != nil {
		return nil, err
	}

	logger, err := NewLogger(&c.Log)
	if err != nil {
		return nil, err
	}

	db, err := InitDB(c)
	if err != nil {
		return nil, err
	}

	db.Logger = logger

	return &Bundle{
		Config: c,
		DB:     db,
		Logger: logger,
	}, nil
}
