[LeetCode] 87. Scramble String 搅乱字符串 

 
Given a string s1, we may represent it as a binary tree by partitioning it to two non-empty substrings recursively.
Below is one possible representation of s1 = "great":
    great
   /    \
  gr    eat
 / \    /  \
g   r  e   at
           / \
          a   t
To scramble the string, we may choose any non-leaf node and swap its two children.
For example, if we choose the node "gr" and swap its two children, it produces a scrambled string "rgeat".
    rgeat
   /    \
  rg    eat
 / \    /  \
r   g  e   at
           / \
          a   t
We say that "rgeat" is a scrambled string of "great".
Similarly, if we continue to swap the children of nodes "eat" and "at", it produces a scrambled string "rgtae".
    rgtae
   /    \
  rg    tae
 / \    /  \
r   g  ta  e
       / \
      t   a
We say that "rgtae" is a scrambled string of "great".
Given two strings s1 and s2 of the same length, determine if s2 is a scrambled string of s1.
Example 1:
Input: s1 = "great", s2 = "rgeat"
Output: true
Example 2:
Input: s1 = "abcde", s2 = "caebd"
Output: false
 
这道题定义了一种搅乱字符串，就是说假如把一个字符串当做一个二叉树的根，然后它的非空子字符串是它的子节点，然后交换某个子字符串的两个子节点，重新爬行回去形成一个新的字符串，这个新字符串和原来的字符串互为搅乱字符串。这道题可以用递归 Recursion 或是动态规划 Dynamic Programming 来做，我们先来看递归的解法，参见网友 uniEagle 的博客，简单的说，就是 s1 和 s2 是 scramble 的话，那么必然存在一个在 s1 上的长度 l1，将 s1 分成 s11 和 s12 两段，同样有 s21 和 s22，那么要么 s11 和 s21 是 scramble 的并且 s12 和 s22 是 scramble 的；要么 s11 和 s22 是 scramble 的并且 s12 和 s21 是 scramble 的。就拿题目中的例子 rgeat 和 great 来说，rgeat 可分成 rg 和 eat 两段， great 可分成 gr 和 eat 两段，rg 和 gr 是 scrambled 的， eat 和 eat 当然是 scrambled。根据这点，我们可以写出代码如下：
 
解法一：

// Recursion
class Solution {
public:
    bool isScramble(string s1, string s2) {
        if (s1.size() != s2.size()) return false;
        if (s1 == s2) return true;
        string str1 = s1, str2 = s2;
        sort(str1.begin(), str1.end());
        sort(str2.begin(), str2.end());
        if (str1 != str2) return false;
        for (int i = 1; i < s1.size(); ++i) {
            string s11 = s1.substr(0, i);
            string s12 = s1.substr(i);
            string s21 = s2.substr(0, i);
            string s22 = s2.substr(i);
            if (isScramble(s11, s21) && isScramble(s12, s22)) return true;
            s21 = s2.substr(s1.size() - i);
            s22 = s2.substr(0, s1.size() - i);
            if (isScramble(s11, s21) && isScramble(s12, s22)) return true;
        }
        return false;
    }
};

 
当然，这道题也可以用动态规划 Dynamic Programming，根据以往的经验来说，根字符串有关的题十有八九可以用 DP 来做，那么难点就在于如何找出状态转移方程。参见网友 Code Ganker 的博客，这其实是一道三维动态规划的题目，使用一个三维数组 dp[i][j][n]，其中i是 s1 的起始字符，j是 s2 的起始字符，而n是当前的字符串长度，dp[i][j][len] 表示的是以i和j分别为 s1 和 s2 起点的长度为 len 的字符串是不是互为 scramble。有了 dp 数组接下来看看状态转移方程，也就是怎么根据历史信息来得到 dp[i][j][len]。判断这个是不是满足，首先是把当前 s1[i...i+len-1] 字符串劈一刀分成两部分，然后分两种情况：第一种是左边和 s2[j...j+len-1] 左边部分是不是 scramble，以及右边和 s2[j...j+len-1] 右边部分是不是 scramble；第二种情况是左边和 s2[j...j+len-1] 右边部分是不是 scramble，以及右边和 s2[j...j+len-1] 左边部分是不是 scramble。如果以上两种情况有一种成立，说明 s1[i...i+len-1] 和 s2[j...j+len-1] 是 scramble 的。而对于判断这些左右部分是不是 scramble 是有历史信息的，因为长度小于n的所有情况都在前面求解过了（也就是长度是最外层循环）。上面说的是劈一刀的情况，对于 s1[i...i+len-1] 有 len-1 种劈法，在这些劈法中只要有一种成立，那么两个串就是 scramble 的。总结起来状态转移方程是：
dp[i][j][len] = || (dp[i][j][k] && dp[i+k][j+k][len-k] || dp[i][j+len-k][k] && dp[i+k][j][len-k])
对于所有 1<=k<len，也就是对于所有 len-1 种劈法的结果求或运算。因为信息都是计算过的，对于每种劈法只需要常量操作即可完成，因此求解递推式是需要 O(len)（因为 len-1 种劈法）。如此总时间复杂度因为是三维动态规划，需要三层循环，加上每一步需要线行时间求解递推式，所以是 O(n^4)。虽然已经比较高了，但是至少不是指数量级的，动态规划还是有很大优势的，空间复杂度是 O(n^3)。代码如下：
 
解法二：

// DP
class Solution {
public:
    bool isScramble(string s1, string s2) {
        if (s1.size() != s2.size()) return false;
        if (s1 == s2) return true;
        int n = s1.size();
        vector<vector<vector<bool>>> dp (n, vector<vector<bool>>(n, vector<bool>(n + 1)));
        for (int len = 1; len <= n; ++len) {
            for (int i = 0; i <= n - len; ++i) {
                for (int j = 0; j <= n - len; ++j) {
                    if (len == 1) {
                        dp[i][j][1] = s1[i] == s2[j];
                    } else {
                        for (int k = 1; k < len; ++k) {
                            if ((dp[i][j][k] && dp[i + k][j + k][len - k]) || (dp[i + k][j][len - k] && dp[i][j + len - k][k])) {
                                dp[i][j][len] = true;
                            }
                        }
                    }                
                }
            }
        }
        return dp[0][0][n];
    }
};

 
上面的代码的实现过程如下，首先按单个字符比较，判断它们之间是否是 scrambled 的。在更新第二个表中第一个值 (gr 和 rg 是否为 scrambled 的)时，比较了第一个表中的两种构成，一种是 g与r, r与g，另一种是 g与g, r与r，其中后者是真，只要其中一个为真，则将该值赋真。其实这个原理和之前递归的原理很像，在判断某两个字符串是否为 scrambled 时，比较它们所有可能的拆分方法的子字符串是否是 scrambled 的，只要有一个种拆分方法为真，则比较的两个字符串一定是 scrambled 的。比较 rge 和 gre 的实现过程如下所示：

     r    g    e
g    x    √    x
r    √    x    x
e    x    x    √


     rg    ge
gr    √    x
re    x    x


     rge
gre   √

 
DP 的另一种写法，参考网友加载中..的博客，思路都一样，代码如下：
 
解法三：

// Still DP
class Solution {
public:
    bool isScramble(string s1, string s2) {
        if (s1.size() != s2.size()) return false;
        if (s1 == s2) return true;
        int n = s1.size();
        vector<vector<vector<bool>>> dp (n, vector<vector<bool>>(n, vector<bool>(n + 1)));
        for (int i = n - 1; i >= 0; --i) {
            for (int j = n - 1; j >= 0; --j) {
                for (int k = 1; k <= n - max(i, j); ++k) {
                    if (s1.substr(i, k) == s2.substr(j, k)) {
                        dp[i][j][k] = true;
                    } else {
                        for (int t = 1; t < k; ++t) {
                            if ((dp[i][j][t] && dp[i + t][j + t][k - t]) || (dp[i][j + k - t][t] && dp[i + t][j][k - t])) {
                                dp[i][j][k] = true;
                                break;
                            }
                        }
                    }
                }
            }
        }
        return dp[0][0][n];
    }
};

 
下面这种解法和第一个解法思路相同，只不过没有用排序算法，而是采用了类似于求异构词的方法，用一个数组来保存每个字母出现的次数，后面判断 Scramble 字符串的方法和之前的没有区别：
 
解法四：

class Solution {
public:
    bool isScramble(string s1, string s2) {
        if (s1 == s2) return true;
        if (s1.size() != s2.size()) return false;
        int n = s1.size(), m[26] = {0};
        for (int i = 0; i < n; ++i) {
            ++m[s1[i] - 'a'];
            --m[s2[i] - 'a'];
        }
        for (int i = 0; i < 26; ++i) {
            if (m[i] != 0) return false;
        }
        for (int i = 1; i < n; ++i) {
            if ((isScramble(s1.substr(0, i), s2.substr(0, i)) && isScramble(s1.substr(i), s2.substr(i))) || (isScramble(s1.substr(0, i), s2.substr(n - i)) && isScramble(s1.substr(i), s2.substr(0, n - i)))) {
                return true;
            }
        }
        return false;
    }
};

 
下面这种解法实际上是解法二的递归形式，我们用了 memo 数组来减少了大量的运算，注意这里的 memo 数组一定要有三种状态，初始化为 -1，区域内为 scramble 是1，不是 scramble 是0，这样就避免了已经算过了某个区间，但由于不是 scramble，从而又进行一次计算，从而会 TLE，感谢网友 bambu 的提供的思路，参见代码如下：
 
解法五：

class Solution {
public:
    bool isScramble(string s1, string s2) {
        if (s1 == s2) return true;
        if (s1.size() != s2.size()) return false;
        int n = s1.size();
        vector<vector<vector<int>>> memo(n, vector<vector<int>>(n, vector<int>(n + 1, -1)));
        return helper(s1, s2, 0, 0, n, memo);
    }
    bool helper(string& s1, string& s2, int idx1, int idx2, int len, vector<vector<vector<int>>>& memo) {
        if (len == 0) return true;
        if (len == 1) memo[idx1][idx2][len] = s1[idx1] == s2[idx2];
        if (memo[idx1][idx2][len] != -1) return memo[idx1][idx2][len];
        for (int k = 1; k < len; ++k) {
            if ((helper(s1, s2, idx1, idx2, k, memo) && helper(s1, s2, idx1 + k, idx2 + k, len - k, memo)) || (helper(s1, s2, idx1, idx2 + len - k, k, memo) && helper(s1, s2, idx1 + k, idx2, len - k, memo))) {
                return memo[idx1][idx2][len] = 1;
            }
        }
        return memo[idx1][idx2][len] = 0;
    }
};