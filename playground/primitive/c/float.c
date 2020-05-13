#include <stdio.h>

// from https://youtu.be/y6qFHu0YKMM?t=688
int main() {
    float x = 0.1;
    float y = 0.2;
    printf("x %.20f y %.20f\n", x, y);
    printf("x + y = %.20f\n", x + y);
    printf("0.3 = %.20f\n", 0.3);
// x 0.10000000149011611938 y 0.20000000298023223877
// x + y = 0.30000001192092895508
// 0.3 = 0.29999999999999998890
}