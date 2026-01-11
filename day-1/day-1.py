from enum import Enum

def do_rotations(start, checkpoint, period, rotation_list):
    crossings = 0
    hits = 0
    distance_to_checkpoint = (start - checkpoint) % period
    for rotation in rotation_list:
        endpoint = distance_to_checkpoint + rotation

        dont_include_starting_position = 0
        if distance_to_checkpoint == 0 and rotation < 0:
            dont_include_starting_position = -1

        crossings_this_rotation, distance_to_checkpoint = divmod(endpoint, period)
        crossings += abs(crossings_this_rotation) + dont_include_starting_position

        if distance_to_checkpoint == 0:
            hits += 1
            if rotation < 0:
                crossings += 1

    return hits, crossings

def calculate_crossings(instructions):
    separated_directions = [dist if turn_direction == 'R' else -dist for (turn_direction, dist) in instructions]
    return [s if s != 0 else None for s in separated_directions]

class PasswordMethod(Enum):
    OnStop = 0
    OnPass = 1

if __name__ == "__main__":
    instruction_list = []

    method = PasswordMethod.OnPass

    with open('input.txt', 'r') as f:
        for line in f:
            line = line.strip()
            if line:
                direction = line[0]
                distance = int(line[1:])
                instruction_list.append((direction, distance))

    rotations = calculate_crossings(instruction_list)
    hits, crossings = do_rotations(50, 0, 100, rotations)
    if method == PasswordMethod.OnPass:
        print(crossings)
    else:
        print(hits)
