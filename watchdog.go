package watchdog

import "io/fs"

var defaultWatchDog = watchDog{
	Root:                 make([]string, 0, 5),
	Depth:                0,
	IgnoreStartupContent: false,
	Handler:              nil,
}

//===========[STRUCTS]====================================================================================================

type watchDog struct {
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

func (w *watchDog) copy() watchDog {
	nw := watchDog{
		Root:                 make([]string, len(w.Root), cap(w.Root)),
		Depth:                w.Depth,
		IgnoreStartupContent: w.IgnoreStartupContent,
		Handler:              w.Handler,
	}

	nw.Root = append(nw.Root, w.Root...)

	return nw
}

//===========[FUNCTIONALITY]====================================================================================================

//NewWatchDog returns newly initiated watchDog
func NewWatchDog() watchDog {
	return defaultWatchDog.copy()
}
