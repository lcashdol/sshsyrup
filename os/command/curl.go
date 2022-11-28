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
	"crypto/md5"
        "encoding/hex"
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
	follow := flag.Bool("L", true, "Follow redirects")
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
	if !*follow {
		// capture the redirect header and use it instead of previous url
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
	//	*out = "index.html"
		*out = GetMD5Hash(string(b))
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

func GetMD5Hash(text string) string {
   hash := md5.Sum([]byte(text))
   return hex.EncodeToString(hash[:])
}

func (wg curl) Where() string {
	return "/usr/bin/curl"
}
