#include <stdio.h>

#define BASE_PRICE 100

int main( int argc, char **argv) {
    static int quantity = 5;
    static double price = 25.00;
    static double total;

    quantity = argc + 2;
    total = price + BASE_PRICE;

    printf("Quantity: %d, Total: %.2f\n", quantity, total);
    return 0;
}