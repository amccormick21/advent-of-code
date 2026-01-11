namespace Joltage
{

    public class Bank(ICollection<int> joltages, int capacity)
    {
        private readonly List<int> joltages = [.. joltages];
        private readonly int capacity = capacity;

        record Pointer(int Index, int Value);
        private static Pointer GetLargestInRange(List<int> list, Range range)
        {
            Pointer largestPointer = new(-1, int.MinValue);
            if (range.Start.Value >= range.End.Value)
            {
                largestPointer = new(range.Start.Value, list[range.Start.Value]);
            }
            for (int i = range.Start.Value; i < range.End.Value; i++)
            {
                if (list[i] > largestPointer.Value)
                {
                    largestPointer = new Pointer(i, list[i]);
                }
            }
            return largestPointer;
        }

        public long LargestDifference()
        {
            if (capacity < 2)
            {
                throw new ArgumentException("Capacity must be greater than two.");
            }
            if (capacity > joltages.Count)
            {
                throw new ArgumentException("Capacity must be less than or equal to the number of joltages.");
            }

            // Find the largest number in the list within the capacity
            Range range = 0..(joltages.Count-capacity+1);
            long totalJoltage = 0;
            for (int i = 0; i < capacity; i++)
            {
                range = range.Start.Value..(joltages.Count - capacity + i + 1);

                var currentPointer = GetLargestInRange(joltages, range);

                // Move the range forward
                range = (currentPointer.Index + 1)..range.End.Value;

                // To get the total, multiply by 10 and add the current pointer
                totalJoltage = 10 * totalJoltage + currentPointer.Value;
            }
            return totalJoltage;
        }
    }

    public class Battery(List<List<int>> grid, int capacity)
    {
        private readonly List<Bank> battery = [.. grid.Select(row => new Bank(row, capacity))];

        public long TotalJoltage()
        {
            return battery.Sum(bank => bank.LargestDifference());
        }
    }
}