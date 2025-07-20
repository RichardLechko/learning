#include <stdio.h>

int main() {
    /* 
    const int arraySize = 128;
    int anArray[arraySize];
    anArray[0] = 123;

    printf("This is an array: "%d", anArray[0]);
    this just prints out: "This is an array: 123"
    */


    const int arraySizes[2] = {123,256}; // this is OK
    const int arraySize = arraySizes[0]; // this is also OK

    int array[ arraySize ]; 
    // this is not OK. this array is created at compile but the length of the array is calculated on run time.
    // there are modern day AVLs, but the idea is that you shouldn't do this.
}