[LeetCode] 59. Spiral Matrix II 螺旋矩阵之二 

 
Given a positive integer n, generate a square matrix filled with elements from 1 to n2 in spiral order.
Example:
Input: 3
Output:
[
 [ 1, 2, 3 ],
 [ 8, 9, 4 ],
 [ 7, 6, 5 ]
]
 
此题跟之前那道 Spiral Matrix 本质上没什么区别，就相当于个类似逆运算的过程，这道题是要按螺旋的顺序来填数，由于给定矩形是个正方形，我们计算环数时用 n / 2 来计算，若n为奇数时，此时最中间的那个点没有被算在环数里，所以最后需要单独赋值，还是下标转换问题是难点，参考之前 Spiral Matrix 的讲解来转换下标吧，参见代码如下：
 
解法一：

class Solution {
public:
    vector<vector<int>> generateMatrix(int n) {
        vector<vector<int>> res(n, vector<int>(n));
        int val = 1, p = n;
        for (int i = 0; i < n / 2; ++i, p -= 2) {
            for (int col = i; col < i + p; ++col)
                res[i][col] = val++;
            for (int row = i + 1; row < i + p; ++row)
                res[row][i + p - 1] = val++;
            for (int col = i + p - 2; col >= i; --col)
                res[i + p - 1][col] = val++;
            for (int row = i + p - 2; row > i; --row)    
                res[row][i] = val++;
        }
        if (n % 2 != 0) res[n / 2][n / 2] = val;
        return res;
    }
};

 
当然我们也可以使用下面这种简化了坐标转换的方法，博主个人还是比较推崇下面这种解法，不容易出错，而且好理解，参见代码如下：
 
解法二：

class Solution {
public:
    vector<vector<int>> generateMatrix(int n) {
        vector<vector<int>> res(n, vector<int>(n));
        int up = 0, down = n - 1, left = 0, right = n - 1, val = 1;
        while (true) {
            for (int j = left; j <= right; ++j) res[up][j] = val++;
            if (++up > down) break;
            for (int i = up; i <= down; ++i) res[i][right] = val++;
            if (--right < left) break;
            for (int j = right; j >= left; --j) res[down][j] = val++;
            if (--down < up) break;
            for (int i = down; i >= up; --i) res[i][left] = val++;
            if (++left > right) break;
        }
        return res;
    }
};
