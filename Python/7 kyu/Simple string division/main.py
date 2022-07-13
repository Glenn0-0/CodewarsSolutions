def solve(st,k):
    longest = len(st) - k
    maxNum = 0
    
    for i in range(len(st) - longest + 1):
        num = int(st[i : i + longest])
        if num > maxNum:
            maxNum = num
    
    return maxNum

if __name__ == "__main__" :
    print(solve("1234", 2))