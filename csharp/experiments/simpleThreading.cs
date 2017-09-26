using System;
using System.Threading;

public class Threading {
    static Object myLock = new Object();

    public static void Main() {
        Thread t1 = new Thread(FirstWorker);
        Thread t2 = new Thread(SecondWorker);

        t1.Start();
        t2.Start();

        t1.Join();
        t2.Join();

        Console.WriteLine("Fully done!");
    }

    static void FirstWorker() {
        Console.WriteLine("First worker started!");

        lock (myLock)
        {
            Console.WriteLine("First worker entered the critical block!");
            Thread.Sleep(1000);
            Console.WriteLine("First worker left the critical block!");
        }

        Console.WriteLine("First worker completed!");
    }

    static void SecondWorker() {
        Console.WriteLine("Second worker started!");

        lock (myLock)
        {
            Console.WriteLine("Second worker entered the critical block!");
            Thread.Sleep(2000);
            Console.WriteLine("Second worker left the critical block!");
        }

        Console.WriteLine("Second worker completed!");
    }
}