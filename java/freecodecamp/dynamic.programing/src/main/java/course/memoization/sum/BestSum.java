package course.memoization.sum;

import course.Util;

import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;
import java.util.TreeMap;
import java.util.function.Supplier;

/**
 * Similar to HowSum
 * BUT
 * -> has to find the smallest combination
 * --> this there is a tie, you can return any on them
 */
public class BestSum {

    /**
     * Time complexity: O(n^m*m)
     * Space complexity: O(m^2)
     */
    public static int[] recursive(int[] array, int target) {
        if (target == 0) return new int[]{};
        if (target < 0) return null;

        int[] shortest = null;

        for (int value : array) {
            int newTarget = target - value;
            int[] arrayReturned = recursive(array, newTarget);
            if (arrayReturned != null) {
                int[] arrayPlusOne = Arrays.copyOf(arrayReturned, arrayReturned.length + 1);
                arrayPlusOne[arrayReturned.length] = value;
                if (shortest == null || arrayPlusOne.length < shortest.length) {
                    shortest = arrayPlusOne;
                }
            }
        }

        return shortest;
    }

    /**
     * Time complexity: O(n^m*m)
     * Space complexity: O(m^2)
     */
    public static int[] recursiveMemoization(int[] array, int target, Map<Integer, int[]> memo) {
        if (target == 0) return new int[]{};
        if (target < 0) return null;

        int key = target; // improves readability
        if (memo.containsKey(key)) return memo.get(key);

        int[] shortest = null;

        for (int value : array) {
            int newTarget = target - value;
            int[] arrayReturned = recursiveMemoization(array, newTarget, memo);
            if (arrayReturned != null) {
                int[] arrayPlusOne = Arrays.copyOf(arrayReturned, arrayReturned.length + 1);
                arrayPlusOne[arrayReturned.length] = value;
                if (shortest == null || arrayPlusOne.length < shortest.length) {
                    shortest = arrayPlusOne;
                }
            }
        }
        memo.put(key, shortest);
        return shortest;
    }

    /**
     * Time complexity: O(n*m^2)
     * Space complexity: O(m^2)
     */
    public static void main(String[] args) {
        System.out.println();
        Map<String, Supplier<String>> calls = new TreeMap<>();
        calls.put("recursive[2, 3](7)", () -> Arrays.toString(recursive(new int[]{2, 3}, 7)));
        calls.put("recursive[1, 4, 5](8)", () -> Arrays.toString(recursive(new int[]{1, 4, 5}, 8)));
        calls.put("recursive[10, 2, 5, 25](100)", () -> Arrays.toString(recursive(new int[]{2, 10, 25}, 100)));
        Util.execute(calls);
        System.out.println();
        System.out.println();
        calls.clear();
        calls.put("recursiveMemoization[2,3](7)", () -> Arrays.toString(recursiveMemoization(new int[]{2, 3}, 7, new HashMap<>())));
        calls.put("recursiveMemoization[2,4](7)", () -> Arrays.toString(recursiveMemoization(new int[]{1, 4, 5}, 8, new HashMap<>())));
        calls.put("recursiveMemoization[7, 14](300)", () -> Arrays.toString(recursiveMemoization(new int[]{2, 10, 25}, 100, new HashMap<>())));
        Util.execute(calls);
    }
}