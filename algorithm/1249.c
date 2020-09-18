#include <stdio.h>
#include <stdlib.h>
#include <string.h>

char *minRemoveToMakeValid(char *s)
{
    int rear = 0;
    int slen = strlen(s);
    int *stack = (int *)malloc(sizeof(int) * slen);
    for (int i = 0; i < slen; i++)
    {
        if (s[i] != '(' && s[i] != ')')
        {
            continue;
        }
        if (s[i] == '(')
        {
            stack[rear++] = i;
        }
        else if (s[i] == ')')
        {
            if (rear == 0)
            {
                s[i] = -1;
            }
            else
            {
                rear--;
            }
        }
    }
    for (int i = 0; i < rear; i++)
    {
        s[stack[i]] = -1;
    }
    int left = 0, right = 0;
    while (right < slen)
    {
        if (s[right] == -1)
        {
            right++;
        }
        else
        {
            s[left++] = s[right++];
        }
    }
    s[left] = '\0';
    return s;
}

int main(int argc, char *argv[])
{
    char a[13] = {'l', 'e', 'e', '(', 't', '(', 'c', ')', 'o', ')', 'd', 'e', ')'};
    printf("%s\n", minRemoveToMakeValid(a));
    return 0;
}