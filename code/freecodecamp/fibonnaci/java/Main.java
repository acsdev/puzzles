package freecodecamp.fibonnaci.java;

// This section will dissect a Dynamic Programming problem called Fibonacci sequence to illustrate the above 2 very important concepts.
//
// Fibonacci sequence is a well-known sequence of numbers that looks something like this 0,1,1,2,3,5,8,13, 21,…..
//
// fib(n) = fib(n-1) + fib(n-2) where fib(0) = 0 and fib(1) = 1
import static java.lang.String.format;
import static java.time.LocalDateTime.now;

import java.time.Duration;
import java.time.LocalDateTime;
import java.util.HashMap;
import java.util.Map;
import java.util.concurrent.TimeUnit;
import java.util.function.Function;
import java.util.stream.IntStream;

class Main {

    public static int withRecursion(int number) { // O(2^n)
        if (number <= 2) return 1;
        int prev = withRecursion(number - 1); // O(n)
        int next = withRecursion(number - 2); // O(n)
        return prev + next;
    }

    public static int withRecursionAndMemoization(int number, Map<Integer, Integer> memoization) {
        if (memoization.containsKey(number)) return memoization.get(number);
        if (number <= 2) return 1;

        int left =  withRecursionAndMemoization(number - 1, memoization);
        int right = withRecursionAndMemoization(number - 2, memoization);
        memoization.put(number, left + right);
        return memoization.get(number);
    }
    
    public static int withTabulation(int number) { // O(n)
        int numberOfFibNumbers = number + 1;
        int[] fib = new int[numberOfFibNumbers];

        fib[0] = 0;
        fib[1] = 1;
        for (int index = 2; index < numberOfFibNumbers; index++) {
            fib[index] = fib[index-1] + fib[index-2];
        }
        return fib[number];
    }

    public static void main(String[] args) {
        Function<LocalDateTime, Number> time = date -> {
            Duration between = Duration.between(date, LocalDateTime.now());
            return  TimeUnit.MICROSECONDS.convert(between);
        };
        IntStream.of(2, 7, 13, 37).forEach(input -> {

            System.out.println(format("B: input %d result %dd in time %d ", 
                input, withRecursion(input), time.apply(now())));

            System.out.println(format("M: input %d result %dd in time %d ", 
                input, withRecursionAndMemoization(input, new HashMap<Integer, Integer>()), time.apply(now())));

            System.out.println(format("T: input %d result %dd in time %d ", 
                input, withTabulation(input), time.apply(now())));
            
                System.out.println();
        });
    }
}
