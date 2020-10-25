#include <stdio.h>
#include <stdlib.h>
#include <string.h>
int directions[4][2] = {{0, 1}, {0, -1}, {1, 0}, {-1, 0}};
int check(char **board, char *word, int **visited, int boardSize, int *boardColSize, int i, int j, int k)
{
    int slen = strlen(word);
    if (k == slen - 1)
    {
        return 1;
    }
    else if (board[i][j] != word[k])
    {
        return 0;
    }
    visited[i][j] = 1;
    int result = 0;
    for (int n = 0; n < 4; n++)
    {
        int newx = i + directions[n][0];
        int newy = j + directions[n][1];
        if (newx >= 0 && newx < boardSize && newy >= 0 && newy < boardColSize[0])
        {
            if (!visited[newx][newy])
            {
                int rc = check(board, word, visited, boardSize, boardColSize, newx, newy, k + 1);
                if (rc)
                {
                    result = 1;
                    break;
                }
            }
        }
    }
    visited[i][j] = 0;
    return result;
}

int exist(char **board, int boardSize, int *boardColSize, char *word)
{
    int **visited = (int **)malloc(sizeof(int *) * boardSize);
    for (int i = 0; i < boardSize; i++)
    {
        visited[i] = (int *)malloc(sizeof(int) * boardColSize[0]);
        memset(visited[i], 0, sizeof(int) * boardColSize[0]);
    }
    for (int i = 0; i < boardSize; i++)
    {
        for (int j = 0; j < boardColSize[0]; j++)
        {
            int rc = check(board, word, visited, boardSize, boardColSize, i, j, 0);
            if (rc)
            {
                return 1;
            }
        }
    }
    return 0;
}