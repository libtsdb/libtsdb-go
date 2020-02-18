#include<stdio.h>

int main() {
    int a = 1024;
    printf("%ld %d\n", sizeof(a), a);
    // cast directly to see endianness
    unsigned char* bytes = (unsigned char*) &a;
    for (int i = 0; i < 4; i++) {
        printf("%u ", bytes[i]);
    }
    printf("\n");
}

// 4 1024
// 0 4 0 0
