[LeetCode] Interleaving String 交织相错的字符串 

 
Given s1, s2, s3, find whether s3 is formed by the interleaving of s1 and s2.
Example 1:
Input: s1 = "aabcc", s2 = "dbbca", s3 = "aadbbcbcac"
Output: true
Example 2:
Input: s1 = "aabcc", s2 = "dbbca", s3 = "aadbbbaccc"
Output: false
 
这道求交织相错的字符串和之前那道 Word Break 的题很类似，就像我之前说的只要是遇到字符串的子序列或是匹配问题直接就上动态规划 Dynamic Programming，其他的都不要考虑，什么递归呀的都是浮云（当然带记忆数组的递归写法除外，因为这也可以算是 DP 的一种），千辛万苦的写了递归结果拿到 OJ 上妥妥 Time Limit Exceeded，能把人气昏了，所以还是直接就考虑 DP 解法省事些。一般来说字符串匹配问题都是更新一个二维 dp 数组，核心就在于找出状态转移方程。那么我们还是从题目中给的例子出发吧，手动写出二维数组 dp 如下：
 

  Ø d b b c a
Ø T F F F F F
a T F F F F F
a T T T T T F
b F T T F T F
c F F T T T T
c F F F T F T

 
首先，这道题的大前提是字符串 s1 和 s2 的长度和必须等于 s3 的长度，如果不等于，肯定返回 false。那么当 s1 和 s2 是空串的时候，s3 必然是空串，则返回 true。所以直接给 dp[0][0] 赋值 true，然后若 s1 和 s2 其中的一个为空串的话，那么另一个肯定和 s3 的长度相等，则按位比较，若相同且上一个位置为 True，赋 True，其余情况都赋 False，这样的二维数组 dp 的边缘就初始化好了。下面只需要找出状态转移方程来更新整个数组即可，我们发现，在任意非边缘位置 dp[i][j] 时，它的左边或上边有可能为 True 或是 False，两边都可以更新过来，只要有一条路通着，那么这个点就可以为 True。那么我们得分别来看，如果左边的为 True，那么我们去除当前对应的 s2 中的字符串 s2[j - 1] 和 s3 中对应的位置的字符相比（计算对应位置时还要考虑已匹配的s1中的字符），为 s3[j - 1 + i], 如果相等，则赋 True，反之赋 False。 而上边为 True 的情况也类似，所以可以求出状态转移方程为：
dp[i][j] = (dp[i - 1][j] && s1[i - 1] == s3[i - 1 + j]) || (dp[i][j - 1] && s2[j - 1] == s3[j - 1 + i]);
其中 dp[i][j] 表示的是 s2 的前 i 个字符和 s1 的前 j 个字符是否匹配 s3 的前 i+j 个字符，根据以上分析，可写出代码如下：
 
解法一：

class Solution {
public:
    bool isInterleave(string s1, string s2, string s3) {
        if (s1.size() + s2.size() != s3.size()) return false;
        int n1 = s1.size(), n2 = s2.size();
        vector<vector<bool>> dp(n1 + 1, vector<bool> (n2 + 1)); 
        dp[0][0] = true;
        for (int i = 1; i <= n1; ++i) {
            dp[i][0] = dp[i - 1][0] && (s1[i - 1] == s3[i - 1]);
        }
        for (int i = 1; i <= n2; ++i) {
            dp[0][i] = dp[0][i - 1] && (s2[i - 1] == s3[i - 1]);
        }
        for (int i = 1; i <= n1; ++i) {
            for (int j = 1; j <= n2; ++j) {
                dp[i][j] = (dp[i - 1][j] && s1[i - 1] == s3[i - 1 + j]) || (dp[i][j - 1] && s2[j - 1] == s3[j - 1 + i]);
            }
        }
        return dp[n1][n2];
    }
};

 
我们也可以把for循环合并到一起，用if条件来处理边界情况，整体思路和上面的解法没有太大的区别，参见代码如下：
 
解法二：

class Solution {
public:
    bool isInterleave(string s1, string s2, string s3) {
        if (s1.size() + s2.size() != s3.size()) return false;
        int n1 = s1.size(), n2 = s2.size();
        vector<vector<bool>> dp(n1 + 1, vector<bool> (n2 + 1, false)); 
        for (int i = 0; i <= n1; ++i) {
            for (int j = 0; j <= n2; ++j) {
                if (i == 0 && j == 0) {
                    dp[i][j] = true;
                } else if (i == 0) {
                    dp[i][j] = dp[i][j - 1] && s2[j - 1] == s3[i + j - 1];
                } else if (j == 0) {
                    dp[i][j] = dp[i - 1][j] && s1[i - 1] == s3[i + j - 1];
                } else {
                    dp[i][j] = (dp[i - 1][j] && s1[i - 1] == s3[i + j - 1]) || (dp[i][j - 1] && s2[j - 1] == s3[i + j - 1]);
                }
            }
        }
        return dp[n1][n2];
    }
};

 
这道题也可以使用带优化的 DFS 来做，我们使用一个 HashSet，用来保存匹配失败的情况，我们分别用变量i，j，和k来记录字符串 s1，s2，和 s3 匹配到的位置，初始化的时候都传入0。在递归函数中，首先根据i和j，算出 key 值，由于我们的 HashSet 中只能放一个数字，而我们要 encode 两个数字i和j，所以通过用i乘以 s3 的长度再加上j来得到 key，此时我们看，如果 key 已经在集合中，直接返回 false，因为集合中存的是无法匹配的情况。然后先来处理 corner case 的情况，如果i等于 s1 的长度了，说明 s1 的字符都匹配完了，此时 s2 剩下的字符和 s3 剩下的字符可以直接进行匹配了，所以我们直接返回两者是否能匹配的 bool 值。同理，如果j等于 s2 的长度了，说明 s2 的字符都匹配完了，此时 s1 剩下的字符和 s3 剩下的字符可以直接进行匹配了，所以我们直接返回两者是否能匹配的 bool 值。如果 s1 和 s2 都有剩余字符，那么当 s1 的当前字符等于 s3 的当前字符，那么调用递归函数，注意i和k都加上1，如果递归函数返回 true，则当前函数也返回 true；还有一种情况是，当 s2 的当前字符等于 s3 的当前字符，那么调用递归函数，注意j和k都加上1，如果递归函数返回 true，那么当前函数也返回 true。如果匹配失败了，则将 key 加入集合中，并返回 false 即可，参见代码如下：
 
解法三：

class Solution {
public:
    bool isInterleave(string s1, string s2, string s3) {
        if (s1.size() + s2.size() != s3.size()) return false;
        unordered_set<int> s;
        return helper(s1, 0, s2, 0, s3, 0, s);
    }
    bool helper(string& s1, int i, string& s2, int j, string& s3, int k, unordered_set<int>& s) {
        int key = i * s3.size() + j;
        if (s.count(key)) return false;
        if (i == s1.size()) return s2.substr(j) == s3.substr(k);
        if (j == s2.size()) return s1.substr(i) == s3.substr(k);
        if ((s1[i] == s3[k] && helper(s1, i + 1, s2, j, s3, k + 1, s)) || 
            (s2[j] == s3[k] && helper(s1, i, s2, j + 1, s3, k + 1, s))) return true;
        s.insert(key);
        return false;
    }
};

 
既然 DFS 可以，那么 BFS 也就坐不住了，也要出来浪一波。这里我们需要用队列 queue 来辅助运算，如果将解法一讲解中的那个二维 dp 数组列出来的 TF 图当作一个迷宫的话，那么 BFS 的目的就是要从 (0, 0) 位置找一条都是T的路径通到 (n1, n2) 位置，这里我们还要使用 HashSet，不过此时保存到是已经遍历过的位置，队列中还是存 key 值，key 值的 encode 方法跟上面 DFS 解法的相同，初始时放个0进去。然后我们进行 while 循环，循环条件除了q不为空，还有一个是k小于 n3，因为匹配完 s3 中所有的字符就结束了。然后由于是一层层的遍历，所以要直接循环 queue 中元素个数的次数，在 for 循环中，对队首元素进行解码，得到i和j值，如果i小于 n1，说明 s1 还有剩余字符，如果 s1 当前字符等于 s3 当前字符，那么把 s1 的下一个位置 i+1 跟j一起加码算出 key 值，如果该 key 值不在于集合中，则加入集合，同时加入队列 queue 中；同理，如果j小于 n2，说明 s2 还有剩余字符，如果 s2 当前字符等于 s3 当前字符，那么把 s2 的下一个位置 j+1 跟i一起加码算出 key 值，如果该 key 值不在于集合中，则加入集合，同时加入队列 queue 中。for 循环结束后，k自增1。最后如果匹配成功的话，那么 queue 中应该只有一个 (n1, n2) 的 key 值，且k此时等于 n3，所以当 queue 为空或者k不等于 n3 的时候都要返回 false，参见代码如下：
 
解法四：

class Solution {
public:
    bool isInterleave(string s1, string s2, string s3) {
        if (s1.size() + s2.size() != s3.size()) return false;
        int n1 = s1.size(), n2 = s2.size(), n3 = s3.size(), k = 0;
        unordered_set<int> s;
        queue<int> q{{0}};
        while (!q.empty() && k < n3) {
            int len = q.size();
            for (int t = 0; t < len; ++t) {
                int i = q.front() / n3, j = q.front() % n3; q.pop();
                if (i < n1 && s1[i] == s3[k]) {
                    int key = (i + 1) * n3 + j;
                    if (!s.count(key)) {
                        s.insert(key);
                        q.push(key);
                    }
                }
                if (j < n2 && s2[j] == s3[k]) {
                    int key = i * n3 + j + 1;
                    if (!s.count(key)) {
                        s.insert(key);
                        q.push(key);
                    }
                }
            }
            ++k;
        }
        return !q.empty() && k == n3;
    }
};
