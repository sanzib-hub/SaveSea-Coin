package blockchain

import (
	"bytes"
	"encoding/binary"
	"strconv"
)

func IntToHex(n int64) []byte {
	return []byte(strconv.FormatInt(n, 10))
}

// Option 2: binary encoding (more compact, less readable)
func IntToBinary(n int64) []byte {
	buf := new(bytes.Buffer)
	_ = binary.Write(buf, binary.BigEndian, n)
	return buf.Bytes()
}
