package gnirut.perfect_matches.java;

class Main {
    public static void main(String[] args) {
        int count = 0;
        // int[] input = {1,1,1,3,2,3};
        int[] input = {4,5,6};
        for (int i = 0; i < input.length; i++) {
            for (int j = i + 1; j < input.length; j++) {
                if ( input[i] == input[j]) {
                    count++;
                }
            }
        }
        System.out.println(count);
    }
}
