// automatically generated by stateify.

// +build go1.9
// +build !go1.17

package tcpip

import (
	"gvisor.dev/gvisor/pkg/state"
)

func (x *StdClock) StateTypeName() string {
	return "pkg/tcpip.StdClock"
}

func (x *StdClock) StateFields() []string {
	return []string{}
}

func (x *StdClock) beforeSave() {}

func (x *StdClock) StateSave(m state.Sink) {
	x.beforeSave()
}

func (x *StdClock) afterLoad() {}

func (x *StdClock) StateLoad(m state.Source) {
}

func init() {
	state.Register((*StdClock)(nil))
}
