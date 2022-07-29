from audioop import mul


def multiple_of_index(arr):
    ints = []
    for i in range(1, len(arr)):
        if arr[i]%i == 0:
            ints.append(arr[i])
    return ints

if __name__ == "__main__":
    print(multiple_of_index([22, -6, 32, 82, 9, 25]))