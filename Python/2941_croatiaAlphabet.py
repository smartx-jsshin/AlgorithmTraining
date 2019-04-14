text = input()

idx = len(text) - 1
word_count = 0

while idx >= 0:
    c = text[idx]
    if c == "=":
        c = text[idx-1]
        if c == "z":
            idx = idx - 1
            if text[idx-1] == "d":
                idx = idx -1
        elif c == "c" or c == "s":
            idx = idx - 1
    elif c == "-":
        c = text[idx-1]
        if c == "c" or c == "d":
            idx = idx - 1
    elif c == "j":
        c = text[idx-1]
        if c == "l" or c == "n":
            idx = idx - 1
    word_count = word_count + 1
    idx = idx - 1
print(word_count)
