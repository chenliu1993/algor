#include <stdio.h>
#include <stdlib.h>
#include "tree.h"

// TODO: maybe rotate?
TreeNode *
MakeEmpty(TreeNode *root) {
    if (root != NULL) {
        MakeEmpty(root->left);
        MakeEmpty(root->right);
        free(root);
    }
    return NULL;
}

TreeNode *
Find(TreeNode *root, int val) {
    if (root ==NULL){
        return NULL;
    }
    if (val<root->val) {
       return Find(root->left, val);
    }else if(val >root->val){
        return Find(root->right, val);
    }
    return root;
}

TreeNode *
FindMin(TreeNode *root) {
    if (root==NULL) {
        return NULL;
    }else if(root->left==NULL) {
        return root;
    }
    return FindMin(root->left);
}

TreeNode *
FindMax(TreeNode *root) {
    if (root==NULL) {
        return NULL;
    }else if(root->right == NULL) {
        return root;
    }
    return FindMax(root->right);
}

TreeNode *
Insert(TreeNode *root, int val) {
    if(root==NULL){
        root=(TreeNode *)malloc(sizeof(TreeNode*));
        root->val = val;
        root->left=NULL;
        root->right=NULL;
        return root;
    }else{
        if(val<root->val) {
            root->left=Insert(root->left, val);
        }
        if(val>root->val) {
            root->right=Insert(root->right, val);
        }
    }
    return root;
}

TreeNode *
Delete(TreeNode *root, int val) {
    TreeNode *temp;
    if(root==NULL) {
        return NULL;
    }else if(val<root->val){
        root->left=Delete(root->left, val);
    }else if(val>root->val){
        root->right=Delete(root->right, val);
    }else if(root->left && root->right) {
        temp =FindMin(root->right);
        root->val=temp->val;
        root->right=Delete(root->right, root->val);
    }else{
        temp = root;
        if(root->left==NULL) {
            root=root->left;
        }else if(root->right==NULL) {
            root=root->right;
        }
        free(temp);
    }
    return root;
}
