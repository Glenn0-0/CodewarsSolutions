def find_height(cubes):
    count = 0
    layer = -1
    while cubes >= count:
        cubes -= count
        count += layer + 2
        layer += 1
    return layer

if __name__ == "__main__":
    print(find_height(32))