package course.memoization;

import course.Util;

import java.util.HashMap;
import java.util.Map;
import java.util.function.Supplier;

public class FibonacciRecursive {

    /**
     * TimeComplexity O(2^n)
     * SpaceComplexity O(2^n)
     *
     * @param number int number
     * @return fibonacci result
     * O(2^n) time **to bad**
     * O(n) space
     */
    public static int recursive(int number) {
        if (number <= 2) return 1;
        return recursive(number - 1) + recursive(number - 2);
    }


    /**
     * TimeComplexity 2n => O(n)
     * SpaceComplexity 2n => O(n)
     *
     * @param number int number
     * @param memoization cache to avoid necessary recursion
     */
    public static int recursiveMemoization(int number, Map<Integer, Integer> memoization) {
        if (memoization.containsKey(number)) return memoization.get(number);
        if (number <= 2) return 1;

        int left =  recursiveMemoization(number - 1, memoization);
        int right = recursiveMemoization(number - 2, memoization);
        memoization.put(number, left + right);
        return memoization.get(number);
    }

    public static void main(String[] args) {
        Map<String, Supplier<Integer>> calls = new HashMap<>();
        calls.put("recursive(50)", () -> FibonacciRecursive.recursive(50));
        calls.put("recursiveMemoization(50)", () -> FibonacciRecursive.recursiveMemoization(50, new HashMap<>()));

        Util.execute(calls);
    }
}
