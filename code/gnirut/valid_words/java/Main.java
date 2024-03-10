package gnirut.valid_words.java;

class Main {

    public static void main(String[] args) {
        
        String s = "cb";
        String[] k = new String[]{"cd", "bd", "cccb", "bcc", "bcdcb"};

        int result = 0;
        next_word:
        for (String word : k) {
            String[] chars = word.split("");
            for(String character : chars) {
                if (!s.contains(character)) {
                    continue next_word;
                }
            }
            result++;
        }
        
        System.out.println(result);
    }
    
}
