// See https://aka.ms/new-console-template for more information
Console.WriteLine("Day 3: Joltage Tests");

var lines = File.ReadAllLines(Path.Combine(Directory.GetCurrentDirectory(), "..", "..", "..", "input.txt"));
var grid = lines.Select(line => line.Select(c => int.Parse(c.ToString())).ToList()).ToList();

const int capacity = 12;
Joltage.Battery battery = new(grid, capacity);
Console.WriteLine(battery.TotalJoltage());