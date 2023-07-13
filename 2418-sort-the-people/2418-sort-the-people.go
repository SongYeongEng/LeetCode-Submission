func sortPeople(names []string, heights []int) []string {
        hashTable := make(map[int]string)
    
        var slice []string

        for i := 0; i < len(heights); i++ {
        key := heights[i]
        value := names[i]
        hashTable[key] = value
    }
    
    // Extract keys into a slice
	var keys []int
	for key := range hashTable {
		keys = append(keys, key)
	}   
    
    sort.Slice(keys, func(i, j int) bool {
		return keys[i] > keys[j]
	})
    
    
    for _, key := range keys {
		value := hashTable[key]
        slice = append(slice,value)
	}
    
    
    return slice
     
}