#include <stdio.h>
#include <stdlib.h>
#include <string.h>
struct ListNode
{
    int val;
    struct ListNode *next;
};

struct ListNode *initList(int *a, int asize)
{
    struct ListNode *head = (struct ListNode *)malloc(sizeof(struct ListNode));
    head->val = a[0];
    head->next = NULL;
    struct ListNode *ptr = head;
    for (int i = 1; i < asize; i++)
    {
        struct ListNode *p = (struct ListNode *)malloc(sizeof(struct ListNode));
        p->val = a[i];
        p->next = NULL;
        ptr->next = p;
        ptr = p;
    }
    return head;
}

void printList(struct ListNode *h)
{
    while (h != NULL)
    {
        printf("%d,", h->val);
        h = h->next;
    }
    printf("\n");
}
struct Stack
{
    struct ListNode *val;
    struct Stack *next;
};
struct ListNode *removeNthFromEnd(struct ListNode *head, int n)
{
    struct ListNode *ptr1, *h;
    h = (struct ListNode *)malloc(sizeof(struct ListNode));
    h->val = 0;
    h->next = head;
    ptr1 = h;
    struct Stack *stack = NULL;
    while (ptr1 != NULL)
    {
        struct Stack *tmp = (struct Stack *)malloc(sizeof(struct Stack));
        tmp->val = ptr1;
        tmp->next = stack;
        stack = tmp;
        ptr1 = ptr1->next;
    }
    while (n)
    {
        n--;
        struct Stack *tmp = stack->next;
        free(stack);
        stack = tmp;
    }
    struct ListNode *prev = stack->val;
    prev->next = prev->next->next;
    struct ListNode *ans = h->next;
    free(h);
    return ans;
}

int main(int argc, char *argv[])
{
    int a[5] = {1, 2, 3, 4, 5};
    struct ListNode *head = initList(a, 5);
    printList(head);

    struct ListNode *rc = removeNthFromEnd(head, 2);

    printList(rc);
    return 0;
}
// struct ListNode *removeNthFromEnd(struct ListNode *head, int n)
// {
//     struct ListNode *first, *rear;
//     struct ListNode *dummy = malloc(sizeof(struct ListNode));
//     dummy->val = 0, dummy->next = head;
//     first = head;
//     rear = dummy;
//     while (n)
//     {
//         first = first->next;
//         n--;
//     }
//     while (first != NULL)
//     {
//         first = first->next;
//         rear = rear->next;
//     }
//     rear->next = rear->next->next;
//     struct ListNode *ans = dummy->next;
//     free(dummy);
//     return ans;
// }
