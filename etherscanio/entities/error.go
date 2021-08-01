package entities

type Error struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Result  string `json:"result"`
}

func (e *Error) Error() string {
	return e.Result
}
