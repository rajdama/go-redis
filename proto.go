package main

import (
	"bytes"
	"fmt"
	"io"
	"log"

	"github.com/tidwall/resp"
)

const (
	CommandSET = "set"
)

type Command interface {
}

type SetCommand struct {
	key, val string
}

func parseCommand(msg string) (Command, error) {
	fmt.Println(msg)
	rd := resp.NewReader(bytes.NewBufferString(msg))
	for {
		v, _, err := rd.ReadValue()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if v.Type() == resp.Array {
			for _, value := range v.Array() {
				switch value.String() {
				case CommandSET:
					if len(v.Array()) != 3 {
						return nil, fmt.Errorf("error in number of set arguments")
					}
					cmd := SetCommand{
						key: v.Array()[1].String(),
						val: v.Array()[2].String(),
					}
					return cmd, nil
				}
			}
		}
	}
	return "foo", nil
}
