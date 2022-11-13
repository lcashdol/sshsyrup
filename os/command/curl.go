package command

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"path"
	"strings"
	"flag"
	"github.com/mkishere/sshsyrup/os"
	"github.com/spf13/afero"
)

type curl struct{}

func init() {
	os.RegisterCommand("curl", curl{})

}
func (wg curl) GetHelp() string {
	return ""
}

func (wg curl) Exec(args []string, sys os.Sys) int {
	flag := flag.NewFlagSet("curl", flag.ContinueOnError)
	out := flag.String("o", "", "curl: option -o: requires parameter")
	quiet := flag.Bool("s", true, "silence output")
	flag.SetOutput(sys.Out())
	err := flag.Parse(args)
	f := flag.Args()
	if len(args) == 0 {
		fmt.Fprintln(sys.Out(), "curl: try 'curl --help' or 'curl --manual' for more information")
		return 1
	}
	url := strings.TrimSpace(f[0])
	if !strings.Contains(url, "://") {
		url = "http://" + url
	}
	if err != nil {
		fmt.Fprintln(sys.Out(), "Malformed URL")
		return 1
	}
	if !*quiet {
		// do nothing just handle the -s
	}
	resp, err := http.Get(url)
	if err != nil {
		// handle error
		fmt.Fprintln(sys.Err(), err)
		return 1
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		//handle
		fmt.Fprintln(sys.Err(), err)
		return 1
	}

	if *out == "" {
       fmt.Fprintf(sys.Out(), "%s", b) // pretend to redirect output to stdout while copying the file to disk
		*out = "index.html"
	}
	af := afero.Afero{sys.FSys()}

	p := *out
	if !path.IsAbs(p) {
		p = path.Join(sys.Getcwd(), p)
	}
	err = af.WriteFile(p, b, 0666)
	if err != nil {
		fmt.Fprintln(sys.Err(), err)
		return 1
	}

	return 0
}

func (wg curl) Where() string {
	return "/usr/bin/curl"
}
