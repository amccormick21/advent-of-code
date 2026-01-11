namespace Day3.Tests;

public class LargestJoltageTests
{
    [Fact]
    public void Length2_Test()
    {
        var expected = 12;
        var actual = new Joltage.Bank([1, 2], 2).LargestDifference();
        Assert.Equal(expected, actual);
    }

    [Fact]
    public void Ascending_Test()
    {
        var expected = 35;
        var actual = new Joltage.Bank([1, 2, 3, 5], 2).LargestDifference();
        Assert.Equal(expected, actual);
    }

    [Fact]
    public void Descending_Test()
    {
        var expected = 53;
        var actual = new Joltage.Bank([5, 3, 2, 1], 2).LargestDifference();
        Assert.Equal(expected, actual);
    }

    [Fact]
    public void Sample_Test()
    {
        var expected = 92;
        var actual = new Joltage.Bank([8,1,8,1,8,1,9,1,1,1,1,2,1,1,1], 2).LargestDifference();
        Assert.Equal(expected, actual);
    }

    [Fact]
    public void SampleMoreValues_Test()
    {
        var expected = 888911112111;
        var actual = new Joltage.Bank([8,1,8,1,8,1,9,1,1,1,1,2,1,1,1], 12).LargestDifference();
        Assert.Equal(expected, actual);
    }
}

public class BatteryTests
{
    [Fact]
    public void SampleBattery_Test()
    {
        var grid = new List<List<int>>()
        {
            new() {1, 2, 3, 4, 6},
            new() {1, 3, 5, 3, 1},
        };

        var expected = 46 + 53;
        var actual = new Joltage.Battery(grid, 2).TotalJoltage();
        Assert.Equal(expected, actual);
    }
}
