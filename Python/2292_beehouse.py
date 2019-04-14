user_input = int(input().encode("utf-8"))

sum = 1
i = 0

while True:
    sum = sum + 6 * i
    if user_input <= sum:
        break
    if sum >= 1000000000:
        break
    i = i+1
print(i+1)
