#include <stdio.h>
#include <stdlib.h>
#include <string.h>
/**
 * Definition for singly-linked list.
 */
struct ListNode
{
    int val;
    struct ListNode *next;
};

struct ListNode *swapPairs(struct ListNode *head)
{
    struct ListNode *first, *p;
    first = (struct ListNode *)calloc(1, sizeof(struct ListNode));
    first->next = head;
    p = first;
    while (head && head->next)
    {
        first->next = head->next;
        struct ListNode *h = head->next->next;
        first->next->next = head;
        head->next = h;
        first = head;
        head = h;
    }
    return p->next;
}