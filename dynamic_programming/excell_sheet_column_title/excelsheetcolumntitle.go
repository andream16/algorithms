func convertToTitle(n int) string {
    const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ";
    
    var (
        res string
        remainder int
    )
    
    for n > 0 {
        remainder = (n-1) % 26
        n = (n-remainder) / 26
        res = string(alphabet[remainder]) + res
    }

    return res

}
