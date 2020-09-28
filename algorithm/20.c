#include <stdio.h>
#include <stdlib.h>
#include <string.h>
char another(char s)
{
    if (s == ')')
    {
        return '(';
    }
    if (s == ']')
    {
        return '[';
    }
    if (s == '}')
    {
        return '{';
    }
    return 0;
}
int isValid(char *s)
{
    int slen = strlen(s), top = 0;
    if (slen % 2 == 1)
    {
        return 0;
    }
    char *rc = (char *)malloc(sizeof(char) * (slen + 1));
    memset(rc, 0, sizeof(char) * (slen + 1));
    for (int i = 0; i < slen; i++)
    {
        if (s[i] == '(' || s[i] == '[' || s[i] == '{')
        {
            rc[top++] = s[i];
        }
        else
        {
            if (top == 0 || s[top - 1] != another(s[i]))
            {
                return 0;
            }
            top--;
        }
    }
    if (top == 0)
    {
        return 1;
    }
    return 0;
}

int main(int argc, char *argv[])
{
    char s[7] = "()[]{}\0";
    printf("%d\n", isValid(s));
    return 0;
}