#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>
#include <ctype.h>
#include <math.h>
int direction[4][2] = {{1, 0}, {-1, 0}, {0, 1}, {0, -1}};
int findPaths(int m, int n, int N, int i, int j)
{
    int mod = 10 * 10 * 10 * 10 * 10 * 10 * 10 + 9;
    int **dp = (int **)malloc(m * sizeof(int *));
    for (int i = 0; i < m; i++)
    {
        dp[i] = (int *)malloc(n * sizeof(int));
        memset(dp[i], 0, n * sizeof(int));
    }
    int count = 0;
    dp[i][j] = 1;
    for (int k = 1; k <= N; k++)
    {
        int **temp = (int **)malloc(m * sizeof(int *));
        for (int i = 0; i < m; i++)
        {
            temp[i] = (int *)malloc(n * sizeof(int));
            memset(temp[i], 0, n * sizeof(int));
        }
        for (int x = 0; x < m; x++)
        {
            for (int y = 0; y < n; y++)
            {
                if (dp[x][y] != 0)
                {
                    for (int d = 0; d < 4; d++)
                    {
                        int newx = x + direction[d][0];
                        int newy = y + direction[d][1];
                        if (newx < 0 || newx >= m || newy < 0 || newy >= n)
                        {
                            count = (count + dp[x][y]) % mod;
                        }
                        else
                        {
                            temp[newx][newy] = (dp[x][y] + temp[newx][newy]) % mod;
                        }
                    }
                }
            }
        }
        dp = temp;
    }
    return count;
}

int main(int argc, char *argv[])
{
    printf("%d\n", findPaths(1, 3, 3, 0, 1));
    return 0;
}