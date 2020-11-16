#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>
#include <ctype.h>
#include <math.h>
int rotateString(char *A, char *B)
{
    int alen = strlen(A);
    int blen = strlen(B);
    if (alen != blen)
    {
        return 0;
    }
    if (alen == 0)
    {
        return 1;
    }
    if (alen == 1)
    {
        if (A[0] == B[0])
        {
            return 1;
        }
        else
        {
            return 0;
        }
    }
    for (int i = 0; i < alen; i++)
    {
        int idx = 0;
        char *tmp = (char *)calloc(alen + 1, sizeof(char));
        for (int m = i + 1; m < alen; m++)
        {
            tmp[idx++] = A[m];
        }
        for (int m = 0; m <= i; m++)
        {
            tmp[idx++] = A[m];
        }
        tmp[idx++] = '\0';
        if (strcmp(tmp, B) == 0)
        {
            return 1;
        }
        else
        {
            free(tmp);
        }
    }
    return 0;
}
int main(int argc, char *argv[])
{
    char a[6] = "abcde\0";
    char b[6] = "deabc\0";
    printf("%d\n", rotateString(a, b));
    return 0;
}
