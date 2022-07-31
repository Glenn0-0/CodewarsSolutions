def reverse_seq(n):
    res = []
    for i in range(n, 0, -1):
        res.append(i)
    return res

if __name__ == "__main__":
    print(reverse_seq(5))