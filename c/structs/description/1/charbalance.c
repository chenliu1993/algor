#include <stdio.h>
#include <stdlib.h>

#include "stack.h"
int
checkBalance(char str[]) {
    ListNode *p1, *p2;
    p1= (ListNode *)malloc(sizeof(ListNode*));
    p2= (ListNode *)malloc(sizeof(ListNode*));
    p1->val=0;
    p1->next=NULL;
    for(int i=0;i<arraysize(str);i++) {
        if (str[i]=='('||str[i]=='['||str[i]=='{') {
            p1 = Push(p1, str[i]-'0');
        }
        if ((str[i]==')'||str[i]==']'||str[i]=='}') && isEmpty(p1)) {
            goto freePtr;
            return 0;
        }
        if(str[i]==')'){
            p2=Pop(p1);
            if(p2->val=='('-'0') {
                goto freePtr;
                return 1;
            }else{
                goto freePtr;
                return 0;
            }
        }
        if(str[i]==']'){
            p2=Pop(p1);
            if(p2->val=='['-'0') {
                goto freePtr;
                return 1;
            }else{
                goto freePtr;
                return 0;
            }
        }
        if(str[i]=='}'){
            p2=Pop(p1);
            if(p2->val=='{'-'0') {
                goto freePtr;
                return 1;
            }else{
                goto freePtr;
                return 0;
            }
        }
    }
freePtr:
    free(p1);
    free(p2);
    return 1;
}

int
main(int argc, char *argv[]) {
    char str[8]={'2','*','(','5','+','1','2',')'};
    int rc=checkBalance(str);
    printf("%d\n",rc);
    return 0;
}