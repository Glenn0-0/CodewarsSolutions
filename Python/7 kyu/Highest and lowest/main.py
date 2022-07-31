def high_and_low(numbers):
    nums = ""
    arr = numbers.split()
    max = arr[0]
    for i in arr: 
        if int(i) > int(max):
            max = i 
    min = arr[0]
    for k in arr: 
        if int(k) < int(min):
            min = k 
    nums = max + " " + min
    return nums

if __name__ == "__main__":
    print(high_and_low("1 2 -3 4 5"))