#include <algorithm>
#include <fstream>
#include <iostream>
#include <numeric>
#include <sstream>
#include <map>
#include <set>
#include <string>
#include <vector>

struct InventoryRange
{
    size_t min;
    size_t max;

    InventoryRange(size_t min, size_t max)
        : min(min), max(max)
    {}

    bool operator<(const InventoryRange& other) const
    {
        if (min == other.min)
        {
            return max < other.max;
        }
        return min < other.min;
    }
};

class Inventory
{
public:    

Inventory()
{
}

std::map<size_t, size_t>::const_iterator Find(size_t id) const
{
    size_t search_id = id;

    // Find the first element with a key greater than the search_id
    auto id_location = inventory_ranges_.upper_bound(search_id);

    if (id_location == inventory_ranges_.begin())
    {
        // No valid range
        return inventory_ranges_.end();
    }

    std::advance(id_location, -1);
    return id_location;
}

bool CheckID(size_t id) const
{    
    auto inventory_element = Find(id);
    if (inventory_element == inventory_ranges_.end())
    {
        return false;
    }

    // Check if the ID is within the range
    return (id >= (*inventory_element).first && id <= (*inventory_element).first + (*inventory_element).second);
}

void AddRange(size_t start, size_t range)
{
    // Is the starting element already in the inventory?
    auto inventory_element_start = Find(start);
    
    if (inventory_element_start == inventory_ranges_.end())
    {
        Insert(start, range);
        return;
    }

    // If it does contain the element, update the end point if necessary
    if (start >= (*inventory_element_start).first && start <= (*inventory_element_start).first + (*inventory_element_start).second)
    {
        // So, the inventory_element_start does contain the element.
        // Update the end point if necessary
        if (start + range >= (*inventory_element_start).first + (*inventory_element_start).second)
        {
            inventory_ranges_[(*inventory_element_start).first] = start + range - (*inventory_element_start).first;
        }
    }
    else
    {
        // Otherwise, just insert a new range
        Insert(start, range);
    }

}

void Insert(size_t start, size_t range)
{
    inventory_ranges_[start] = range;
}

size_t TotalValidIDs() const
{
    return std::accumulate(inventory_ranges_.begin(), inventory_ranges_.end(), 0ULL, [](size_t sum, const std::pair<const size_t, size_t>& element) {
        return sum + element.second + 1;
    });
}

private:

std::map<size_t, size_t> inventory_ranges_;
};

class InventoryInput
{
public:
    void AddRange(InventoryRange range)
    {
        if (range.min <= range.max)
        {
            ranges_.insert(range);
        }
    }

    void AddID(size_t id)
    {
        ids_to_check.push_back(id);
    }

    Inventory CreateInventory() const
    {
        Inventory inventory;

        for (const auto range : ranges_)
        {
            size_t width = range.max - range.min;
            inventory.AddRange(range.min, width);
        }

        return inventory;
    }

    size_t CountValidIDs(const Inventory& inventory) const
    {
        size_t count = 0;
        std::for_each(ids_to_check.begin(), ids_to_check.end(), [&](size_t id) { count += (inventory.CheckID(id) ? 1 : 0); });
        return count;
    }

private:
    std::set<InventoryRange> ranges_;
    std::vector<size_t> ids_to_check;
};


size_t ParseNumber(const std::string& str)
{
    std::istringstream iss(str);
    size_t number;
    iss >> number;
    return number;
}

class Parser
{
public:
    Parser(std::string file_name, char sep)
        : fs(file_name),
        sep_(sep)
    {
    }

    ~Parser()
    {
        fs.close();
    }

    InventoryInput Read()
    {
        ReadRanges();
        ReadIDs();
        return inventoryInput;
    }

private:

    void ReadRanges()
    {
        bool blank_line = false;
        while (!blank_line)
        {    
            std::string line = "";
            std::getline(fs, line);

            if (line == "")
            {
                blank_line = true;
                break;
            }

            // Split the string into before and after the hyphen
            auto pos_sep = std::distance(line.begin(), std::find(line.begin(), line.end(), sep_));
            auto left_side = line.substr(0, pos_sep);
            auto right_side = line.substr(pos_sep + 1, line.length() - pos_sep - 1);
            
            size_t min = ParseNumber(left_side);
            size_t max = ParseNumber(right_side);
        
            inventoryInput.AddRange(InventoryRange(min, max));
            
        }
    }

    void ReadIDs()
    {
        size_t id;
        while (!fs.eof())
        {
            std::string line;
            std::getline(fs, line);
            if (line != "")
            {
                std::istringstream iss(line);
                iss >> id;
                inventoryInput.AddID(id);
            }
        }
    }

    InventoryInput inventoryInput;
    std::fstream fs;
    char sep_;
};

int main(int argc, char** argv)
{
    if (argc < 2)
    {
        std::cerr << "Usage: " << argv[0] << " <input_file>" << std::endl;
        return 1;
    }

    Parser parser(argv[1], '-');
    auto input = parser.Read();

    Inventory inventory = input.CreateInventory();
    
    std::cout << "Valid IDs: " << input.CountValidIDs(inventory) << std::endl;
    std::cout << "Total fresh ingredients: " << inventory.TotalValidIDs() << std::endl;
}