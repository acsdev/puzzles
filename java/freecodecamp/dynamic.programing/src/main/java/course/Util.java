package course;

import java.time.Duration;
import java.time.LocalDateTime;
import java.util.Map;
import java.util.concurrent.TimeUnit;
import java.util.function.Supplier;

public class Util {

   public static <T> void execute(Map<String, Supplier<T>> calls) {
        calls.entrySet().stream().distinct().forEach( entry -> {
            LocalDateTime now = LocalDateTime.now();
            T value = entry.getValue().get();
            Duration between = Duration.between(now, LocalDateTime.now());
            System.out.printf("%s: value ( %s ) in %s nano seconds", entry.getKey(), value, TimeUnit.NANOSECONDS.convert(between));
            System.out.println();
        });
    }
}
