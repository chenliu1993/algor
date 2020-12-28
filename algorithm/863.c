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
#define MAX_TREE_SIZE 501
struct TreeNode **father;
int flag[MAX_TREE_SIZE] = {0};
int count;

void GetFather(struct TreeNode *node, struct TreeNode **father)
{
    if (node->left != NULL)
    {
        father[node->left->val] = node;
        GetFather(node->left, father);
    }

    if (node->right != NULL)
    {
        father[node->right->val] = node;
        GetFather(node->right, father);
    }
}

void GetResult(struct TreeNode *node, int distance, int k, int *rc)
{
    if (distance == k)
    {
        rc[count++] = node->val;
        return;
    }
    if (node->left != NULL && flag[node->left->val] == 0)
    {
        flag[node->left->val] = 1;
        GetResult(node->left, distance + 1, k, rc);
    }
    if (node->right != NULL && flag[node->right->val] == 0)
    {
        flag[node->right->val] = 1;
        GetResult(node->right, distance + 1, k, rc);
    }
    if (father[node->val] != NULL && flag[father[node->val]->val] == 0)
    {
        flag[father[node->val]->val] = 1;
        GetResult(father[node->val], distance + 1, k, rc);
    }
}

int *distanceK(struct TreeNode *root, struct TreeNode *target, int K, int *returnSize)
{
    if (root == NULL || target == NULL || K < 0 || K > 1000)
    {
        *returnSize = 0;
        return NULL;
    }

    father = (struct TreeNode **)malloc(MAX_TREE_SIZE * sizeof(struct TreeNode *));
    if (father != NULL)
    {
        memset(father, 0, MAX_TREE_SIZE * sizeof(struct TreeNode *));
    }
    memset(flag, 0, MAX_TREE_SIZE * sizeof(int));
    int *result = (int *)malloc(MAX_TREE_SIZE * sizeof(int));
    if (result != NULL)
    {
        memset(result, -1, MAX_TREE_SIZE * sizeof(int));
    }
    GetFather(root, father);
    count = 0;
    flag[target->val] = 1;
    GetResult(target, 0, K, result);

    *returnSize = count;
    return result;
}