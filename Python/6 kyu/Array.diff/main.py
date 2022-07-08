def array_diff(a, b):
    for elem in b:
        while elem in a: a.remove(elem)
    return a

print(array_diff([1, 2, 2, 4, 5, 5, 5, 1, 7, 0, 9, 0, 2, 3],[2, 5, 9]))