package app

import (
	"bytes"
	"encoding/gob"
)

type UserData struct {
	Name string
	Args []any
}

func NewUserData(name string, args []any) *UserData {
	return &UserData{
		Name: name,
		Args: args,
	}
}

func (u *UserData) Encode() ([]byte, error) {
	var buf bytes.Buffer
	bufEncode := gob.NewEncoder(&buf)
	if err := bufEncode.Encode(u); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (u *UserData) Decode(b []byte) error {
	buf := bytes.NewBuffer(b)
	bufDec := gob.NewDecoder(buf)
	if err := bufDec.Decode(u); err != nil {
		return err
	}
	return nil
}
