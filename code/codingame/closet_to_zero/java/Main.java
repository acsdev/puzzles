package codingame.closet_to_zero.java;

// In this exercise, you have to analyze records of temperature to find the closest to zero.
//    
// Sample temperatures. Here, -1 is the closest to 0.
//  
// Write a program that prints the temperature closest to 0 among input data.
//  
// INPUT:
// Line 1: N, the number of temperatures to analyse
// Line 2: The N temperatures expressed as integers ranging from -273 to 5526
//  
// OUTPUT:
// Display 0 (zero) if no temperature is provided
// Otherwise, display the temperature closest to 0, knowing that if two numbers are equally close to zero, positive integer has to be considered closest to zero (for instance, if the temperatures are -5 to 5, then display 5)
//  
// CONSTRAINTS:
// 0 ≤ N < 10000
//  
// EXAMPLE: 
//     Input
//     5
//     1 -2 -8 4 5
//     Output
//     1
// 

import java.util.Map;
import java.util.Scanner;
import java.util.function.Consumer;

class Main {
    
    public static void main(String[] args) {

        var cases = Map.of(
        "7 5 9 1 4", 1,
        "-237", -237,
        "5526", 5526,
        "15 -7 -9 -14 -12", -7,
        "-10 -10", -10,
        "", 0,
        "15 7 9 14 7 12",7
        );
        
        cases.forEach((k, v) -> {
            var entries = k.isEmpty() ? 0 : k.split(" ").length;
            var input = entries +"\n" + k.replace(" ", "\n");

            Consumer<Integer> print = (result) -> {
                System.out.println("({" + input.replace("\n", " ") + "} -> " + v + ") = result was " + result);
                System.out.println(result);
            };

            Scanner in = new Scanner(input);
            int N = in.nextInt();
            Integer neg = null;
            Integer pos = null;
            if (N != 0) {    
                for (int i = 0; i < N; i++) {
                    int t = in.nextInt();
                    if (t < 0 && (neg == null || t >= neg)) {
                        neg = t;
                    }
                    if (t > 0 && (pos == null || t <= pos)) {
                        pos = t;
                    }
                }
                if (neg == null || pos == null || neg + pos == 0) {
                    if (neg != null) print.accept(neg);
                    else print.accept(pos);
                } else {
                    if ( Math.abs(neg) > pos) {
                        print.accept(pos);
                    } else {
                        print.accept(neg);
                    }
                }
            } else {
                print.accept(0);
            }
            in.close(); 
        });
    }
}