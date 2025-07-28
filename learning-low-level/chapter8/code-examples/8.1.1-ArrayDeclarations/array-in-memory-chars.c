#include <stdio.h>

static char array[8] = {0,1,2,3,4,5,6,7};

/* 
	.file	"array-in-memory.c"
	.text
	.data
	.align 8
	.type	array, @object
	.size	array, 8
array:
	.string	""
	.ascii	"\001\002\003\004\005\006\007"
	.section	.rodata
.LC0:
	.string	"%d\n"
	.text
	.globl	main
	.type	main, @function

since a char is 1 byte, it makes sense that the align is 8 bytes.
then the OS loads the values into memeory: "\001\002\003\004\005\006\007"

the number of bytes an array consumes is the number of elements multiplied by the number of bytes per element.
*/

int main( void ) {
    printf( "%d\n", array[0]);
}