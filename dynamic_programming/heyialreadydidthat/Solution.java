import java.util.Arrays;
import java.util.Comparator;
import java.util.HashMap;

public class Solution {

    public static int solution(String n, int b) {
        HashMap<String, Integer> hs =new HashMap<>();
        hs.put(n, 0);
        return solutionRec(n, b, 0, hs);
    }

    public static int solutionRec(String s, int b, int counter, HashMap<String, Integer> results) {
        counter += 1;
        int l = s.length();
        String[] xy = getXY(s.split(""));
        Integer x = Integer.parseInt(xy[0], b);
        Integer y = Integer.parseInt(xy[1], b);
        String currRes = Integer.toString(x-y, b);
        if (l != currRes.length()) {
            currRes = String.format("%0"+l+"d", Integer.parseInt(currRes, b));
        }
        Integer idx = results.get(currRes);
        if (idx != null) {
            return counter - idx;
        }
        results.put(currRes, counter);
        return solutionRec(currRes, b, counter, results);
    }

    private static String[] getXY(String[] xArr) {
        int l = xArr.length;

        String[] yArr = Arrays.copyOf(xArr, l);

        xArr = sort(xArr, true);
        yArr = sort(yArr, false);

        StringBuilder x = new StringBuilder();
        StringBuilder y = new StringBuilder();

        for (int i=l-1; i>=0; i--) {
            x.append(xArr[i]);
            y.append(yArr[i]);
        }

        return new String[]{x.toString(), y.toString()};

    }

    private static String[] sort(String[] arr, Boolean asc) {
        Arrays.sort(arr, new Comparator<String>() {
            public int compare(String x, String y) {
                if (asc) {
                    return x.compareTo(y);
                }
                return y.compareTo(x);
            }
        });
        return arr;
    }

}
