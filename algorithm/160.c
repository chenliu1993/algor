#include <stdio.h>
#include <stdlib.h>
#include <string.h>
struct ListNode
{
    int val;
    struct ListNode *next;
};
// struct ListNode *getIntersectionNode(struct ListNode *headA, struct ListNode *headB)
// {
//     struct ListNode *fast, *slow, *rear, *ptr;
//     ptr = headA;
//     while (ptr->next != NULL)
//     {
//         ptr = ptr->next;
//     }
//     ptr->next = headA;
//     slow = headB;
//     fast = headB->next->next;
//     while (fast != slow)
//     {
//         slow = slow->next;
//         fast = fast->next->next;
//     }
//     return slow;
// }
struct ListNode *getIntersectionNode(struct ListNode *headA, struct ListNode *headB)
{
    struct ListNode *h1, *h2;
    h1 = headA;
    h2 = headB;
    while (h1 != h2)
    {
        if (h1 != NULL)
        {
            h1 = h1->next;
        }
        else
        {
            h1 = headB;
        }
        if (h2 != NULL)
        {
            h2 = h2->next;
        }
        else
        {
            h2 = headA;
        }
    }
    return h1;
}
