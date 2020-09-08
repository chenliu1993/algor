#include <stdio.h>
#include <stdlib.h>
#include "stack.h"

ListNode *
Push(ListNode *head, int val) {
    ListNode *p;
    p=(ListNode *)malloc(sizeof(ListNode*));
    p->val=val;
    p->next=head->next;
    head->next=p;
    return head;
}

ListNode *
CreateStack(int a[]) {
    ListNode *head;
    head = (ListNode *)malloc(sizeof(ListNode*));
    head->val=0;
    head->next=NULL;
    for(int i=0; i<arraysize(a);i++){
        head = Push(head, a[i]);
    }
    return head;
}

ListNode *
Pop(ListNode *head) {
    if (head->next==NULL){
        return NULL;
    }
    ListNode *tgt;
    tgt = head->next;
    head->next=tgt->next;
    tgt->next=NULL;
    return tgt;
}

int
isEmpty(ListNode *head){
    return head->next == NULL;
}