def multiplication_table(size):
    table = []
    for i in range(1, size + 1):
        row = []
        for k in range(i, i*size+1):
            if k%i == 0:
                row.append(k)
        table.append(row)
    return table

print(multiplication_table(5))