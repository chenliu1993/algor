#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int checkValidString(char *s)
{
    if (s == NULL)
    {
        return 1;
    }
    int slen = strlen(s);
    int minlen, maxlen;
    minlen = maxlen = 0;
    for (int i = 0; i < slen; i++)
    {
        if (s[i] == '(')
        {
            minlen++;
            maxlen++;
        }
        if (s[i] == '*')
        {
            minlen--;
            minlen = minlen < 0 ? 0 : minlen;
            maxlen++;
        }
        if (s[i] == ')')
        {
            minlen--;
            minlen = minlen < 0 ? 0 : minlen;
            maxlen--;
        }
        if (maxlen < 0)
        {
            return 0;
        }
    }
    if (minlen == 0)
    {
        return 1;
    }
    return 0;
}
int main(int argc, char *argv[])
{
    char s[12] = "(((******))\0";
    printf("%d\n", checkValidString(s));
    return 0;
}