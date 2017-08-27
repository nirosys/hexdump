package hexdump

import (
	"bytes"
	"fmt"
)

const (
	columns        int = 80
	bytesPerRow    int = (columns - 1) / 4
	asciiStartChar int = bytesPerRow*3 + 1
)

var hex []byte = []byte{
	'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
	'A', 'B', 'C', 'D', 'E', 'F',
}

func ByteToHexString(b []byte) string {
	var buffer bytes.Buffer

	for i := 0; i < len(b); i++ {
		buffer.WriteString(fmt.Sprintf("0x%2x ", b[i]))
	}

	buffer.Truncate(len(b)*5 - 1)

	return buffer.String()
}

func ByteToHexAsciiString(b []byte) string {
	var entire bytes.Buffer
	var line []byte = make([]byte, columns)

	for i := 0; i < columns; i++ {
		line[i] = 0x20
	}

	for i := 0; i < len(b); i++ {
		lineByte := i % bytesPerRow
		line[lineByte*3] = hex[b[i]>>4]
		line[(lineByte*3)+1] = hex[b[i]&0x0F]

		line[asciiStartChar+lineByte] = '.'
		if b[i] > 33 && b[i] < 127 {
			line[asciiStartChar+lineByte] = b[i]
		}

		if lineByte == bytesPerRow-1 {
			entire.Write(line[:asciiStartChar+bytesPerRow])
			entire.WriteString("\n")
		}
	}

	if r := len(b) % bytesPerRow; r > 0 {
		for i := r; i < bytesPerRow; i++ {
			line[i*3] = 0x20
			line[i*3+1] = 0x20
			line[asciiStartChar+i] = 0x20
		}
		entire.Write(line[:asciiStartChar+r])
	}

	return entire.String()
}
