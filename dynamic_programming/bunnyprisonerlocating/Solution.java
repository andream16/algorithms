public class Solution {
    public static String solution(long x, long y) {
        return String.valueOf((((x+y-1)*(x+y-2)) / 2) + x);
    }
}
