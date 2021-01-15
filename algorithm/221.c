#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>
#include <ctype.h>
#include <math.h>
#define MIN(x, y) (x < y ? x : y)
#define MAX(x, y) (x < y ? y : x)
int maximalSquare(char **matrix, int matrixSize, int *matrixColSize)
{
    int m = matrixSize;
    int n = *matrixColSize;
    int maxSide = INT_MIN;
    int **dp = (int **)malloc(m * sizeof(int *));
    for (int i = 0; i < m; i++)
    {
        dp[i] = (int *)calloc(n, sizeof(int));
    }
    for (int i = 0; i < m; i++)
    {
        for (int j = 0; j < n; j++)
        {
            if (matrix[i][j] == '1')
            {
                if (i == 0 || j == 0)
                {
                    dp[i][j] = 1;
                }
                else
                {
                    dp[i][j] = MIN(dp[i - 1][j], MIN(dp[i][j - 1], dp[i - 1][j - 1])) + 1;
                }
                maxSide = MAX(maxSide, dp[i][j]);
            }
        }
    }
    return maxSide * maxSide;
}

int main(int argc, char *argv[])
{
    int matrixSize = 4;
    int *matrixColSize = (int *)calloc(1, sizeof(int));
    *matrixColSize = 5;
    char **matrix = (char **)malloc(4 * sizeof(char *));
    char m0[6] = "10100\0";
    char m1[6] = "10111\0";
    char m2[6] = "11111\0";
    char m3[6] = "10010\0";
    matrix[0] = m0;
    matrix[1] = m1;
    matrix[2] = m2;
    matrix[3] = m3;
    printf("%d\n", maximalSquare(matrix, matrixSize, matrixColSize));
    return 0;
}