from copy import copy


testFalse = [
    [5, 3, 4, 6, 7, 8, 9, 1, 2],
	[6, 7, 2, 1, 9, 0, 3, 4, 8],
	[1, 0, 0, 3, 4, 2, 5, 6, 0],
	[8, 5, 9, 7, 6, 1, 0, 2, 0],
	[4, 2, 6, 8, 5, 3, 7, 9, 1],
	[7, 1, 3, 9, 2, 4, 8, 5, 6],
	[9, 0, 1, 5, 3, 7, 2, 1, 4],
	[2, 8, 7, 4, 1, 9, 6, 3, 5],
	[3, 0, 0, 4, 8, 1, 1, 7, 9],
]

testTrue = [
	[5, 3, 4, 6, 7, 8, 9, 1, 2],
	[6, 7, 2, 1, 9, 5, 3, 4, 8],
	[1, 9, 8, 3, 4, 2, 5, 6, 7],
	[8, 5, 9, 7, 6, 1, 4, 2, 3],
	[4, 2, 6, 8, 5, 3, 7, 9, 1],
	[7, 1, 3, 9, 2, 4, 8, 5, 6],
	[9, 6, 1, 5, 3, 7, 2, 8, 4],
	[2, 8, 7, 4, 1, 9, 6, 3, 5],
	[3, 4, 5, 2, 8, 6, 1, 7, 9],
]

def validArray(row):
    crow = row.copy()
    crow.sort()
    if crow == [1,2,3,4,5,6,7,8,9]:
        return True
    return False

def validLines(board):
    for row in board:
        if not validArray(row):
            return False
    return True

def validColumns(board):
    for col in range(0, 9):
        column = []
        for row in range(0, 9):
            column.append(board[row][col])
        
        if not validArray(column):
            return False

    return True

def validSquare(board, a, b):
    square = []
    for x in range(a, a+3):
        for y in range(b, b+3):
            square.append(board[x][y])
    return validArray(square)

def validSquares(board):
    for i in range(0, 7, 3):
        for j in range(0, 7, 3):
            if not validSquare(board, i, j):
                return False
    return True

def ValidateSolution(board):
    if validLines(board) and validColumns(board) and validSquares(board):
        return True
    return False

if __name__ == "__main__":
    print(ValidateSolution(testFalse))
    print(ValidateSolution(testTrue))