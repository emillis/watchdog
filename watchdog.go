package watchdog

import (
	"io/fs"
	"sync"
)

//===========[STATIC]====================================================================================================

//These are the default values for the WatchDog
var defaultWatchDog = WatchDog{
	root:                 make([]string, 0, 5),
	depth:                0,
	ignoreStartupContent: false,
	handler:              nil,
}

//===========[STRUCTS]====================================================================================================

type WatchDog struct {
	//Root folders to start monitoring
	root []string

	//Depth allows you to limit depth of file monitoring in nested folders
	depth int

	//ignoreStartupContent - if this is set to true, handler function will not fire during the initial indexing
	//of the root folder
	ignoreStartupContent bool

	//handler will be invoked each time a change is detected withing the root folder
	handler func(file fs.File)

	mx sync.RWMutex
}

//copy makes a perfect copy of the WatchDog
func (w *WatchDog) copy() WatchDog {
	w.mx.RLock()
	defer w.mx.RUnlock()

	nw := WatchDog{
		root:                 make([]string, len(w.root), cap(w.root)),
		depth:                w.depth,
		ignoreStartupContent: w.ignoreStartupContent,
		handler:              w.handler,
		mx:                   sync.RWMutex{},
	}

	nw.root = append(nw.root, w.root...)

	return nw
}

//Root returns the slice of root folders that the watchdog is monitoring
func (w *WatchDog) Root() []string {
	w.mx.RLock()
	defer w.mx.RUnlock()

	r := make([]string, 0, len(w.root))

	r = append(r, w.root...)

	return r
}

//SetRoot sets a slice of root locations that the WatchDog is going to be monitoring
func (w *WatchDog) SetRoot(root []string) {
	w.mx.Lock()
	defer w.mx.Unlock()

	w.root = make([]string, 0, len(root))

	w.root = append(w.root, root...)
}

//Depth returns the limit of folder depth starting from the root folder
func (w *WatchDog) Depth() int {
	w.mx.RLock()
	defer w.mx.RUnlock()

	return w.depth
}

//SetDepth sets a new limit of scanning depth starting with the root folder
func (w *WatchDog) SetDepth(val int) {
	w.mx.Lock()
	defer w.mx.Unlock()

	w.depth = val
}

//IgnoreStartupContent returns boolean which indicates whether the contents that are present within the root folder
//during the initial scan of the locations should invoke the handler function
func (w *WatchDog) IgnoreStartupContent() bool {
	w.mx.RLock()
	defer w.mx.RUnlock()

	return w.ignoreStartupContent
}

//SetIgnoreStartupContent decides whether the files present in the root folder structure should invoke the handler
//function during the initial scan of the locations
func (w *WatchDog) SetIgnoreStartupContent(val bool) {
	w.mx.Lock()
	defer w.mx.Unlock()

	w.ignoreStartupContent = val
}

//SetHandler sets new handler function
func (w *WatchDog) SetHandler(h func(fs.File)) {
	w.mx.Lock()
	defer w.mx.Unlock()

	w.handler = h
}

//Start begins watching root directories
func (w *WatchDog) Start() {

}

//Stop stops watching the root directories
func (w *WatchDog) Stop() {

}

//===========[FUNCTIONALITY]====================================================================================================

//NewWatchDog returns newly initiated WatchDog
func NewWatchDog() WatchDog {
	return defaultWatchDog.copy()
}
