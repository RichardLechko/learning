/* int a = 10; // 00001010
int b = 6;  // 00000110

WriteLine($"a = {a}");
WriteLine($"b = {b}");
WriteLine($"a & b = {a & b}");
WriteLine($"a | b = {a | b}");
WriteLine($"a ^ b = {a ^ b}");

WriteLine($"a << 3 = {a << 3}");
WriteLine($"a * 8 = {a * 8}");
WriteLine($"b >> 1 = {b >> 1}");

WriteLine();
WriteLine("Outputting integers as binary:");
WriteLine($"a      =    {ToBinaryString(a)}");
WriteLine($"b      =    {ToBinaryString(b)}");
WriteLine($"a & b  =    {ToBinaryString(a & b)}");
WriteLine($"a | b  =    {ToBinaryString(a | b)}");
WriteLine($"a ^ b  =    {ToBinaryString(a ^ b)}");
WriteLine($"a << 3 =    {ToBinaryString(a << 3)}");
WriteLine($"b >> 1 =    {ToBinaryString(b >> 1)}");

static string ToBinaryString(int value)
{
    return Convert.ToString(value, toBase: 2).PadLeft(8, '0');
} */

int age = 50;
char firstDigit = age.ToString()[0];
WriteLine($"The {nameof(age)} variable uses {sizeof(int)} bytes of memory.");