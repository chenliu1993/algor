
from itertools import groupby

# Option 2

# Python3 code to demonstrate working of 
# Values Frequency grouping of K in dictionaries
# Using groupby() + dictionary comprehension
from itertools import groupby
 
# initializing list
test_list = [{'gfg' : 3, 'best' : 4}, {'gfg' : 3, 'best' : 5}, 
             {'gfg' : 4, 'best' : 4}, {'gfg' : 7, 'best' : 4} ]
 
# printing original list
print("The original list is : " + str(test_list))
 
# initializing K 
K = 'gfg'
 
# groupby() used to group values and len() to compute Frequency
# print(groupby(test_list, lambda sub: sub[K]))
res = [{key: len(list(val))} for key, val in groupby(test_list, lambda sub: sub[K])]
 
 
for key, val in groupby(test_list, lambda sub: sub[K]):
    print(key, list(val))
    
# printing result 
print("The Values Frequency : " + str(res))
   
   
    

    
    