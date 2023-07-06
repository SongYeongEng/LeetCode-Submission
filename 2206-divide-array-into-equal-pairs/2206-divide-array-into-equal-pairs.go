func divideArray(nums []int) bool {
    
    //make paires
    
    //Check %2 = 0 means all even number
    
    //How many paires len list/2
    
  
    HashMap := make(map[int]int)
    
 
    
    
    
    //Put all number into the hash map
    //the nif exist the nonly ad dif not add new number if not update
    
    //use another for loop to check
    
    for value := range nums{
        if HashMap[nums[value]] == 0{
            
            HashMap[nums[value]] = 1
        }else{
            
            HashMap[nums[value]] += 1
        }
        
        
    } 
    
    for key, value := range HashMap {
        
        fmt.Printf("Key: %d, Value: %d\n", key, value)
        
        if value %2 != 0{
            return false
        }
    }
    
    return true
    
}