def to_camel_case(text):
    res = ""
    arr = []
    text = text.replace("-", "_")
    arr = text.split("_")

    for i in range(len(arr)):
        if i == 0:
            res += arr[i]
        else:
            res += arr[i].capitalize()
    return res

print(to_camel_case("the-very_beginning"))