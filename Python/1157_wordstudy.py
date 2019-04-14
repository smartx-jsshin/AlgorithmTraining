text = input().upper()

count_dict = {}
for i in text:
    count = count_dict.get(i)
    if count is None:
        count_dict[i] = 1
    else:
        count_dict[i] = count + 1

sorted_dict = [(k, count_dict[k])for k in sorted(count_dict, key=count_dict.get, reverse=True)]

(fk, fv) = sorted_dict[0]
if len(sorted_dict) == 1:
    print(fk)
else:
    (sk, sv) = sorted_dict[1]
    if fv == sv:
        print("?")
    else:
        print(fk)
