#ifndef _TREE_H
#define _TREE_H
/* This is a binary search tree */
typedef struct _TreeNode {
    int val;
    struct _TreeNode *left;
    struct _TreeNode *right;
} TreeNode;
TreeNode *MakeEmpty(TreeNode *root);
TreeNode *Find(TreeNode *root, int val);
TreeNode *FindMin(TreeNode *root);
TreeNode *FindMax(TreeNode *root);
TreeNode *Insert(TreeNode *root, int val);
TreeNode *Delete(TreeNode *root, int val);
#endif