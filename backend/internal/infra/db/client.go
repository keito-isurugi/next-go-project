package db

import (
	"context"
	"fmt"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/keito-isurugi/next-go-project/internal/infra/env"
)

type Client interface {
	Conn(ctx context.Context) *gorm.DB
	Close() error
}

type client struct {
	db *gorm.DB
}

func NewClient(e *env.DB, zapLogger *zap.Logger) (Client, error) {
	gormLogger := initGormLogger(zapLogger)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo",
		e.PostgresHost,
		e.PostgresUser,
		e.PostgresPassword,
		e.PostgresDatabase,
		e.PostgresPort,
	)
	db, err := gorm.Open(
		postgres.Open(dsn),
		&gorm.Config{
			Logger: gormLogger,
		},
	)

	if err != nil {
		return nil, err
	}
	fmt.Println(db)
	db.Logger = db.Logger.LogMode(gormLogger.LogLevel)

	return &client{
		db: db,
	}, nil
}

func (c *client) Conn(ctx context.Context) *gorm.DB {
	return c.db.WithContext(ctx)
}

func (c *client) Close() error {
	db, err := c.db.DB()
	if err != nil {
		return err
	}
	return db.Close()
}
