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

int dfs(struct TreeNode *root, long long low, long long high)
{
    if (root == NULL)
    {
        return 1;
    }
    if (root->val <= low || root->val >= high)
    {
        return 0;
    }
    return dfs(root->left, low, root->val) && dfs(root->right, root->val, high);
}

int isValidBST(struct TreeNode *root)
{
    return dfs(root, LONG_MIN, LONG_MAX);
}