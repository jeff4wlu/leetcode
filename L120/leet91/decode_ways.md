[LeetCode] 91. Decode Ways 解码方法 

 
A message containing letters from A-Z is being encoded to numbers using the following mapping:
'A' -> 1
'B' -> 2
...
'Z' -> 26
Given a non-empty string containing only digits, determine the total number of ways to decode it.
Example 1:
Input: "12"
Output: 2
Explanation: It could be decoded as "AB" (1 2) or "L" (12).
Example 2:
Input: "226"
Output: 3
Explanation: It could be decoded as "BZ" (2 26), "VF" (22 6), or "BBF" (2 2 6).
 
这道题要求解码方法，跟之前那道 Climbing Stairs 非常的相似，但是还有一些其他的限制条件，比如说一位数时不能为0，两位数不能大于 26，其十位上的数也不能为0，除去这些限制条件，跟爬梯子基本没啥区别，也勉强算特殊的斐波那契数列，当然需要用动态规划 Dynamci Programming 来解。建立一维 dp 数组，其中 dp[i] 表示s中前i个字符组成的子串的解码方法的个数，长度比输入数组长多多1，并将 dp[0] 初始化为1。现在来找状态转移方程，dp[i] 的值跟之前的状态有着千丝万缕的联系，就拿题目中的例子2来分析吧，当 i=1 时，对应s中的字符是 s[0]='2'，只有一种拆分方法，就是2，注意 s[0] 一定不能为0，这样的话无法拆分。当 i=2 时，对应s中的字符是 s[1]='2'，由于 s[1] 不为0，那么其可以被单独拆分出来，就可以在之前 dp[i-1] 的每种情况下都加上一个单独的2，这样 dp[i] 至少可以有跟 dp[i-1] 一样多的拆分情况，接下来还要看其能否跟前一个数字拼起来，若拼起来的两位数小于等于26，并且大于等于 10（因为两位数的高位不能是0），那么就可以在之前 dp[i-2] 的每种情况下都加上这个二位数，所以最终 dp[i] = dp[i-1] + dp[i-2]，是不是发现跟斐波那契数列的性质吻合了。所以0是个很特殊的存在，若当前位置是0，则一定无法单独拆分出来，即不能加上 dp[i-1]，就只能看否跟前一个数字组成大于等于 10 且小于等于 26 的数，能的话可以加上 dp[i-2]，否则就只能保持为0了。具体的操作步骤是，在遍历的过程中，对每个数字首先判断其是否为0，若是则将 dp[i] 赋为0，若不是，赋上 dp[i-1] 的值，然后看数组前一位是否存在，如果存在且满足前一位是1，或者和当前位一起组成的两位数不大于 26，则当前 dp[i] 值加上 dp[i - 2]。最终返回 dp 数组的最后一个值即可，代码如下：
 
C++ 解法一：

class Solution {
public:
    int numDecodings(string s) {
        if (s.empty() || s[0] == '0') return 0;
        vector<int> dp(s.size() + 1, 0);
        dp[0] = 1;
        for (int i = 1; i < dp.size(); ++i) {
            dp[i] = (s[i - 1] == '0') ? 0 : dp[i - 1];
            if (i > 1 && (s[i - 2] == '1' || (s[i - 2] == '2' && s[i - 1] <= '6'))) {
                dp[i] += dp[i - 2];
            }
        }
        return dp.back();
    }
};

 
Java 解法一：

class Solution {
    public int numDecodings(String s) {
        if (s.isEmpty() || s.charAt(0) == '0') return 0;
        int[] dp = new int[s.length() + 1];
        dp[0] = 1;
        for (int i = 1; i < dp.length; ++i) {
            dp[i] = (s.charAt(i - 1) == '0') ? 0 : dp[i - 1];
            if (i > 1 && (s.charAt(i - 2) == '1' || (s.charAt(i - 2) == '2' && s.charAt(i - 1) <= '6'))) {
                dp[i] += dp[i - 2];
            }
        }
        return dp[s.length()];
    }
}

 
下面这种方法跟上面的方法的思路一样，只是写法略有不同：
 
C++ 解法二：

class Solution {
public:
    int numDecodings(string s) {
        if (s.empty() || s[0] == '0') return 0;
        vector<int> dp(s.size() + 1, 0);
        dp[0] = 1;
        for (int i = 1; i < dp.size(); ++i) {
            if (s[i - 1] != '0') dp[i] += dp[i - 1];
            if (i >= 2 && s.substr(i - 2, 2) <= "26" && s.substr(i - 2, 2) >= "10") {
                dp[i] += dp[i - 2];
            }
        }
        return dp.back();
    }
};

 
Java  解法二：

class Solution {
    public int numDecodings(String s) {
        if (s.isEmpty() || s.charAt(0) == '0') return 0;
        int[] dp = new int[s.length() + 1];
        dp[0] = 1;
        for (int i = 1; i < dp.length; ++i) {
            if (s.charAt(i - 1) != '0') dp[i] += dp[i - 1];
            if (i >= 2 && (s.substring(i - 2, i).compareTo("10") >= 0 && s.substring(i - 2, i).compareTo("26") <= 0)) {
                dp[i] += dp[i - 2];
            }
        }
        return dp[s.length()];
    }
}

 
我们再来看一种空间复杂度为 O(1) 的解法，用两个变量 a, b 来分别表示 s[i-1] 和 s[i-2] 的解码方法，然后从 i=1 开始遍历，也就是字符串的第二个字符，判断如果当前字符为 '0'，说明当前字符不能单独拆分出来，只能和前一个字符一起，先将 a 赋为0，然后看前面的字符，如果前面的字符是1或者2时，就可以更新 a = a + b，然后 b = a - b，其实 b 赋值为之前的 a，如果不满足这些条件的话，那么 b = a，参见代码如下：
 
C++ 解法三：

class Solution {
public:
    int numDecodings(string s) {
        if (s.empty() || s[0] == '0') return 0;
        int a = 1, b = 1, n = s.size();
        for (int i = 1; i < n; ++i) {
            if (s[i] == '0') a = 0;
            if (s[i - 1] == '1' || (s[i - 1] == '2' && s[i] <= '6')) {
                a = a + b;
                b = a - b;
            } else {
                b = a;
            }
        }
        return a;
    }
};

 
Java 解法三：

class Solution {
    public int numDecodings(String s) {
        if (s.isEmpty() || s.charAt(0) == '0') return 0;
        int a = 1, b = 1, n = s.length();
        for (int i = 1; i < n; ++i) {
            if (s.charAt(i) == '0') a = 0;
            if (s.charAt(i - 1) == '1' || (s.charAt(i - 1) == '2' && s.charAt(i) <= '6')) {
                a = a + b;
                b = a - b;
            } else {
                b = a;
            }
        }
        return a;
    }
}
