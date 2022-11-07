package watchdog

import "io/fs"

type Requirements struct {
	Root                 []string
	Depth                int
	IgnoreStartupContent bool

	//How often to scan the root folders (ms)
	ScanFrequency uint32

	//This function will get invoked for each file detected
	Handler func(fs.File)
}
