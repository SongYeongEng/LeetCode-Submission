func kWeakestRows(mat [][]int, k int) []int {
    
    // return number of soulder compare to k return lmao if same then use ij
    //how to conunt how many
    
   // how mayn row are there
    //how to acess [][]golang
    
    var record []int
    var result []int
    var count int
    
     for _, row := range mat {
        for _, num := range row {
            if num == 1 {
            count++
            }
        }
         
         record = append(record,count)
         count = 0
    }
    
    
    //sorting
    //sort it check if same number start from zero if yes return weakest
    
    sorted := make([]int, len(record))
	copy(sorted, record)
    sort.Ints(sorted)
    
	fmt.Println(sorted)
    fmt.Println(record)

  for i := 0; i < k; i++ {
      for j:= 0; j< len(record);j++{
          outerLoopFlag := false
          if sorted[i] == record [j]{
             
            
            for _, value := range result {
                if value == j {
                    outerLoopFlag = true
                }
            }
              
              if outerLoopFlag {
                  continue
              } else{
             result = append(result,j)
             break;
              }
              

          }
      }
}

    return result
    
    
}