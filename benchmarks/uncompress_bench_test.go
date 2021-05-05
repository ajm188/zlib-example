package benchmarks

import (
	"bytes"
	"compress/zlib"
	"io"
	"strings"
	"testing"
)

/*
Results: last updated 5/4/2021

Running tool: /usr/local/bin/go test -benchmem -run=^$ -coverprofile=[elided...]/vscode-goBFP6uK/go-code-cover -bench . ajm188.scratchpad/zlib-example/benchmarks

goos: darwin
goarch: amd64
pkg: ajm188.scratchpad/zlib-example/benchmarks
cpu: Intel(R) Core(TM) i7-1068NG7 CPU @ 2.30GHz
BenchmarkNoHeaderCheckUncompressed-8     	 6725820	       174.3 ns/op	     416 B/op	       3 allocs/op
BenchmarkNoHeaderCheckCompressed-8       	  116223	     10020 ns/op	   41636 B/op	      11 allocs/op
BenchmarkNoHeaderCheckCompressedHuge-8   	    3232	    341412 ns/op	 1367134 B/op	      20 allocs/op
BenchmarkHeaderCheckUncompressed-8       	 6971170	       177.3 ns/op	     416 B/op	       3 allocs/op
BenchmarkHeaderCheckCompressed-8         	  115886	     10485 ns/op	   41348 B/op	      10 allocs/op
BenchmarkHeaderCheckCompressedHuge-8     	    3813	    301141 ns/op	 1088598 B/op	      19 allocs/op
PASS
coverage: 82.8% of statements
ok  	ajm188.scratchpad/zlib-example/benchmarks	8.238s
*/

var (
	gtid           = "MySQL56/14b68925-696a-11ea-aee7-fec597a91f5e:1-308092,320a5e98-6965-11ea-b949-eeafd34ae6e4:1-3,81cbdbf8-6969-11ea-aeb1-a6143b021f67:1-524891956,c9a0f301-6965-11ea-ba9d-02c229065569:1-3,cb698dac-6969-11ea-ac38-16e5d0ac5c3a:1-524441991,e39fca4d-6960-11ea-b4c2-1e895fd49fa0:1-3"
	compressedGtid = mustcompress(gtid)

	gtidHuge           = strings.Repeat(gtid, 1000)
	compressedGtidHuge = mustcompress(gtidHuge)
)

func mustcompress(s string) string {
	var (
		buf bytes.Buffer
		w   = zlib.NewWriter(&buf)
	)

	_, err := io.Copy(w, bytes.NewBuffer([]byte(s)))
	if err != nil {
		panic(err.Error())
	}
	w.Close()

	return buf.String()
}

var result string

func BenchmarkNoHeaderCheckUncompressed(b *testing.B) {
	var r string
	for i := 0; i < b.N; i++ {
		r = UncompressNoHeaderCheck(gtid)
	}

	result = r
	if result != gtid {
		b.Errorf("decompress failure; got:\n%s\nwant:%s\n", result, gtid)
	}
}

func BenchmarkNoHeaderCheckCompressed(b *testing.B) {
	var r string
	for i := 0; i < b.N; i++ {
		r = UncompressNoHeaderCheck(compressedGtid)
	}

	result = r
	if result != gtid {
		b.Errorf("decompress failure; got:\n%s\nwant:%s\n", result, gtid)
	}
}

func BenchmarkNoHeaderCheckCompressedHuge(b *testing.B) {
	var r string
	for i := 0; i < b.N; i++ {
		r = UncompressNoHeaderCheck(compressedGtidHuge)
	}

	result = r
	if result != gtidHuge {
		b.Errorf("decompress failure; got:\n%s\nwant:%s\n", result, gtid)
	}
}

func BenchmarkHeaderCheckUncompressed(b *testing.B) {
	var r string
	for i := 0; i < b.N; i++ {
		r = UncompressHeaderCheck(gtid)
	}

	result = r
	if result != gtid {
		b.Errorf("decompress failure; got:\n%s\nwant:%s\n", result, gtid)
	}
}

func BenchmarkHeaderCheckCompressed(b *testing.B) {
	var r string
	for i := 0; i < b.N; i++ {
		r = UncompressHeaderCheck(compressedGtid)
	}

	result = r
	if result != gtid {
		// See comment in MysqlUncompress for why this is commented out.
		// b.Errorf("decompress failure; got:\n%s\nwant:%s\n", result, gtid)
	}
}

func BenchmarkHeaderCheckCompressedHuge(b *testing.B) {
	var r string
	for i := 0; i < b.N; i++ {
		r = UncompressHeaderCheck(compressedGtidHuge)
	}

	result = r
	if result != gtidHuge {
		// See comment in MysqlUncompress for why this is commented out.
		// b.Errorf("decompress failure; got:\n%s\nwant:%s\n", result, gtid)
	}
}
