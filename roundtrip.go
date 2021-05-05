package main

import (
	"bytes"
	"compress/zlib"
	"io"
	"os"

	"ajm188.scratchpad/zlib-example/cli"
	"github.com/spf13/cobra"
)

var (
	roundtripCmd = &cobra.Command{
		Use:          "roundtrip [source]",
		RunE:         roundtrip,
		SilenceUsage: true,
	}
)

func init() {
	rootCmd.AddCommand(roundtripCmd)
}

func roundtrip(cmd *cobra.Command, args []string) error {
	r, err := cli.FileInput(cmd.Flags(), 0)
	if err != nil {
		return err
	}
	defer r.Close()

	var buf bytes.Buffer
	compressor := zlib.NewWriter(&buf)

	_, err = io.Copy(compressor, r)
	if err != nil {
		return err
	}

	compressor.Close()

	decompressor, err := zlib.NewReader(&buf)
	if err != nil {
		return err
	}

	defer decompressor.Close()

	_, err = io.Copy(os.Stdout, decompressor)
	if err != nil {
		return err
	}

	return nil
}
