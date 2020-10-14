#include <stdio.h>
#include <stdlib.h>
#include <string.h>
typedef struct _node
{
    int count;
    char c;
} node;

int compareFunc(const void *a, const void *b)
{
    return ((*(node *)a).count > (*(node *)b).count) ? 1 : -1;
}
char *frequencySort(char *s)
{
    int slen = strlen(s), ptr = 0;
    node *count = (node *)calloc(128, sizeof(node));
    for (int i = 0; i < slen; i++)
    {
        count[s[i]].c = s[i];
        count[s[i]].count++;
        printf("%d\n", s[i]);
    }
    qsort(count, 128, sizeof(node), compareFunc);
    char *rc = (char *)malloc(sizeof(char) * (slen + 1));
    memset(rc, 0, sizeof(char) * (slen + 1));
    for (int i = 128 - 1; i >= 0; i--)
    {
        if (count[i].count)
        {
            int sum = count[i].count;
            while (sum > 0)
            {
                rc[ptr++] = count[i].c;
                sum--;
            }
        }
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
    char a[5] = "Aabb\0";
    printArray(a, 4);
    char *b = frequencySort(a);
    printArray(b, 4);
    return 0;
}