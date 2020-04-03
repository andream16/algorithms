/*
Hey, I Already Did That!
========================

Commander Lambda uses an automated algorithm to assign minions randomly to tasks, in order to keep her minions on their toes. But you've noticed a flaw in the algorithm - it eventually loops back on itself, so that instead of assigning new minions as it iterates, it gets stuck in a cycle of values so that the same minions end up doing the same tasks over and over again. You think proving this to Commander Lambda will help you make a case for your next promotion. 

You have worked out that the algorithm has the following process: 

1) Start with a random minion ID n, which is a nonnegative integer of length k in base b
2) Define x and y as integers of length k.  x has the digits of n in descending order, and y has the digits of n in ascending order
3) Define z = x - y.  Add leading zeros to z to maintain length k if necessary
4) Assign n = z to get the next minion ID, and go back to step 2

For example, given minion ID n = 1211, k = 4, b = 10, then x = 2111, y = 1112 and z = 2111 - 1112 = 0999. Then the next minion ID will be n = 0999 and the algorithm iterates again: x = 9990, y = 0999 and z = 9990 - 0999 = 8991, and so on.

Depending on the values of n, k (derived from n), and b, at some point the algorithm reaches a cycle, such as by reaching a constant value. For example, starting with n = 210022, k = 6, b = 3, the algorithm will reach the cycle of values [210111, 122221, 102212] and it will stay in this cycle no matter how many times it continues iterating. Starting with n = 1211, the routine will reach the integer 6174, and since 7641 - 1467 is 6174, it will stay as that value no matter how many times it iterates.

Given a minion ID as a string n representing a nonnegative integer of length k in base b, where 2 <= k <= 9 and 2 <= b <= 10, write a function solution(n, b) which returns the length of the ending cycle of the algorithm above starting with n. For instance, in the example above, solution(210022, 3) would return 3, since iterating on 102212 would return to 210111 when done in base 3. If the algorithm reaches a constant, such as 0, then the length is 1.

-- Java cases -- 
- Input: Solution.solution('210022', 3) Output: 3
- Input: Solution.solution('1211', 10) Output: 1
*/

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
