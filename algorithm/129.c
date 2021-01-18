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

int getCount(struct TreeNode *root, int prevSum)
{
    if (root == NULL)
    {
        return 0;
    }
    int sum = prevSum * 10 + root->val;
    if (root->left == NULL && root->right == NULL)
    {
        return sum;
    }
    else
    {
        return getCount(root->left, sum) + getCount(root->right, sum);
    }
}

int sumNumbers(struct TreeNode *root)
{
    return getCount(root, 0);
}