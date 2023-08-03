package main

//#include "l2.h"
//#include <stdlib.h>
import "C"
import "unsafe"

func main() {
	out := C.createMsg()
	C.printMsg(out)
	C.free(unsafe.Pointer(out))
}
