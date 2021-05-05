package benchmarks

import (
	"bytes"
	"compress/zlib"
	"io"
	"strings"
	"testing"
)

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
