// INPUT = ZY, RES = 0
//
// Z = 90; A = 65; Z-A = 25; RES = (RES*26) + Y-A +1 = 0 + 25 + 1 = 26
// Y = 89; A = 65; Y-A = 24; RES = (RES*26) + Y-A +1 = 676 + 24 + 1 = 701
func titleToNumber(s string) int {
    var n int
    for _, c := range s {
        n *= 26
        n += int(c-'A')+1
        fmt.Println(n)
    }
    return n
}
