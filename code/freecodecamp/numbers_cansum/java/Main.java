package freecodecamp.numbers_cansum.java;

/**
Problem Statement
++++++++++++++++++
Given a target sum and an unsorted array of positive integers.
Whether it is possible to generate the targetSum using the numbers in the array? — canSum(targetSum, integers)
canSum(8, [2,3,5]) => true, [2,2,2,2] or [2, 3, 3] or [3, 5]
canSum(8, [3,6,9]) => false
*/


import static java.lang.String.format;
import static java.time.LocalDateTime.now;

import java.time.Duration;
import java.time.LocalDateTime;
import java.util.Arrays;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.concurrent.TimeUnit;
import java.util.function.Function;

class Main {

    /*
     * Time complexity: O(n) * O(n) => O(n^2)
     */
    public static boolean withRecursion(int[] array, int target) {
        if (target == 0) return true;
        if (target < 0) return false;

        for (int value : array) { //O(n)
            int newTarget = target - value;
            if (withRecursion(array, newTarget)) {//O(n)
                return true;
            }
        }
        return false;
    }

    /**
     * 
     */
    public static boolean withRecursionAndMemoization(int[] array, int target, Map<Integer, Boolean> memo) {
        if (target == 0) return true;
        if (target < 0) return false;

        int key = target; // improves readability
        if (memo.containsKey(key)) return memo.get(key);

        for (int value : array) {
            key = target - value;
            boolean recursive = withRecursionAndMemoization(array, key, memo);
            if (recursive) {
                memo.put(key, true);
                return true;
            }
        }
        memo.put(key, false);
        return false;
    }

    
    public static boolean withTabulation(int[] array, int target) {
        int inArray = target + 1;
        boolean[] canSum = new boolean[inArray];        

        canSum[0] = true;
        begin:
        for (int index = 0; index < canSum.length; index++) {
            boolean value = canSum[index];
            if (value) {
                for (int j = 0; j < array.length; j++) {
                    int factor = array[j];
                    int jump = factor + index;
                    if (canSum.length > jump) {
                        canSum[jump] = true;
                        if (jump == target) {
                            break begin;
                        }
                    }
                }
            }
        } 

        return canSum[target];
    }

    public static void main(String[] args) {
        Function<LocalDateTime, Number> time = date -> {
            Duration between = Duration.between(date, LocalDateTime.now());
            return  TimeUnit.MICROSECONDS.convert(between);
        };
        
        List.of(
            Pair.of(8, new int[]{2, 3, 5}), 
            Pair.of(8, new int[]{3, 6, 9})).forEach(pair -> {

            System.out.println(format("B: input %d(%s) result %s in time %d ", 
                pair.target, Arrays.toString(pair.values), withRecursion(pair.values, pair.target), time.apply(now())));

            var memo = new HashMap<Integer, Boolean>();
            System.out.println(format("M: input %d(%s) result %s in time %d ", 
                pair.target, Arrays.toString(pair.values), withRecursionAndMemoization(pair.values, pair.target, memo), time.apply(now())));

            System.out.println(format("T: input %d(%s) result %s in time %d ", 
                pair.target, Arrays.toString(pair.values), withTabulation(pair.values, pair.target), time.apply(now())));
            
                System.out.println();
        });
    }

    public static class Pair {
        int target;
        int[] values;
        static Pair of(int target, int[] values) {
            Pair p = new Pair();
            p.target = target;
            p.values = values;
            return p;
        }
    }
}
