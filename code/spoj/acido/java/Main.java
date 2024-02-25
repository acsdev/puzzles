package spoj.acido.java;

import java.io.BufferedReader;
import java.io.File;
import java.io.FileReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.util.function.Function;

class Main {

    private static final char B = 'B';
    private static final char S = 'S';
    private static final char C = 'C';
    private static final char F = 'F';

    static class CharStack {
        final StringBuilder sb = new StringBuilder();

        public void push(char ch) {
            sb.append(ch);
        }

        public char peek() {
            int last = sb.length() - 1;
            return sb.charAt(last);
        }

        public char pop() {
            int last = sb.length() - 1;
            char ch = sb.charAt(last);
            sb.setLength(last);
            return ch;
        }

        public int size() {
            return sb.length();
        }
    }

    public static void main(String[] args) {
        try {

            Function<String[], BufferedReader> getInput = (array) -> {
                if (array.length == 1) {
                    try {
                        File input = new File(array[0]).getAbsoluteFile();
                        return new BufferedReader(new FileReader(input));
                    } catch (IOException e) {
                        throw new RuntimeException(e);
                    }
                } 
                return new BufferedReader(new InputStreamReader(System.in));
            };
            
            final BufferedReader reader = getInput.apply(args);
            try (reader) {
                while (true) {
                    String line = reader.readLine();
                    if (line.isEmpty() || line.isBlank()) {
                        break;
                    } 
                    
                    char[] letters = line.toCharArray();
                    int count = 0;
                    
                    CharStack stack = new CharStack();
                    for (char letter : letters) {
                        if (stack.size() == 0) {
                            stack.push(letter);
                        } else {
                            char topValue = stack.peek();
                            boolean bs = topValue == B && letter == S || topValue == S && letter == B;
                            boolean cf = topValue == C && letter == F || topValue == F && letter == C;
                            if (bs || cf) {
                                stack.pop();
                                count++;
                            } else {
                                stack.push(letter);
                            }
                        }
                    }
                    System.out.println(count);
                }
            }
        } catch (Exception ex) {
            System.err.println(ex.getMessage());
        }
    }
}