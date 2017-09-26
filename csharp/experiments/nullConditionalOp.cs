using System;

public class NullConditionalOp {
    public static string Truncate(string value, int length) {
        return value?.Substring(0, Math.Min(value.Length, length));
    }

    public static void Main() {
        var res = Truncate(null, 2);
        Console.WriteLine(res == null ? "Yep, null!" : res);

        res = Truncate("some", 2);
        Console.WriteLine(res == null ? "Yep, null!" : res);
    }
}