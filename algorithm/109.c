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

struct ListNode
{
    int val;
    struct ListNode *next;
};

struct ListNode *getMid(struct ListNode *left, struct ListNode *right)
{
    struct ListNode *fast = left;
    struct ListNode *slow = left;
    while (fast != right && fast->next != right)
    {
        fast = fast->next->next;
        slow = slow->next;
    }
    free(fast);
    return slow;
}

struct TreeNode *buildTree(struct ListNode *left, struct ListNode *right)
{
    if (left == right)
    {
        return NULL;
    }
    struct ListNode *mid = getMid(left, right);
    struct TreeNode *root = (struct TreeNode *)malloc(sizeof(struct TreeNode));
    root->val = mid->val;
    root->left = NULL;
    root->right = NULL;
    root->left = buildTree(left, mid);
    root->right = buildTree(mid->next, right);
    return root;
};

struct TreeNode *sortedListToBST(struct ListNode *head)
{
    return buildTree(head, NULL);
}