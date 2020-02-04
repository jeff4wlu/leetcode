[LeetCode] Number of Islands 岛屿的数量 

 
Given a 2d grid map of '1's (land) and '0's (water), count the number of islands. An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically. You may assume all four edges of the grid are all surrounded by water.
Example 1:
Input:
11110
11010
11000
00000

Output: 1
Example 2:
Input:
11000
11000
00100
00011

Output: 3
Credits:
Special thanks to @mithmatt for adding this problem and creating all test cases.
 
这道求岛屿数量的题的本质是求矩阵中连续区域的个数，很容易想到需要用深度优先搜索 DFS 来解，我们需要建立一个 visited 数组用来记录某个位置是否被访问过，对于一个为 ‘1’ 且未被访问过的位置，我们递归进入其上下左右位置上为 ‘1’ 的数，将其 visited 对应值赋为 true，继续进入其所有相连的邻位置，这样可以将这个连通区域所有的数找出来，并将其对应的 visited 中的值赋 true，找完相邻区域后，我们将结果 res 自增1，然后我们在继续找下一个为 ‘1’ 且未被访问过的位置，以此类推直至遍历完整个原数组即可得到最终结果，代码如下：
 
解法一：

class Solution {
public:
    int numIslands(vector<vector<char>>& grid) {
        if (grid.empty() || grid[0].empty()) return 0;
        int m = grid.size(), n = grid[0].size(), res = 0;
        vector<vector<bool>> visited(m, vector<bool>(n));
        for (int i = 0; i < m; ++i) {
            for (int j = 0; j < n; ++j) {
                if (grid[i][j] == '0' || visited[i][j]) continue;
                helper(grid, visited, i, j);
                ++res;
            }
        }
        return res;
    }
    void helper(vector<vector<char>>& grid, vector<vector<bool>>& visited, int x, int y) {
        if (x < 0 || x >= grid.size() || y < 0 || y >= grid[0].size() || grid[x][y] == '0' || visited[x][y]) return;
        visited[x][y] = true;
        helper(grid, visited, x - 1, y);
        helper(grid, visited, x + 1, y);
        helper(grid, visited, x, y - 1);
        helper(grid, visited, x, y + 1);
    }
};

 
当然，这种类似迷宫遍历的题目 DFS 和 BFS 两对好基友肯定是形影不离的，那么 BFS 搞起。其实也很简单，就是在遍历到 ‘1’ 的时候，且该位置没有被访问过，那么就调用一个 BFS 即可，借助队列 queue 来实现，现将当前位置加入队列，然后进行 while 循环，将队首元素提取出来，并遍历其周围四个位置，若没有越界的话，就将 visited 中该邻居位置标记为 true，并将其加入队列中等待下次遍历即可，参见代码如下：
 
解法二：

class Solution {
public:
    int numIslands(vector<vector<char>>& grid) {
        if (grid.empty() || grid[0].empty()) return 0;
        int m = grid.size(), n = grid[0].size(), res = 0;
        vector<vector<bool>> visited(m, vector<bool>(n));
        vector<int> dirX{-1, 0, 1, 0}, dirY{0, 1, 0, -1};
        for (int i = 0; i < m; ++i) {
            for (int j = 0; j < n; ++j) {
                if (grid[i][j] == '0' || visited[i][j]) continue;
                ++res;
                queue<int> q{{i * n + j}};
                while (!q.empty()) {
                    int t = q.front(); q.pop();
                    for (int k = 0; k < 4; ++k) {
                        int x = t / n + dirX[k], y = t % n + dirY[k];
                        if (x < 0 || x >= m || y < 0 || y >= n || grid[x][y] == '0' || visited[x][y]) continue;
                        visited[x][y] = true;
                        q.push(x * n + y);
                    }
                }
            }
        }
        return res;
    }
};

 
类似题目：
Number of Islands II
Surrounded Regions
Walls and Gates
Number of Connected Components in an Undirected Graph 
Number of Distinct Islands 
Max Area of Island