// automatically generated by stateify.

package signalfd

import (
	"gvisor.dev/gvisor/pkg/state"
)

func (s *SignalOperations) StateTypeName() string {
	return "pkg/sentry/kernel/signalfd.SignalOperations"
}

func (s *SignalOperations) StateFields() []string {
	return []string{
		"target",
		"mask",
	}
}

func (s *SignalOperations) beforeSave() {}

func (s *SignalOperations) StateSave(stateSinkObject state.Sink) {
	s.beforeSave()
	stateSinkObject.Save(0, &s.target)
	stateSinkObject.Save(1, &s.mask)
}

func (s *SignalOperations) afterLoad() {}

func (s *SignalOperations) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &s.target)
	stateSourceObject.Load(1, &s.mask)
}

func init() {
	state.Register((*SignalOperations)(nil))
}
