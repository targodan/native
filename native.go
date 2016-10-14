package native

import (
	"encoding/binary"
	"unsafe"
)

// ByteOrder contains the native binary.ByteOrder.
var ByteOrder binary.ByteOrder

func init() {
	var i int = 0x1
	check := (*[2]int)(unsafe.Pointer(&i))[0]
	if check == 0 {
		ByteOrder = binary.BigEndian
	} else {
		ByteOrder = binary.LittleEndian
	}
}
