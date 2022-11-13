package command

import (
	"fmt"
	"path"
	"github.com/mkishere/sshsyrup/os"
	"github.com/spf13/afero"
)

type rm struct{}

func init() {
	os.RegisterCommand("rm", rm{})
}

func (i rm) GetHelp() string {
	return ""
}

func (i rm) Exec(args []string, sys os.Sys) int {
	// Implement rm command, we need to handle common args like -rf here eventually
	af := afero.Afero{sys.FSys()}

        filename := args[0]

	       if !path.IsAbs(filename) {
                filename = path.Join(sys.Getcwd(), filename)
        }

	err := af.Remove(filename)
	if (err != nil) {
	fmt.Fprintf(sys.Out(), "rm: %s: No such file or directory\n",filename)
}
	return 0
}

func (i rm) Where() string {
	return "/bin/rm"
}
