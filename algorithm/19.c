#include <stdio.h>
#include <stdlib.h>
#include <string.h>
struct ListNode
{
    int val;
    struct ListNode *next;
};
// struct ListNode *removeNthFromEnd(struct ListNode *head, int n)
// {
//     struct ListNode *h = head, *ptr1 = head, *ptr2, *ptr3, *temp;
//     int count = 0;
//     ptr2 = ptr1->next;
//     ptr3 = head;
//     while (ptr3 != NULL)
//     {
//         ptr3 = ptr3->next;
//         count++;
//     }
//     n = count - n;
//     while (n)
//     {
//         ptr2 = ptr1;
//         ptr1 = ptr1->next;
//         n--;
//     }
//     temp = ptr1->next;
//     ptr2->next = temp;
//     ptr1->next = NULL;
//     return h;
// }
struct ListNode *removeNthFromEnd(struct ListNode *head, int n)
{
    struct ListNode *ptr1 = head, *ptr2, *h;
    h = (struct ListNode *)malloc(sizeof(struct ListNode));
    h->val = 0;
    h->next = head;
    int rear = 0;
    struct ListNode *stack = (struct ListNode *)calloc(1000, sizeof(struct ListNode));
    while (ptr1 != NULL)
    {
        stack[rear++] = *ptr1;
        ptr1 = ptr1->next;
    }
    rear--;
    while (n)
    {
        n--;
        rear--;
    }
    ptr1 = &stack[--rear];
    ptr1->next = ptr1->next->next;
    struct ListNode *ans = h->next;
    free(h);
    return ans;
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
