func generate(numRows int) [][]int {

        /*
        1.Create Slice 1,2, 3 for template,4 save
        2.2 for loop one traverser through n another traverse through horizontal
        3.insert 1 in front and back other just traverse thourgh list done.
        */
        
        var slice2D [][]int

        slice1 := []int{1}
        slice2 := []int{1, 1}
        slice3 := []int{}

        slice2D = append(slice2D, slice1)
    
        if numRows == 1{
            return slice2D
        }
        
        slice2D = append(slice2D, slice2)
    
        if numRows == 2{
            return slice2D
        }
        
    
        //Start from 3 and treverser through the list save repeat
        
        for i := 1; i < numRows - 1 ; i ++ {
        
            thirdElement := slice2D[i]
            length := len(thirdElement)
            
            slice3 = append(slice3, 1)
            
            for j := 0; j < length - 1 ; j++{
                
                sum := slice2D[i][j] + slice2D[i][j+1]
                slice3 = append(slice3, sum)
                
            }
            
            slice3 = append(slice3, 1)
            slice2D = append(slice2D, slice3)
            
            slice3 = []int{}
        
    }
        
        return slice2D
}