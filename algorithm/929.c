#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>
#include <ctype.h>
int findAddr(char **diffs, int diffSize, char *tgt)
{
    for (int i = 0; i < diffSize; i++)
    {
        if (strcmp(diffs[i], tgt) == 0)
        {
            return 1;
        }
    }
    return 0;
}
int numUniqueEmails(char **emails, int emailsSize)
{
    char **diffs = (char **)malloc(100 * sizeof(char));
    for (int i = 0; i < 100; i++)
    {
        diffs[i] = (char *)calloc(100, sizeof(char));
    }
    int index, at, elen, count = 0;
    for (int i = 0; i < emailsSize; i++)
    {
        index = 0;
        elen = strlen(emails[i]);
        for (int j = 0; j < elen; j++)
        {
            if (emails[i][j] == '@')
            {
                at = j;
                break;
            }
        }
        char *tmp = (char *)calloc(100, sizeof(char));
        for (int j = 0; j < at; j++)
        {
            if (emails[i][j] == '.')
            {
                continue;
            }
            if (emails[i][j] == '+')
            {
                break;
            }
            tmp[index++] = emails[i][j];
        }
        for (int j = at; j < elen; j++)
        {
            tmp[index++] = emails[i][j];
        }
        if (findAddr(diffs, 100, tmp) != 1)
        {
            diffs[count++] = tmp;
        }
    }
    return count;
}

int main(int argc, char *argv[])
{
    int emailsSize = 3;
    char **emails = (char **)malloc(emailsSize * sizeof(char *));
    for (int i = 0; i < emailsSize; i++)
    {
        emails[i] = (char *)calloc(100, sizeof(char));
    }
    char email1[29] = "test.email+alex@leetcode.com\0";
    char email2[35] = "test.e.mail+bob.cathy@leetcode.com\0";
    char email3[31] = "testemail+david@lee.tcode.com\0";
    emails[0] = email1;
    emails[1] = email2;
    emails[2] = email3;
    printf("%d\n", numUniqueEmails(emails, emailsSize));
    return 0;
}