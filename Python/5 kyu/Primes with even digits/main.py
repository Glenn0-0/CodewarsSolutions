from math import sqrt

def isPrime(num):
    if num <= 1: return False
    if num == 2 or num == 3: return True
    if num%2 == 0 or num%3 == 0: return False
    for i in range(5, int(sqrt(num) + 1), 6):
        if (num % i == 0) or (num % (i + 2) == 0): return False
    return True 

def countEvens(n):
    count = 0
    while n > 0:
        if (n%10) % 2 == 0:
            count += 1
        n //= 10
    return count

def countNums(n):
    count = 0
    while n > 9:
        count += 1
        n //= 10
    count += 1
    if n == 1:
        count -= 1
    return count
    

def f(n):
    for i in range(n - 1, 0, -1):
        if isPrime(i) and (countEvens(i) == (countNums(i) - 1)):
            return i
    return 0

if __name__ == "__main__":
    print(f(2856))