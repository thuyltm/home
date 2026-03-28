package main

/*
#cgo CFLAGS: -g -Wall -I.
#cgo LDFLAGS: -L. -ladd
#include "libadd.h"
#include <stdlib.h>
*/
import "C"
import "fmt"

func main() {
	x := C.Add(1, 2)
	fmt.Println("Sum:", x)
	name := C.CString("Gopher")
	C.hello(name)
}

/*
CFLAGS: Tells the C compiler to look for header files
LDFLAGS: Tells the linker where to find the library
*/
