#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int canTransform(char *start, char *end)
{
    int i, j, slen = strlen(start);
    i = j = 0;
    while (i < slen && j < slen)
    {
        while (j < slen && end[j] == 'X')
            j++;
        while (i < slen && start[i] == 'X')
            i++;
        if (start[i] != end[j] || start[i] == 'L' && i < j || start[i] == 'R' && i > j)
        {
            return 0;
        }
        else
        {
            i++;
            j++;
        }
    }
    while (i < slen && start[i] == 'X')
        i++;
    while (j < slen && end[j] == 'X')
        j++;
    if (i != j)
        return 0;
    return 1;
}

int main(int argc, char *argv[])
{
    char start[10] = "RXXLRXRXL\0";
    char end[10] = "XRLXXRRLX\0";
    printf("%d\n", canTransform(start, end));
    return 0;
}