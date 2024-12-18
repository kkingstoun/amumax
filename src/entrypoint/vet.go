package entrypoint

import (
	"flag"
	"fmt"
	"os"

	"github.com/MathieuMoalic/amumax/src/engine_old"
	"github.com/MathieuMoalic/amumax/src/log_old"
)

// check all input files for errors, don't run.
func vet() {
	status := 0
	for _, f := range flag.Args() {
		src, ioerr := os.ReadFile(f)
		log_old.Log.PanicIfError(ioerr)
		engine_old.World.EnterScope() // avoid name collisions between separate files
		_, err := engine_old.World.Compile(string(src))
		engine_old.World.ExitScope()
		if err != nil {
			fmt.Println(f, ":", err)
			status = 1
		} else {
			fmt.Println(f, ":", "OK")
		}
	}
	os.Exit(status)
}
