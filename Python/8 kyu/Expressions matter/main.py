def expression_matter(a, b, c):
    v1 = a*b*c
    v2 = a*b+c
    v3 = a+b*c
    v4 = a+b+c
    v5 = (a+b)*c
    v6 = a*(b+c)
    m = max(v1, v2, v3, v4, v5, v6)
    return m

if __name__ == "__main__":
    print(expression_matter(3, 4, 5))