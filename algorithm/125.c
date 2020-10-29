#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>
#include <ctype.h>
int isPalindrome(char *s)
{
    if (s == NULL)
    {
        return 1;
    }
    int slen = strlen(s);
    int left = 0, right = slen - 1;
    while (left < right)
    {

        if (!isalnum(s[left]))
        {
            left++;
            continue;
        }
        if (!isalnum(s[right]))
        {
            right--;
            continue;
        }
        if (tolower(s[left]) == tolower(s[right]))
        {
            left++;
            right--;
        }
        else
        {
            return 0;
        }
    }
    return 1;
}

int main(int argc, char *argv[])
{
    char s[11] = "race a car\0";
    printf("%d\n", isPalindrome(s));
    return 0;
}