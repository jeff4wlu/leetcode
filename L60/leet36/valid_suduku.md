[LeetCode] 36. Valid Sudoku 验证数独 

 
Determine if a 9x9 Sudoku board is valid. Only the filled cells need to be validated according to the following rules:
Each row must contain the digits 1-9without repetition.
Each column must contain the digits 1-9 without repetition.
Each of the 9 3x3 sub-boxes of the grid must contain the digits 1-9 without repetition.

A partially filled sudoku which is valid.
The Sudoku board could be partially filled, where empty cells are filled with the character '.'.
Example 1:
Input:
[
  ["5","3",".",".","7",".",".",".","."],
  ["6",".",".","1","9","5",".",".","."],
  [".","9","8",".",".",".",".","6","."],
  ["8",".",".",".","6",".",".",".","3"],
  ["4",".",".","8",".","3",".",".","1"],
  ["7",".",".",".","2",".",".",".","6"],
  [".","6",".",".",".",".","2","8","."],
  [".",".",".","4","1","9",".",".","5"],
  [".",".",".",".","8",".",".","7","9"]
]
Output: true
Example 2:
Input:
[
  ["8","3",".",".","7",".",".",".","."],
  ["6",".",".","1","9","5",".",".","."],
  [".","9","8",".",".",".",".","6","."],
  ["8",".",".",".","6",".",".",".","3"],
  ["4",".",".","8",".","3",".",".","1"],
  ["7",".",".",".","2",".",".",".","6"],
  [".","6",".",".",".",".","2","8","."],
  [".",".",".","4","1","9",".",".","5"],
  [".",".",".",".","8",".",".","7","9"]
]
Output: false
Explanation: Same as Example 1, except with the 5 in the top left corner being 
    modified to 8. Since there are two 8's in the top left 3x3 sub-box, it is invalid.
Note:
A Sudoku board (partially filled) could be valid but is not necessarily solvable.
Only the filled cells need to be validated according to the mentioned rules.
The given board contain only digits 1-9and the character '.'.
The given board size is always 9x9.
 
这道题让验证一个方阵是否为数独矩阵，想必数独游戏我们都玩过，就是给一个 9x9 大小的矩阵，可以分为9个 3x3 大小的矩阵，要求是每个小矩阵内必须都是1到9的数字不能有重复，同时大矩阵的横纵列也不能有重复数字，是一种很经典的益智类游戏，经常在飞机上看见有人拿着纸笔在填数，感觉是消磨时光的利器。这道题给了一个残缺的二维数组，让我们判断当前的这个数独数组是否合法，即要满足上述的条件。判断标准是看各行各列是否有重复数字，以及每个小的 3x3 的小方阵里面是否有重复数字，如果都无重复，则当前矩阵是数独矩阵，但不代表待数独矩阵有解，只是单纯的判断当前未填完的矩阵是否是数独矩阵。那么根据数独矩阵的定义，在遍历每个数字的时候，就看看包含当前位置的行和列以及 3x3 小方阵中是否已经出现该数字，这里需要三个 boolean 型矩阵，大小跟原数组相同，分别记录各行，各列，各小方阵是否出现某个数字，其中行和列标志下标很好对应，就是小方阵的下标需要稍稍转换一下，具体代码如下：
 
解法一：

class Solution {
public:
    bool isValidSudoku(vector<vector<char>>& board) {
        vector<vector<bool>> rowFlag(9, vector<bool>(9));
        vector<vector<bool>> colFlag(9, vector<bool>(9));
        vector<vector<bool>> cellFlag(9, vector<bool>(9));
        for (int i = 0; i < 9; ++i) {
            for (int j = 0; j < 9; ++j) {
                if (board[i][j] == '.') continue;
                int c = board[i][j] - '1';
                if (rowFlag[i][c] || colFlag[c][j] || cellFlag[3 * (i / 3) + j / 3][c]) return false;
                rowFlag[i][c] = true;
                colFlag[c][j] = true;
                cellFlag[3 * (i / 3) + j / 3][c] = true;
            }
        }
        return true;
    }
};

 
我们也可以对空间进行些优化，只使用一个 HashSet 来记录已经存在过的状态，将每个状态编码成为一个字符串，能将如此大量信息的状态编码成一个单一的字符串还是需要有些技巧的。对于每个1到9内的数字来说，其在每行每列和每个小区间内都是唯一的，将数字放在一个括号中，每行上的数字就将行号放在括号左边，每列上的数字就将列数放在括号右边，每个小区间内的数字就将在小区间内的行列数分别放在括号的左右两边，这样每个数字的状态都是独一无二的存在，就可以在 HashSet 中愉快地查找是否有重复存在啦，参见代码如下：
 
解法二：

class Solution {
public:
    bool isValidSudoku(vector<vector<char>>& board) {
        unordered_set<string> st;
        for (int i = 0; i < 9; ++i) {
            for (int j = 0; j < 9; ++j) {
                if (board[i][j] == '.') continue;
                string t = "(" + to_string(board[i][j]) + ")";
                string row = to_string(i) + t, col = t + to_string(j), cell = to_string(i / 3) + t + to_string(j / 3);
                if (st.count(row) || st.count(col) || st.count(cell)) return false;
                st.insert(row);
                st.insert(col);
                st.insert(cell);
            }
        }
        return true;
    }
};
