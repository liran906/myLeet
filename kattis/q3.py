import sys

stdread = sys.stdin.readline
stdwrite = sys.stdout.write

n = int(stdread())

pwd = []
filepaths = set()
for _ in range(n):
    cmd, arg = stdread().split()
    if cmd == 'cd':
        if arg != '..':
            pwd.append(arg)
        else:
            pwd.pop()
    elif cmd == 'nano':
        if len(pwd) == 0:
            filepaths.add(arg)
        else:
            filepaths.add('/'.join(pwd) + '/' + arg)

sorted_filepaths = sorted(list(filepaths))
for filepath in sorted_filepaths:
    stdwrite(f"git add {filepath}\n")
stdwrite('git commit\n')
stdwrite('git push')