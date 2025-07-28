#include <stdio.h>

typedef struct {
    int a;
    int b;
    char c;
} NineBytes;

static NineBytes shortArray[2] = {{2,3,1}, {4,5,1}};

/* 

        .file   "nine-byte-struct.c"
        .text
        .data
        .align 16
        .type   shortArray, @object
        .size   shortArray, 24
shortArray:
        .long   2
        .long   3
        .byte   1
        .zero   3
        .long   4
        .long   5
        .byte   1
        .zero   3
        .section        .rodata
.LC0:
        .string "%d\n"
        .text
        .globl  main
        .type   main, @function



*/

int main( void ) {
    printf( "%d\n", shortArray[0].a);
}