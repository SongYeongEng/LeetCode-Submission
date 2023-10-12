class Solution(object):
    def twoSum(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """
        
        my_hashmap = {}
        
        for i in range(len(nums)):
            my_hashmap[nums[i]] = i
        
        for i in range(len(nums)):
            complement = target - nums[i]
            
            if complement in my_hashmap and my_hashmap[complement] !=i:
                return [i, my_hashmap[complement]]
            
            