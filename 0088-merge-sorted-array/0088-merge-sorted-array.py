class Solution(object):
    def merge(self, nums1, m, nums2, n):
        """
        :type nums1: List[int]
        :type m: int
        :type nums2: List[int]
        :type n: int
        :rtype: None Do not return anything, modify nums1 in-place instead.
        """
        
        #Merge sorting
        
        #Sorted ascending
        #nums 1 array got empty space so need replace instead of insterting
        #length = m+n
        
        
        #if = then replace if not insert
        total = m+n
        i = 0
        j = m
        k = 0
        
        value  = 0
        
        #valuie over m is 0
        
        while m in range(len(nums1)):
            nums1.pop()
            print nums1    
    
        #treverse through the list
        #might needt ot use two for loop
        
        for i in range(n):
            nums1.insert(i,nums2[i])
            
        nums1.sort()
        
        