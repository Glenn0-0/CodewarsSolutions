from math import sqrt

def is_prime(num):
    if num <= 1: return False
    if num == 2 or num == 3: return True
    if num%2 == 0 or num%3 == 0: return False
    for i in range(5, int(sqrt(num) + 1), 6):
        if (num % i == 0) or (num % (i + 2) == 0): return False
    return True 

if __name__ == "__main__":
    print(is_prime(7))