#include <stdio.h>
#include <stdlib.h>
#include <assert.h>

#include "list.h"

ListNode *
newNode(int val) {
    ListNode *p=(ListNode *)malloc(sizeof(ListNode*));
    p->val = val;
    p->next = NULL;
    return p;
}

int
isEmpty(ListNode *tgt) {
    return tgt->next == NULL;
}

int
isLast(ListNode *tgt) {
    return tgt->next == NULL;
}

ListNode *
findEle(ListNode *tgt, int val) {
    ListNode *p;
    p = tgt->next;
    while(p!=NULL && p->val != val) {
        p=p->next;
    }
    return p;
}


ListNode *
findPrev(ListNode *tgt, int val) {
    ListNode *p;
    p=tgt;
    while(p->next!=NULL&&p->next->val != val) {
        p=p->next;
    }
    return p;
}

void
Delete(ListNode *tgt, int val) {
    ListNode *p, *current;
    p = findPrev(tgt, val);
    if (!isLast(p)) {
        current = p->next;
        p->next = current->next;
        free(current);
    }
}

void
Insert(ListNode *tgt, int val, int posi) {
    ListNode *p, *current;
    current = newNode(0);
    assert(current!=NULL);

    p = findPrev(tgt, val);

    current->val = val;
    current->next=p->next;
    p->next=current;
}
