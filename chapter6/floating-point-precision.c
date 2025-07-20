#include <stdio.h>
#include <stdlib.h>

/* static double OnePointZero_c = 1.0; */

int main( int argc, char **argv, char **envp) {

    
    static int j;
    static double i = 1.0;
    static double a[8] = {1,2,3,4,5,6,7};

    j = 0;
    a[j] = i+1.0;
   

    /* 
    static float i;
    i = 1.0;
    */

    /* static double i;
    i = OnePointZero_c; */
}