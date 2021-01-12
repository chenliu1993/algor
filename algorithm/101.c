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

int Symmetric(struct TreeNode *left, struct TreeNode *right)
{
    if (left == NULL && right == NULL)
    {
        return 1;
    }
    if (left != NULL && right == NULL)
    {
        return 0;
    }
    if (left == NULL && right != NULL)
    {
        return 0;
    }
    if (left->val != right->val)
    {
        return 0;
    }
    return Symmetric(left->left, right->right) && Symmetric(left->right, right->left);
}

int isSymmetric(struct TreeNode *root)
{
    if (root == NULL)
    {
        return 1;
    }
    return Symmetric(root->left, root->right);
}