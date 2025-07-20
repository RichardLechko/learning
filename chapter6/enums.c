#include <stdio.h>

/* 
enum display {crt, lcd, led, plasma, oled };
by default the values are set: 0 to n with an increment of 1
*/

/* 
#define CRT 0
#define LCD 1
#define LED 2
#define PLASMA 3
#define OLED 4

this is the same thing as:

enum displays {crt, lcd, led, plasma, oled};
*/

enum displays {crt = 10, lcd = 20, led = 30, plasma = 40, oled = 50};

int main() {
    enum displays myDisplays = lcd;

    printf("Display type: %d\n", myDisplays);

    if (myDisplays == lcd) {
        printf("Using LCD Display!\n");
    }

    return 0;
}