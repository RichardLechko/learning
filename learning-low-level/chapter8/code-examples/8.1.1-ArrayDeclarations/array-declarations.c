#include <stdio.h>
static int intArray[8] = {1,2,3,4,5,6,7,8};
// since this is a composite data type and can not be encoded directly into machine instruction, it gets thrown into the DATA segment

static int array2[8];
// this declares an array with 8 slots that are uninitialized
// this array gets loaded into the BSS segment, then when the OS loads the BSS segment into memory, the OS zeros out all the memory associated with that segment.

/* 
intArray:
	.long	1
	.long	2
	.long	3
	.long	4
	.long	5
	.long	6
	.long	7
	.long	8
	.local	array2
	.comm	array2,32,32
	.section	.rodata

you can see .long means 32 bits. intArray has 8 ints which is 4 * 8 = 32 bytes. so we use .long to load the values into memeory.
the ".comm" communicates to the program that this is an uninitialized array, but we should allocate 32 bytes of memory.
*/

int main( int argc, char **argv ) {
    int i;
    for (i = 0; i < 8; ++i) {
        array2[i] = intArray[i];
    } 
    for (i = 7; i >= 0; --i) {
        printf( "%d\n", array2[i]);
    }
    return 0;
}