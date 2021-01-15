#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>
#include <ctype.h>
#include <math.h>

struct Node
{
    int val;
    struct Node *left;
    struct Node *right;
    struct Node *next;
};

#define NODES_SIZE 4097
struct Node *genNext(struct Node *root)
{
    if (root == NULL)
    {
        return root;
    }
    int front, rear, curRear;
    struct Node **queue = (struct Node **)malloc(NODES_SIZE * sizeof(struct Node *));
    front = 0;
    rear = 0;
    queue[rear++] = root;
    while (front != rear)
    {
        curRear = rear;
        for (int i = front; i < curRear; i++)
        {
            struct Node *cur = queue[i];
            if (cur->left != NULL)
            {
                queue[rear++] = cur->left;
                queue[rear++] = cur->right;
            }
            if (i != curRear - 1)
            {
                queue[i]->next = queue[i + 1];
            }
        }
        queue[curRear - 1]->next = NULL;
        front = curRear;
    }
    return root;
}

struct Node *connect(struct Node *root)
{
    return genNext(root);
}