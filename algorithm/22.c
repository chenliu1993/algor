#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>
#include <ctype.h>
#include <math.h>

char **ans;
int left;
int right;
int rear;

void parenthesisUtil(int n, int *returnSize, int k, char *ele)
{
    if (left == n && right == n)
    {
        (*returnSize)++;
        ans[(*returnSize) - 1] = (char *)malloc(2 * n * sizeof(char));
        memcpy(ans[(*returnSize) - 1], ele, 2 * n * sizeof(char));
        return;
    }
    if (left < n)
    {
        ele[rear++] = '(';
        left++;
        parenthesisUtil(n, returnSize, k + 1, ele);
        rear--;
        left--;
    }
    if (left > right && right < n)
    {
        ele[rear++] = ')';
        right++;
        parenthesisUtil(n, returnSize, k + 1, ele);
        rear--;
        right--;
    }
}

char **generateParenthesis(int n, int *returnSize)
{
    char *ele = (char *)malloc(2 * n * sizeof(char));
    ans = (char **)malloc(2001 * sizeof(char *));
    left = 0;
    right = 0;
    rear = 0;
    (*returnSize) = 0;
    parenthesisUtil(n, returnSize, 0, ele);
    return ans;
}

int main(int argc, char *argv[])
{
    int n = 3;
    int *returnSize = (int *)calloc(1, sizeof(int));
    char **res = generateParenthesis(n, returnSize);
    for (int i = 0; i < (*returnSize); i++)
    {
        printf("%s\n", res[i]);
    }
    return 0;
}