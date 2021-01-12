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
int isSameTree(struct TreeNode *p, struct TreeNode *q)
{
    if (p == NULL && q == NULL)
    {
        return 1;
    }
    if (p == NULL && q != NULL)
    {
        return 0;
    }
    if (p != NULL && q == NULL)
    {
        return 0;
    }
    if (p->val != q->val)
    {
        return 0;
    }
    return isSameTree(p->left, q->left) && isSameTree(p->right, q->right);
}