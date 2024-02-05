package aba12c;

import java.io.File;
import java.io.IOException;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.Scanner;
import java.util.function.Function;

public class Main {

    // private record Data(int nrFriends, int nrKilograms, List<Integer> packList, Map<Integer, Integer> KgPrice) {
    // }

    public static class Data {
        int nrFriends;
        int nrKilograms;
        List<Integer> packList;
        Map<Integer, Integer> KgPrice;
        Data(int nrFriends, int nrKilograms, List<Integer> packList, Map<Integer, Integer> kgPrice) {
            this.nrFriends = nrFriends;
            this.nrKilograms = nrKilograms;
            this.packList = packList;
            this.KgPrice = kgPrice;
        }
    }

    public static void main(String[] args) {

        Function<String[], Scanner> getScanner = (array) -> {
            if (array.length == 1) {
                try {
                    File input = new File(array[0]).getAbsoluteFile();
                    return new Scanner(input);
                } catch (IOException e) {
                    throw new RuntimeException(e);
                }
            } 
            return new Scanner(System.in);
        };
        
        final Scanner scanner = getScanner.apply(args);
        try (scanner) {
            scanner.hasNextLine();
            int scenarios = Integer.parseInt(scanner.nextLine());
            for (int scenario = 0; scenario < scenarios; scenario++) {
                
                scanner.hasNextLine();
                String[] firstLine = scanner.nextLine().split(" ");
                scanner.hasNextLine();
                String[] secondLine = scanner.nextLine().split(" ");

                int nrFriends = Integer.parseInt(firstLine[0]);
                int nrKilograms = Integer.parseInt(firstLine[1]);
                Map<Integer, Integer> kgPrice = new HashMap<>();
                List<Integer> packList = new ArrayList<>();
                for (int index = 0; index < secondLine.length ; index++) {
                    int priceAsInt = Integer.parseInt(secondLine[index]);
                    if (priceAsInt < 0) continue;
                    
                    int kg = index + 1;
                    packList.add(kg);
                    kgPrice.put(kg, priceAsInt);
                }

                var data = new Data(nrFriends, nrKilograms, packList, kgPrice);
                System.out.println(solution(data));
            }
        }
    }

    private static int solution(Data data) {
        int value = withRecursion(data.nrKilograms, data.packList, data.KgPrice, new HashMap<Integer, Integer>());
        return value;
    }

    
    private static int withRecursion(
        int amountKg,
        List<Integer> kgList, 
        Map<Integer, Integer> kgPrice,
        Map<Integer, Integer> memo) {

        if (amountKg == 0) {
            return 0;
        }

        if (amountKg < 0) {
            return -1;
        }

        if (memo.containsKey(amountKg)) {
            return memo.get(amountKg);
        }

        int minPrice = -1;
        for (Integer kg : kgList) {
            int subAmountKg = amountKg - kg;
            int subPrice = withRecursion(subAmountKg, kgList, kgPrice, memo);
            if (subPrice >= 0) {
                int price = kgPrice.get(kg);
                int combinedPrice = price + subPrice;
                if (combinedPrice < minPrice || minPrice == -1) {
                    minPrice = combinedPrice;
                }
            }
        }

        memo.put(amountKg, minPrice);
        return minPrice;
    }
}