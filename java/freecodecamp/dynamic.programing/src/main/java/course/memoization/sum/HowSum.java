package course.memoization.sum;

import course.Util;

import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;
import java.util.TreeMap;
import java.util.function.Supplier;

/**
 * Given an array of size N and a value target
 * Determine one combination where the target can be found using the numbers inside the array
 * -> if not possible, return null
 * -> all values in the array are positive
 * -> all values can be use more the one time
 */
public class HowSum {

    /**
     *  Time complexity: O(n^m*m)
     *  Space complexity: O(m)
     */
    public static int[] recursive(int[] array, int target) {
        if (target == 0) return new int[]{};
        if (target < 0) return null;

        for (int value : array) {
            int newTarget = target - value;
            int[] arrayReturned = recursive(array, newTarget);
            if (arrayReturned != null) {
                int[] arrayPlusOne = Arrays.copyOf(arrayReturned, arrayReturned.length + 1);
                arrayPlusOne[arrayReturned.length] = value;
                return arrayPlusOne;
            }
        }
        return null;
    }

    /**
     *  Time complexity O(n*m^2)
     *  Space complexity O(m^2)
     */
    public static int[] recursiveMemoization(int[] array, int target, Map<Integer, int[]> memo) {
        if (target == 0) return new int[]{};
        if (target < 0) return null;

        int key = target; // improves readability
        if (memo.containsKey(key)) return memo.get(key);

        for (int value : array) {
            key = target - value;
            int[] arrayReturned = recursiveMemoization(array, key, memo);
            if (arrayReturned != null) {
                int[] arrayPlusOne = Arrays.copyOf(arrayReturned, arrayReturned.length + 1);
                arrayPlusOne[arrayReturned.length] = value;
                memo.put(key, arrayPlusOne);
                return arrayPlusOne;
            }
        }
        memo.put(key, null);
        return null;
    }
    public static void main(String[] args) {
        System.out.println();
        Map<String, Supplier<String>> calls = new TreeMap<>();
        calls.put("recursive[2, 3](7)", () -> Arrays.toString(recursive(new int[]{2, 3}, 7)));
        calls.put("recursive[2, 4](7)", () -> Arrays.toString(recursive(new int[]{2, 4}, 7)));
        calls.put("recursive[7, 14](300)", () -> Arrays.toString(recursive(new int[]{7, 14}, 300)));
        Util.execute(calls);
        System.out.println();
        System.out.println();
        calls.clear();
        calls.put("recursiveMemoization[2,3](7)", () -> Arrays.toString(recursiveMemoization(new int[]{2, 3}, 7, new HashMap<>())));
        calls.put("recursiveMemoization[2,4](7)", () -> Arrays.toString(recursiveMemoization(new int[]{2, 4}, 7, new HashMap<>())));
        calls.put("recursiveMemoization[7, 14](300)", () -> Arrays.toString(recursiveMemoization(new int[]{7, 14}, 300, new HashMap<>())));
        Util.execute(calls);
    }
}
