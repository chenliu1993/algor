#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int countSubstrings(char *s)
{
    int n = strlen(s), ans = 0;
    for (int i = 0; i < 2 * n - 1; ++i)
    {
        int l = i / 2, r = i / 2 + i % 2;
        while (l >= 0 && r < n && s[l] == s[r])
        {
            --l;
            ++r;
            ++ans;
        }
    }
    return ans;
}

int main(int argc, char *argv[])
{
    char s[4] = "aaa\0";
    printf("%d\n", countSubstrings(s));
    return 0;
}