// +build ignore

package main

import (
	"fmt"

	"github.com/aler9/gomavlib"
	"github.com/aler9/gomavlib/dialects/ardupilotmega"
)

func main() {
	// create a node which
	// - communicates with a TCP endpoint in client mode
	// - understands ardupilotmega dialect
	// - writes messages with given system id
	node, err := gomavlib.NewNode(gomavlib.NodeConf{
		Endpoints: []gomavlib.EndpointConf{
			gomavlib.EndpointTcpClient{"1.2.3.4:5600"},
		},
		Dialect:     ardupilotmega.Dialect,
		OutVersion:  gomavlib.V2, // change to V1 if you're unable to communicate with the target
		OutSystemId: 10,
	})
	if err != nil {
		panic(err)
	}
	defer node.Close()

	// print every message we receive
	for evt := range node.Events() {
		if frm, ok := evt.(*gomavlib.EventFrame); ok {
			fmt.Printf("received: id=%d, %+v\n", frm.Message().GetId(), frm.Message())
		}
	}
}
