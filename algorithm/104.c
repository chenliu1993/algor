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

// needs to use fmax to reduce time overlimit
#define MAX(x, y) (x < y ? y : x)
int maxDepth(struct TreeNode *root)
{
    if (root == NULL)
    {
        return 0;
    }
    return MAX(maxDepth(root->left), maxDepth(root->right)) + 1;
}