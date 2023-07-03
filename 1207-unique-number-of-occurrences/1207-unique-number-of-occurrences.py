class Solution(object):
    def uniqueOccurrences(self, arr):
        """
        :type arr: List[int]
        :rtype: bool
        """
        
        #how pust isnside dictionarry
        
        ihash = {}
        value = 0
        ilist = []
        value2 =0
        
        #put all into hash
        
        for num in arr:
            
            #if exist the nadd
            if num in ihash:
                value = ihash[num] + 1
                ihash[num] = value
            else:
                ihash[num] = 1
        value = 0
        
        for value in ihash.values():
            if value in ilist:
                return False
            ilist.append(value)
          
            
        return True
    
       