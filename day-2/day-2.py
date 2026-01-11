import math


def get_order_of_magnitude(number):
    order = 0
    while math.pow(10, order) <= number:
        order += 1
    return order


def get_factors_of_number(number, only_half_factor=False):
    factors = []
    if only_half_factor:
        div, mod = divmod(number, 2)
        if mod == 0:
            factors.append(div)
    else:
        factors.append(1)
        test_factor = 2
        while test_factor <= number // 2:
            if number % test_factor == 0:
                factors.append(test_factor)
            test_factor += 1
    return factors


def get_all_one_above_orders(number):
    """ Find all the numbers which can be multiplied to give
    repeating sequences. For example:
    11 * 1-digit number -> 2
    111 * 1-digit number -> 3
    101 * 2-digit number -> 4
    10101 * 2-digit number -> 6
    1001 * 3-digit number -> 6
    1010101 * 2-digit number -> 8
    10001 * 4-digit number -> 8
    1001001 * 3-digit number -> 9
    100001 * 5-digit number -> 10
    10000001 * 6-digit number -> 12
    """
    multipliers = []

    factors = get_factors_of_number(number, only_half_factor=False)

    for factor in factors:
        index = 0
        multiplier = 0
        while index < number:
            multiplier += math.pow(10, index)
            index += factor

        if multiplier != 1:
            multipliers.append(multiplier)
    return multipliers


def sum_invalid_ids(range_min, range_max):

    orders_of_magnitude_in_range = get_order_of_magnitude(range_min)
    multipliers = get_all_one_above_orders(orders_of_magnitude_in_range)
    invalid_ids = []

    for multiplier in multipliers:
        repeating_sequence_min = math.ceil(range_min / float(multiplier))
        repeating_sequence_max = math.floor(range_max / float(multiplier))

        invalid_ids.extend([i * multiplier for i in range(repeating_sequence_min, repeating_sequence_max + 1)])
        invalid_ids = list(set(invalid_ids))
    return invalid_ids


def split_ranges(ranges):
    gated_ranges = []
    for (range_min, range_max) in ranges:
        magnitude_min = get_order_of_magnitude(range_min)-1
        magnitude_max = get_order_of_magnitude(range_max)-1

        if magnitude_max > magnitude_min:
            gated_ranges.append((range_min, math.pow(10, magnitude_max) - 1))
            gated_ranges.append((math.pow(10, magnitude_max), range_max))
        else:
            gated_ranges.append((range_min, range_max))

    return gated_ranges


def sum_all_ranges(ranges):
    trimmed_ranges = split_ranges(ranges)
    invalid_ids_sum = 0

    for min, max in trimmed_ranges:
        invalid_ids = sum_invalid_ids(min, max)
        invalid_ids_sum += sum(invalid_ids)
        print(invalid_ids)

    return invalid_ids_sum


def parse_file(file_path):
    ranges = []
    with open(file_path, 'r') as f:
        for line in f:
            split_line = line.split(',')
            for split_item in split_line:
                ranges.append(split_item)

    parsed_ranges = []
    for range in ranges:
        range_parts = range.split('-')
        range_min = int(range_parts[0])
        range_max = int(range_parts[1])
        parsed_ranges.append((range_min, range_max))

    return parsed_ranges


if __name__ == "__main__":

    ranges = parse_file('input.txt')
    print(sum_all_ranges(ranges))


