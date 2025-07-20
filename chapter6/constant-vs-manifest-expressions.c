#include <stdio.h>
#include <stdlib.h>

#define THOUSAND 1000
/* this is a manifest constant. use these */

int main() {

    /* 
    int i;
    int oneThousand = 1000; // this allocates 4 bytes
    i = oneThousand; // this allocates 4 bytes
    return i;

    we used 8 bytes, not optimal.
    */

    int i;
    i = THOUSAND; // this ONLY allocated 4 bytes
    // we used 4 bytes for this instead of 8.

    return i;
}