package validate

type Error struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
	Msg   string `json:"tip"`
}

func (e Error) Error() string {
	return e.Msg
}

func newError(field, tag, msg string) error {
	return &Error{
		Field: field,
		Tag:   tag,
		Msg:   msg,
	}
}
