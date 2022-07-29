def goesLower(arr, i):
    for k in range(i, len(arr)):
        if arr[k] < arr[i]:
            return True
        elif arr[k] > arr[i]:
            return False
    return False

def pick_peaks(arr):
    pos = []
    peaks = []
    for ind in range(1, len(arr) - 1):
        if arr[ind] > arr [ind - 1] and arr[ind] >= arr[ind + 1] and goesLower(arr, ind):
            pos.append(ind)
            peaks.append(arr[ind])
    return {"pos": pos, "peaks": peaks}

if __name__ == "__main__":
    print(pick_peaks([3,2,3,6,4,1,2,3,2,1,2,2,2,1]))