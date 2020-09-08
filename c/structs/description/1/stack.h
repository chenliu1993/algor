#ifndef _STACK_H
#define _STACK_H
#define arraysize(a) \
    (sizeof(a)/sizeof*(a))

typedef struct _ListNode {
    int val;
    struct _ListNode *next;
} ListNode;

ListNode *Push(ListNode *head, int val);
ListNode *CreateStack(int a[]);
ListNode *Pop(ListNode *head);
int isEmpty(ListNode *head);

#endif /* _STACK_H */


