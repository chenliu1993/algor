#include <stdio.h>
#include <stdlib.h>
#define arraysize(a) \
    (sizeof(a)/sizeof*(a))
/* https://www.cnblogs.com/yorkyang/p/10876604.html */
typedef struct _ListNode {
    int val;
    struct _ListNode *next;
    // This is a c++ feature like new.
    // _ListNode(int a): val(a), next(NULL) {}
} ListNode;

ListNode *
initList(int a[], int len) {
    ListNode *head, *p;
    p = (ListNode *)malloc(sizeof(ListNode*));
    p->val = a[0];
    p->next = (ListNode *)malloc(sizeof(ListNode*));
    head = p;
    p=p->next;
    for(int i=1;i<len;i++) {
        p->val = a[i];
        p->next = (ListNode *)malloc(sizeof(ListNode*));
        p = p->next;
    }
    p=NULL;
    return head;
}

int
existLoop(ListNode *head) {
    ListNode *slow, *fast;
    slow=fast=head;
    while(slow!=NULL&&fast->next!=NULL) {
        slow=slow->next;
        fast=fast->next->next;
        if(slow == fast) {
            if (fast->next==NULL) {
                return 0;
            } else{
                return 1;
            }      
        }
    }
    return 0;
}

int
twoListCross(ListNode *h1, ListNode *h2) {
    ListNode *slow;
    slow=h1;
    while(slow->next!=NULL){
        slow=slow->next;
    }
    slow->next = h1->next;
    return existLoop(h2);
}

ListNode *
findLoopStart(ListNode *head) {
    ListNode *slow, *fast, *ptr1, *ptr2;
    slow=fast=head;
    ptr2=head;
    while(slow!=NULL&&fast->next!=NULL) {
        slow=slow->next;
        fast=fast->next->next;
        if(slow==fast){
            if (fast->next == NULL) {
                return NULL;
            }else{
                ptr1=slow;
                break;
            }
        }
    }
    while(ptr1!=ptr2){
        ptr1=ptr1->next;
        ptr2=ptr2->next;
    }
    return ptr1;
}

int
lenOfLoop1(ListNode *head) {
    ListNode *start, *p;
    int len, lenStart;
    start = findLoopStart(head);
    len=lenStart=0;
    p=head;
    while(p!=NULL) {
        p=p->next;
        len += 1;
        if(p==start){
            lenStart = len;
        }
    }
    return len-lenStart+1;
}

int
lenOfLoop2(ListNode *head) {
    ListNode *start, *slow, *fast;
    int len;
    start = findLoopStart(head);
    if (start==NULL) {return 0;}
    slow=fast=start;
    len=0;
    while(slow!=NULL&&fast->next!=NULL) {
        slow=slow->next;
        fast=fast->next->next;
        len +=1;
        if(slow==fast){
            break;
        }
    }
    return len;
}

ListNode *
findFastNodeOnLoop(ListNode *head, ListNode *nodeOnLoop) {
    ListNode *slow, *fast;
    slow=fast=nodeOnLoop;
    while(fast->next!=NULL){
        slow=slow->next;
        fast=fast->next->next;
        if(fast==nodeOnLoop || fast->next==nodeOnLoop) {
            return slow;
        }
    }
    // logically this will never be touched.
    return NULL;
}

ListNode *
headInsert(ListNode *head, int a[]) {
    ListNode *p;
    for(int i=0;i<arraysize(a);i++){
        p=(ListNode *)malloc(sizeof(ListNode*));
        p->val=a[i];
        p->next = head->next;
        head->next=p;
    }
    return head;
}

ListNode *
rearInsert(ListNode *head, int a[]) {
    ListNode *p, *end;
    p=head;
    while(p->next!=NULL) {
        p=p->next;
    }
    end = p;
    for(int i=0;i<arraysize(a);i++){
        p=(ListNode *)malloc(sizeof(ListNode*));
        p->val=a[i];
        end->next=p;
        end=p;
    }
    return head;
}