// +build ignore

package main

import (
	"fmt"
	"github.com/gswly/gomavlib"
	"github.com/gswly/gomavlib/dialects/ardupilotmega"
)

func main() {
	// create a node which
	// - communicates through multiple endpoints
	// - understands ardupilotmega dialect
	// - writes messages with given system id and component id
	node, err := gomavlib.NewNode(gomavlib.NodeConf{
		Endpoints: []gomavlib.EndpointConf{
			gomavlib.EndpointSerial{"/dev/ttyUSB0:57600"},
			gomavlib.EndpointUdpClient{"1.2.3.4:5900"},
		},
		Dialect:     ardupilotmega.Dialect,
		SystemId:    10,
		ComponentId: 1,
	})
	if err != nil {
		panic(err)
	}
	defer node.Close()

	for evt := range node.Events() {
		if frm, ok := evt.(*gomavlib.NodeEventFrame); ok {
			fmt.Printf("received: id=%d, %+v\n", frm.Message().GetId(), frm.Message())

			// route frame to every other channel
			node.WriteFrameExcept(frm.Channel, frm.Frame)
		}
	}
}
