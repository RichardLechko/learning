#include <stdio.h>

static int array[8] = {0,0,0,0,0,0,0,0};

/* 
.file   "array-in-memory-ints.c"
        .text
        .local  array
        .comm   array,32,32
        .section        .rodata
.LC0:
        .string "%d\n"
        .text
        .globl  main
        .type   main, @function

so this is interesting to see. the compiler with the -O0 flag is treating this array as if it was unininitialized.
so the compiler communicates to the OS with ".comm array,32,32" saying that this array has no data to be loaded in and that the OS needs to allocate 32 bytes of memeory.
when the OS loads this array into memeory, all the elements get zerod out.

*/

int main( void ) {
    printf( "%d\n", array[0]);
}