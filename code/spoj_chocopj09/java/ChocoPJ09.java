import java.io.File;
import java.io.IOException;
import java.util.Scanner;
import java.util.function.Function;

class ChocoPJ09 
{
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
            String[] input = scanner.nextLine().split(" ");
            int numberOfBalls = Integer.parseInt(input[0]);
            int maximumBallsPerTime = Integer.parseInt(input[1]);
            
            if (maximumBallsPerTime >= numberOfBalls) {
                System.out.println("Paula");
                System.exit(0);
            }

            boolean value = numberOfBalls % (maximumBallsPerTime + 1) != 0;
            System.out.println( value ? "Paula" : "Carlos");
        }
    }
}