[LeetCode] 245. Shortest Word Distance III 最短单词距离之三 

 
Given a list of words and two words word1 and word2, return the shortest distance between these two words in the list.
word1 and word2 may be the same and they represent two individual words in the list.
Example:
Assume that words = ["practice", "makes", "perfect", "coding", "makes"].
Input: word1 = “makes”, word2 = “coding”
Output: 1
Input: word1 = "makes", word2 = "makes"
Output: 3
Note:
You may assume word1 and word2 are both in the list.
 
这道题还是让我们求最短单词距离，有了之前两道题 Shortest Word Distance II 和 Shortest Word Distance 的基础，就大大降低了题目本身的难度。这里增加了一个条件，就是说两个单词可能会相同，所以在第一题中的解法的基础上做一些修改，博主最先想的解法是基于第一题中的解法二，由于会有相同的单词的情况，那么 p1 和 p2 就会相同，这样结果就会变成0，显然不对，所以要对 word1 和 word2 是否的相等的情况分开处理，如果相等了，由于 p1 和 p2 会相同，所以需要一个变量t来记录上一个位置，这样如果t不为 -1，且和当前的 p1 不同，可以更新结果，如果 word1 和 word2 不等，那么还是按原来的方法做，参见代码如下：
 
解法一：

class Solution {
public:
    int shortestWordDistance(vector<string>& words, string word1, string word2) {
        int p1 = -1, p2 = -1, res = INT_MAX;
        for (int i = 0; i < words.size(); ++i) {
            int t = p1;
            if (words[i] == word1) p1 = i;
            if (words[i] == word2) p2 = i;
            if (p1 != -1 && p2 != -1) {
                if (word1 == word2 && t != -1 && t != p1) {
                    res = min(res, abs(t - p1));
                } else if (p1 != p2) {
                    res = min(res, abs(p1 - p2));
                }
            }
        }
        return res;
    }
};

 
上述代码其实可以优化一下，我们并不需要变量t来记录上一个位置，将 p1 初始化为数组长度，p2 初始化为数组长度的相反数，然后当 word1 和 word2 相等的情况，用 p1 来保存 p2 的结果，p2 赋为当前的位置i，这样就可以更新结果了，如果 word1 和 word2 不相等，则还跟原来的做法一样，这种思路真是挺巧妙的，参见代码如下：
 
解法二：

class Solution {
public:
    int shortestWordDistance(vector<string>& words, string word1, string word2) {
        int p1 = words.size(), p2 = -words.size(), res = INT_MAX;
        for (int i = 0; i < words.size(); ++i) {
            if (words[i] == word1) p1 = word1 == word2 ? p2 : i;
            if (words[i] == word2) p2 = i;
            res = min(res, abs(p1 - p2));
        }
        return res;
    }
};

 
我们再来看一种更进一步优化的方法，只用一个变量 idx，这个 idx 的作用就相当于记录上一次的位置，当前 idx 不等 -1 时，说明当前i和 idx 不同，然后在 word1 和 word2 相同或者 words[i] 和 words[idx] 相同的情况下更新结果，最后别忘了将 idx 赋为i，参见代码如下；
 
解法三：

class Solution {
public:
    int shortestWordDistance(vector<string>& words, string word1, string word2) {
        int idx = -1, res = INT_MAX;
        for (int i = 0; i < words.size(); ++i) {
            if (words[i] == word1 || words[i] == word2) {
                if (idx != -1 && (word1 == word2 || words[i] != words[idx])) {
                    res = min(res, i - idx);
                }
                idx = i;
            }
        }
        return res;
    }
};

 
Github 同步地址：
https://github.com/grandyang/leetcode/issues/245
 
类似题目：
Shortest Word Distance II
Shortest Word Distance