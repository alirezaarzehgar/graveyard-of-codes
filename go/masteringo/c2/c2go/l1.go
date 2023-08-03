package main

//#include <stdio.h>
//#include <unistd.h>
//#include <string.h>
//#include <stdlib.h>
//void CallC()
//{
//	printf("Hello, World\n");
//}
import "C"
import "unsafe"

func fmain() {
	C.CallC()
	a := C.CString("Hello")
	defer C.free(unsafe.Pointer(a))
}
