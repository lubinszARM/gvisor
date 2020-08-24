// automatically generated by stateify.

package proc

import (
	"gvisor.dev/gvisor/pkg/state"
)

func (x *FilesystemType) StateTypeName() string {
	return "pkg/sentry/fsimpl/proc.FilesystemType"
}

func (x *FilesystemType) StateFields() []string {
	return []string{}
}

func (x *FilesystemType) beforeSave() {}

func (x *FilesystemType) StateSave(m state.Sink) {
	x.beforeSave()
}

func (x *FilesystemType) afterLoad() {}

func (x *FilesystemType) StateLoad(m state.Source) {
}

func (x *subtasksInode) StateTypeName() string {
	return "pkg/sentry/fsimpl/proc.subtasksInode"
}

func (x *subtasksInode) StateFields() []string {
	return []string{
		"InodeNotSymlink",
		"InodeDirectoryNoNewChildren",
		"InodeAttrs",
		"OrderedChildren",
		"AlwaysValid",
		"locks",
		"fs",
		"task",
		"pidns",
		"cgroupControllers",
	}
}

func (x *subtasksInode) beforeSave() {}

func (x *subtasksInode) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.InodeNotSymlink)
	m.Save(1, &x.InodeDirectoryNoNewChildren)
	m.Save(2, &x.InodeAttrs)
	m.Save(3, &x.OrderedChildren)
	m.Save(4, &x.AlwaysValid)
	m.Save(5, &x.locks)
	m.Save(6, &x.fs)
	m.Save(7, &x.task)
	m.Save(8, &x.pidns)
	m.Save(9, &x.cgroupControllers)
}

func (x *subtasksInode) afterLoad() {}

func (x *subtasksInode) StateLoad(m state.Source) {
	m.Load(0, &x.InodeNotSymlink)
	m.Load(1, &x.InodeDirectoryNoNewChildren)
	m.Load(2, &x.InodeAttrs)
	m.Load(3, &x.OrderedChildren)
	m.Load(4, &x.AlwaysValid)
	m.Load(5, &x.locks)
	m.Load(6, &x.fs)
	m.Load(7, &x.task)
	m.Load(8, &x.pidns)
	m.Load(9, &x.cgroupControllers)
}

func (x *taskInode) StateTypeName() string {
	return "pkg/sentry/fsimpl/proc.taskInode"
}

func (x *taskInode) StateFields() []string {
	return []string{
		"InodeNotSymlink",
		"InodeDirectoryNoNewChildren",
		"InodeNoDynamicLookup",
		"InodeAttrs",
		"OrderedChildren",
		"locks",
		"task",
	}
}

func (x *taskInode) beforeSave() {}

func (x *taskInode) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.InodeNotSymlink)
	m.Save(1, &x.InodeDirectoryNoNewChildren)
	m.Save(2, &x.InodeNoDynamicLookup)
	m.Save(3, &x.InodeAttrs)
	m.Save(4, &x.OrderedChildren)
	m.Save(5, &x.locks)
	m.Save(6, &x.task)
}

func (x *taskInode) afterLoad() {}

func (x *taskInode) StateLoad(m state.Source) {
	m.Load(0, &x.InodeNotSymlink)
	m.Load(1, &x.InodeDirectoryNoNewChildren)
	m.Load(2, &x.InodeNoDynamicLookup)
	m.Load(3, &x.InodeAttrs)
	m.Load(4, &x.OrderedChildren)
	m.Load(5, &x.locks)
	m.Load(6, &x.task)
}

func (x *fdDirInode) StateTypeName() string {
	return "pkg/sentry/fsimpl/proc.fdDirInode"
}

func (x *fdDirInode) StateFields() []string {
	return []string{
		"InodeNotSymlink",
		"InodeDirectoryNoNewChildren",
		"InodeAttrs",
		"OrderedChildren",
		"AlwaysValid",
		"fdDir",
	}
}

func (x *fdDirInode) beforeSave() {}

func (x *fdDirInode) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.InodeNotSymlink)
	m.Save(1, &x.InodeDirectoryNoNewChildren)
	m.Save(2, &x.InodeAttrs)
	m.Save(3, &x.OrderedChildren)
	m.Save(4, &x.AlwaysValid)
	m.Save(5, &x.fdDir)
}

func (x *fdDirInode) afterLoad() {}

func (x *fdDirInode) StateLoad(m state.Source) {
	m.Load(0, &x.InodeNotSymlink)
	m.Load(1, &x.InodeDirectoryNoNewChildren)
	m.Load(2, &x.InodeAttrs)
	m.Load(3, &x.OrderedChildren)
	m.Load(4, &x.AlwaysValid)
	m.Load(5, &x.fdDir)
}

func (x *fdSymlink) StateTypeName() string {
	return "pkg/sentry/fsimpl/proc.fdSymlink"
}

func (x *fdSymlink) StateFields() []string {
	return []string{
		"InodeAttrs",
		"InodeNoopRefCount",
		"InodeSymlink",
		"task",
		"fd",
	}
}

func (x *fdSymlink) beforeSave() {}

func (x *fdSymlink) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.InodeAttrs)
	m.Save(1, &x.InodeNoopRefCount)
	m.Save(2, &x.InodeSymlink)
	m.Save(3, &x.task)
	m.Save(4, &x.fd)
}

func (x *fdSymlink) afterLoad() {}

func (x *fdSymlink) StateLoad(m state.Source) {
	m.Load(0, &x.InodeAttrs)
	m.Load(1, &x.InodeNoopRefCount)
	m.Load(2, &x.InodeSymlink)
	m.Load(3, &x.task)
	m.Load(4, &x.fd)
}

func (x *fdInfoDirInode) StateTypeName() string {
	return "pkg/sentry/fsimpl/proc.fdInfoDirInode"
}

func (x *fdInfoDirInode) StateFields() []string {
	return []string{
		"InodeNotSymlink",
		"InodeDirectoryNoNewChildren",
		"InodeAttrs",
		"OrderedChildren",
		"AlwaysValid",
		"fdDir",
	}
}

func (x *fdInfoDirInode) beforeSave() {}

func (x *fdInfoDirInode) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.InodeNotSymlink)
	m.Save(1, &x.InodeDirectoryNoNewChildren)
	m.Save(2, &x.InodeAttrs)
	m.Save(3, &x.OrderedChildren)
	m.Save(4, &x.AlwaysValid)
	m.Save(5, &x.fdDir)
}

func (x *fdInfoDirInode) afterLoad() {}

func (x *fdInfoDirInode) StateLoad(m state.Source) {
	m.Load(0, &x.InodeNotSymlink)
	m.Load(1, &x.InodeDirectoryNoNewChildren)
	m.Load(2, &x.InodeAttrs)
	m.Load(3, &x.OrderedChildren)
	m.Load(4, &x.AlwaysValid)
	m.Load(5, &x.fdDir)
}

func (x *fdInfoData) StateTypeName() string {
	return "pkg/sentry/fsimpl/proc.fdInfoData"
}

func (x *fdInfoData) StateFields() []string {
	return []string{
		"DynamicBytesFile",
		"AtomicRefCount",
		"task",
		"fd",
	}
}

func (x *fdInfoData) beforeSave() {}

func (x *fdInfoData) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.DynamicBytesFile)
	m.Save(1, &x.AtomicRefCount)
	m.Save(2, &x.task)
	m.Save(3, &x.fd)
}

func (x *fdInfoData) afterLoad() {}

func (x *fdInfoData) StateLoad(m state.Source) {
	m.Load(0, &x.DynamicBytesFile)
	m.Load(1, &x.AtomicRefCount)
	m.Load(2, &x.task)
	m.Load(3, &x.fd)
}

func (x *auxvData) StateTypeName() string {
	return "pkg/sentry/fsimpl/proc.auxvData"
}

func (x *auxvData) StateFields() []string {
	return []string{
		"DynamicBytesFile",
		"task",
	}
}

func (x *auxvData) beforeSave() {}

func (x *auxvData) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.DynamicBytesFile)
	m.Save(1, &x.task)
}

func (x *auxvData) afterLoad() {}

func (x *auxvData) StateLoad(m state.Source) {
	m.Load(0, &x.DynamicBytesFile)
	m.Load(1, &x.task)
}

func (x *cmdlineData) StateTypeName() string {
	return "pkg/sentry/fsimpl/proc.cmdlineData"
}

func (x *cmdlineData) StateFields() []string {
	return []string{
		"DynamicBytesFile",
		"task",
		"arg",
	}
}

func (x *cmdlineData) beforeSave() {}

func (x *cmdlineData) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.DynamicBytesFile)
	m.Save(1, &x.task)
	m.Save(2, &x.arg)
}

func (x *cmdlineData) afterLoad() {}

func (x *cmdlineData) StateLoad(m state.Source) {
	m.Load(0, &x.DynamicBytesFile)
	m.Load(1, &x.task)
	m.Load(2, &x.arg)
}

func (x *commInode) StateTypeName() string {
	return "pkg/sentry/fsimpl/proc.commInode"
}

func (x *commInode) StateFields() []string {
	return []string{
		"DynamicBytesFile",
		"task",
	}
}

func (x *commInode) beforeSave() {}

func (x *commInode) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.DynamicBytesFile)
	m.Save(1, &x.task)
}

func (x *commInode) afterLoad() {}

func (x *commInode) StateLoad(m state.Source) {
	m.Load(0, &x.DynamicBytesFile)
	m.Load(1, &x.task)
}

func (x *commData) StateTypeName() string {
	return "pkg/sentry/fsimpl/proc.commData"
}

func (x *commData) StateFields() []string {
	return []string{
		"DynamicBytesFile",
		"task",
	}
}

func (x *commData) beforeSave() {}

func (x *commData) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.DynamicBytesFile)
	m.Save(1, &x.task)
}

func (x *commData) afterLoad() {}

func (x *commData) StateLoad(m state.Source) {
	m.Load(0, &x.DynamicBytesFile)
	m.Load(1, &x.task)
}

func (x *idMapData) StateTypeName() string {
	return "pkg/sentry/fsimpl/proc.idMapData"
}

func (x *idMapData) StateFields() []string {
	return []string{
		"DynamicBytesFile",
		"task",
		"gids",
	}
}

func (x *idMapData) beforeSave() {}

func (x *idMapData) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.DynamicBytesFile)
	m.Save(1, &x.task)
	m.Save(2, &x.gids)
}

func (x *idMapData) afterLoad() {}

func (x *idMapData) StateLoad(m state.Source) {
	m.Load(0, &x.DynamicBytesFile)
	m.Load(1, &x.task)
	m.Load(2, &x.gids)
}

func (x *mapsData) StateTypeName() string {
	return "pkg/sentry/fsimpl/proc.mapsData"
}

func (x *mapsData) StateFields() []string {
	return []string{
		"DynamicBytesFile",
		"task",
	}
}

func (x *mapsData) beforeSave() {}

func (x *mapsData) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.DynamicBytesFile)
	m.Save(1, &x.task)
}

func (x *mapsData) afterLoad() {}

func (x *mapsData) StateLoad(m state.Source) {
	m.Load(0, &x.DynamicBytesFile)
	m.Load(1, &x.task)
}

func (x *smapsData) StateTypeName() string {
	return "pkg/sentry/fsimpl/proc.smapsData"
}

func (x *smapsData) StateFields() []string {
	return []string{
		"DynamicBytesFile",
		"task",
	}
}

func (x *smapsData) beforeSave() {}

func (x *smapsData) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.DynamicBytesFile)
	m.Save(1, &x.task)
}

func (x *smapsData) afterLoad() {}

func (x *smapsData) StateLoad(m state.Source) {
	m.Load(0, &x.DynamicBytesFile)
	m.Load(1, &x.task)
}

func (x *taskStatData) StateTypeName() string {
	return "pkg/sentry/fsimpl/proc.taskStatData"
}

func (x *taskStatData) StateFields() []string {
	return []string{
		"DynamicBytesFile",
		"task",
		"tgstats",
		"pidns",
	}
}

func (x *taskStatData) beforeSave() {}

func (x *taskStatData) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.DynamicBytesFile)
	m.Save(1, &x.task)
	m.Save(2, &x.tgstats)
	m.Save(3, &x.pidns)
}

func (x *taskStatData) afterLoad() {}

func (x *taskStatData) StateLoad(m state.Source) {
	m.Load(0, &x.DynamicBytesFile)
	m.Load(1, &x.task)
	m.Load(2, &x.tgstats)
	m.Load(3, &x.pidns)
}

func (x *statmData) StateTypeName() string {
	return "pkg/sentry/fsimpl/proc.statmData"
}

func (x *statmData) StateFields() []string {
	return []string{
		"DynamicBytesFile",
		"task",
	}
}

func (x *statmData) beforeSave() {}

func (x *statmData) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.DynamicBytesFile)
	m.Save(1, &x.task)
}

func (x *statmData) afterLoad() {}

func (x *statmData) StateLoad(m state.Source) {
	m.Load(0, &x.DynamicBytesFile)
	m.Load(1, &x.task)
}

func (x *statusData) StateTypeName() string {
	return "pkg/sentry/fsimpl/proc.statusData"
}

func (x *statusData) StateFields() []string {
	return []string{
		"DynamicBytesFile",
		"task",
		"pidns",
	}
}

func (x *statusData) beforeSave() {}

func (x *statusData) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.DynamicBytesFile)
	m.Save(1, &x.task)
	m.Save(2, &x.pidns)
}

func (x *statusData) afterLoad() {}

func (x *statusData) StateLoad(m state.Source) {
	m.Load(0, &x.DynamicBytesFile)
	m.Load(1, &x.task)
	m.Load(2, &x.pidns)
}

func (x *ioData) StateTypeName() string {
	return "pkg/sentry/fsimpl/proc.ioData"
}

func (x *ioData) StateFields() []string {
	return []string{
		"DynamicBytesFile",
		"ioUsage",
	}
}

func (x *ioData) beforeSave() {}

func (x *ioData) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.DynamicBytesFile)
	m.Save(1, &x.ioUsage)
}

func (x *ioData) afterLoad() {}

func (x *ioData) StateLoad(m state.Source) {
	m.Load(0, &x.DynamicBytesFile)
	m.Load(1, &x.ioUsage)
}

func (x *oomScoreAdj) StateTypeName() string {
	return "pkg/sentry/fsimpl/proc.oomScoreAdj"
}

func (x *oomScoreAdj) StateFields() []string {
	return []string{
		"DynamicBytesFile",
		"task",
	}
}

func (x *oomScoreAdj) beforeSave() {}

func (x *oomScoreAdj) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.DynamicBytesFile)
	m.Save(1, &x.task)
}

func (x *oomScoreAdj) afterLoad() {}

func (x *oomScoreAdj) StateLoad(m state.Source) {
	m.Load(0, &x.DynamicBytesFile)
	m.Load(1, &x.task)
}

func (x *exeSymlink) StateTypeName() string {
	return "pkg/sentry/fsimpl/proc.exeSymlink"
}

func (x *exeSymlink) StateFields() []string {
	return []string{
		"InodeAttrs",
		"InodeNoopRefCount",
		"InodeSymlink",
		"task",
	}
}

func (x *exeSymlink) beforeSave() {}

func (x *exeSymlink) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.InodeAttrs)
	m.Save(1, &x.InodeNoopRefCount)
	m.Save(2, &x.InodeSymlink)
	m.Save(3, &x.task)
}

func (x *exeSymlink) afterLoad() {}

func (x *exeSymlink) StateLoad(m state.Source) {
	m.Load(0, &x.InodeAttrs)
	m.Load(1, &x.InodeNoopRefCount)
	m.Load(2, &x.InodeSymlink)
	m.Load(3, &x.task)
}

func (x *mountInfoData) StateTypeName() string {
	return "pkg/sentry/fsimpl/proc.mountInfoData"
}

func (x *mountInfoData) StateFields() []string {
	return []string{
		"DynamicBytesFile",
		"task",
	}
}

func (x *mountInfoData) beforeSave() {}

func (x *mountInfoData) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.DynamicBytesFile)
	m.Save(1, &x.task)
}

func (x *mountInfoData) afterLoad() {}

func (x *mountInfoData) StateLoad(m state.Source) {
	m.Load(0, &x.DynamicBytesFile)
	m.Load(1, &x.task)
}

func (x *mountsData) StateTypeName() string {
	return "pkg/sentry/fsimpl/proc.mountsData"
}

func (x *mountsData) StateFields() []string {
	return []string{
		"DynamicBytesFile",
		"task",
	}
}

func (x *mountsData) beforeSave() {}

func (x *mountsData) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.DynamicBytesFile)
	m.Save(1, &x.task)
}

func (x *mountsData) afterLoad() {}

func (x *mountsData) StateLoad(m state.Source) {
	m.Load(0, &x.DynamicBytesFile)
	m.Load(1, &x.task)
}

func (x *ifinet6) StateTypeName() string {
	return "pkg/sentry/fsimpl/proc.ifinet6"
}

func (x *ifinet6) StateFields() []string {
	return []string{
		"DynamicBytesFile",
		"stack",
	}
}

func (x *ifinet6) beforeSave() {}

func (x *ifinet6) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.DynamicBytesFile)
	m.Save(1, &x.stack)
}

func (x *ifinet6) afterLoad() {}

func (x *ifinet6) StateLoad(m state.Source) {
	m.Load(0, &x.DynamicBytesFile)
	m.Load(1, &x.stack)
}

func (x *netDevData) StateTypeName() string {
	return "pkg/sentry/fsimpl/proc.netDevData"
}

func (x *netDevData) StateFields() []string {
	return []string{
		"DynamicBytesFile",
		"stack",
	}
}

func (x *netDevData) beforeSave() {}

func (x *netDevData) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.DynamicBytesFile)
	m.Save(1, &x.stack)
}

func (x *netDevData) afterLoad() {}

func (x *netDevData) StateLoad(m state.Source) {
	m.Load(0, &x.DynamicBytesFile)
	m.Load(1, &x.stack)
}

func (x *netUnixData) StateTypeName() string {
	return "pkg/sentry/fsimpl/proc.netUnixData"
}

func (x *netUnixData) StateFields() []string {
	return []string{
		"DynamicBytesFile",
		"kernel",
	}
}

func (x *netUnixData) beforeSave() {}

func (x *netUnixData) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.DynamicBytesFile)
	m.Save(1, &x.kernel)
}

func (x *netUnixData) afterLoad() {}

func (x *netUnixData) StateLoad(m state.Source) {
	m.Load(0, &x.DynamicBytesFile)
	m.Load(1, &x.kernel)
}

func (x *netTCPData) StateTypeName() string {
	return "pkg/sentry/fsimpl/proc.netTCPData"
}

func (x *netTCPData) StateFields() []string {
	return []string{
		"DynamicBytesFile",
		"kernel",
	}
}

func (x *netTCPData) beforeSave() {}

func (x *netTCPData) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.DynamicBytesFile)
	m.Save(1, &x.kernel)
}

func (x *netTCPData) afterLoad() {}

func (x *netTCPData) StateLoad(m state.Source) {
	m.Load(0, &x.DynamicBytesFile)
	m.Load(1, &x.kernel)
}

func (x *netTCP6Data) StateTypeName() string {
	return "pkg/sentry/fsimpl/proc.netTCP6Data"
}

func (x *netTCP6Data) StateFields() []string {
	return []string{
		"DynamicBytesFile",
		"kernel",
	}
}

func (x *netTCP6Data) beforeSave() {}

func (x *netTCP6Data) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.DynamicBytesFile)
	m.Save(1, &x.kernel)
}

func (x *netTCP6Data) afterLoad() {}

func (x *netTCP6Data) StateLoad(m state.Source) {
	m.Load(0, &x.DynamicBytesFile)
	m.Load(1, &x.kernel)
}

func (x *netUDPData) StateTypeName() string {
	return "pkg/sentry/fsimpl/proc.netUDPData"
}

func (x *netUDPData) StateFields() []string {
	return []string{
		"DynamicBytesFile",
		"kernel",
	}
}

func (x *netUDPData) beforeSave() {}

func (x *netUDPData) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.DynamicBytesFile)
	m.Save(1, &x.kernel)
}

func (x *netUDPData) afterLoad() {}

func (x *netUDPData) StateLoad(m state.Source) {
	m.Load(0, &x.DynamicBytesFile)
	m.Load(1, &x.kernel)
}

func (x *netSnmpData) StateTypeName() string {
	return "pkg/sentry/fsimpl/proc.netSnmpData"
}

func (x *netSnmpData) StateFields() []string {
	return []string{
		"DynamicBytesFile",
		"stack",
	}
}

func (x *netSnmpData) beforeSave() {}

func (x *netSnmpData) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.DynamicBytesFile)
	m.Save(1, &x.stack)
}

func (x *netSnmpData) afterLoad() {}

func (x *netSnmpData) StateLoad(m state.Source) {
	m.Load(0, &x.DynamicBytesFile)
	m.Load(1, &x.stack)
}

func (x *netRouteData) StateTypeName() string {
	return "pkg/sentry/fsimpl/proc.netRouteData"
}

func (x *netRouteData) StateFields() []string {
	return []string{
		"DynamicBytesFile",
		"stack",
	}
}

func (x *netRouteData) beforeSave() {}

func (x *netRouteData) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.DynamicBytesFile)
	m.Save(1, &x.stack)
}

func (x *netRouteData) afterLoad() {}

func (x *netRouteData) StateLoad(m state.Source) {
	m.Load(0, &x.DynamicBytesFile)
	m.Load(1, &x.stack)
}

func (x *netStatData) StateTypeName() string {
	return "pkg/sentry/fsimpl/proc.netStatData"
}

func (x *netStatData) StateFields() []string {
	return []string{
		"DynamicBytesFile",
		"stack",
	}
}

func (x *netStatData) beforeSave() {}

func (x *netStatData) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.DynamicBytesFile)
	m.Save(1, &x.stack)
}

func (x *netStatData) afterLoad() {}

func (x *netStatData) StateLoad(m state.Source) {
	m.Load(0, &x.DynamicBytesFile)
	m.Load(1, &x.stack)
}

func (x *tasksInode) StateTypeName() string {
	return "pkg/sentry/fsimpl/proc.tasksInode"
}

func (x *tasksInode) StateFields() []string {
	return []string{
		"InodeNotSymlink",
		"InodeDirectoryNoNewChildren",
		"InodeAttrs",
		"OrderedChildren",
		"AlwaysValid",
		"locks",
		"fs",
		"pidns",
		"selfSymlink",
		"threadSelfSymlink",
		"cgroupControllers",
	}
}

func (x *tasksInode) beforeSave() {}

func (x *tasksInode) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.InodeNotSymlink)
	m.Save(1, &x.InodeDirectoryNoNewChildren)
	m.Save(2, &x.InodeAttrs)
	m.Save(3, &x.OrderedChildren)
	m.Save(4, &x.AlwaysValid)
	m.Save(5, &x.locks)
	m.Save(6, &x.fs)
	m.Save(7, &x.pidns)
	m.Save(8, &x.selfSymlink)
	m.Save(9, &x.threadSelfSymlink)
	m.Save(10, &x.cgroupControllers)
}

func (x *tasksInode) afterLoad() {}

func (x *tasksInode) StateLoad(m state.Source) {
	m.Load(0, &x.InodeNotSymlink)
	m.Load(1, &x.InodeDirectoryNoNewChildren)
	m.Load(2, &x.InodeAttrs)
	m.Load(3, &x.OrderedChildren)
	m.Load(4, &x.AlwaysValid)
	m.Load(5, &x.locks)
	m.Load(6, &x.fs)
	m.Load(7, &x.pidns)
	m.Load(8, &x.selfSymlink)
	m.Load(9, &x.threadSelfSymlink)
	m.Load(10, &x.cgroupControllers)
}

func (x *statData) StateTypeName() string {
	return "pkg/sentry/fsimpl/proc.statData"
}

func (x *statData) StateFields() []string {
	return []string{
		"dynamicBytesFileSetAttr",
	}
}

func (x *statData) beforeSave() {}

func (x *statData) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.dynamicBytesFileSetAttr)
}

func (x *statData) afterLoad() {}

func (x *statData) StateLoad(m state.Source) {
	m.Load(0, &x.dynamicBytesFileSetAttr)
}

func (x *loadavgData) StateTypeName() string {
	return "pkg/sentry/fsimpl/proc.loadavgData"
}

func (x *loadavgData) StateFields() []string {
	return []string{
		"dynamicBytesFileSetAttr",
	}
}

func (x *loadavgData) beforeSave() {}

func (x *loadavgData) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.dynamicBytesFileSetAttr)
}

func (x *loadavgData) afterLoad() {}

func (x *loadavgData) StateLoad(m state.Source) {
	m.Load(0, &x.dynamicBytesFileSetAttr)
}

func (x *meminfoData) StateTypeName() string {
	return "pkg/sentry/fsimpl/proc.meminfoData"
}

func (x *meminfoData) StateFields() []string {
	return []string{
		"dynamicBytesFileSetAttr",
	}
}

func (x *meminfoData) beforeSave() {}

func (x *meminfoData) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.dynamicBytesFileSetAttr)
}

func (x *meminfoData) afterLoad() {}

func (x *meminfoData) StateLoad(m state.Source) {
	m.Load(0, &x.dynamicBytesFileSetAttr)
}

func (x *uptimeData) StateTypeName() string {
	return "pkg/sentry/fsimpl/proc.uptimeData"
}

func (x *uptimeData) StateFields() []string {
	return []string{
		"dynamicBytesFileSetAttr",
	}
}

func (x *uptimeData) beforeSave() {}

func (x *uptimeData) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.dynamicBytesFileSetAttr)
}

func (x *uptimeData) afterLoad() {}

func (x *uptimeData) StateLoad(m state.Source) {
	m.Load(0, &x.dynamicBytesFileSetAttr)
}

func (x *versionData) StateTypeName() string {
	return "pkg/sentry/fsimpl/proc.versionData"
}

func (x *versionData) StateFields() []string {
	return []string{
		"dynamicBytesFileSetAttr",
	}
}

func (x *versionData) beforeSave() {}

func (x *versionData) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.dynamicBytesFileSetAttr)
}

func (x *versionData) afterLoad() {}

func (x *versionData) StateLoad(m state.Source) {
	m.Load(0, &x.dynamicBytesFileSetAttr)
}

func (x *filesystemsData) StateTypeName() string {
	return "pkg/sentry/fsimpl/proc.filesystemsData"
}

func (x *filesystemsData) StateFields() []string {
	return []string{
		"DynamicBytesFile",
	}
}

func (x *filesystemsData) beforeSave() {}

func (x *filesystemsData) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.DynamicBytesFile)
}

func (x *filesystemsData) afterLoad() {}

func (x *filesystemsData) StateLoad(m state.Source) {
	m.Load(0, &x.DynamicBytesFile)
}

func (x *mmapMinAddrData) StateTypeName() string {
	return "pkg/sentry/fsimpl/proc.mmapMinAddrData"
}

func (x *mmapMinAddrData) StateFields() []string {
	return []string{
		"DynamicBytesFile",
		"k",
	}
}

func (x *mmapMinAddrData) beforeSave() {}

func (x *mmapMinAddrData) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.DynamicBytesFile)
	m.Save(1, &x.k)
}

func (x *mmapMinAddrData) afterLoad() {}

func (x *mmapMinAddrData) StateLoad(m state.Source) {
	m.Load(0, &x.DynamicBytesFile)
	m.Load(1, &x.k)
}

func (x *hostnameData) StateTypeName() string {
	return "pkg/sentry/fsimpl/proc.hostnameData"
}

func (x *hostnameData) StateFields() []string {
	return []string{
		"DynamicBytesFile",
	}
}

func (x *hostnameData) beforeSave() {}

func (x *hostnameData) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.DynamicBytesFile)
}

func (x *hostnameData) afterLoad() {}

func (x *hostnameData) StateLoad(m state.Source) {
	m.Load(0, &x.DynamicBytesFile)
}

func (x *tcpSackData) StateTypeName() string {
	return "pkg/sentry/fsimpl/proc.tcpSackData"
}

func (x *tcpSackData) StateFields() []string {
	return []string{
		"DynamicBytesFile",
		"stack",
		"enabled",
	}
}

func (x *tcpSackData) beforeSave() {}

func (x *tcpSackData) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.DynamicBytesFile)
	m.Save(1, &x.stack)
	m.Save(2, &x.enabled)
}

func (x *tcpSackData) afterLoad() {}

func (x *tcpSackData) StateLoad(m state.Source) {
	m.Load(0, &x.DynamicBytesFile)
	m.LoadWait(1, &x.stack)
	m.Load(2, &x.enabled)
}

func (x *tcpRecoveryData) StateTypeName() string {
	return "pkg/sentry/fsimpl/proc.tcpRecoveryData"
}

func (x *tcpRecoveryData) StateFields() []string {
	return []string{
		"DynamicBytesFile",
		"stack",
	}
}

func (x *tcpRecoveryData) beforeSave() {}

func (x *tcpRecoveryData) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.DynamicBytesFile)
	m.Save(1, &x.stack)
}

func (x *tcpRecoveryData) afterLoad() {}

func (x *tcpRecoveryData) StateLoad(m state.Source) {
	m.Load(0, &x.DynamicBytesFile)
	m.LoadWait(1, &x.stack)
}

func (x *tcpMemData) StateTypeName() string {
	return "pkg/sentry/fsimpl/proc.tcpMemData"
}

func (x *tcpMemData) StateFields() []string {
	return []string{
		"DynamicBytesFile",
		"dir",
		"stack",
	}
}

func (x *tcpMemData) beforeSave() {}

func (x *tcpMemData) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.DynamicBytesFile)
	m.Save(1, &x.dir)
	m.Save(2, &x.stack)
}

func (x *tcpMemData) afterLoad() {}

func (x *tcpMemData) StateLoad(m state.Source) {
	m.Load(0, &x.DynamicBytesFile)
	m.Load(1, &x.dir)
	m.LoadWait(2, &x.stack)
}

func init() {
	state.Register((*FilesystemType)(nil))
	state.Register((*subtasksInode)(nil))
	state.Register((*taskInode)(nil))
	state.Register((*fdDirInode)(nil))
	state.Register((*fdSymlink)(nil))
	state.Register((*fdInfoDirInode)(nil))
	state.Register((*fdInfoData)(nil))
	state.Register((*auxvData)(nil))
	state.Register((*cmdlineData)(nil))
	state.Register((*commInode)(nil))
	state.Register((*commData)(nil))
	state.Register((*idMapData)(nil))
	state.Register((*mapsData)(nil))
	state.Register((*smapsData)(nil))
	state.Register((*taskStatData)(nil))
	state.Register((*statmData)(nil))
	state.Register((*statusData)(nil))
	state.Register((*ioData)(nil))
	state.Register((*oomScoreAdj)(nil))
	state.Register((*exeSymlink)(nil))
	state.Register((*mountInfoData)(nil))
	state.Register((*mountsData)(nil))
	state.Register((*ifinet6)(nil))
	state.Register((*netDevData)(nil))
	state.Register((*netUnixData)(nil))
	state.Register((*netTCPData)(nil))
	state.Register((*netTCP6Data)(nil))
	state.Register((*netUDPData)(nil))
	state.Register((*netSnmpData)(nil))
	state.Register((*netRouteData)(nil))
	state.Register((*netStatData)(nil))
	state.Register((*tasksInode)(nil))
	state.Register((*statData)(nil))
	state.Register((*loadavgData)(nil))
	state.Register((*meminfoData)(nil))
	state.Register((*uptimeData)(nil))
	state.Register((*versionData)(nil))
	state.Register((*filesystemsData)(nil))
	state.Register((*mmapMinAddrData)(nil))
	state.Register((*hostnameData)(nil))
	state.Register((*tcpSackData)(nil))
	state.Register((*tcpRecoveryData)(nil))
	state.Register((*tcpMemData)(nil))
}
