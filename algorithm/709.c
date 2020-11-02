#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>
#include <ctype.h>
char *toLowerCase(char *str)
{
    int slen = strlen(str);
    if (slen == 0 || str == NULL)
    {
        return NULL;
    }
    char *rc = (char *)malloc((slen + 1) * sizeof(char));
    memset(rc, 0, (slen + 1) * sizeof(char));
    for (int i = 0; i < slen; i++)
    {
        if (str[i] >= 'A' && str[i] <= 'Z')
        {
            rc[i] = str[i] - 'A' + 'a';
        }
        else
        {
            rc[i] = str[i];
        }
    }
    rc[slen] = '\0';
    return rc;
}
int main(int argc, char *argv[])
{
    char a[6] = "Hello\0";
    char *rc = toLowerCase(a);
    printf("%s\n", rc);
    return 0;
}