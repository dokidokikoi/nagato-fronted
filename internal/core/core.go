package core

type Data[T interface{}] struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

type List[T interface{}] struct {
	List  []T   `json:"list"`
	Total int64 `json:"total"`
}
