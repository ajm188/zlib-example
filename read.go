package main

import (
	"compress/zlib"
	"fmt"
	"io"
	"os"

	"ajm188.scratchpad/zlib-example/cli"
	"github.com/spf13/cobra"
)

var (
	readCmd = &cobra.Command{
		Use:          "read [source]",
		RunE:         read,
		SilenceUsage: true,
	}
)

func read(cmd *cobra.Command, args []string) error {
	r, err := cli.FileInput(cmd.Flags(), 0)
	if err != nil {
		return fmt.Errorf("error opening file: %w", err)
	}

	decompressor, err := zlib.NewReader(r)
	if err != nil {
		return err
	}

	defer decompressor.Close()

	_, err = io.Copy(os.Stdout, decompressor)
	return err
}
