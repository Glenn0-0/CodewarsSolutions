def scoreCount(word):
    score = 0
    letters = "0abcdefghijklmnopqrstuvwxyz"
    for letter in word:
        score += letters.index(letter)
    return score

def high(x):
    arr = x.split()
    max = 0
    res = ""
    for word in arr:
        score = scoreCount(word)
        if score > max:
            max = score
            res = word
    return res

print(high("man i need a taxi up to ubud"))