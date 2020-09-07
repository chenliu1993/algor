#include <stdio.h>
#define MAXSIZE 10
#define arraysize(re) \
    (sizeof(re) / sizeof *(re))

/* arraysize is used compare the size of an array. */

int
main(int argc, char *argv[]) {
    int a[MAXSIZE];
    a[0]=0;
    a[1]=1;
    a[2]=2;

    printf("%d\n", *(a+1));
    printf("%ld, %ld\n", sizeof(a), sizeof *(a));
    printf("the arraysize is %ld\n", arraysize(a));
    return 0;
}