[LeetCode] Flatten 2D Vector 压平二维向量 

 
Implement an iterator to flatten a 2d vector.
For example,
Given 2d vector =
[
  [1,2],
  [3],
  [4,5,6]
]
 
By calling next repeatedly until hasNext returns false, the order of elements returned by next should be: [1,2,3,4,5,6].
Hint:
How many variables do you need to keep track?
Two variables is all you need. Try with x and y.
Beware of empty rows. It could be the first few rows.
To write correct code, think about the invariant to maintain. What is it?
The invariant is x and y must always point to a valid point in the 2d vector. Should you maintain your invariant ahead of time or right when you need it?
Not sure? Think about how you would implement hasNext(). Which is more complex?
Common logic in two different places should be refactored into a common method.
Follow up:
As an added challenge, try to code it using only iterators in C++ or iterators in Java.
 
这道题让我们压平一个二维向量数组，并且实现一个 iterator 的功能，包括 next 和 hasNext 函数，那么最简单的方法就是将二维数组按顺序先存入到一个一维数组里，然后此时只要维护一个变量i来记录当前遍历到的位置，hasNext 函数看当前坐标是否小于元素总数，next 函数即为取出当前位置元素，坐标后移一位，参见代码如下：                      
 
解法一：

class Vector2D {
public:
    Vector2D(vector<vector<int>>& vec2d) {
        for (auto a : vec2d) {
            v.insert(v.end(), a.begin(), a.end());
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

 
下面我们来看另一种解法，不直接转换为一维数组，而是维护两个变量x和y，我们将x和y初始化为0，对于 hasNext 函数，我们检查当前x是否小于总行数，y是否和当前行的列数相同，如果相同，说明要转到下一行，则x自增1，y初始化为0，若此时x还是小于总行数，说明下一个值可以被取出来，那么在 next 函数就可以直接取出行为x，列为y的数字，并将y自增1，参见代码如下：
 
解法二：

class Vector2D {
public:
    Vector2D(vector<vector<int>>& vec2d): data(vec2d), x(0), y(0) {}

    int next() {
        hasNext();
        return data[x][y++];
    }
    bool hasNext() {
        while (x < data.size() && y == data[x].size()) {
            ++x; 
            y = 0;
        }
        return x < data.size();
    }    
private:
    vector<vector<int>> data;
    int x, y;
};

 
题目中的 Follow up 让我们用 interator 来做，C++中 iterator 不像 Java 中的那么强大，自己本身并没有包含 next 和 hasNext 函数，所以我们得自己来实现，我们将x定义为行的 iterator，再用个 end 指向二维数组的末尾，定义一个整型变量y来指向列位置，实现思路和上一种解法完全相同，只是写法略有不同，参见代码如下：
 
解法三：

class Vector2D {
public:
    Vector2D(vector<vector<int>>& vec2d): x(vec2d.begin()), end(vec2d.end()) {}
    
    int next() {
        hasNext();
        return (*x)[y++];
    }
    bool hasNext() {
        while (x != end && y == (*x).size()) {
            ++x; 
            y = 0;
        }
        return x != end;
    }
private:
    vector<vector<int>>::iterator x, end;
    int y = 0;
};

 
类似题目：
Binary Search Tree Iterator
Zigzag Iterator
Peeking Iterator
Flatten Nested List Iterator