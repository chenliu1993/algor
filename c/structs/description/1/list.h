#ifndef _LIST_H

typedef struct _ListNode {
    int val;
    ListNode *next;
} ListNode;

ListNode *newNode(int val);
void makeEmpty(ListNode *tgt);
int isEmpty(ListNode *tgt);
int isLast(ListNode *tgt);
ListNode *findEle(ListNode *tgt, int val);
ListNode *findPrev(ListNode *tgt, int val);
void Delete(ListNode *tgt, int val);
void Insert(ListNode *tgt, int val, int posi);

#endif /* _LIST_H */