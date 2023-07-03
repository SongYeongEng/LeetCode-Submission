class Solution(object):
    def findDifference(self, nums1, nums2):
        """
        :type nums1: List[int]
        :type nums2: List[int]
        :rtype: List[List[int]]
        """
        
        #Not yet learn space power
        
        l1 = []
        l2 = []
        l3 = []
        
        for num in nums1:
            if num not in nums2:
                if num not in l1:
                    l1.append(num)

                

        for num in nums2:
            if num not in nums1:
                 if num not in l2:
                    l2.append(num)

                
        l3=[l1,l2]
        
        return l3