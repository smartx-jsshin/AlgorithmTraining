import sys

num_cases = int(sys.stdin.readline().rstrip())

i = 0

while i < num_cases:
    a, b = sys.stdin.readline().rstrip().split(' ')
    print(int(a)+int(b))
    i = i+1
