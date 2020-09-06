#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <sys/socket.h>
#include <netinet/in.h>
#include <arpa/inet.h>

void main() {
    char *src;
    const char *ch;
    // register int rc;
    struct in_addr addrToConv;
    src = (char *)malloc(20 *sizeof(src));
    memset(&addrToConv, 0, sizeof(addrToConv));
    scanf("%s", &addrToConv.s_addr);
    // rc = inet_pton(AF_INET,src,(void *)&addrToConv);
    // if (rc>0) {
    //     printf("addr converted result is %lu\n", addrToConv.s_addr);
    // }else{
    //     perror("wrong converted");
    // }
    ch = inet_ntop(AF_INET, (void *)&addrToConv.s_addr, src, INET_ADDRSTRLEN);
    if (ch==NULL){
        perror("wrong converted");
    }else{
        printf("addr converted result is: %s\n", src);
    }
    free(src);
    return;
}