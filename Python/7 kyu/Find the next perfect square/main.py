from math import sqrt

def find_next_square(sq):
    num = sqrt(sq)
    if num == int(num):
        return (num + 1)**2
    return -1

if __name__ == "__main__":
    print(int(find_next_square(49)))