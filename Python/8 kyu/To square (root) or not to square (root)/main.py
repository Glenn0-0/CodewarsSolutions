import math

def square_or_square_root(arr):
    for i in range(len(arr)):
        sqrt = math.sqrt(arr[i])
        if sqrt == int(sqrt):
            arr[i] = int(sqrt)
        else:
            arr[i] *= arr[i]
    return arr

if __name__ == "__main__":
    print(square_or_square_root([10, 9, 37, 36, 1]))