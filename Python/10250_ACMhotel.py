num = int(input().encode("utf-8"))

idx = 0
while idx < num:
    val = input().split(' ')

    height = int(val[0])
    width = int(val[1])
    guest = int(val[2])

    col = int((guest-1) / height) + 1
    row = (guest-1) % height + 1

    res = "{}{:02d}".format(row, col)
    print(res)
    idx = idx + 1