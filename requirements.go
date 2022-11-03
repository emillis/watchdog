package watchdog

import "io/fs"

type Requirements struct {
	Root                 []string
	Depth                int
	IgnoreStartupContent bool
	Handler              func(fs.File)
}
