// automatically generated by stateify.

package fsutil

import (
	"gvisor.dev/gvisor/pkg/state"
)

func (s *DirtySet) StateTypeName() string {
	return "pkg/sentry/fs/fsutil.DirtySet"
}

func (s *DirtySet) StateFields() []string {
	return []string{
		"root",
	}
}

func (s *DirtySet) beforeSave() {}

// +checklocksignore
func (s *DirtySet) StateSave(stateSinkObject state.Sink) {
	s.beforeSave()
	var rootValue *DirtySegmentDataSlices = s.saveRoot()
	stateSinkObject.SaveValue(0, rootValue)
}

func (s *DirtySet) afterLoad() {}

// +checklocksignore
func (s *DirtySet) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.LoadValue(0, new(*DirtySegmentDataSlices), func(y interface{}) { s.loadRoot(y.(*DirtySegmentDataSlices)) })
}

func (n *Dirtynode) StateTypeName() string {
	return "pkg/sentry/fs/fsutil.Dirtynode"
}

func (n *Dirtynode) StateFields() []string {
	return []string{
		"nrSegments",
		"parent",
		"parentIndex",
		"hasChildren",
		"maxGap",
		"keys",
		"values",
		"children",
	}
}

func (n *Dirtynode) beforeSave() {}

// +checklocksignore
func (n *Dirtynode) StateSave(stateSinkObject state.Sink) {
	n.beforeSave()
	stateSinkObject.Save(0, &n.nrSegments)
	stateSinkObject.Save(1, &n.parent)
	stateSinkObject.Save(2, &n.parentIndex)
	stateSinkObject.Save(3, &n.hasChildren)
	stateSinkObject.Save(4, &n.maxGap)
	stateSinkObject.Save(5, &n.keys)
	stateSinkObject.Save(6, &n.values)
	stateSinkObject.Save(7, &n.children)
}

func (n *Dirtynode) afterLoad() {}

// +checklocksignore
func (n *Dirtynode) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &n.nrSegments)
	stateSourceObject.Load(1, &n.parent)
	stateSourceObject.Load(2, &n.parentIndex)
	stateSourceObject.Load(3, &n.hasChildren)
	stateSourceObject.Load(4, &n.maxGap)
	stateSourceObject.Load(5, &n.keys)
	stateSourceObject.Load(6, &n.values)
	stateSourceObject.Load(7, &n.children)
}

func (d *DirtySegmentDataSlices) StateTypeName() string {
	return "pkg/sentry/fs/fsutil.DirtySegmentDataSlices"
}

func (d *DirtySegmentDataSlices) StateFields() []string {
	return []string{
		"Start",
		"End",
		"Values",
	}
}

func (d *DirtySegmentDataSlices) beforeSave() {}

// +checklocksignore
func (d *DirtySegmentDataSlices) StateSave(stateSinkObject state.Sink) {
	d.beforeSave()
	stateSinkObject.Save(0, &d.Start)
	stateSinkObject.Save(1, &d.End)
	stateSinkObject.Save(2, &d.Values)
}

func (d *DirtySegmentDataSlices) afterLoad() {}

// +checklocksignore
func (d *DirtySegmentDataSlices) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &d.Start)
	stateSourceObject.Load(1, &d.End)
	stateSourceObject.Load(2, &d.Values)
}

func (s *FileRangeSet) StateTypeName() string {
	return "pkg/sentry/fs/fsutil.FileRangeSet"
}

func (s *FileRangeSet) StateFields() []string {
	return []string{
		"root",
	}
}

func (s *FileRangeSet) beforeSave() {}

// +checklocksignore
func (s *FileRangeSet) StateSave(stateSinkObject state.Sink) {
	s.beforeSave()
	var rootValue *FileRangeSegmentDataSlices = s.saveRoot()
	stateSinkObject.SaveValue(0, rootValue)
}

func (s *FileRangeSet) afterLoad() {}

// +checklocksignore
func (s *FileRangeSet) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.LoadValue(0, new(*FileRangeSegmentDataSlices), func(y interface{}) { s.loadRoot(y.(*FileRangeSegmentDataSlices)) })
}

func (n *FileRangenode) StateTypeName() string {
	return "pkg/sentry/fs/fsutil.FileRangenode"
}

func (n *FileRangenode) StateFields() []string {
	return []string{
		"nrSegments",
		"parent",
		"parentIndex",
		"hasChildren",
		"maxGap",
		"keys",
		"values",
		"children",
	}
}

func (n *FileRangenode) beforeSave() {}

// +checklocksignore
func (n *FileRangenode) StateSave(stateSinkObject state.Sink) {
	n.beforeSave()
	stateSinkObject.Save(0, &n.nrSegments)
	stateSinkObject.Save(1, &n.parent)
	stateSinkObject.Save(2, &n.parentIndex)
	stateSinkObject.Save(3, &n.hasChildren)
	stateSinkObject.Save(4, &n.maxGap)
	stateSinkObject.Save(5, &n.keys)
	stateSinkObject.Save(6, &n.values)
	stateSinkObject.Save(7, &n.children)
}

func (n *FileRangenode) afterLoad() {}

// +checklocksignore
func (n *FileRangenode) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &n.nrSegments)
	stateSourceObject.Load(1, &n.parent)
	stateSourceObject.Load(2, &n.parentIndex)
	stateSourceObject.Load(3, &n.hasChildren)
	stateSourceObject.Load(4, &n.maxGap)
	stateSourceObject.Load(5, &n.keys)
	stateSourceObject.Load(6, &n.values)
	stateSourceObject.Load(7, &n.children)
}

func (f *FileRangeSegmentDataSlices) StateTypeName() string {
	return "pkg/sentry/fs/fsutil.FileRangeSegmentDataSlices"
}

func (f *FileRangeSegmentDataSlices) StateFields() []string {
	return []string{
		"Start",
		"End",
		"Values",
	}
}

func (f *FileRangeSegmentDataSlices) beforeSave() {}

// +checklocksignore
func (f *FileRangeSegmentDataSlices) StateSave(stateSinkObject state.Sink) {
	f.beforeSave()
	stateSinkObject.Save(0, &f.Start)
	stateSinkObject.Save(1, &f.End)
	stateSinkObject.Save(2, &f.Values)
}

func (f *FileRangeSegmentDataSlices) afterLoad() {}

// +checklocksignore
func (f *FileRangeSegmentDataSlices) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &f.Start)
	stateSourceObject.Load(1, &f.End)
	stateSourceObject.Load(2, &f.Values)
}

func (s *FrameRefSet) StateTypeName() string {
	return "pkg/sentry/fs/fsutil.FrameRefSet"
}

func (s *FrameRefSet) StateFields() []string {
	return []string{
		"root",
	}
}

func (s *FrameRefSet) beforeSave() {}

// +checklocksignore
func (s *FrameRefSet) StateSave(stateSinkObject state.Sink) {
	s.beforeSave()
	var rootValue *FrameRefSegmentDataSlices = s.saveRoot()
	stateSinkObject.SaveValue(0, rootValue)
}

func (s *FrameRefSet) afterLoad() {}

// +checklocksignore
func (s *FrameRefSet) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.LoadValue(0, new(*FrameRefSegmentDataSlices), func(y interface{}) { s.loadRoot(y.(*FrameRefSegmentDataSlices)) })
}

func (n *FrameRefnode) StateTypeName() string {
	return "pkg/sentry/fs/fsutil.FrameRefnode"
}

func (n *FrameRefnode) StateFields() []string {
	return []string{
		"nrSegments",
		"parent",
		"parentIndex",
		"hasChildren",
		"maxGap",
		"keys",
		"values",
		"children",
	}
}

func (n *FrameRefnode) beforeSave() {}

// +checklocksignore
func (n *FrameRefnode) StateSave(stateSinkObject state.Sink) {
	n.beforeSave()
	stateSinkObject.Save(0, &n.nrSegments)
	stateSinkObject.Save(1, &n.parent)
	stateSinkObject.Save(2, &n.parentIndex)
	stateSinkObject.Save(3, &n.hasChildren)
	stateSinkObject.Save(4, &n.maxGap)
	stateSinkObject.Save(5, &n.keys)
	stateSinkObject.Save(6, &n.values)
	stateSinkObject.Save(7, &n.children)
}

func (n *FrameRefnode) afterLoad() {}

// +checklocksignore
func (n *FrameRefnode) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &n.nrSegments)
	stateSourceObject.Load(1, &n.parent)
	stateSourceObject.Load(2, &n.parentIndex)
	stateSourceObject.Load(3, &n.hasChildren)
	stateSourceObject.Load(4, &n.maxGap)
	stateSourceObject.Load(5, &n.keys)
	stateSourceObject.Load(6, &n.values)
	stateSourceObject.Load(7, &n.children)
}

func (f *FrameRefSegmentDataSlices) StateTypeName() string {
	return "pkg/sentry/fs/fsutil.FrameRefSegmentDataSlices"
}

func (f *FrameRefSegmentDataSlices) StateFields() []string {
	return []string{
		"Start",
		"End",
		"Values",
	}
}

func (f *FrameRefSegmentDataSlices) beforeSave() {}

// +checklocksignore
func (f *FrameRefSegmentDataSlices) StateSave(stateSinkObject state.Sink) {
	f.beforeSave()
	stateSinkObject.Save(0, &f.Start)
	stateSinkObject.Save(1, &f.End)
	stateSinkObject.Save(2, &f.Values)
}

func (f *FrameRefSegmentDataSlices) afterLoad() {}

// +checklocksignore
func (f *FrameRefSegmentDataSlices) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &f.Start)
	stateSourceObject.Load(1, &f.End)
	stateSourceObject.Load(2, &f.Values)
}

func init() {
	state.Register((*DirtySet)(nil))
	state.Register((*Dirtynode)(nil))
	state.Register((*DirtySegmentDataSlices)(nil))
	state.Register((*FileRangeSet)(nil))
	state.Register((*FileRangenode)(nil))
	state.Register((*FileRangeSegmentDataSlices)(nil))
	state.Register((*FrameRefSet)(nil))
	state.Register((*FrameRefnode)(nil))
	state.Register((*FrameRefSegmentDataSlices)(nil))
}
