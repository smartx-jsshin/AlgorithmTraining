num = int(input().encode("utf-8"))

out_idx = 0
while out_idx < num:
    a = input().split(' ')
    dist = int(a[1]) - int(a[0])

    in_idx = 1
    res = 0

    while True:
        if dist <= in_idx:
            res = res + 1
            break
        elif dist <= (2 * in_idx):
            res = res + 2
            break
        dist = dist - (2 * in_idx)
        res = res + 2
        in_idx = in_idx + 1

    print(res)
    out_idx = out_idx + 1
