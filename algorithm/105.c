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

int *constructArray(int *a, int start, int end)
{
    int len = end - start + 1;
    int *rc = (int *)calloc(len, sizeof(int));
    int j = start;
    for (int i = 0; i < len; i++)
    {
        rc[i] = a[j++];
    }
    return rc;
}

struct TreeNode *buildTree(int *preorder, int preorderSize, int *inorder, int inorderSize)
{
    if (preorder == NULL || preorderSize == 0 || inorder == NULL || inorderSize == 0)
    {
        return NULL;
    }
    struct TreeNode *root = (struct TreeNode *)malloc(sizeof(struct TreeNode));
    root->val = preorder[0];
    for (int i = 0; i < preorderSize; i++)
    {
        if (inorder[i] == preorder[0])
        {
            int *preFirst = constructArray(preorder, 1, i);
            int *preLast = constructArray(preorder, i + 1, preorderSize - 1);
            int *inFirst = constructArray(inorder, 0, i - 1);
            int *inLast = constructArray(inorder, i + 1, inorderSize - 1);
            root->left = buildTree(preFirst, i, inFirst, i);
            root->right = buildTree(preLast, preorderSize - 1 - i, inLast, inorderSize - 1 - i);
        }
    }
    return root;
}