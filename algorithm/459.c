#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>
#include <ctype.h>
int repeatedSubstringPattern(char *s)
{
    int n = strlen(s);
    for (int i = 1; i * 2 <= n; ++i)
    {
        if (n % i == 0)
        {
            int match = 1;
            for (int j = i; j < n; ++j)
            {
                if (s[j] != s[j - i])
                {
                    match = 0;
                    break;
                }
            }
            if (match)
            {
                return 1;
            }
        }
    }
    return 0;
}
int main(int argc, char *argv[])
{
    char a[7] = "abcabc\0";
    printf("%d\n", repeatedSubstringPattern(a));
    return 0;
}