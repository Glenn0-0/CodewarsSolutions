def move_zeros(array):
    for elem in array:
        if elem == 0:
            array.remove(0)
            array.append(0)
    return array

print(move_zeros([0, 2, 4, 5, 0, 0, 0, 6, 1]))