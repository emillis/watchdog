package watchdog

import "io/fs"

var defaultWatchDog = WatchDog{
	Root:                 make([]string, 0, 5),
	Depth:                0,
	IgnoreStartupContent: false,
	Handler:              nil,
}

//===========[STRUCTS]====================================================================================================

type WatchDog struct {
	//Root folders to start monitoring
	Root []string

	//Depth allows you to limit depth of file monitoring in nested folders
	Depth int

	//IgnoreStartupContent - if this is set to true, handler function will not fire during the initial indexing
	//of the root folder
	IgnoreStartupContent bool

	//Handler will be invoked each time a change is detected withing the root folder
	Handler func(file fs.File)
}

func (w *WatchDog) copy() WatchDog {
	nw := WatchDog{
		Root:                 make([]string, len(w.Root), cap(w.Root)),
		Depth:                w.Depth,
		IgnoreStartupContent: w.IgnoreStartupContent,
		Handler:              w.Handler,
	}

	nw.Root = append(nw.Root, w.Root...)

	return nw
}

//===========[FUNCTIONALITY]====================================================================================================

//NewWatchDog returns newly initiated WatchDog
func NewWatchDog() WatchDog {
	return defaultWatchDog.copy()
}
