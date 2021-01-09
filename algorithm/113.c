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

int **ans;
int rear;

void dfs(struct TreeNode *root, int sum, int cur, int *ele, int *returnSize, int **returnColumnSize)
{
    if (cur == sum && root->left == NULL && root->right == NULL)
    {
        (*returnSize)++;
        ans[(*returnSize) - 1] = malloc(rear * sizeof(int));
        memcpy(ans[(*returnSize) - 1], ele, rear * sizeof(int));
        int temp = rear;
        returnColumnSize[(*returnSize) - 1] = &temp;
        return;
    }
    if (root->left != NULL)
    {
        ele[rear++] = root->val;
        dfs(root->left, sum, cur + root->val, ele, returnSize, returnColumnSize);
        rear--;
    }
    if (root->right != NULL)
    {
        ele[rear++] = root->val;
        dfs(root->right, sum, cur + root->val, ele, returnSize, returnColumnSize);
        rear--;
    }
}

int **pathSum(struct TreeNode *root, int sum, int *returnSize, int **returnColumnSizes)
{
    if (root == NULL)
    {
        return NULL;
    }
    rear = 0;
    ans = (int **)malloc(2001 * sizeof(int *));
    int *ele = (int *)malloc(2001 * sizeof(int));
    dfs(root, sum, 0, ele, returnSize, returnColumnSizes);
}

// int** ret;
// int retSize;
// int* retColSize;

// int* path;
// int pathSize;

// void dfs(struct TreeNode* root, int sum) {
//     if (root == NULL) {
//         return;
//     }
//     path[pathSize++] = root->val;
//     sum -= root->val;
//     if (root->left == NULL && root->right == NULL && sum == 0) {
//         int* tmp = malloc(sizeof(int) * pathSize);
//         memcpy(tmp, path, sizeof(int) * pathSize);
//         ret[retSize] = tmp;
//         retColSize[retSize++] = pathSize;
//     }
//     dfs(root->left, sum);
//     dfs(root->right, sum);
//     pathSize--;
// }

// int** pathSum(struct TreeNode* root, int sum, int* returnSize, int** returnColumnSizes) {
//     ret = malloc(sizeof(int*) * 2001);
//     retColSize = malloc(sizeof(int) * 2001);
//     path = malloc(sizeof(int) * 2001);
//     retSize = pathSize = 0;
//     dfs(root, sum);
//     *returnColumnSizes = retColSize;
//     *returnSize = retSize;
//     return ret;
// }
