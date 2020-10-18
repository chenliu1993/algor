#include <stdio.h>
#include <stdlib.h>
#include <string.h>
char *convertToBase7(int num)
{
    if (num == 0)
        return "0";
    char *retstr = malloc(12 * sizeof(char));
    memset(retstr, 0, 12 * sizeof(char));
    if (num < 0)
    {
        strcat(retstr, "-");
    }
    num = abs(num);
    char arr[12];
    memset(arr, 0, 12 * sizeof(char));
    int i = 0;
    while (num > 0)
    {
        arr[i] = (char)(num % 7 + '0');
        num = num / 7;
        i++;
    }
    i--;
    for (; i >= 0; i--)
    {
        strncat(retstr, arr + i, 1);
    }
    return retstr;
}

int main(int argc, char *argv[])
{
    printf("%s\n", convertToBase7(250));
    return 0;
}