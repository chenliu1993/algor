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

#define MAX(x, y) (x < y ? y : x)

int rob(struct TreeNode *root)
{
    if (root == NULL)
    {
        return 0;
    }
    int res1 = root->val;
    if (root->left != NULL)
    {
        res1 += (rob(root->left->left) + rob(root->left->right));
    }
    if (root->right != NULL)
    {
        res1 += (rob(root->right->left) + rob(root->right->right));
    }

    int res2 = rob(root->left) + rob(root->right);
    return MAX(res1, res2);
}
