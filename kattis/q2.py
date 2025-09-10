import sys
import heapq
stdread = sys.stdin.readline
stdwrite = sys.stdout.write

n = int(stdread())
tasks = []
max_heap = []
for _ in range(n):
    tasks.append(list(map(int, stdread().split())))

tasks.sort(key=lambda x: x[1])
total_time = 0
drinks = 0
for t, ddl in tasks:
    heapq.heappush(max_heap, -t)
    total_time += t
    while total_time > ddl:
        longest = -heapq.heappop(max_heap)
        total_time -= longest
        longest //= 2
        heapq.heappush(max_heap, -longest)
        total_time += longest
        drinks += 1
stdwrite(str(drinks))