package data

import (
	"nagato/internal/model"

	db "github.com/dokidokikoi/go-common/db/base"
)

type downloads struct {
	db.PgModel[model.Download]
}

func newDownloads(store *Store) *downloads {
	return &downloads{PgModel: db.PgModel[model.Download]{DB: store.db}}
}
