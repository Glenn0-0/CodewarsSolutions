from audioop import mul


def multi_table(number):
    res = ""
    for i in range(1, 11):
        if i != 10:
            res += str(i) + " * " + str(number) + " = " + str(i*number) + "\n"
        else:
            res += str(i) + " * " + str(number) + " = " + str(i*number)
    return res

if __name__ == "__main__":
    print(multi_table(4))