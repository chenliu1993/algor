#include <stdio.h>
#include <stdlib.h>
#include <string.h>
char *restoreString(char *s, int *indices, int indicesSize)
{
    int j = 0, i = 0;
    char *rc = (char *)calloc(indicesSize + 1, sizeof(char));
    while (i < indicesSize)
    {
        rc[indices[i++]] = s[j++];
    }
    return rc;
}

void printArray(char *a, int asize)
{
    for (int i = 0; i < asize; i++)
    {
        printf("%c,", a[i]);
    }
    printf("\n");
}
int main(int argc, char *argv[])
{
    char a[9] = "codeleet\0";
    int idx[8] = {4, 5, 6, 7, 0, 2, 1, 3};
    printArray(a, 8);
    printArray(restoreString(a, idx, 8), 8);
    return 0;
}