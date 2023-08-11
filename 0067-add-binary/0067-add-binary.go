func addBinary(a string, b string) string {
    
  
    lenB := len(b) - 1
    charsA := []rune(a)
    charsB := []rune(b)

    for i := len(a) - 1; i >= 0; i-- {
        fmt.Print(i)
        fmt.Println("CharsB",string(charsB))
        
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
                    fmt.Println("outj",string(charsB))
                    lenB--
                    break
                } else {
                    fmt.Println("jtop",string(charsB))
                    fmt.Println("jindex",j)
                    charsB[j] = '0'
                    if j == 0 {
                        charsB = append([]rune{'1'}, charsB...)
                        fmt.Println("j",string(charsB))
                        break
                    }
                }
            }
        }
    }
    return string(charsB)
}