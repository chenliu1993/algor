#include <stdio.h>
#include <stdlib.h>
#define arraysize(a) \
    (sizeof(a)/sizeof*(a))
typedef struct _ListNode {
    int  val;
    struct _ListNode *next;
    _ListNode(int a): val(a), next(NULL) {}
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

ListNode *
findLastKth(ListNode *l, int k) {
    ListNode *h1,*h2;
    h1 = h2 = l;
    for(int i=0;i<=k;i++){
        h1 = h1->next;
    }
    while(h1!=NULL){
        h2=h2->next;
        h1=h1->next;
    }
    return h2;
} 

int
main(int argc, char *argv[]){
    int k = 3;
    int a[8] = {4,-3,5,-2,-1,2,6,-2};
    ListNode *h = initList(a, arraysize(a));
    ListNode *tgt = findLastKth(h, k);
    printf("the last %dth node is %d", k, tgt->val);
    return 0;
}
