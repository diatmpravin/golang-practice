package main

import (
	"fmt"
	"unsafe"
)

type INotifyInfo struct {
	Wd     int32  // Watch descriptor
	Mask   uint32 // Watch mask
	Cookie uint32 // Cookie to synchronize two events
	Len    uint32 // Length (including NULs) of name
}

func main() {
	var info INotifyInfo
	const infoSize = unsafe.Sizeof(info)
	fmt.Println(infoSize)
}
