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

struct TreeNode **generateTrees(int n, int *returnSize)
{
    if (n == 0)
    {
        (*returnSize) = 0;
        return NULL;
    }
    return buildTrees(1, n, returnSize);
}

struct TreeNode **buildTrees(int start, int end, int *returnSize)
{
    if (start > end)
    {
        (*returnSize) = 1;
        struct TreeNode **rc = (struct TreeNode **)malloc(sizeof(struct TreeNode *));
        rc[0] = NULL;
        return rc;
    }
    *returnSize = 0;
    struct TreeNode *curTreeNode = (struct TreeNode *)malloc(sizeof(struct TreeNode));
    struct TreeNode **allTrees = malloc(0);
    for (int i = start; i <= end; i++)
    {
        int leftTreeSize = 0;
        struct TreeNode *leftTree = buildTree(start, i - 1, &leftTreeSize);
        int rightTreeSize = 0;
        struct TreeNode *rightTree = buildTree(i + 1, end, &rightTreeSize);
        for (int left = 0; left < leftTreeSize; left++)
        {
            for (int right = 0; right < rightTreeSize; right++)
            {
                curTreeNode->val = i;
                curTreeNode->left = &leftTree[left];
                curTreeNode->right = &rightTree[right];
                (*returnSize)++;
                allTrees = realloc(allTrees, sizeof(struct TreeNode *) * (*returnSize));
                allTrees[(*returnSize) - 1] = curTreeNode;
            }
        }
        free(leftTree);
        free(rightTree);
    }
    return allTrees;
}