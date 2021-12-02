package database

import (
	"context"
	"time"
)

func NewDatabaseContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 10*time.Second)
}
