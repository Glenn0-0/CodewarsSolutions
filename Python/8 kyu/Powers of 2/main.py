def powers_of_two(n):
    arr = []
    for i in range(n+1):
        arr.append(2**i)
    return arr

if __name__ == "__main__":
    print(powers_of_two(4))