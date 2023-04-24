// mumax3 main command
package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path"
	"time"

	"github.com/MathieuMoalic/amumax/cuda"
	"github.com/MathieuMoalic/amumax/engine"
	"github.com/MathieuMoalic/amumax/script"
	"github.com/MathieuMoalic/amumax/util"
)

var (
	flag_failfast = flag.Bool("failfast", false, "If one simulation fails, stop entire batch immediately")
	flag_test     = flag.Bool("test", false, "Cuda test (internal)")
	flag_version  = flag.Bool("v", true, "Print version")
	flag_vet      = flag.Bool("vet", false, "Check input files for errors, but don't run them")
	// more flags in engine/gofiles.go
)

func main() {
	go checkUpdate()
	flag.Parse()
	log.SetPrefix("")
	log.SetFlags(0)

	cuda.Init(*engine.Flag_gpu)

	cuda.Synchronous = *engine.Flag_sync
	if *flag_version {
		printVersion()
	}

	// used by bootstrap launcher to test cuda
	// successful exit means cuda was initialized fine
	if *flag_test {
		fmt.Println(cuda.GPUInfo)
		os.Exit(0)
	}

	defer engine.Close() // flushes pending output, if any

	if *flag_vet {
		vet()
		return
	}
	// go api.Start()
	switch flag.NArg() {
	case 0:
		if *engine.Flag_interactive {
			runInteractive()
		}
	case 1:
		runFileAndServe(flag.Arg(0))
	default:
		RunQueue(flag.Args())
	}
}

type Release struct {
	TagName string `json:"tag_name"`
}

func checkUpdate() {
	resp, err := http.Get("https://api.github.com/repos/mathieumoalic/amumax/releases/latest")
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	var release Release
	err = json.Unmarshal(body, &release)
	if err != nil {
		return
	}
	if release.TagName != engine.VERSION {
		exePath, err := os.Executable()
		if err != nil {
			return
		}
		fmt.Println("New amumax version available, run the following command to update amumax:")
		fmt.Println()
		fmt.Println("curl -L https://github.com/mathieumoalic/amumax/releases/latest/download/amumax >", exePath)
		fmt.Println()
	}
}

func runInteractive() {
	fmt.Println("//no input files: starting interactive session")
	// setup outut dir
	now := time.Now()
	outdir := fmt.Sprintf("./mumax-%v-%02d-%02d_%02dh%02d.out", now.Year(), int(now.Month()), now.Day(), now.Hour(), now.Minute())
	engine.InitIO(outdir, outdir, *engine.Flag_forceclean)

	engine.Timeout = 365 * 24 * time.Hour // basically forever

	// set up some sensible start configuration
	engine.Eval(`SetGridSize(128, 64, 1)
		SetCellSize(4e-9, 4e-9, 4e-9)
		Msat = 1e6
		Aex = 10e-12
		alpha = 1
		m = RandomMag()`)
	addr := goServeGUI()
	openbrowser("http://127.0.0.1" + addr)
	engine.RunInteractive()
}

func runFileAndServe(fname string) {
	if path.Ext(fname) == ".go" {
		runGoFile(fname)
	} else {
		runScript(fname)
	}
}

func runScript(fname string) {
	if _, err := os.Stat(fname); errors.Is(err, os.ErrNotExist) {
		util.Fatal("Error: File `", fname, "` does not exist")
	}
	outDir := util.NoExt(fname) + ".zarr"
	if *engine.Flag_od != "" {
		outDir = *engine.Flag_od
	}
	engine.InitIO(fname, outDir, *engine.Flag_forceclean)

	fname = engine.InputFile

	var code *script.BlockStmt
	var err2 error
	if fname != "" {
		// first we compile the entire file into an executable tree
		code, err2 = engine.CompileFile(fname)
		util.FatalErr(err2)
	}

	// now the parser is not used anymore so it can handle web requests
	goServeGUI()

	if *engine.Flag_interactive {
		openbrowser("http://127.0.0.1" + *engine.Flag_webui_addr)
	}

	// start executing the tree, possibly injecting commands from web gui
	engine.EvalFile(code)

	if *engine.Flag_interactive {
		engine.RunInteractive()
	}
}

func runGoFile(fname string) {

	// pass through flags
	flags := []string{"run", fname}
	flag.Visit(func(f *flag.Flag) {
		if f.Name != "o" {
			flags = append(flags, fmt.Sprintf("-%v=%v", f.Name, f.Value))
		}
	})

	if *engine.Flag_od != "" {
		flags = append(flags, fmt.Sprintf("-o=%v", *engine.Flag_od))
	}

	cmd := exec.Command("go", flags...)
	log.Println("go", flags)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		os.Exit(1)
	}
}

// start Gui server and return server address
func goServeGUI() string {
	if *engine.Flag_webui_addr == "" {
		log.Println(`//not starting GUI (-http="")`)
		return ""
	}
	addr := engine.GoServe(*engine.Flag_webui_addr)
	return addr
}

// print version to stdout
func printVersion() {
	engine.LogOut(engine.UNAME)
	engine.LogOut(fmt.Sprintf("GPU info: %s, using cc=%d PTX", cuda.GPUInfo, cuda.UseCC))
}
