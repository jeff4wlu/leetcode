[LeetCode] 187. Repeated DNA Sequences 求重复的DNA序列 

 
All DNA is composed of a series of nucleotides abbreviated as A, C, G, and T, for example: "ACGAATTCCG". When studying DNA, it is sometimes useful to identify repeated sequences within the DNA.
Write a function to find all the 10-letter-long sequences (substrings) that occur more than once in a DNA molecule.
Example:
Input: s = "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"

Output: ["AAAAACCCCC", "CCCCCAAAAA"]
 
看到这道题想到这应该属于 CS 的一个重要分支生物信息 Bioinformatics 研究的内容，研究 DNA 序列特征的重要意义自然不用多说，但是对于我们广大码农来说，还是专注于算法吧，此题还是用位操作 Bit Manipulation 来求解，计算机由于其二进制存储的特点可以很巧妙的解决一些问题，像之前的 Single Number 和 Single Number II 都是很巧妙利用位操作来求解。此题由于构成输入字符串的字符只有四种，分别是 A, C, G, T，下面来看下它们的 ASCII 码用二进制来表示：
A: 0100 0001　　C: 0100 0011　　G: 0100 0111　　T: 0101 0100
由于目的是利用位来区分字符，当然是越少位越好，通过观察发现，每个字符的后三位都不相同，故而可以用末尾三位来区分这四个字符。而题目要求是 10 个字符长度的串，每个字符用三位来区分，10 个字符需要30位，在 32 位机上也 OK。为了提取出后 30 位，还需要用个 mask，取值为 0x7ffffff，用此 mask 可取出后27位，再向左平移三位即可。算法的思想是，当取出第十个字符时，将其存在 HashMap 里，和该字符串出现频率映射，之后每向左移三位替换一个字符，查找新字符串在 HashMap 里出现次数，如果之前刚好出现过一次，则将当前字符串存入返回值的数组并将其出现次数加一，如果从未出现过，则将其映射到1。为了能更清楚的阐述整个过程，就用题目中给的例子来分析整个过程：
首先取出前九个字符 AAAAACCCC，根据上面的分析，用三位来表示一个字符，所以这九个字符可以用二进制表示为 001001001001001011011011011，然后继续遍历字符串，下一个进来的是C，则当前字符为 AAAAACCCCC，二进制表示为 001001001001001011011011011011，然后将其存入 HashMap 中，用二进制的好处是可以用一个 int 变量来表示任意十个字符序列，比起直接存入字符串大大的节省了内存空间，然后再读入下一个字符C，则此时字符串为 AAAACCCCCA，还是存入其二进制的表示形式，以此类推，当某个序列之前已经出现过了，将其存入结果 res 中即可，参见代码如下：
 
解法一：

class Solution {
public:
    vector<string> findRepeatedDnaSequences(string s) {
        vector<string> res;
        if (s.size() <= 10) return res;
        int mask = 0x7ffffff, cur = 0;
        unordered_map<int, int> m;
        for (int i = 0; i < 9; ++i) {
            cur = (cur << 3) | (s[i] & 7);
        }
        for (int i = 9; i < s.size(); ++i) {
            cur = ((cur & mask) << 3) | (s[i] & 7);
            if (m.count(cur)) {
                if (m[cur] == 1) res.push_back(s.substr(i - 9, 10));
                ++m[cur]; 
            } else {
                m[cur] = 1;
            }
        }
        return res;
    }
};

 
上面的方法可以写的更简洁一些，这里可以用 HashSet 来代替 HashMap，只要当前的数已经在 HashSet 中存在了，就将其加入 res 中，这里 res 也定义成 HashSet，这样就可以利用 HashSet 的不能有重复项的特点，从而得到正确的答案，最后将 HashSet 转为 vector 即可，参见代码如下：
 
解法二：

class Solution {
public:
    vector<string> findRepeatedDnaSequences(string s) {
        unordered_set<string> res;
        unordered_set<int> st;
        int cur = 0;
        for (int i = 0; i < 9; ++i) cur = cur << 3 | (s[i] & 7);
        for (int i = 9; i < s.size(); ++i) {
            cur = ((cur & 0x7ffffff) << 3) | (s[i] & 7);
            if (st.count(cur)) res.insert(s.substr(i - 9, 10));
            else st.insert(cur);
        }
        return vector<string>(res.begin(), res.end());
    }
};

 
上面的方法都是用三位来表示一个字符，这里可以用两位来表示一个字符，00 表示A，01 表示C，10 表示G，11 表示T，那么总共需要 20 位就可以表示十个字符流，其余的思路跟上面的方法完全相同，注意这里的 mask 只需要表示 18 位，所以变成了 0x3ffff，参见代码如下：
 
解法三：

class Solution {
public:
    vector<string> findRepeatedDnaSequences(string s) {
        unordered_set<string> res;
        unordered_set<int> st;
        unordered_map<int, int> m{{'A', 0}, {'C', 1}, {'G', 2}, {'T', 3}};
        int cur = 0;
        for (int i = 0; i < 9; ++i) cur = cur << 2 | m[s[i]];
        for (int i = 9; i < s.size(); ++i) {
            cur = ((cur & 0x3ffff) << 2) | (m[s[i]]);
            if (st.count(cur)) res.insert(s.substr(i - 9, 10));
            else st.insert(cur);
        }
        return vector<string>(res.begin(), res.end());
    }
};

 
如果不需要考虑节省内存空间，那可以直接将 10个 字符组成字符串存入 HashSet 中，那么也就不需要 mask 啥的了，但是思路还是跟上面的方法相同:
 
解法四：

class Solution {
public:
    vector<string> findRepeatedDnaSequences(string s) {
        unordered_set<string> res, st;
        for (int i = 0; i + 9 < s.size(); ++i) {
            string t = s.substr(i, 10);
            if (st.count(t)) res.insert(t);
            else st.insert(t);
        }
        return vector<string>{res.begin(), res.end()};
    }
};
