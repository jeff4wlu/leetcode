[LeetCode] 214. Shortest Palindrome 最短回文串 

 
Given a string s, you are allowed to convert it to a palindrome by adding characters in front of it. Find and return the shortest palindrome you can find by performing this transformation.
Example 1:
Input: "aacecaaa"
Output: "aaacecaaa"
Example 2:
Input: "abcd"
Output: "dcbabcd"
Credits:
Special thanks to @ifanchu for adding this problem and creating all test cases. Thanks to @Freezen for additional test cases.
 
这道题让我们求最短的回文串，LeetCode 中关于回文串的其他的题目有 Palindrome Number，Validate Palindrome，Palindrome Partitioning，Palindrome Partitioning II 和 Longest Palindromic Substring。题目让在给定字符串s的前面加上最少个字符，使之变成回文串，来看题目中给的两个例子，最坏的情况下是s中没有相同的字符，那么最小需要添加字符的个数为 s.size() - 1 个，第一个例子的字符串包含一个回文串，只需再在前面添加一个字符即可，还有一点需要注意的是，前面添加的字符串都是从s的末尾开始，一位一位往前添加的，那么只需要知道从s末尾开始需要添加到前面的个数。
首先还是先将待处理的字符串s翻转得到t，然后比较原字符串s和翻转字符串t，从第一个字符开始逐一比较，如果相等，说明s本身就是回文串，不用添加任何字符，直接返回即可；如果不相等，s去掉最后一位，t去掉第一位，继续比较，以此类推直至有相等，或者循环结束，这样就能将两个字符串在正确的位置拼接起来了，代码请参见评论区5楼。但这种写法却会超时 TLE，无法通过 OJ，所以需要用一些比较巧妙的方法来解。
这里使用双指针来找出字符串s的最长回文前缀的大概范围，这里所谓的最长回文前缀是指从开头开始到某个位置为止是回文串，比如 "abbac" 这个子串，可以知道前四个字符组成的回文串 "abba" 就是最长回文前缀。方法是指针i和j分别指向s串的开头和末尾，若 s[i] 和 s[j] 相等，则i自增1，j自减1，否则只有j自减1。需要注意的是，这样遍历一遍后，得到的范围 [0, i) 中的子串并不一定就是最大回文前缀，也可能还需要添加字符，举个例子来说，对于 "adcba"，在 for 循环执行之后，i=2，可以发现前面的 "ad" 并不是最长回文前缀，其本身甚至不是回文串，需要再次调用递归函数来填充使其变为回文串，但可以确定的是可以添加最少的字符数让其变为回文串。而且可以确定的是后面剩余的部分肯定不属于回文前缀，此时提取出剩下的字符，翻转一下加到最前面，而对范围 [0, i) 内的子串再次递归调用本函数，这样，在子函数最终会组成最短的回文串，从而使得整个的回文串就是最短的，参见代码如下：
 
C++ 解法一：

class Solution {
public:
    string shortestPalindrome(string s) {
        int i = 0, n = s.size();
        for (int j = n - 1; j >= 0; --j) {
            if (s[i] == s[j]) ++i;
        }
        if (i == n) return s;
        string rem = s.substr(i);
        reverse(rem.begin(), rem.end());
        return rem + shortestPalindrome(s.substr(0, i)) + s.substr(i);
    }
};

 
Java 解法一：

public class Solution {
    public String shortestPalindrome(String s) {
        int i = 0, n = s.length();
        for (int j = n - 1; j >= 0; --j) {
            if (s.charAt(i) == s.charAt(j)) ++i;
        }
        if (i == n) return s;
        String rem = s.substring(i);
        String rem_rev = new StringBuilder(rem).reverse().toString();
        return rem_rev + shortestPalindrome(s.substring(0, i)) + rem;
    }
}

 
其实这道题的最快的解法是使用 KMP 算法，KMP 算法是一种专门用来匹配字符串的高效的算法，具体方法可以参见博主之前的这篇博文 KMP Algorithm 字符串匹配算法KMP小结。把s和其转置r连接起来，中间加上一个其他字符，形成一个新的字符串t，还需要一个和t长度相同的一位数组 next，其中 next[i] 表示从 t[i] 到开头的子串的相同前缀后缀的个数，具体可参考 KMP 算法中解释。最后把不相同的个数对应的字符串添加到s之前即可，代码如下：
 
C++ 解法二：

class Solution {
public:
    string shortestPalindrome(string s) {
        string r = s;
        reverse(r.begin(), r.end());
        string t = s + "#" + r;
        vector<int> next(t.size(), 0);
        for (int i = 1; i < t.size(); ++i) {
            int j = next[i - 1];
            while (j > 0 && t[i] != t[j]) j = next[j - 1];
            next[i] = (j += t[i] == t[j]);
        }
        return r.substr(0, s.size() - next.back()) + s;
    }
};

 
Java 解法二：

public class Solution {
    public String shortestPalindrome(String s) {
        String r = new StringBuilder(s).reverse().toString();
        String t = s + "#" + r;
        int[] next = new int[t.length()];
        for (int i = 1; i < t.length(); ++i) {
            int j = next[i - 1];
            while (j > 0 && t.charAt(i) != t.charAt(j)) j = next[j - 1];
            j += (t.charAt(i) == t.charAt(j)) ? 1 : 0;
            next[i] = j;
        }
        return r.substring(0, s.length() - next[t.length() - 1]) + s;
    }
}

 
从上面的 Java 和 C++ 的代码中，可以看出来 C++ 和 Java 在使用双等号上的明显的不同，感觉 Java 对于双等号对使用更加苛刻一些，比如 Java 中的双等号只对 primitive 类数据结构(比如 int, char 等)有效，但是即便有效，也不能将结果直接当1或者0来用。而对于一些从 Object 派生出来的类，比如 Integer 或者 String 等，不能直接用双等号来比较，而是要用其自带的 equals() 函数来比较，因为双等号判断的是不是同一个对象，而不是他们所表示的值是否相同。同样需要注意的是，Stack 的 peek() 函数取出的也是对象，不能直接和另一个 Stack 的 peek() 取出的对象直接双等，而是使用 equals 或者先将其中一个强行转换成 primitive 类，再和另一个强行比较。
下面这种方法的写法比较简洁，虽然不是明显的 KMP 算法，但是也有其的影子在里面。这种 Java 写法也是在找相同的前缀后缀，但是并没有每次把前缀后缀取出来比较，而是用两个指针分别指向对应的位置比较，然后用 end 指向相同后缀的起始位置，最后再根据 end 的值来拼接两个字符串。有意思的是这种方法对应的 C++ 写法会 TLE，跟上面正好相反，那么我们是否能得出 Java 的 substring 操作略慢，而 C++ 的 reverse 略慢呢，博主也仅仅是猜测而已。
 
Java 解法三：

public class Solution {
    public String shortestPalindrome(String s) {
        int i = 0, end = s.length() - 1, j = end;
        char arr = s.toCharArray();
        while (i < j) {
            if (arr[i] == arr[j]) {
                ++i; --j;
            } else {
                i = 0; --end; j = end;
            }
        }
        return new StringBuilder(s.substring(end + 1)).reverse().toString() + s;
    }
}

 
Github 同步地址：
https://github.com/grandyang/leetcode/issues/214
 
类似题目：
Longest Palindromic Subsequence
Implement strStr()
Palindrome Pairs 