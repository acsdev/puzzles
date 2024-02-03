package course.memoization.sum;

import course.Util;

import java.util.HashMap;
import java.util.Map;
import java.util.TreeMap;
import java.util.function.Supplier;

/**
 * Given an array of size N and a value target
 * Determine if is possible to find the target using the numbers inside the array
 * -> all values in the array are positive
 * -> all values can be use more the one time
 */
public class CanSum {

    /**
     *  Time complexity O(n^m)
     *  Sapce complexity O(m)
     */
    public static boolean recursive(int[] array, int target) {
        if (target == 0) return true;
        if (target < 0) return false;

        for (int value : array) {
            int newTarget = target - value;
            if (recursive(array, newTarget)) {
                return true;
            }
        }
        return false;
    }

    /**
     *  Time complexity O(m*n)
     *  Space complexity O(m)
     */
    public static boolean recursiveMemoization(int[] array, int target, Map<Integer, Boolean> memo) {
        if (target == 0) return true;
        if (target < 0) return false;

        int key = target; // improves readability
        if (memo.containsKey(key)) return memo.get(key);

        for (int value : array) {
            key = target - value;
            boolean recursive = recursiveMemoization(array, key, memo);
            if (recursive) {
                memo.put(key, true);
                return true;
            }
        }
        memo.put(key, false);
        return false;
    }
    public static void main(String[] args) {
        System.out.println();
        Map<String, Supplier<Boolean>> calls = new TreeMap<>();
        calls.put("recursive[2, 3](7)", () -> recursive(new int[]{2, 3}, 7));
        calls.put("recursive[2, 4](7)", () -> recursive(new int[]{2, 4}, 7));
        calls.put("recursive[7, 14](300)", () -> recursive(new int[]{7, 14}, 300));
        Util.execute(calls);
        System.out.println();
        System.out.println();
        calls.clear();
        calls.put("recursiveMemoization[2,3](7)", () -> recursiveMemoization(new int[]{2, 3}, 7, new HashMap<>()));
        calls.put("recursiveMemoization[2,4](7)", () -> recursiveMemoization(new int[]{2, 4}, 7, new HashMap<>()));
        calls.put("recursiveMemoization[7, 14](300)", () -> recursiveMemoization(new int[]{7, 14}, 300, new HashMap<>()));
        Util.execute(calls);
    }
}
