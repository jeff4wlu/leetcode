[LeetCode] 48. Rotate Image 旋转图像 

 
You are given an n x n 2D matrix representing an image.
Rotate the image by 90 degrees (clockwise).
Note:
You have to rotate the image in-place, which means you have to modify the input 2D matrix directly. DO NOT allocate another 2D matrix and do the rotation.
Example 1:
Given input matrix = 
[
  [1,2,3],
  [4,5,6],
  [7,8,9]
],

rotate the input matrix in-place such that it becomes:
[
  [7,4,1],
  [8,5,2],
  [9,6,3]
]
Example 2:
Given input matrix =
[
  [ 5, 1, 9,11],
  [ 2, 4, 8,10],
  [13, 3, 6, 7],
  [15,14,12,16]
], 

rotate the input matrix in-place such that it becomes:
[
  [15,13, 2, 5],
  [14, 3, 4, 1],
  [12, 6, 8, 9],
  [16, 7,10,11]
]
 
在计算机图像处理里，旋转图片是很常见的，由于图片的本质是二维数组，所以也就变成了对数组的操作处理，翻转的本质就是某个位置上数移动到另一个位置上，比如用一个简单的例子来分析：
 
1  2  3　　　 　　 7  4  1　

4  5  6　　-->　　 8  5  2　　

7  8  9 　　　 　　9  6  3
 
对于90度的翻转有很多方法，一步或多步都可以解，先来看一种直接的方法，这种方法是按顺时针的顺序去覆盖前面的数字，从四个顶角开始，然后往中间去遍历，每次覆盖的坐标都是同理，如下：
(i, j)  <-  (n-1-j, i)  <-  (n-1-i, n-1-j)  <-  (j, n-1-i)
这其实是个循环的过程，第一个位置又覆盖了第四个位置，这里i的取值范围是 [0, n/2)，j的取值范围是 [i, n-1-i)，至于为什么i和j是这个取值范围，为啥i不用遍历 [n/2, n)，若仔细观察这些位置之间的联系，不难发现，实际上j列的范围 [i, n-1-i) 顺时针翻转 90 度，正好就是i行的 [n/2, n) 的位置，这个方法每次循环换四个数字，如下所示：
 
1  2  3               7  2  1            7  4  1

4  5  6      -->      4  5  6　　 -->  　 8  5  2　　

7  8  9               9  8  3　　　　   　 9  6  3
 
解法一：

class Solution {
public:
    void rotate(vector<vector<int>>& matrix) {
        int n = matrix.size();
        for (int i = 0; i < n / 2; ++i) {
            for (int j = i; j < n - 1 - i; ++j) {
                int tmp = matrix[i][j];
                matrix[i][j] = matrix[n - 1 - j][i];
                matrix[n - 1 - j][i] = matrix[n - 1 - i][n - 1 - j];
                matrix[n - 1 - i][n - 1 - j] = matrix[j][n - 1 - i];
                matrix[j][n - 1 - i] = tmp;
            }
        }
    }
};

 
还有一种解法，首先以从对角线为轴翻转，然后再以x轴中线上下翻转即可得到结果，如下图所示(其中蓝色数字表示翻转轴)：
 
1  2  3　　　 　　 9  6  3　　　　　     7  4  1

4  5  6　　-->　　 8  5  2　　 -->   　 8  5  2　　

7  8  9 　　　 　　7  4  1　　　　　     9  6  3
 
解法二：

class Solution {
public:
    void rotate(vector<vector<int>>& matrix) {
        int n = matrix.size();
        for (int i = 0; i < n - 1; ++i) {
            for (int j = 0; j < n - i; ++j) {
                swap(matrix[i][j], matrix[n - 1- j][n - 1 - i]);
            }
        }
        reverse(matrix.begin(), matrix.end());
    }
};

 
最后再来看一种方法，这种方法首先对原数组取其转置矩阵，然后把每行的数字翻转可得到结果，如下所示（其中蓝色数字表示翻转轴，Github 上可能无法显示颜色，请参见博客园上的帖子）：
 
1  2  3　　　 　　 1  4  7　　　　　     7  4  1

4  5  6　　-->　　 2  5  8　　 -->  　  8  5  2　　

7  8  9 　　　 　　3  6  9　　　　       9  6  3
 
解法三：

class Solution {
public:
    void rotate(vector<vector<int>>& matrix) {
        int n = matrix.size();
        for (int i = 0; i < n; ++i) {
            for (int j = i + 1; j < n; ++j) {
                swap(matrix[i][j], matrix[j][i]);
            }
            reverse(matrix[i].begin(), matrix[i].end());
        }
    }
};