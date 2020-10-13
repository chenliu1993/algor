#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>

#define MIN(x, y) (x < y ? x : y)
int findSmallest(char *s, int ssize, char c, int position)
{
    int leftSmallest, rightSmallest;
    leftSmallest = rightSmallest = INT_MAX;
    for (int i = position; i >= 0; i--)
    {
        if (s[i] == c)
        {
            leftSmallest = position - i;
            break;
        }
    }
    for (int i = position; i < ssize; i++)
    {
        if (s[i] == c)
        {
            rightSmallest = i - position;
            break;
        }
    }
    return MIN(leftSmallest, rightSmallest);
}
int *shortestToChar(char *S, char C, int *returnSize)
{
    int slen = strlen(S);
    int *rc = (int *)calloc(slen, sizeof(int));
    for (int i = 0; i < slen; i++)
    {
        rc[i] = findSmallest(S, slen, C, i);
    }
    return rc;
}
void printArray(int *a, int asize)
{
    for (int i = 0; i < asize; i++)
    {
        printf("%d,", a[i]);
    }
    printf("\n");
}
int main(int argc, char *argv[])
{
    int *returnSize = (int *)calloc(1, sizeof(int));
    char a[13] = "loveleetcode\0";
    char b = 'e';
    int *c = shortestToChar(a, b, returnSize);
    printArray(c, 12);
    return 0;
}