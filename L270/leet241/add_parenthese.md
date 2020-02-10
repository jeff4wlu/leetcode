[LeetCode] 241. Different Ways to Add Parentheses 添加括号的不同方式 

 
Given a string of numbers and operators, return all possible results from computing all the different possible ways to group numbers and operators. The valid operators are +, - and *.
Example 1:
Input: "2-1-1"
Output: [0, 2]
Explanation: 
((2-1)-1) = 0 
(2-(1-1)) = 2
Example 2:
Input: "2*3-4*5"
Output: [-34, -14, -10, -10, 10]
Explanation: 
(2*(3-(4*5))) = -34 
((2*3)-(4*5)) = -14 
((2*(3-4))*5) = -10 
(2*((3-4)*5)) = -10 
(((2*3)-4)*5) = 10
 
这道题让给了一个可能含有加减乘的表达式，让我们在任意位置添加括号，求出所有可能表达式的不同值。这道题乍一看感觉还蛮难的，给人的感觉是既要在不同的位置上加括号，又要计算表达式的值，结果一看还是道 Medium 的题，直接尼克杨问号脸？！遇到了难题不要害怕，从最简单的例子开始分析，慢慢的找规律，十有八九就会在分析的过程中灵光一现，找到了破题的方法。这道题貌似默认输入都必须是合法的，虽然题目中没有明确的指出这一点，所以我们也就不必进行 valid 验证了。先从最简单的输入开始，若 input 是空串，那就返回一个空数组。若 input 是一个数字的话，那么括号加与不加其实都没啥区别，因为不存在计算，但是需要将字符串转为整型数，因为返回的是一个整型数组。当然，input 是一个单独的运算符这种情况是不存在的，因为前面说了这道题默认输入的合法的。下面来看若 input 是数字和运算符的时候，比如 "1+1" 这种情况，那么加不加括号也没有任何影响，因为只有一个计算，结果一定是2。再复杂一点的话，比如题目中的例子1，input 是 "2-1-1" 时，就有两种情况了，(2-1)-1 和 2-(1-1)，由于括号的不同，得到的结果也不同，但如果我们把括号里的东西当作一个黑箱的话，那么其就变为 ()-1  和 2-()，其最终的结果跟括号内可能得到的值是息息相关的，那么再 general 一点，实际上就可以变成 () ? () 这种形式，两个括号内分别是各自的表达式，最终会分别计算得到两个整型数组，中间的问号表示运算符，可以是加，减，或乘。那么问题就变成了从两个数组中任意选两个数字进行运算，瞬间变成我们会做的题目了有木有？而这种左右两个括号代表的黑盒子就交给递归去计算，像这种分成左右两坨的 pattern 就是大名鼎鼎的分治法 Divide and Conquer 了，是必须要掌握的一个神器。类似的题目还有之前的那道 Unique Binary Search Trees II 用的方法一样，用递归来解，划分左右子树，递归构造。
好，继续来说这道题，我们不用新建递归函数，就用其本身来递归就行，先建立一个结果 res 数组，然后遍历 input 中的字符，根据上面的分析，我们希望在每个运算符的地方，将 input 分成左右两部分，从而扔到递归中去计算，从而可以得到两个整型数组 left 和 right，分别表示作用两部分各自添加不同的括号所能得到的所有不同的值，此时我们只要分别从两个数组中取数字进行当前的运算符计算，然后把结果存到 res 中即可。当然，若最终结果 res 中还是空的，那么只有一种情况，input 本身就是一个数字，直接转为整型存入结果 res 中即可，参见代码如下：
 
解法一：

class Solution {
public:
    vector<int> diffWaysToCompute(string input) {
        vector<int> res;
        for (int i = 0; i < input.size(); ++i) {
            if (input[i] == '+' || input[i] == '-' || input[i] == '*') {
                vector<int> left = diffWaysToCompute(input.substr(0, i));
                vector<int> right = diffWaysToCompute(input.substr(i + 1));
                for (int j = 0; j < left.size(); ++j) {
                    for (int k = 0; k < right.size(); ++k) {
                        if (input[i] == '+') res.push_back(left[j] + right[k]);
                        else if (input[i] == '-') res.push_back(left[j] - right[k]);
                        else res.push_back(left[j] * right[k]);
                    }
                }
            }
        }
        if (res.empty()) res.push_back(stoi(input));
        return res;
    }
};

 
我们也可以使用 HashMap 来保存已经计算过的情况，这样可以减少重复计算，从而提升运算速度，以空间换时间，岂不美哉，参见代码如下：
 
解法二：

class Solution {
public:
    unordered_map<string, vector<int>> memo;
    vector<int> diffWaysToCompute(string input) {
        if (memo.count(input)) return memo[input];
        vector<int> res;
        for (int i = 0; i < input.size(); ++i) {
            if (input[i] == '+' || input[i] == '-' || input[i] == '*') {
                vector<int> left = diffWaysToCompute(input.substr(0, i));
                vector<int> right = diffWaysToCompute(input.substr(i + 1));
                for (int j = 0; j < left.size(); ++j) {
                    for (int k = 0; k < right.size(); ++k) {
                        if (input[i] == '+') res.push_back(left[j] + right[k]);
                        else if (input[i] == '-') res.push_back(left[j] - right[k]);
                        else res.push_back(left[j] * right[k]);
                    }
                }
            }
        }
        if (res.empty()) res.push_back(stoi(input));
        memo[input] = res;
        return res;
    }
};

 
当然，这道题还可以用动态规划 Dynamic Programming 来做，但明显没有分治法来的简单，但是既然论坛里这么多陈独秀同学，博主还是要给以足够的尊重的。这里用一个三维数组 dp，其中 dp[i][j] 表示在第i个数字到第j个数字之间范围内的子串添加不同括号所能得到的不同值的整型数组，所以是个三位数组，需要注意的是我们需要对 input 字符串进行预处理，将数字跟操作分开，加到一个字符串数组 ops 中，并统计数字的个数 cnt，用这个 cnt 来初始化 dp 数组的大小，并同时要把 dp[i][j] 的数组中都加上第i个数字，通过 ops[i*2] 取得，当然还需要转为整型数。既然 dp 是个三维数组，那么肯定要用3个 for 循环来更新，这里采用的更新顺序跟之前那道 Burst Balloons 是一样的，都是从小区间往大区间更新，由于小区间之前更新过，所以我们将数字分为两部分 [i, j] 和 [j, i+len]，然后分别取出各自的数组 dp[i][j] 和 dp[j][i+len]，把对应的运算符也取出来，现在又变成了两个数组中任取两个数字进行运算，又整两个 for 循环，所以总共整了5个 for 循环嵌套，啊呀妈呀，看这整的，看不晕你算我输，参见代码如下：
 
解法三：

class Solution {
public:
    vector<int> diffWaysToCompute(string input) {
        if (input.empty()) return {};
        vector<string> ops;
        int n = input.size();
        for (int i = 0; i < n; ++i) {
            int j = i;
            while (j < n && isdigit(input[j])) ++j;
            ops.push_back(input.substr(i, j - i));
            if (j < n) ops.push_back(input.substr(j, 1));
            i = j;
        }
        int cnt = (ops.size() + 1) / 2;
        vector<vector<vector<int>>> dp(cnt, vector<vector<int>>(cnt, vector<int>()));
        for (int i = 0; i < cnt; ++i) dp[i][i].push_back(stoi(ops[i * 2]));
        for (int len = 0; len < cnt; ++len) {
            for (int i = 0; i < cnt - len; ++i) {
                for (int j = i; j < i + len; ++j) {
                    vector<int> left = dp[i][j], right = dp[j + 1][i + len];
                    string op = ops[j * 2 + 1];
                    for (int num1 : left) {
                        for (int num2 : right) {
                            if (op == "+") dp[i][i + len].push_back(num1 + num2);
                            else if (op == "-") dp[i][i + len].push_back(num1 - num2);
                            else dp[i][i + len].push_back(num1 * num2);
                        }
                    }
                }
            }
        }
        return dp[0][cnt - 1];
    }
};

 
类似题目：
Remove Invalid Parentheses
Longest Valid Parentheses
Generate Parentheses
Valid Parentheses
Unique Binary Search Trees II
Basic Calculator
Expression Add Operators