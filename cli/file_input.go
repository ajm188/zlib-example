package cli

import (
	"io"
	"os"

	"github.com/spf13/pflag"
)

func FileInput(flags *pflag.FlagSet, n int) (io.ReadCloser, error) {
	if flags.NArg() <= n {
		return os.Stdin, nil
	}

	return os.Open(flags.Arg(n))
}
