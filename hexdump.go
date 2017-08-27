package hexdump

import (
	"bytes"
	"fmt"
)

const (
	// Default number of columns. Currently not configurable.
	columns int = 80

	// Number of bytes from the slice that will fit on a single line.
	bytesPerRow int = (columns - 1) / 4

	// The offset into the line buffer where the ASCII representation starts.
	asciiStartChar int = bytesPerRow*3 + 1
)

// Array of hex characters for byte to string conversion.
var hex []byte = []byte{
	'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
	'A', 'B', 'C', 'D', 'E', 'F',
}

// Converts a byte slice into a single line string of space separated
// hexadecimal values for each entry in the byte slice.
// For example:
//    s := hexdump.ByteToHexString([]byte{0x41, 0x42})
//    // s == "0x41 0x42"
func ByteToHexString(b []byte) string {
	var buffer bytes.Buffer

	for i := 0; i < len(b); i++ {
		buffer.WriteString(fmt.Sprintf("0x%2x ", b[i]))
	}

	buffer.Truncate(len(b)*5 - 1)

	return buffer.String()
}

// Returns a string representation of the provided by slice in the typical
// canonical Hex+ASCII format. The string is formatted into an 80 column
// output, splitting the slice into as many bytes (hex + ascii) as can fit
// on a single 80 column line.
//
// Non-printable ASCII values are represented with '.' characters.
//
// For example:
//    s := hexdump.ByteToHexAsciiString([]byte{0x41, 0x42, 0x43})
//    // s == "41 42 43  ABC"
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
