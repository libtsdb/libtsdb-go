#include <stdio.h>

// NOTE: it is from CSAPP 3 Figure 2.4

void show_bytes(unsigned char* start, size_t len) {
    for (int i = 0; i < len; i++) {
        printf(" %.2x", start[i]);
    }
    printf("\n");
}

void show_int(int x) {
    show_bytes((unsigned char*) &x, sizeof(int));
}

void show_float(float x) {
    show_bytes((unsigned char*) &x, sizeof(float));
}

int main() {
    show_int(1024);
    show_int(-1024);
    show_float(12345.0);
    show_float(-12345.0);
//  00 04 00 00
//  00 fc ff ff
//  00 e4 40 46
//  00 e4 40 c6

// ff ff ff ff
    show_int(-1);
    return 0;
}