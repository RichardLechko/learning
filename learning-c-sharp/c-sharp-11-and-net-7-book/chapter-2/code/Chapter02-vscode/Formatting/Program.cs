/* int numberOfApples = 12;
decimal pricePerApple = 0.35M; */

/* Console.WriteLine(
    format: "{0} apples cost {1:C}",
    arg0: numberOfApples,
    arg1: pricePerApple * numberOfApples);

Console.WriteLine(
    format: "{0} {1} lived in {2}.",
    arg0: "Roger", arg1: "Cevung",
    arg2: "Stockholm, Stockholm");

Console.WriteLine(
    format: "{0} {1} lived in {2}, {3} and worked in the {4} team at {5}.",
    "Roger", "Cevung", "Stockholm", "Sweden", "Education", "Optimizely"); */

/* Console.WriteLine($"{numberOfApples} apples cost {pricePerApple * numberOfApples:C}");

const string firstname = "Omar";
const string lastname = "Rudberg";
const string fullname = $"{firstname} {lastname}";
Console.WriteLine(fullname); */

/* string applesText = "Apples";
int applesCount = 1234;
string bananasText = "Bananas";
int bananasCount = 56789;

Console.WriteLine(
    format: "{0,-10} {1,6}",
    arg0: "Name",
    arg1: "Count");
Console.WriteLine(
    format: "{0,-10} {1,6:N0}",
    arg0: applesText,
    arg1: applesCount);
Console.WriteLine(
    format: "{0,-10} {1,6:N0}",
    arg0: bananasText,
    arg1: bananasCount); */

/* Write("Type your first name and press ENTER: ");
string firstName = Console.ReadLine()!;

Write("Type your age and press ENTER: ");
string age = Console.ReadLine()!;

WriteLine($"Hello {firstName}, you look good for {age}."); */

/* Write("Press any key combination: ");
ConsoleKeyInfo key = ReadKey();
WriteLine();
WriteLine("Key: {0}, Char: {1}, Modifiers: {2}",
    arg0: key.Key,
    arg1: key.KeyChar,
    arg2: key.Modifiers); */

using System;

namespace Arguments
{
    class Program
    {
        static void Main(string[] args)
        {
            Console.WriteLine("Hello World!");
        }
    }
}