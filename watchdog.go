package watchdog

import (
	"io/fs"
	"path/filepath"
	"sync"
	"time"
)

//===========[STATIC]====================================================================================================

//These are the default values for the WatchDog
var defaultRequirements = &Requirements{
	Root:                 make([]string, 0, 5),
	Depth:                0,
	IgnoreStartupContent: false,
	ScanFrequency:        3000,
	OperatingMode:        Burst,
	Handler:              nil,
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

	//How often to scan the root folders (ms)
	scanFrequency uint32

	//handler will be invoked each time a change is detected withing the root folder
	handler func(file fs.File)

	stopChan chan struct{}

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
func (w *WatchDog) SetDepth(val uint32) {
	w.mx.Lock()
	defer w.mx.Unlock()

	w.depth = int(val)
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
func NewWatchDog(req *Requirements) (*WatchDog, error) {
	if req == nil {
		req = defaultRequirements
	}

	if err := checkRequirements(req); err != nil {
		return nil, err
	}

	wd := &WatchDog{
		root:                 make([]string, cap(req.Root), len(req.Root)),
		depth:                int(req.Depth),
		ignoreStartupContent: req.IgnoreStartupContent,
		scanFrequency:        req.ScanFrequency,
		handler:              req.Handler,
		stopChan:             make(chan struct{}),
		mx:                   sync.RWMutex{},
	}

	wd.root = append(wd.root, req.Root...)

	return wd, nil
}

//
func checkRequirements(req *Requirements) error {
	return nil
}

//
func newWatcherRoutine(results chan fs.File, wd *WatchDog) {
	go func() {
		root := wd.Root()
		delay := time.Duration(wd.scanFrequency) //TODO: change this to method

		for {
			switch {
			case <-wd.stopChan:

			}

			for i := 0; i < len(root); i++ {
				filepath.WalkDir(root[i], func(path string, d fs.DirEntry, err error) error {
					if d.IsDir() {
						return nil
					}

					return nil
				})
			}

			time.Sleep(time.Millisecond * delay)
		}
	}()
}
