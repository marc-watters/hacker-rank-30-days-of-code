import java.io.*;
import java.util.*;

public class Solution {

    public static void main(String[] args) {
        /* Enter your code here. Read input from STDIN. Print output to STDOUT. Your class should be named Solution. */
        try(Scanner scanner = new Scanner(System.in)) {
            while (scanner.hasNext()) {
                List<Object> list = new ArrayList<>();

                int inputs = scanner.nextInt();
                for (int i = 0; i < inputs; i++) {
                    list.add(scanner.next());
                }

                printArray(list);
            }
        }
    }

    public static <T> void printArray(List<Object> arr) {
        for (int i = 0; i < arr.size(); i++) {
            System.out.println(arr.get(i));
        }
    }
}
