#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>
#include <ctype.h>
#include <math.h>

struct TreeNode
{
    int val;
    struct TreeNode *left;
    struct TreeNode *right;
};

#define MAXSIZE 10000
int **rc;
int *currentPath;
int pathTop;
void dfs(struct TreeNode *root, int sum, int *returnSize, int **returnColumnSizes)
{
    rc[pathTop++] = root->val;
    sum -= root->val;
    if (!root->left && !root->right)
    {
        if (!sum)
        {
            rc[*returnSize] = (int *)malloc(sizeof(int) * pathTop);
            memcpy(rc[*returnSize], currentPath, sizeof(int) * pathTop);
            (*returnColumnSizes)[*returnSize] = pathTop;
            (*returnSize)++;
        }
    }
    else
    {
        if (root->left)
        {
            dfs(root->left, sum, returnSize, returnColumnSizes);
        }
        if (root->right)
        {
            dfs(root->right, sum, returnSize, returnColumnSizes);
        }
    }
    pathTop--;
}

int **pathSum(struct TreeNode *root, int sum, int *returnSize, int **returnColumnSizes)
{
    rc = (int **)malloc(MAXSIZE * sizeof(int *));
    currentPath = (int *)malloc(MAXSIZE * sizeof(int));
    pathTop = 0;
    *returnSize = 0;
    *returnColumnSizes = (int *)malloc(1000 * sizeof(int *));
    if (root != NULL)
    {
        dfs(root, sum, returnSize, returnColumnSizes);
    }
    return rc;
}