#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>
#include <ctype.h>
#include <math.h>

void printGrid(char **a, int aSize, int aColSize)
{
    for (int i = 0; i < aSize; i++)
    {
        for (int j = 0; j < aColSize; j++)
        {
            printf("%c,", a[i][j]);
        }
        printf("\n");
    }
    printf("\n");
}

int rear;
void DFS(char *s, int slen, int count, int *visited, char **rc, char *ele)
{
    if (count == slen)
    {
        rc[rear] = (char *)malloc(slen * sizeof(char));
        memcpy(rc[rear++], ele, slen * sizeof(char));
        return;
    }
    for (int i = 0; i < slen; i++)
    {
        if (visited[i] != 0)
        {
            continue;
        }
        ele[count] = s[i];
        visited[i] = 1;
        DFS(s, slen, count + 1, visited, rc, ele);
        visited[i] = 0;
    }
}

char **permutation(char *s, int *returnSize)
{
    int slen = strlen(s);
    if (s == NULL || slen == 0)
    {
        (*returnSize) = 0;
        return NULL;
    }
    int sum = 1;
    for (int i = 1; i <= slen; i++)
    {
        sum *= i;
    }
    int *visited = (int *)calloc(slen, sizeof(int));
    char **rc = (char **)malloc(sum * sizeof(char *));

    char *ele = (char *)malloc(slen * sizeof(char));
    rear = 0;
    DFS(s, slen, 0, visited, rc, ele);
    (*returnSize) = sum;
    return rc;
}

int main(int argc, char *argv[])
{
    char s[4] = "abc\0";
    int *returnSize = (int *)calloc(1, sizeof(int));
    char **rc = permutation(s, returnSize);
    printGrid(rc, *returnSize, 3);
    return 0;
}