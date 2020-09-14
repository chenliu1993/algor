#include <stdio.h>
#include <stdlib.h>
#define MAXSIZE 20000
typedef struct TreeNode
{
    int val;
    struct TreeNode *left;
    struct TreeNode *right;
} TreeNode;

TreeNode *
CreateTree(int *a, int asize, int start)
{
    TreeNode *root = (TreeNode *)malloc(sizeof(TreeNode));
    if (start >= asize)
    {
        return NULL;
    }
    else
    {
        root->val = a[start];
        root->left = CreateTree(a, asize, start + 1);
        root->right = CreateTree(a, asize, start + 2);
    }
    return root;
}

void iterateTree(TreeNode *root)
{
    if (root != NULL)
    {
        printf("%d,", root->val);
    }
    else
    {
        printf("\n");
        return;
    }
    iterateTree(root->left);
    iterateTree(root->right);
}

/* https://blog.csdn.net/lingchen2348/article/details/52774535 */

void iterateTreeAsLevel(TreeNode *root, TreeNode **front, TreeNode **rear)
{
    *front = root;
    if ((*front)->left != NULL)
    {
        rear = rear + 1;
        *rear = (*front)->left;
    }
    if ((*front)->right != NULL)
    {
        rear = rear + 1;
        *rear = (*front)->right;
    }
    if (front != rear)
    {
        front = front + 1;
        iterateTreeAsLevel(*front, front, rear);
    }
}

int main(int argc, char *argv[])
{
    int asize = 5;
    int a[5] = {3, 9, 20, 15, 7};
    TreeNode *root = CreateTree(a, 5, 0);
    iterateTree(root);
    TreeNode **h, **r;
    TreeNode *queue[asize];
    h = r = queue;
    iterateTreeAsLevel(root, h, r);
    for (int i = 0; i < asize; i++)
    {
        printf("%d,", queue[i]->val);
    }
    printf("\n");
    return 0;
}
