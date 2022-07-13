def solution(roman):
    sum = 0
    nums = ["IV", "IX", "XL", "XC", "CD", "CM", "I", "V", "X", "L", "C", "D", "M"]
    romans = {"IV": 4, "IX": 9, "XL": 40, "XC": 90, "CD": 400, "CM": 900, "I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
    
    for num in nums:
        sum += roman.count(num)*romans[num]
        roman = roman.replace(num, "")
    return sum

if __name__ == "__main__":
    print(solution("XXIV"))