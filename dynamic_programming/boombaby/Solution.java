package solution;

import java.math.BigInteger;

public class Solution {
    public static String solution(String x, String y) {

        BigInteger nMach = new BigInteger(x);
        BigInteger nFacula = new BigInteger(y);
        BigInteger iterations = new BigInteger("0");
        final String impossible = "impossible";
        final BigInteger one = BigInteger.ONE;
        final BigInteger zero = new BigInteger("0");

        if (nMach.equals(nFacula)) {
            if (nMach.equals(one)) return "0";
            return impossible;
        }

        while (nMach.compareTo(one) > 0 && nFacula.compareTo(one) > 0) {
            if (nFacula.compareTo(nMach) > 0) {
                if (nFacula.mod(nMach).equals(zero)) return impossible;
                BigInteger diff = nFacula.divide(nMach);
                iterations = iterations.add(diff);
                nFacula = nFacula.mod(nMach);
                continue;
            }
            if (nMach.mod(nFacula).equals(zero)) return impossible;
            BigInteger diff = nMach.divide(nFacula);
            iterations = iterations.add(diff);
            nMach = nMach.mod(nFacula);
        }

        iterations = iterations.add(nMach.max(nFacula));
        iterations = iterations.subtract(one);

        return String.valueOf(iterations);
    }
}
