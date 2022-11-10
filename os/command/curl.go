package command

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"path"
	"strings"

	"github.com/mkishere/sshsyrup/os"
	"github.com/spf13/afero"
	"github.com/spf13/pflag"
)

type curl struct{}

func init() {
	os.RegisterCommand("curl", curl{})

}
func (wg curl) GetHelp() string {
	return ""
}

func (wg curl) Exec(args []string, sys os.Sys) int {
	flag := pflag.NewFlagSet("arg", pflag.ContinueOnError)
	out := flag.String("o", "", "Write documents to FILE.")
	flag.SetOutput(sys.Out())
	err := flag.Parse(args)
	f := flag.Args()
	if len(args) == 0 {
		fmt.Fprintln(sys.Out(), "curl: missing URL\nUsage: curl [OPTION]... [URL]...\n\nTry `curl --help' for more options.")
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
/*	if !*quiet {
		if urlobj.Scheme != "http" && urlobj.Scheme != "https" {
			fmt.Fprintf(sys.Out(), "Resolving %v (%v)... failed: Name or service not known.\n", urlobj.Scheme, urlobj.Scheme)
			fmt.Fprintf(sys.Out(), "curl: unable to resolve host address ‘%v’\n", urlobj.Scheme)
			return 1
		}
		fmt.Fprintf(sys.Out(), "--%v--  %v\n", printTs(), url)
	}*/
	/*ip, err := net.LookupIP(urlobj.Hostname())
	if err != nil {
		// handle error
		fmt.Fprintln(sys.Err(), err)
	}*/
/*	if !*quiet {
		fmt.Fprintf(sys.Out(), "Resolving %v (%v)... %v\n", urlobj.Hostname(), urlobj.Hostname(), ip)
	}*/
	resp, err := http.Get(url)
	if err != nil {
		// handle error
		fmt.Fprintln(sys.Err(), err)
		return 1
	}
/*	if !*quiet {
		fmt.Fprintf(sys.Out(), "Connecting to %v (%v)|%v|:80... connected\n", urlobj.Hostname(), urlobj.Hostname(), ip[0])
		mimeType := resp.Header.Get("Content-Type")
		fmt.Fprintln(sys.Out(), "HTTP request sent, awaiting response... 200 OK")
		fmt.Fprintf(sys.Out(), "Length: unspecified [%v]\n", mimeType[:strings.LastIndex(mimeType, ";")])
	}*/
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		//handle
		fmt.Fprintln(sys.Err(), err)
		return 1
	}
	if *out == "" {
		*out = "index.html"
	}
/*	if !*quiet {
		fmt.Fprintf(sys.Out(), "Saving to: ‘%v’\n\n", *out)
		fmt.Fprintf(sys.Out(), "[ <=>%v ] %v       --.-K/s   in 0.1s\n", strings.Repeat(" ", sys.Width()-38), format(len(b)))
	}*/
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
		fmt.Fprintf(sys.Out(), "%s", b)
	return 0
}

func (wg curl) Where() string {
	return "/usr/bin/curl"
}
