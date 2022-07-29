def partitions(n):
    G = [int(g_pow == 0) for g_pow in range(n + 1)]

    for k in range(1, n):
        G = [G[g_pow] if g_pow - k < 0 else G[g_pow] + G[g_pow - k] for g_pow in range(n + 1)]

    return G[n] + 1

if __name__ == "__main__":
    print(partitions(5))