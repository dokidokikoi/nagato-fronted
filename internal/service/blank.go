package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"nagato/internal/core"
	"nagato/internal/data"
	"nagato/internal/model"
	"nagato/internal/tools"
	"net/http"
)

type BlankSrv interface {
	CreateBlank(example model.CreateBlank) (int, error)
	List() ([]model.Blank, int, error)
	Get(id uint) (*model.Blank, int, error)
}

type blankService struct {
	store *data.Store
}

func (b blankService) CreateBlank(example model.CreateBlank) (int, error) {
	req, _ := json.Marshal(example)
	_, code, err := tools.SendReq(http.MethodPost, "blank", bytes.NewBuffer(req), map[string]string{"Authorization": Authorization})
	return code, err
}

func (b blankService) List() ([]model.Blank, int, error) {
	data, code, err := tools.SendReq(http.MethodGet, "blank", nil, map[string]string{"Authorization": Authorization})
	if code != 200 || err != nil {
		return nil, code, err
	}

	var res core.Data[core.List[model.Blank]]
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, -1, err
	}
	return res.Data.List, code, err
}

func (b blankService) Get(id uint) (*model.Blank, int, error) {
	data, code, err := tools.SendReq(http.MethodGet, fmt.Sprintf("blank/%d", id), nil, map[string]string{"Authorization": Authorization})
	if code != 200 || err != nil {
		return nil, code, err
	}

	var res core.Data[model.Blank]
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, -1, err
	}
	return &res.Data, code, err
}

func newBlankService(s *data.Store) BlankSrv {
	return &blankService{store: s}
}
