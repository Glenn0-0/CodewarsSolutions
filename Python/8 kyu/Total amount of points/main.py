def points(games):
    points = 0
    for score in games:
        nums = score.split(":")
        if int(nums[0]) > int(nums[1]):
            points += 3
        elif int(nums[0]) == int(nums[1]):
            points += 1
    return points

if __name__ == "__main__":
    print(points(['1:0','2:0','3:0','4:0','2:1','3:1','4:1','3:2','4:2','4:3']))