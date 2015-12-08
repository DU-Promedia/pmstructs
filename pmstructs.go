package pmstructs

import (
	"flag"
)

var (
	debugMode *bool
)

func main() {
	debugMode = flag.Bool("debug", false, "Set grabberd to debug mode")
	flag.Parse()
}
