package watchdog

import (
	"io/fs"
	"sync"
)

var defaultWatchDog = watchDog{
	root:                 make([]string, 0, 5),
	Depth:                0,
	IgnoreStartupContent: false,
	Handler:              nil,
}

//===========[STRUCTS]====================================================================================================

type watchDog struct {
	//Root folders to start monitoring
	root []string

	//Depth allows you to limit depth of file monitoring in nested folders
	Depth int

	//IgnoreStartupContent - if this is set to true, handler function will not fire during the initial indexing
	//of the root folder
	IgnoreStartupContent bool

	//Handler will be invoked each time a change is detected withing the root folder
	Handler func(file fs.File)

	mx sync.RWMutex
}

func (w *watchDog) copy() watchDog {
	w.mx.RLock()
	defer w.mx.RUnlock()

	nw := watchDog{
		root:                 make([]string, len(w.root), cap(w.root)),
		Depth:                w.Depth,
		IgnoreStartupContent: w.IgnoreStartupContent,
		Handler:              w.Handler,
	}

	nw.root = append(nw.root, w.root...)

	return nw
}

func (w *watchDog) Root() []string {
	w.mx.RLock()
	defer w.mx.RUnlock()

	r := make([]string, 0, len(w.root))

	r = append(r, w.root...)

	return r
}

func (w *watchDog) SetRoot(root []string) {
	w.mx.Lock()
	defer w.mx.Unlock()

	w.root = make([]string, 0, len(root))

	w.root = append(w.root, root...)
}

//===========[FUNCTIONALITY]====================================================================================================

//NewWatchDog returns newly initiated watchDog
func NewWatchDog() watchDog {
	return defaultWatchDog.copy()
}
