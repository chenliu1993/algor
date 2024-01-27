test_list = [{'name': 'test1', 'age': 10}, {'name': 'test2', 'age': 20}, {'name': 'test3', 'age': 30},
             {'name': 'test1', 'age': 10}, {'name': 'test1', 'age': 20}, {'name': 'test2', 'age': 60}]


# Option 1 
def organize(arr):
    sub = {}
    for i in range(len(arr)):
        sub[arr[i]['name']]=[]
    
    for i in range(len(arr)):
        sub[arr[i]['name']].append(arr[i]['age'])
    
    return sub

def find_occurrences(dictionary, keyword):
    ans = {}
    target_list = dictionary[keyword]
    for i in range(len(target_list)):
        ans[target_list[i]] = 0

    for i in range(len(target_list)):
        ans[target_list[i]]=ans[target_list[i]]+1
        
    return ans