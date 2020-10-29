#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>
#include <ctype.h>
/** time over limit
int validPalindromeSingle(char *s, int avoid, int nodel)
{
    int slen = strlen(s);
    int left = 0, right = slen - 1;
    while (left < right)
    {

        if (left == avoid && nodel == 0)
        {
            left++;
        }
        if (right == avoid && nodel == 0)
        {
            right--;
        }
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

int validPalindrome(char *s)
{
    if (s == NULL)
    {
        return 1;
    }
    int slen = strlen(s);
    if (validPalindromeSingle(s, 0, 1))
    {
        return 1;
    }
    for (int i = 0; i < slen; i++)
    {
        if (validPalindromeSingle(s, i, 0))
        {
            return 1;
        }
    }
    return 0;
}**/
int validPalindromeSingle(char *s, int left, int right)
{
    int slen = strlen(s);
    int start = left, end = right;
    while (start < end)
    {
        if (s[start] == s[end])
        {
            start++;
            end--;
        }
        else
        {
            return 0;
        }
    }
    return 1;
}
int validPalindrome(char *s)
{
    if (s == NULL)
    {
        return 1;
    }
    int slen = strlen(s);
    int left = 0, right = slen - 1;
    while (left < right)
    {
        if (s[left] == s[right])
        {
            left++;
            right--;
        }
        else
        {
            return validPalindromeSingle(s, left, right - 1) || validPalindromeSingle(s, left + 1, right);
        }
    }
    return 1;
}
int main(int argc, char *argv[])
{
    char a[5] = "abca\0";
    printf("%d\n", validPalindrome(a));
    return 0;
}