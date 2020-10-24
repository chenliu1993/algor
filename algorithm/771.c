#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <limits.h>
int numJewelsInStones(char *J, char *S)
{
    int jlen = strlen(J);
    int slen = strlen(S);
    int count = 0;
    for (int i = 0; i < slen; i++)
    {
        for (int j = 0; j < jlen; j++)
        {
            if (S[i] == J[j])
            {
                count++;
            }
        }
    }
    return count;
}
int main(int argc, char *argv[])
{
    char J[3] = "aA\0";
    char S[8] = "aAAbbbb\0";
    printf("%d\n", numJewelsInStones(J, S));
    return 0;
}