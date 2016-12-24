package resource

import (
	_ "github.com/go-sql-driver/mysql"

	"github.com/bearchit/kiss/sql"
	"github.com/serenize/snaker"
)

func InitDB(c *Config) (*sql.DB, error) {
	return sql.OpenMySQL(&sql.Config{
		Host:         c.DB.Host,
		Port:         c.DB.Port,
		Name:         c.DB.Name,
		User:         c.DB.User,
		Password:     c.DB.Password,
		Charset:      c.DB.Charset,
		Location:     c.DB.Location,
		MaxIdleConns: c.DB.MaxIdleConns,
		MaxOpenConns: c.DB.MaxOpenConns,
		MapperFunc:   snaker.CamelToSnake,
	})
}
