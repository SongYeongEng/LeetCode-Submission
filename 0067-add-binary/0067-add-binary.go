func addBinary(a string, b string) string {
    
  
    lenB := len(b) - 1
    charsA := []rune(a)
    charsB := []rune(b)

    for i := len(a) - 1; i >= 0; i-- {
        
        if lenB == -1{
            charsB = append([]rune{charsA[i]}, charsB...)
            continue
        }
        
        if charsB[lenB] == '0' {
            charsB[lenB] = charsA[i]
            lenB--
            continue
        }

        if charsB[lenB] == '1' && charsA[i] == '0' {
            lenB--
            continue
        }

        if charsB[lenB] == '1' && charsA[i] == '1' {
            for j := lenB; j >= 0; j-- {
                if charsB[j] == '0' {
                    charsB[j] = '1'
                    lenB--
                    break
                } else {
                    charsB[j] = '0'
                    if j == 0 {
                        charsB = append([]rune{'1'}, charsB...)
                        break
                    }
                }
            }
        }
    }
    return string(charsB)
}