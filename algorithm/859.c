#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>
#include <ctype.h>
int buddyStrings(char *A, char *B)
{
    int alen = strlen(A);
    int blen = strlen(B);
    if (alen != blen)
    {
        return 0;
    }
    int diff = 0;
    int *cnt = (int *)calloc(alen, sizeof(int));
    for (int i = 0; i < alen; i++)
    {
        if (A[i] != B[i])
        {
            cnt[diff++] = i;
        }
    }
    if (diff > 2 || diff == 1)
    {
        return 0;
    }
    else if (diff == 0)
    {
        // if there is two char same, then it means we can just switch these two.
        // Notify that diff==0 means A==B
        for (int i = 0; i < alen; i++)
        {
            for (int j = i + 1; j < blen; j++)
            {
                if (A[i] == A[j])
                {
                    return 1;
                }
            }
        }
    }
    else if (diff == 2 && A[cnt[0]] == B[cnt[1]] && A[cnt[1]] == B[cnt[0]])
    {
        return 1;
    }
    return 0;
}