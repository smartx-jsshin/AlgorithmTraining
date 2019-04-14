input_val = int(input().encode("utf-8"))

layer = 1
sum = 0
res = None

while True:
    sum = sum + layer
    if input_val <= sum:
        remainder = sum - input_val
        if layer % 2 == 0:
            res = "{}/{}".format(layer-remainder, remainder+1)
            break
        else:
            res = "{}/{}".format(remainder+1, layer-remainder)
            break
    layer = layer + 1
print(res)
