#include <stdio.h>

/* 1-4 */

void
PrintDigit(int a) {
    printf("%d\n", a);
}

void
PrintOut(int a) {
    if (a>=0) {
        if (a>=10) {
            PrintOut(a/10);
        }
        PrintDigit(a%10);
    } else {
        a = -1 *a;
        printf("%d\n",-1);
        PrintOut(a);
    }
    
}

int
main(int argc, char *argv[]) {
    PrintOut(-76234);
    return 0;
}