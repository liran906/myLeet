import sys
stdr = sys.stdin.readline
stdw = sys.stdout.write

n = int(stdr())
lst = []
for _ in range(n):
    line = stdr().rstrip()[::-1]
    lst.append(line)

for line in reversed(lst):
    stdw(line)