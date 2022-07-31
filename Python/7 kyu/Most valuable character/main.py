def solve(st):
    alphabet = "abcdefghijklmnopqrstuvwxyz"
    letters = ""
    for i in st:
        if i not in letters: letters += i
    maxDiff = -1
    lt = ""
    for elem in letters:
        diff = st.rindex(elem) - st.index(elem)
        if diff > maxDiff:
            maxDiff = diff
            lt = elem
        elif diff == maxDiff:
            if alphabet.index(elem) < alphabet.index(lt):
                lt = elem
    return lt

if __name__ == "__main__":
    print(solve("dgljirsljsefj"))