package main

/*
#cgo CFLAGS: -g -Wall -I.
#cgo LDFLAGS: -L. -ladd
#cgo LDFLAGS: -luuid
#include "libadd.h"
#include <uuid/uuid.h>
#include <stdlib.h>
char* _go_uuid() {
	uuid_t uuid;
	uuid_generate_random(uuid);
	char *str = malloc(37);
	uuid_unparse_lower(uuid, str);
	return str;
}
*/
import "C"
import "fmt"

func uuid() string {
	return C.GoString(C._go_uuid())
}
func main() {
	x := C.Add(1, 2)
	fmt.Println("Sum:", x)
	name := C.CString("Gopher")
	C.hello(name)
	fmt.Println(uuid())
}
