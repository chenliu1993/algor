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

int *reconstructArray(int *a, int start, int end)
{
    int *b = (int *)malloc((end - start + 1) * sizeof(int));
    for (int i = start; i <= end; i++)
    {
        b[i - start] = a[i];
    }
    return b;
}

struct TreeNode *buildTree(int *inorder, int inorderSize, int *postorder, int postorderSize)
{
    if (inorder == NULL || inorderSize == 0 || postorder == NULL || postorderSize == 0)
    {
        return NULL;
    }
    struct TreeNode *node = (struct TreeNode *)malloc(sizeof(struct TreeNode));
    node->val = postorder[postorderSize - 1];
    for (int i = 0; i < inorderSize; i++)
    {
        if (inorder[i] == postorder[postorderSize - 1])
        {
            int *inorderLeft = reconstructArray(inorder, 0, i - 1);
            int *inorderRight = reconstructArray(inorder, i + 1, inorderSize - 1);
            int *postorderLeft = reconstructArray(postorder, 0, i - 1);
            int *postorderRight = reconstructArray(postorder, i, postorderSize - 2);
            node->left = buildTree(inorderLeft, i, postorderLeft, i);
            node->right = buildTree(inorderRight, inorderSize - 1 - i, postorderRight, postorderSize - 1 - i);
        }
    }
    return node;
}