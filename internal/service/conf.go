package service

import (
	"context"
	"nagato/internal/data"
	"nagato/internal/model"
)

type ConfSrv interface {
	GetConfig(ctx context.Context) (*model.Config, error)
}

type confService struct {
	store *data.Store
}

func (s confService) GetConfig(ctx context.Context) (*model.Config, error) {
	c, err := s.store.Configs().Get(context.Background(), &model.Config{InUse: true}, nil)
	model.SetConfig(c)
	if err == nil {
		return c, nil
	}

	config := model.GetConfig()
	config.InUse = true
	err = s.store.Configs().Create(ctx, config, nil)
	if err != nil {
		return nil, err
	}

	return model.GetConfig(), nil
}

func newConfService(s *data.Store) ConfSrv {
	return &confService{store: s}
}
