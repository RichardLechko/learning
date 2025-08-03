// See https://aka.ms/new-console-template for more information
/* using System;

namespace Basics
{
    class Program
    {
        static void Main(string[] args)
        {
            Console.WriteLine("Hello World!");
        }
    }
} */

/* using System.Reflection;

System.Data.DataSet ds;
HttpClient client;

Assembly? myApp = Assembly.GetEntryAssembly();

if (myApp == null) return;

foreach (AssemblyName name in myApp.GetReferencedAssemblies())
{
    Assembly a = Assembly.Load(name);

    int methodCount = 0;

    foreach (TypeInfo t in a.DefinedTypes)
    {
        methodCount += t.GetMethods().Count();
    }

    Console.WriteLine(
        "{0:N0} types with {1:N9} methods in {2} assembly.",
        arg0: a.DefinedTypes.Count(),
        arg1: methodCount,
        arg2: name.Name);
} */

/* double heightInMetres = 1.88;
Console.WriteLine($"The variable {nameof(heightInMetres)} has the value {heightInMetres}."); */

/* char letter = 'A';
char digit = '1';
char symbol = '$';
char userChoice = GetSomeKeyStroke(); */

/* string firstName = "Bob";
string lastName = "Smith";
string phoneNumber = "(215) 555-4256";

string horizontalLine = new('-', count: 74);

string address = GetAddressFromDatabase(id: 563);

string grinningEmoji = char.ConvertFromUtf32(0x1F600);

Console.OutputEncoding = System.Text.Encoding.UTF8;
Console.WriteLine(grinningEmoji); */

/* string filePath = @"C:\televisions\sony\bravia.txt";
Console.WriteLine(filePath); */

/* string xml = """
             <person age="50">
                <first_name>Mark</first_name>
             </person>
             """;
Console.WriteLine(xml); */

var person = new { FirstName = "Alice", Age = 56 };

string json = $$"""
              {
                "first_name": "{{person.FirstName}}",
                "age": "{{person.Age}}",
                "calculation": "{{1 + 2}}",
              }
              """;
Console.WriteLine(json);