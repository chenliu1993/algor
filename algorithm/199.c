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
int *ans;
#define MAXSIZE 2001
void getRightPoints(struct TreeNode *root, int *returnSize)
{
    if (root == NULL)
    {
        return;
    }
    int front, rear, curRear;
    struct TreeNode **queue = (struct TreeNode **)malloc(MAXSIZE * sizeof(struct TreeNode *));
    front = 0;
    rear = 0;
    queue[rear++] = root;
    while (front != rear)
    {
        curRear = rear;
        for (int i = front; i < curRear; i++)
        {
            struct TreeNode *cur = queue[i];
            if (cur->left != NULL)
            {
                queue[rear++] = cur->left;
            }
            if (cur->right != NULL)
            {
                queue[rear++] = cur->right;
            }
            if (i == curRear - 1)
            {
                (*returnSize)++;
                ans[(*returnSize) - 1] = cur->val;
            }
        }
        front = curRear;
    }
}

int *rightSideView(struct TreeNode *root, int *returnSize)
{
    ans = (int *)malloc(MAXSIZE * sizeof(int));
    (*returnSize) = 0;
    getRightPoints(root, returnSize);
    return ans;
}