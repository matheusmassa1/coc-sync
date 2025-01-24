package storage

import "context"

type DB interface {
	Save(ctx context.Context, table string, data map[string]interface{}) error
}
