package app

import (
	"errors"
	"net"
	"reflect"
)

type Client struct {
	conn net.Conn
}

func NewClient(conn net.Conn) *Client {
	return &Client{
		conn: conn,
	}
}

func (c *Client) CallRPC(name string, fptr any) chan error {
	fn := reflect.ValueOf(fptr).Elem()
	if fn.Type().Kind() != reflect.Func {
		panic(errors.New("err fptr not func"))
	}
	errChan := make(chan error, 1)
	f := func(args []reflect.Value) []reflect.Value {
		inArgs := make([]interface{}, 0, len(args))
		for _, arg := range args {
			inArgs = append(inArgs, arg.Interface())
		}
		session := NewSession(c.conn)
		rpcData := NewUserData(name, inArgs)
		b, err := rpcData.Encode()
		if err != nil {
			errChan <- errors.New("Error Encode err :" + err.Error())
			return nil
		}
		if err = session.Write(b); err != nil {
			errChan <- err
			return nil
		}
		responseBytes, err := session.Read()
		if err != nil {
			errChan <- err
			return nil
		}
		responseRPC := new(UserData)
		if err := responseRPC.Decode(responseBytes); err != nil {
			errChan <- err
			return nil
		}
		outArgs := make([]reflect.Value, 0, len(responseRPC.Args))
		for i, arg := range responseRPC.Args {
			if arg == nil {
				outArgs = append(outArgs, reflect.Zero(fn.Type().Out(i)))
				continue
			}
			outArgs = append(outArgs, reflect.ValueOf(arg))
		}
		errChan <- err
		return outArgs
	}
	v := reflect.MakeFunc(fn.Type(), f)
	fn.Set(v)
	return errChan
}
