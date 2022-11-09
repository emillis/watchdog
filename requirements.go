package watchdog

import "io/fs"

//===========[STATIC]====================================================================================================

//OperatingMode allows to choose different modes of operation of the WatchDog
type OperatingMode string

const (
	//Sequential mode, Handler will be caller for each new file
	Sequential OperatingMode = "sequential"
	//Burst mode, Handler will be called every ScanFrequency amount of time with the files being accumulated
	//until that time
	Burst OperatingMode = "burst"
)

//===========[STRUCTS]====================================================================================================

type Requirements struct {
	//A list of root folders where to scan for files
	Root []string

	//How deep down the root folders should the scan take place
	Depth uint32

	//Whether to ignore the files that are present in the root directory during the initial scan
	IgnoreStartupContent bool

	//How often to scan the root folders (ms)
	ScanFrequency uint32

	//OperatingMode allows to choose different modes of operation of the WatchDog
	OperatingMode OperatingMode

	//This function will get invoked for each file detected
	Handler func(fs.File)
}
