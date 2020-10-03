#include <stdio.h>
#include <stdlib.h>
#include <string.h>

/**
 * BFS:
 *  while(queue not empty) {
 *      for(){
 *          in queue;
 *          out queue;
 *      }
 *      update conditions;
 *  }
 * */

int direction[4][2] = {{1, 0}, {-1, 0}, {0, 1}, {0, -1}};

void solve(char **board, int boardSize, int *boardColSize)
{
    int colsize = boardColSize[0];
    if (boardSize == 0)
    {
        return;
    }
    int **visited = (int **)calloc(boardSize * colsize, sizeof(int *));
    for (int i = 0; i < boardSize * colsize; i++)
    {
        visited[i] = (int *)calloc(2, sizeof(int));
    }
    int l = 0, r = 0;
    for (int i = 0; i < boardSize; i++)
    {
        if (board[i][0] == 'O')
        {
            board[i][0] = 'A';
            visited[r][0] = i;
            visited[r++][1] = 0;
        }
        if (board[i][colsize - 1] == 'O')
        {
            board[i][colsize - 1] = 'A';
            visited[r][0] = i;
            visited[r++][1] = colsize - 1;
        }
    }
    for (int i = 0; i < colsize; i++)
    {
        if (board[0][i] == 'O')
        {
            board[0][i] = 'A';
            visited[r][0] = 0;
            visited[r++][1] = i;
        }
        if (board[boardSize - 1][i] == 'O')
        {
            board[boardSize - 1][i] = 'A';
            visited[r][0] = boardSize - 1;
            visited[r++][1] = i;
        }
    }

    while (l < r)
    {
        int x = visited[l][0], y = visited[l][1];
        l++;
        for (int i = 0; i < 4; i++)
        {
            int mx = x + direction[i][0];
            int my = y + direction[i][1];
            if (mx < 0 || my < 0 || mx >= boardSize || my >= colsize || board[mx][my] != 'O')
            {
                continue;
            }
            board[mx][my] = 'A';
            visited[r][0] = mx;
            visited[r++][1] = my;
        }
    }
    for (int i = 0; i < boardSize; i++)
    {
        for (int j = 0; j < colsize; j++)
        {
            if (board[i][j] == 'A')
            {
                board[i][j] = 'O';
            }
            else if (board[i][j] == 'O')
            {
                board[i][j] = 'X';
            }
        }
    }
    for (int i = 0; i < boardSize * colsize; i++)
    {
        free(visited[i]);
    }
    free(visited);
}

void printBoard(char **board, int row, int col)
{

    for (int i = 0; i < row; i++)
    {
        for (int j = 0; j < col; j++)
        {
            printf("%c,", board[i][j]);
        }
        printf("\n");
    }
}

int main(int argc, char *argv[])
{
    char **board = (char **)calloc(4, sizeof(char *));
    for (int i = 0; i < 4; i++)
    {
        board[i] = (char *)calloc(5, sizeof(char));
    }

    char c1[5] = "XXXX\0";
    char c2[5] = "XOOX\0";
    char c3[5] = "XXOX\0";
    char c4[5] = "XOXX\0";

    board[0] = c1;
    board[1] = c2;
    board[2] = c3;
    board[3] = c4;
    printBoard(board, 4, 4);
    printf("\n");
    int colsize[4] = {4, 4, 4, 4};
    solve(board, 4, colsize);
    printBoard(board, 4, 4);
    return 0;
}