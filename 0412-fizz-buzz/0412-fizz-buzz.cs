public class Solution {
    public IList<string> FizzBuzz(int n) {
        
        string[] fb = new string[n];
        
        for (int i = 1; i <=n ; i++){
            
            if (i % 3 == 0 && i % 5 == 0)
            {
                fb [i-1] = "FizzBuzz";
            }
            else if (i % 3 == 0)
            {
                fb [i-1] = "Fizz";
            }
            else if (i % 5 == 0)
            {
                fb [i-1] = "Buzz";
            }
            else
            {
                fb [i-1] = i.ToString();
            }
                
            
        }
        
        return fb;
        
    }
    
    
}