#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>
#include <ctype.h>
#define MIN(x, y) (x < y ? x : y)
int videoStitching(int **clips, int clipsSize, int *clipsColSize, int T)
{
    int *dp = (int *)calloc(T + 1, sizeof(int));
    for (int i = 1; i < T + 1; i++)
    {
        dp[i] = INT_MAX;
    }
    for (int i = 1; i < T + 1; i++)
    {
        for (int j = 0; j < clipsSize; j++)
        {
            if (clips[j][0] < i && i <= clips[j][1])
            {
                dp[i] = MIN(dp[i], dp[clips[j][0]] + 1);
            }
        }
    }
    return dp[T] == INT_MAX ? -1 : dp[T];
}

int main(int argc, char *argv[])
{
    int c1[2] = {0, 1};
    int c2[2] = {1, 2};
    int T = 5;
    int clipSize = 2;
    int **c = (int **)malloc(clipSize * sizeof(int *));
    c[0] = c1;
    c[1] = c2;
    int *clipColSize = (int *)calloc(clipSize, sizeof(int));
    for (int i = 0; i < clipSize; i++)
    {
        clipColSize[i] = 2;
    }
    printf("%d\n", videoStitching(c, clipSize, clipColSize, T));
    return 0;
}