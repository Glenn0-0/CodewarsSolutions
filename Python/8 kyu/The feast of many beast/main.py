def feast(beast, dish):
    return beast[0] == dish[0] and beast[-1] == dish[-1]

if __name__ == "__main__":
    print(feast("great blue heron", "garlic naan"))