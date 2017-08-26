package hexdump

import (
	"bytes"
	"fmt"
)

func ByteToHexString(b []byte) string {
	var buffer bytes.Buffer

	for i := 0; i < len(b); i++ {
		buffer.WriteString(fmt.Sprintf("0x%2x ", b[i]))
	}

	buffer.Truncate(len(b)*5 - 1)

	return buffer.String()
}

func ByteToHexAsciiString(b []byte) string {
	var hex bytes.Buffer
	var ascii bytes.Buffer
	var buffer bytes.Buffer
	var columns int = 80

	for i := 0; i < len(b); i++ {
		hex.WriteString(fmt.Sprintf("%02x ", b[i]))
		if b[i] > 33 && b[i] < 127 {
			ascii.WriteString(fmt.Sprintf("%c", b[i]))
		} else {
			ascii.WriteByte(0x2E)
		}
		if columns-(hex.Len()+ascii.Len()+2) < 3 {
			buffer.Write(hex.Bytes())
			buffer.WriteString("  ")
			buffer.Write(ascii.Bytes())
			buffer.WriteString("\n")
			hex.Reset()
			ascii.Reset()
		}
	}

	if hex.Len() > 0 {
		buffer.Write(hex.Bytes())
		blank := []byte{0x20, 0x20, 0x20}
		moar := (((columns-1)/4)*3 - hex.Len()) / 3
		for i := 0; i < moar; i++ {
			buffer.Write(blank)
		}

		buffer.WriteString("  ")
		buffer.Write(ascii.Bytes())
		buffer.WriteString("\n")
	}
	buffer.Truncate(buffer.Len() - 1)

	return buffer.String()
}
