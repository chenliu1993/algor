#include <stdlib.h>
#include <stdio.h>

/*
    https://blog.csdn.net/shanghx_123/article/details/79679726
    if pack(4) then it is the default settings.
*/
#pragma pack(2)
typedef struct _TestStruct {
    char a;
    int b;
    char c;
    short d;
} TestStruct;

int
main(int argc, char *argv[]) {
    printf("the size of stuct is %lu, the align is %lu\n", sizeof(TestStruct), __alignof__(TestStruct));
    return 0;
}