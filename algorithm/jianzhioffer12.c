#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>
#include <ctype.h>
#include <math.h>

int dfs(char **board, int m, int n, char *word, int k, int x, int y)
{
    if (x < 0 || x >= m)
    {
        return 0;
    }
    if (y < 0 || y >= n)
    {
        return 0;
    }
    if (word[k] != board[x][y])
    {
        return 0;
    }
    if (k == strlen(word) - 1)
    {
        return 1;
    }
    char temp = board[x][y];
    board[x][y] = ' ';
    int res = dfs(board, m, n, word, k + 1, x - 1, y) || dfs(board, m, n, word, k + 1, x + 1, y) ||
              dfs(board, m, n, word, k + 1, x, y - 1) || dfs(board, m, n, word, k + 1, x, y + 1);
    board[x][y] = temp;
    return res;
}

int exist(char **board, int boardSize, int *boardColSize, char *word)
{
    int m = boardSize;
    int n = boardColSize[0];
    for (int i = 0; i < m; i++)
    {
        for (int j = 0; j < n; j++)
        {
            if (dfs(board, m, n, word, 0, i, j) == 1)
            {
                return 1;
            }
        }
    }
    return 0;
}

int main(int argc, char *argv[])
{
    char **board = (char **)malloc(3 * sizeof(char *));
    char b0[5] = {"ABCE\0"};
    char b1[5] = {"SFCS\0"};
    char b2[5] = {"ADEE\0"};
    char word[7] = {"ABCCED\0"};
    board[0] = b0;
    board[1] = b1;
    board[2] = b2;
    int boardSize = 3;
    int *boardColSize = (int *)calloc(1, sizeof(int));
    boardColSize[0] = 4;
    printf("%d\n", exist(board, boardSize, boardColSize, word));
    return 0;
}