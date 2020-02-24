[LeetCode] 281. Zigzag Iterator 之字形迭代器 

 
Given two 1d vectors, implement an iterator to return their elements alternately.
Example:
Input:
v1 = [1,2]
v2 = [3,4,5,6] 

Output: [1,3,2,4,5,6]

Explanation: By calling next repeatedly until hasNext returns false, 
             the order of elements returned by next should be: [1,3,2,4,5,6].
Follow up: What if you are given k 1d vectors? How well can your code be extended to such cases?
Clarification for the follow up question:
The "Zigzag" order is not clearly defined and is ambiguous for k > 2 cases. If "Zigzag" does not look right to you, replace "Zigzag" with "Cyclic". For example:
Input:
[1,2,3]
[4,5,6,7]
[8,9]

Output: [1,4,8,2,5,9,3,6,7].
 
这道题让我们写一个之字形迭代器，跟之前那道 Flatten 2D Vector 有些类似，那道题是横向打印，这道题是纵向打印，虽然方向不同，但是实现思路都是大同小异。博主最先想到的方法是用两个变量i和j分别记录两个向量的当前元素位置，初始化为0，然后当 i<=j 时，则说明需要打印 v1 数组的元素，反之则打印 v2 数组中的元素。在 hasNext 函数中，当i或j打印等于对应数组的长度时，将其赋为一个特大值，这样不影响打印另一个数组的值，只有当i和j都超过格子数组的长度时，返回 false，参见代码如下：
 
解法一：

class ZigzagIterator {
public:
    ZigzagIterator(vector<int>& v1, vector<int>& v2) {
        v.push_back(v1);
        v.push_back(v2);
        i = j = 0;
    }
    int next() {
        return i <= j ? v[0][i++] : v[1][j++];
    }
    bool hasNext() {
        if (i >= v[0].size()) i = INT_MAX;
        if (j >= v[1].size()) j = INT_MAX;
        return i < v[0].size() || j < v[1].size();
    }
private:
    vector<vector<int>> v;
    int i, j;
};

 
下面来看另一种解法，这种解法直接在初始化的时候就两个数组按照之字形的顺序存入另一个一位数组中了，那么就按顺序打印新数组中的值即可，参见代码如下：
 
解法二：

class ZigzagIterator {
public:
    ZigzagIterator(vector<int>& v1, vector<int>& v2) {
        int n1 = v1.size(), n2 = v2.size(), n = max(n1, n2);
        for (int i = 0; i < n; ++i) {
            if (i < n1) v.push_back(v1[i]);
            if (i < n2) v.push_back(v2[i]);
        }
    }
    int next() {
        return v[i++];
    }
    bool hasNext() {
        return i < v.size();
    }
private:
    vector<int> v;
    int i = 0;
};

 
由于题目中的 Follow up 让将输入换成k个数组的情况，那么上面的解法一就不适用了，解法二的空间复杂度比较高，所以需要一种更高效的方法。这里可以采用 queue 加 iterator 的方法，用一个 queue 里面保存 iterator 的 pair，在初始化的时候，有几个数组就生成几个 pair 放到 queue 中，每个 pair 保存该数组的首位置和尾位置的 iterator，在 next() 函数中，取出 queue 队首的一个 pair，如果当前的 iterator 不等于 end()，将其下一个位置的 iterator 和 end 存入队尾，然后返回当前位置的值。在 hasNext() 函数中，只需要看 queue 是否为空即可，参见代码如下：
 
解法三：

class ZigzagIterator {
public:
    ZigzagIterator(vector<int>& v1, vector<int>& v2) {
        if (!v1.empty()) q.push(make_pair(v1.begin(), v1.end()));
        if (!v2.empty()) q.push(make_pair(v2.begin(), v2.end()));
    }
    int next() {
        auto it = q.front().first, end = q.front().second;
        q.pop();
        if (it + 1 != end) q.push(make_pair(it + 1, end));
        return *it;
    }
    bool hasNext() {
        return !q.empty();
    }
private:
    queue<pair<vector<int>::iterator, vector<int>::iterator>> q;
};

 
Github 同步地址：
https://github.com/grandyang/leetcode/issues/281
 
类似题目：
Flatten 2D Vector