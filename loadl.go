package main

// #cgo CFLAGS: -I/data/src/MemoryModule
// #cgo LDFLAGS: /data/src/MemoryModule/build/MemoryModule.a
// #include "loadl.h"
import "C"
import "fmt"


func main() {

	// LoadFromMemory(argv[1])
        fmt.Println("Loading DLL")
        sc_c:=C.CString("my.dll")
	C.LoadFromMemory(sc_c)
}
