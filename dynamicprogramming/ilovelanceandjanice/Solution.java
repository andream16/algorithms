public class Solution {
    static String alphabet = "abcdefghijklmnopqrstuvwxyz";
    public static String solution(String x) {
        String msg = "";
        for (int i = 0; i < x.length(); i++) {
            char c = x.charAt(i);
            if (c >= 'a' && c <= 'z') {
                msg += alphabet.charAt(35 - Character.getNumericValue(c));
                continue;
            }
            msg += c;
        }
        return msg;
    }
}
