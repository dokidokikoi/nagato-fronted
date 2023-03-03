package data

import (
	"nagato/internal/model"

	db "github.com/dokidokikoi/go-common/db/base"
)

type configs struct {
	db.PgModel[model.Config]
}

func newConfigs(store *Store) *configs {
	return &configs{PgModel: db.PgModel[model.Config]{DB: store.db}}
}
