import java.io.File;
import java.io.IOException;
import java.util.Scanner;
import java.util.function.Function;

class Main {
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
            while (scanner.hasNextLine()) {
                int number = Integer.parseInt(scanner.nextLine());
                if (number == 42) break;

                System.out.println(number);
            }
        }
    }
}