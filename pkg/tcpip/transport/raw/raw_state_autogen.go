// automatically generated by stateify.

package raw

import (
	"gvisor.dev/gvisor/pkg/state"
	"gvisor.dev/gvisor/pkg/tcpip/buffer"
)

func (p *rawPacket) StateTypeName() string {
	return "pkg/tcpip/transport/raw.rawPacket"
}

func (p *rawPacket) StateFields() []string {
	return []string{
		"rawPacketEntry",
		"data",
		"timestampNS",
		"senderAddr",
	}
}

func (p *rawPacket) beforeSave() {}

func (p *rawPacket) StateSave(stateSinkObject state.Sink) {
	p.beforeSave()
	var dataValue buffer.VectorisedView = p.saveData()
	stateSinkObject.SaveValue(1, dataValue)
	stateSinkObject.Save(0, &p.rawPacketEntry)
	stateSinkObject.Save(2, &p.timestampNS)
	stateSinkObject.Save(3, &p.senderAddr)
}

func (p *rawPacket) afterLoad() {}

func (p *rawPacket) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &p.rawPacketEntry)
	stateSourceObject.Load(2, &p.timestampNS)
	stateSourceObject.Load(3, &p.senderAddr)
	stateSourceObject.LoadValue(1, new(buffer.VectorisedView), func(y interface{}) { p.loadData(y.(buffer.VectorisedView)) })
}

func (ep *endpoint) StateTypeName() string {
	return "pkg/tcpip/transport/raw.endpoint"
}

func (ep *endpoint) StateFields() []string {
	return []string{
		"TransportEndpointInfo",
		"waiterQueue",
		"associated",
		"hdrIncluded",
		"rcvList",
		"rcvBufSize",
		"rcvBufSizeMax",
		"rcvClosed",
		"sndBufSize",
		"sndBufSizeMax",
		"closed",
		"connected",
		"bound",
		"linger",
		"owner",
	}
}

func (ep *endpoint) StateSave(stateSinkObject state.Sink) {
	ep.beforeSave()
	var rcvBufSizeMaxValue int = ep.saveRcvBufSizeMax()
	stateSinkObject.SaveValue(6, rcvBufSizeMaxValue)
	stateSinkObject.Save(0, &ep.TransportEndpointInfo)
	stateSinkObject.Save(1, &ep.waiterQueue)
	stateSinkObject.Save(2, &ep.associated)
	stateSinkObject.Save(3, &ep.hdrIncluded)
	stateSinkObject.Save(4, &ep.rcvList)
	stateSinkObject.Save(5, &ep.rcvBufSize)
	stateSinkObject.Save(7, &ep.rcvClosed)
	stateSinkObject.Save(8, &ep.sndBufSize)
	stateSinkObject.Save(9, &ep.sndBufSizeMax)
	stateSinkObject.Save(10, &ep.closed)
	stateSinkObject.Save(11, &ep.connected)
	stateSinkObject.Save(12, &ep.bound)
	stateSinkObject.Save(13, &ep.linger)
	stateSinkObject.Save(14, &ep.owner)
}

func (ep *endpoint) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &ep.TransportEndpointInfo)
	stateSourceObject.Load(1, &ep.waiterQueue)
	stateSourceObject.Load(2, &ep.associated)
	stateSourceObject.Load(3, &ep.hdrIncluded)
	stateSourceObject.Load(4, &ep.rcvList)
	stateSourceObject.Load(5, &ep.rcvBufSize)
	stateSourceObject.Load(7, &ep.rcvClosed)
	stateSourceObject.Load(8, &ep.sndBufSize)
	stateSourceObject.Load(9, &ep.sndBufSizeMax)
	stateSourceObject.Load(10, &ep.closed)
	stateSourceObject.Load(11, &ep.connected)
	stateSourceObject.Load(12, &ep.bound)
	stateSourceObject.Load(13, &ep.linger)
	stateSourceObject.Load(14, &ep.owner)
	stateSourceObject.LoadValue(6, new(int), func(y interface{}) { ep.loadRcvBufSizeMax(y.(int)) })
	stateSourceObject.AfterLoad(ep.afterLoad)
}

func (l *rawPacketList) StateTypeName() string {
	return "pkg/tcpip/transport/raw.rawPacketList"
}

func (l *rawPacketList) StateFields() []string {
	return []string{
		"head",
		"tail",
	}
}

func (l *rawPacketList) beforeSave() {}

func (l *rawPacketList) StateSave(stateSinkObject state.Sink) {
	l.beforeSave()
	stateSinkObject.Save(0, &l.head)
	stateSinkObject.Save(1, &l.tail)
}

func (l *rawPacketList) afterLoad() {}

func (l *rawPacketList) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &l.head)
	stateSourceObject.Load(1, &l.tail)
}

func (e *rawPacketEntry) StateTypeName() string {
	return "pkg/tcpip/transport/raw.rawPacketEntry"
}

func (e *rawPacketEntry) StateFields() []string {
	return []string{
		"next",
		"prev",
	}
}

func (e *rawPacketEntry) beforeSave() {}

func (e *rawPacketEntry) StateSave(stateSinkObject state.Sink) {
	e.beforeSave()
	stateSinkObject.Save(0, &e.next)
	stateSinkObject.Save(1, &e.prev)
}

func (e *rawPacketEntry) afterLoad() {}

func (e *rawPacketEntry) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &e.next)
	stateSourceObject.Load(1, &e.prev)
}

func init() {
	state.Register((*rawPacket)(nil))
	state.Register((*endpoint)(nil))
	state.Register((*rawPacketList)(nil))
	state.Register((*rawPacketEntry)(nil))
}
