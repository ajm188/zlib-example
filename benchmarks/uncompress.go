package benchmarks

import (
	"bytes"
	"compress/zlib"
	"encoding/binary"
	"io"
)

// MysqlUncompress will uncompress a binary string in the format stored by mysql's compress() function
// The first four bytes represent the size of the original string passed to compress()
// Remaining part is the compressed string using zlib, which we uncompress here using golang's zlib library
func MysqlUncompress(input string) []byte {
	// consistency check
	inputBytes := []byte(input)
	if len(inputBytes) < 5 {
		return nil
	}

	// determine length
	dataLength := uint32(inputBytes[0]) + uint32(inputBytes[1])<<8 + uint32(inputBytes[2])<<16 + uint32(inputBytes[3])<<24
	dataLengthBytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(dataLengthBytes, dataLength)
	dataLength = binary.LittleEndian.Uint32(dataLengthBytes)

	// uncompress using zlib

	// (@ajm188) Chopping off the first 4 bytes causes a zlib.ErrHeader error.
	// inputData := inputBytes[4:]
	// inputDataBuf := bytes.NewBuffer(inputData)
	// reader, err := zlib.NewReader(inputDataBuf)
	// if err != nil {
	// 	log.Print(err.Error())
	// 	return nil
	// }

	reader, err := zlib.NewReader(bytes.NewBuffer(inputBytes))
	if err != nil {
		return nil
	}
	var outputBytes bytes.Buffer
	io.Copy(&outputBytes, reader)
	if outputBytes.Len() == 0 {
		return nil
	}

	// (@ajm188): somehow, the length checksumming is never matching for me,
	// so the end-of-benchmark assertion for BenchmarkHeaderCheckCompressed
	// never passes
	if dataLength != uint32(outputBytes.Len()) { // double check that the stored and uncompressed lengths match
		return nil
	}
	return outputBytes.Bytes()
}

func UncompressHeaderCheck(input string) string {
	if b := MysqlUncompress(input); b != nil {
		return string(b)
	}

	return input
}

func UncompressNoHeaderCheck(input string) string {
	decompressor, err := zlib.NewReader(bytes.NewBuffer([]byte(input)))
	if err != nil {
		if err == zlib.ErrHeader {
			return input
		}

		panic(err.Error())
	}

	var buf bytes.Buffer
	io.Copy(&buf, decompressor)
	decompressor.Close()

	return buf.String()
}
