/* object height = 1.88;
object name = "Amir";
Console.WriteLine($"{name} is {height} metres tall.");

//int length1 = name.Length;
int length2 = ((string)name).Length;
Console.WriteLine($"{name} has {length2} characters."); */

/* dynamic something = "Ahmed";

something = 12;
something = new[] {3, 5, 7};
Console.WriteLine($"Length is {something.Length}"); */

/* using System.Xml;

var population = 67_000_000;
var weight = 1.88;
var price = 4.99M;
var fruit = "Apples";
var letter = 'Z';
var happy = true;

// good use of var because it avoids the repeated type
var xml1 = new XmlDocument();
XmlDocument xml2 = new XmlDocument();

// bad use of var because we can not tell the type
var file1 = File.CreateText("something1.txt");
StreamWriter file2 = File.CreateText("something2.txt"); */

int number = 13;

Console.WriteLine($"number has been set to: {number}");

number = default;
Console.WriteLine($"number has been reset to its default: {number}");
Console.WriteLine($"default(int) = {default(int)}");
Console.WriteLine($"default(bool) = {default(bool)}");
Console.WriteLine($"default(DateTime) = {default(DateTime)}");
Console.WriteLine($"default(string) = {default(string)}");