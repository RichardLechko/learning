#include <stdio.h>

static int i = 5;
static int j = 6;

int main( int argc, char **argv) {
    i = j + 3;
    j = i + 2;
    printf("%d %d\n", i, j);
    return 0;
}