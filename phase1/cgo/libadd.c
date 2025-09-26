#include "libadd.h"
#include <stdio.h>

int Add(int a, int b) {
    printf("Welcome from external C function\n");
    return a+b;
}

void hello(char* name) {
    printf("Hello %s\n", name);
}