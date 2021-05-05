package main

import (
	"compress/zlib"
	"fmt"
	"io"
	"strings"

	"ajm188.scratchpad/zlib-example/cli"
	"github.com/spf13/cobra"
)

var (
	writeCmd = &cobra.Command{
		Use:          "write [source]",
		RunE:         write,
		SilenceUsage: true,
	}
)

func write(cmd *cobra.Command, args []string) (err error) {
	r, err := cli.FileInput(cmd.Flags(), 0)
	if err != nil {
		return err
	}
	defer r.Close()

	buf := &strings.Builder{}
	compressor := zlib.NewWriter(buf)

	_, err = io.Copy(compressor, r)
	if err != nil {
		return fmt.Errorf("error copying file: %w", err)
	}

	compressor.Close()
	fmt.Println(buf.String())

	return nil
}
