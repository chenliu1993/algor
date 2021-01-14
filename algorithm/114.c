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

int rear;
void getPreOrder(struct TreeNode *root, struct TreeNode ***link)
{
    if (root == NULL)
    {
        return;
    }
    rear++;
    (*link) = realloc((*link), rear * sizeof(struct TreeNode *));
    (*link)[rear - 1] = (struct TreeNode *)malloc(sizeof(struct TreeNode));
    (*link)[rear - 1] = root;
    getPreOrder(root->left, link);
    getPreOrder(root->right, link);
}
void flatten(struct TreeNode *root)
{
    struct TreeNode **link = (struct TreeNode **)malloc(0);
    rear = 0;
    getPreOrder(root, link);
    for (int i = 1; i < rear; i++)
    {
        struct TreeNode *node = link[i - 1];
        struct TreeNode *cur = link[i];
        node->left = NULL;
        node->right = cur;
    }
    free(link);
}