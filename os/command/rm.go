package command

import (
	"fmt"
	"github.com/mkishere/sshsyrup/os"
)

type rm struct{}

func init() {
	os.RegisterCommand("rm", rm{})
}

func (i rm) GetHelp() string {
	return ""
}

func (i rm) Exec(args []string, sys os.Sys) int {
        filename := args[0]
	fmt.Fprintf(sys.Out(), "rm: %s: No such file or directory\n",filename)
	return 0
}

func (i rm) Where() string {
	return "/bin/uname"
}
