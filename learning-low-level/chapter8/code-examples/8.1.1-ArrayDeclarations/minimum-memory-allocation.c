#include <stdio.h>

typedef struct {
    long a;
    char b;
} FiveBytes;

static FiveBytes shortArray[2] = {{2,3}, {4,5}};

/* 

        .file   "minimum-memory-allocation.c"
        .text
        .data
        .align 32
        .type   shortArray, @object
        .size   shortArray, 32
shortArray:
        .quad   2
        .byte   3
        .zero   7
        .quad   4
        .byte   5
        .zero   7
        .section        .rodata
.LC0:
        .string "%ld\n"
        .text
        .globl  main
        .type   main, @function

so this is also interesting. the compiler allocated 32 bytes for this struct.

.quad   2
.byte   3

this is shortArray[0].a for the quad and shortArray[0].b for the byte.

.zero   7

this is padding, to align next element to 4 bytes.

.quad   4
.byte   5

this is shortArray[1].a for the quad and shortArray[1].b for the byte.

.zero   7

this is padding, end of an array.

*/

int main( void ) {
    printf( "%ld\n", shortArray[0].a);
}