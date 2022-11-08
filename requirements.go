package watchdog

import "io/fs"

//===========[STATIC]====================================================================================================

type OperatingMode string

const (
	Sequential OperatingMode = "sequential"
	Burst      OperatingMode = "burst"
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

	OperatingMode OperatingMode

	//This function will get invoked for each file detected
	Handler func(fs.File)
}
