[LeetCode] 22. Generate Parentheses 生成括号 

 
Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.
For example, given n = 3, a solution set is:
[
  "((()))",
  "(()())",
  "(())()",
  "()(())",
  "()()()"
]
  
在 LeetCode 中有关括号的题共有七道，除了这一道的另外六道是 Score of Parentheses，Valid Parenthesis String， Remove Invalid Parentheses，Different Ways to Add Parentheses，Valid Parentheses 和 Longest Valid Parentheses。这道题给定一个数字n，让生成共有n个括号的所有正确的形式，对于这种列出所有结果的题首先还是考虑用递归 Recursion 来解，由于字符串只有左括号和右括号两种字符，而且最终结果必定是左括号3个，右括号3个，所以这里定义两个变量 left 和 right 分别表示剩余左右括号的个数，如果在某次递归时，左括号的个数大于右括号的个数，说明此时生成的字符串中右括号的个数大于左括号的个数，即会出现 ')(' 这样的非法串，所以这种情况直接返回，不继续处理。如果 left 和 right 都为0，则说明此时生成的字符串已有3个左括号和3个右括号，且字符串合法，则存入结果中后返回。如果以上两种情况都不满足，若此时 left 大于0，则调用递归函数，注意参数的更新，若 right 大于0，则调用递归函数，同样要更新参数，参见代码如下：
 
C++ 解法一：

class Solution {
public:
    vector<string> generateParenthesis(int n) {
        vector<string> res;
        generateParenthesisDFS(n, n, "", res);
        return res;
    }
    void generateParenthesisDFS(int left, int right, string out, vector<string> &res) {
        if (left > right) return;
        if (left == 0 && right == 0) res.push_back(out);
        else {
            if (left > 0) generateParenthesisDFS(left - 1, right, out + '(', res);
            if (right > 0) generateParenthesisDFS(left, right - 1, out + ')', res);
        }
    }
};

 
Java 解法一：

public class Solution {
    public List<String> generateParenthesis(int n) {
        List<String> res = new ArrayList<String>();
        helper(n, n, "", res);
        return res;
    }
    void helper(int left, int right, String out, List<String> res) {
        if (left < 0 || right < 0 || left > right) return;
        if (left == 0 && right == 0) {
            res.add(out);
            return;
        }
        helper(left - 1, right, out + "(", res);
        helper(left, right - 1, out + ")", res);
    }
}

 
再来看那一种方法，这种方法是 CareerCup 书上给的方法，感觉也是满巧妙的一种方法，这种方法的思想是找左括号，每找到一个左括号，就在其后面加一个完整的括号，最后再在开头加一个 ()，就形成了所有的情况，需要注意的是，有时候会出现重复的情况，所以用set数据结构，好处是如果遇到重复项，不会加入到结果中，最后我们再把set转为vector即可，参见代码如下：：
n＝1:    ()
n=2:    (())    ()()
n=3:    (()())    ((()))    ()(())    (())()    ()()()   
 
C++ 解法二：

class Solution {
public:
    vector<string> generateParenthesis(int n) {
        unordered_set<string> st;
        if (n == 0) st.insert("");
        else {
            vector<string> pre = generateParenthesis(n - 1);
            for (auto a : pre) {
                for (int i = 0; i < a.size(); ++i) {
                    if (a[i] == '(') {
                        a.insert(a.begin() + i + 1, '(');
                        a.insert(a.begin() + i + 2, ')');
                        st.insert(a);
                        a.erase(a.begin() + i + 1, a.begin() + i + 3);
                    }
                }
                st.insert("()" + a);
            }
        }
        return vector<string>(st.begin(), st.end());
    }
};

 
Java 解法二:

public class Solution {
    public List<String> generateParenthesis(int n) {
        Set<String> res = new HashSet<String>();
        if (n == 0) {
            res.add("");
        } else {
            List<String> pre = generateParenthesis(n - 1);
            for (String str : pre) {
                for (int i = 0; i < str.length(); ++i) {
                    if (str.charAt(i) == '(') {
                        str = str.substring(0, i + 1) + "()" + str.substring(i + 1, str.length());
                        res.add(str);
                        str = str.substring(0, i + 1) +  str.substring(i + 3, str.length());
                    }
                }
                res.add("()" + str);
            }
        }
        return new ArrayList(res);
    }
}
