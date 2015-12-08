package pmstructs

import (
	"flag"
)

var (
	debugMode *bool
)

func main() {
	debugMode = flag.Bool("debug", false, "Set application to debug mode")
	flag.Parse()
}
