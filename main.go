package main

import (
	"fmt"

	"github.com/benhoyt/goland-repro/common"
)

type Client struct {
	facade interface{}
}

type ClientSidecar struct {
	*Client
	*common.CharmsClient
}

func (c *ClientSidecar) foo() {
	fmt.Println(c.facade)
}

func main() {
	c := &ClientSidecar{Client: &Client{facade: 42}}
	c.foo()
}
