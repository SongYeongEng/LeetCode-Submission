func relativeSortArray(arr1 []int, arr2 []int) []int {
    
   //All 2 is in 1
    
    hashTable := make(map[int]int)
    
    var slice []int
    
    for _, value := range arr1 {
		
        if hashTable[value] != 0{
            
            hashTable[value] += 1
            
        }else{
            hashTable[value] = 1
            
        }
        
	}
    
    for _, value := range arr2{
        
        j := hashTable[value]
        k := 0 
        for k < j{
            
            slice = append(slice,value)
            k++
        }
        delete(hashTable, value)
    }
    
    
        //Get key then sort
    keys := make([]int, 0, len(hashTable))
        for key := range hashTable {
            keys = append(keys, key)
        }
    
    sort.Ints(keys)

    
    for _, value := range keys{
        
        j := hashTable[value]
        k := 0 
        for k < j{
            
            slice = append(slice,value)
            k++
        }
        
    }
    
    

    
    return slice
}