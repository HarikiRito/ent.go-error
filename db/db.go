package db

import (
	"abc/ent"
	"entgo.io/ent/dialect/sql"
)

func Open() (*ent.Client, error) {
	drv, err := sql.Open("mysql", "<mysql-dsn>")
	if err != nil {
		return nil, err
	}
	// Get the underlying sql.DB object of the driver.
	db := drv.DB()
	db.SetMaxIdleConns(10)
	return ent.NewClient(ent.Driver(drv)), nil
}
