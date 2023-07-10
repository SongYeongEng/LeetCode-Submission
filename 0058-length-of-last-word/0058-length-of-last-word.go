func lengthOfLastWord(s string) int {
    
    
    count := 0
    
    for i := len(s) -1; i >= 0; i--{
        
        if s[i] !=  ' '{
            
            for j:= i;j >= 0;j--{
                
                fmt.Print("Success")
                
                if s[j] ==  ' '{
                    
                    return count
                    
                }
                
               count ++
                
                
                
            }
        
            return count
        }
        
        
    }
    return count
    
}