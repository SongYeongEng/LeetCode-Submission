class Solution(object):
    def romanToInt(self, s):
        """
        :type s: str
        :rtype: int
        """
        
        
        hash_map = { 'I':1,'V':5,'X':10,'L':50,'C':100,'D':500,'M':1000 }
        
        #sum and tmp
        
        total = 0
        tmp = 0
        tmp2 = 0
        value = 0
        j = 1
        #understand question
        
   
        
        #Largest to lowest
        #How do i know what come first
        if(len(s) == 1):
            
            total = hash_map[s[0]]
            return total
        
        #iterate through the list
        for i in range(len(s)-1):
        #compare currenty with next if lower then minus
            
                tmp = hash_map[s[i]]
                tmp2 = hash_map[s[i+1]]
                
                
                #  use two to check if smaller only add
                
                if tmp >= tmp2:
                    total += tmp
                elif tmp2 > tmp:
                    total -= tmp
        
        #final case wont be bigger so just add 
        total += hash_map[s[-1]]
        
        return total        
                
                