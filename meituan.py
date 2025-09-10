from collections import deque

class Solution:
    def calcTotalSumOfAllPat(self, m, n):
        # Write Code Here 
        if n == 1 or m == 1:
            return 1
        dir = [(1, 0), (0, 1)]
        count = 0

        frontier = deque()
        frontier.append((0, 0))
        while frontier:
            current = frontier.popleft()
            if current == (m - 1, n - 1):
                count += 1
            for d in dir:
                r, c = current[0] + d[0], current[1] + d[1]
                if r < m and c < n:
                    frontier.append((r, c))
        return count

    def calcTotalSumOfAllPat(self, m, n):
        if n == 1 or m == 1:
            return 1
        grid = [[0] * n for _ in range(m)]
        grid[m-1][n-1] = 1

        for r in range(m-1, -1, -1):
            for c in range(n-1, -1, -1):
                if r == m-1 and c == n-1:
                    continue
                if r == m-1:
                    grid[r][c] = grid[r][c+1]
                elif c == n-1:
                    grid[r][c] = grid[r+1][c]
                else:
                    grid[r][c] = grid[r+1][c] + grid[r][c+1]
        return grid[0][0]