#include <stdio.h>

int main() {
    char *str1 = "Hello World";
    char *str2 = "Hello World";
    char *str3 = "Hello World";
    
    printf("str1 address: %p\n", (void*)str1);
    printf("str2 address: %p\n", (void*)str2);
    printf("str3 address: %p\n", (void*)str3);
    
    if (str1 == str2 && str2 == str3) {
        printf("Optimization: Single copy (GOOD)\n");
    } else {
        printf("No optimization: Multiple copies (BAD)\n");
    }
    
    return 0;
}