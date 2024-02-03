package course.memoization;

import course.Util;

import java.util.HashMap;
import java.util.Map;
import java.util.function.Supplier;

/**
 *  Problem in a 2D grid
 *  You begin in the top-left corner and your goal is to travel to the bottom-right corner
 *  You can only move down or right
 *  In how many paths can you travel to the goal on a grid with x * y dimensions
 */
public class GridTraveller {
    /**
     * Time Complexity  O(2^n+m)
     * Space Complexity O(n+m)
     */
    public static int recursive(int x, int y) {
        if (x == 0 || y == 0) return 0;
        if (x == 1 && y == 1) return 1;
        //
        return recursive(x - 1, y) + recursive(x, y - 1);
    }

    public static int recursiveMemoization(int x, int y, HashMap<String, Integer> memo) {
        if (x == 0 || y == 0) return 0;
        if (x == 1 && y == 1) return 1;

        int stepX = x - 1;
        int stepY = y - 1;
        String key = String.format("%s,%s", stepX, stepY);
        if (memo.containsKey(key)) return memo.get(key);
        //
        int value = recursiveMemoization(stepX, y, memo) + recursiveMemoization(x, stepY, memo);
        memo.put(key, value);
        return value;
    }

    /**
     * Time Complexity  O(2^n+m)
     * Space Complexity O(n+m)
     */
    public static void main(String[] args) {
        Map<String, Supplier<Integer>> calls = new HashMap<>();
        calls.put("recursive(2,3)", () -> recursive(2, 3));
        calls.put("withRecursiveMemoization(2,3)", () -> recursiveMemoization(2, 3, new HashMap<>()));
        calls.put("recursive(23,25)", () -> recursive(18, 17));
        calls.put("withRecursiveMemoization(23,25)", () -> recursiveMemoization(18, 17, new HashMap<>()));

        Util.execute(calls);
    }

}
