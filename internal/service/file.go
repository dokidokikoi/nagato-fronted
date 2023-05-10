package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"nagato/internal/core"
	"nagato/internal/data"
	"nagato/internal/model"
	"nagato/internal/tools"
	"net/http"
	"os"
)

type FileSrv interface {
	GetFolderByUid(uuid string) (*model.Matter, int, error)
	Upload(name, filePath, puuid string) (int, error)
	CreateDir(name, puuid string) (int, error)
}

type fileService struct {
	store *data.Store
}

func (f fileService) GetFolderByUid(uuid string) (*model.Matter, int, error) {
	data, code, err := tools.SendReq(http.MethodGet, fmt.Sprintf("matter/%s", uuid), nil, map[string]string{"Authorization": Authorization})
	if err != nil {
		return nil, code, err
	}

	var res core.Data[model.Matter]
	fmt.Println(string(data))
	err = json.Unmarshal(data, &res)
	return &res.Data, code, err
}

func (f fileService) Upload(name, filePath, puuid string) (int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return -1, err
	}

	info, err := file.Stat()
	if err != nil {
		return -1, err
	}

	sha256, err := tools.CalculateHash(file)
	if err != nil {
		return -1, err
	}

	body := model.UploadToken{
		Name:    name,
		Sha256:  sha256,
		Size:    uint(info.Size()),
		Privacy: false,
		PUUID:   puuid,
	}
	req, _ := json.Marshal(body)
	data, code, err := tools.SendReq(http.MethodPost, "file", bytes.NewBuffer(req), map[string]string{"Authorization": Authorization})
	if err != nil {
		return code, err
	}
	if code != 200 {
		return code, err
	}

	var token core.Data[string]
	json.Unmarshal(data, &token)

	file, err = os.Open(filePath)
	count := int(info.Size() / 240000)
	if info.Size()%240000 != 0 {
		count++
	}
	var buf = make([]byte, 240000)
	for i := 0; i < count; i++ {
		n, err := io.ReadFull(file, buf)
		if err != nil && err != io.EOF && err != io.ErrUnexpectedEOF {
			return code, err
		}
		buf = buf[:n]

		_, code, err = tools.SendReq(
			http.MethodPut,
			fmt.Sprintf("file/temp/%s", token.Data),
			bytes.NewBuffer(buf),
			map[string]string{
				"range":         fmt.Sprintf("bytes= %d-", i*240000),
				"Authorization": Authorization,
			},
		)
		if err != nil {
			return code, err
		}
		if code != 200 {
			return code, err
		}
	}

	return code, err
}

func (f fileService) CreateDir(name, puuid string) (int, error) {
	type CreateDir struct {
		Name    string `json:"name" binding:"required"`
		Privacy bool   `json:"privacy"`
		PUUID   string `json:"puuid"`
	}
	body := CreateDir{
		Name:    name,
		Privacy: false,
		PUUID:   puuid,
	}
	req, _ := json.Marshal(body)
	_, code, err := tools.SendReq(http.MethodPost, "matter/dir", bytes.NewBuffer(req), map[string]string{"Authorization": Authorization})
	return code, err
}

func newFileService(s *data.Store) FileSrv {
	return &fileService{store: s}
}
