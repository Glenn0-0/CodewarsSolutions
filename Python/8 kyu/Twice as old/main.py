def twice_as_old(dad_years_old, son_years_old):
    for n in range(200):
        if (dad_years_old - n) == 2*(son_years_old - n) or (dad_years_old + n) == 2*(son_years_old + n):
            return n

if __name__ == "__main__":
    print(twice_as_old(29, 7))