package main

/*
#cgo LDFLAGS: -luuid
#include <stdio.h>
#include <stdlib.h>
#include <uuid/uuid.h>
void printMessage(const char* s) {
    printf("%s\n", s);
}
char* _go_uuid() {
	uuid_t uuid;
	uuid_generate_random(uuid);
	char *str = malloc(37);
	uuid_unparse_lower(uuid, str);
	return str;
}
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func uuid() string {
	return C.GoString(C._go_uuid())
}

func main() {
	// Convert a Go string to a C string
	message := C.CString("Hello from C!")
	// Schedule the C string to be freed when the main function returns
	defer C.free(unsafe.Pointer(message))

	// Call the C function
	C.printMessage(message)
	fmt.Println(uuid())
}
