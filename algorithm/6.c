#include <stdio.h>
#include <stdlib.h>
#include <string.h>

char *convert(char *s, int numRows)
{
    if (s == NULL || numRows == 1)
    {
        return s;
    }
    int slen = strlen(s);
    char *rc = (char *)malloc(sizeof(char) * (slen + 1));
    memset(rc, 0, sizeof(char) * (slen + 1));
    int rows = numRows, cols = numRows - 2, k = 0, col = cols;
    for (int i = 0; i < numRows; i++)
    {
        if (i == 0 || i == numRows - 1)
        {
            for (int j = i; j < slen; j += (rows + cols))
            {
                rc[k++] = s[j];
            }
        }
        else
        {
            for (int j = i; j < slen; j += (rows + cols))
            {
                rc[k++] = s[j];
                if (j + 2 * col < slen)
                {
                    rc[k++] = s[j + 2 * col];
                }
            }
            col--;
        }
    }
    printf("\n");
    rc[slen] = '\0';
    return rc;
}

int main(int argc, char *argv[])
{
    char s[15] = "PAYPALISHIRING\0";
    printf("%s\n", convert(s, 3));
    return 0;
}