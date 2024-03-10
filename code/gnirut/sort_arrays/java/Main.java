package gnirut.sort_arrays.java;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.HashSet;
import java.util.List;
import java.util.Map;
import java.util.Set;
import java.util.Map.Entry;


class Main {

    public static void main(String[] args) {
        
        int[] array = new int[]{-1, 1, -6, 4, 5, -6, 1, 4, 1};

        // group value by ocurrences
        Map<Integer, Integer> valueByOcurrences = new HashMap<>();
        for (int i = 0; i < array.length; i++) {
            
            int item = array[i];
            int value = valueByOcurrences.getOrDefault(item, 0);
            
            value += 1;
            
            valueByOcurrences.put(item, value);
        }

        // group occurences by values
        Map<Integer, Set<Integer>> occurrences = new HashMap<>();
        for (Entry<Integer,Integer> entry : valueByOcurrences.entrySet()) {
            Set<Integer> items = occurrences.getOrDefault(entry.getValue(), new HashSet<>());
            items.add(entry.getKey());
            occurrences.put(entry.getValue(), items);
        }

        // reorder array with items
        int index = 0;
        for (Entry<Integer, Set<Integer>> entry : occurrences.entrySet()) {
            List<Integer> values = new ArrayList<>(entry.getValue());
            for (int i = values.size() - 1; i >= 0; i--) {
                int count = entry.getKey();
                for (int j = 0; j < count; j++) {
                    array[index] = values.get(i);
                    index++;
                }
            }

        }

        System.out.println(array);

    }
}
