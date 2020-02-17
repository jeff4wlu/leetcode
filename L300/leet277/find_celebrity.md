[LeetCode] 277. Find the Celebrity 寻找名人 

 
Suppose you are at a party with n people (labeled from 0 to n - 1) and among them, there may exist one celebrity. The definition of a celebrity is that all the other n - 1 people know him/her but he/she does not know any of them.
Now you want to find out who the celebrity is or verify that there is not one. The only thing you are allowed to do is to ask questions like: "Hi, A. Do you know B?" to get information of whether A knows B. You need to find out the celebrity (or verify there is not one) by asking as few questions as possible (in the asymptotic sense).
You are given a helper function bool knows(a, b)which tells you whether A knows B. Implement a function int findCelebrity(n). There will be exactly one celebrity if he/she is in the party. Return the celebrity's label if there is a celebrity in the party. If there is no celebrity, return -1.
 
Example 1:

Input: graph = [
  [1,1,0],
  [0,1,0],
  [1,1,1]
]
Output: 1
Explanation: There are three persons labeled with 0, 1 and 2. graph[i][j] = 1 means person i knows person j, otherwise graph[i][j] = 0 means person i does not know person j. The celebrity is the person labeled as 1 because both 0 and 2 know him but 1 does not know anybody.
Example 2:

Input: graph = [
  [1,0,1],
  [1,1,0],
  [0,1,1]
]
Output: -1
Explanation: There is no celebrity.
 
Note:
The directed graph is represented as an adjacency matrix, which is an n x n matrix where a[i][j] = 1 means person i knows person j while a[i][j] = 0 means the contrary.
Remember that you won't have direct access to the adjacency matrix.
 
这道题让我们在一群人中寻找名人，所谓名人就是每个人都认识他，他却不认识任何人，限定了只有1个或0个名人，给定了一个 API 函数，输入a和b，用来判断a是否认识b，让我们尽可能少的调用这个函数，来找出人群中的名人。我最先想的方法是建立个一维数组用来标记每个人的名人候选状态，开始均初始化为 true，表示每个人都是名人候选人，然后我们一个人一个人的验证其是否为名人，对于候选者i，我们遍历所有其他人j，如果i认识j，或者j不认识i，说明i不可能是名人，那么我们标记其为 false，然后验证下一个候选者，反之如果i不认识j，或者j认识i，说明j不可能是名人，标记之。对于每个候选者i，如果遍历了一圈而其候选者状态仍为 true，说明i就是名人，返回即可，如果遍历完所有人没有找到名人，返回 -1，参见代码如下：
 
解法一：

bool knows(int a, int b);
class Solution {
public:
    int findCelebrity(int n) {
        vector<bool> candidate(n, true);
        for (int i = 0; i < n; ++i) {
            for (int j = 0; j < n; ++j) {
                if (candidate[i] && i != j) {
                    if (knows(i, j) || !knows(j, i)) {
                        candidate[i] = false;
                        break;
                    } else {
                        candidate[j] = false;
                    }
                }
            }
            if (candidate[i]) return i;
        }
        return -1;
    }
};

 
我们其实可以不用一维数组来标记每个人的状态，我们对于不是名人的i，直接 break，继续检查下一个，但是由于我们没有标记后面的候选人的状态，所以有可能会重复调用一些 knows 函数，所以下面这种方法虽然省了空间，但是调用 knows 函数的次数可能会比上面的方法次数要多，参见代码如下：
 
解法二：

bool knows(int a, int b);
class Solution {
public:
    int findCelebrity(int n) {
        for (int i = 0, j = 0; i < n; ++i) {
            for (j = 0; j < n; ++j) {
                if (i != j && (knows(i, j) || !knows(j, i))) break;
            }
            if (j == n) return i;
        }
        return -1;
    }
};

 
下面这种方法是网上比较流行的一种方法，设定候选人 res 为0，原理是先遍历一遍，对于遍历到的人i，若候选人 res 认识i，则将候选人 res 设为i，完成一遍遍历后，我们来检测候选人 res 是否真正是名人，我们如果判断不是名人，则返回 -1，如果并没有冲突，返回 res，参见代码如下：
 
解法三：

bool knows(int a, int b);
class Solution {
public:
    int findCelebrity(int n) {
        int res = 0;
        for (int i = 0; i < n; ++i) {
            if (knows(res, i)) res = i;
        }
        for (int i = 0; i < n; ++i) {
            if (res != i && (knows(res, i) || !knows(i, res))) return -1;
        }
        return res;
    }
};

 
由热心网友 fgvlty 提醒，我们还可以进一步减少 API 的调用量，找候选者的方法跟上面相同，但是在验证的时候，分为两段，先验证候选者前面的所有人，若候选者认识任何人，或者任何人不认识候选者，直接返回 -1。再验证候选者后面的人，这时候只需要验证是否有人不认识候选者就可以了，因为我们在最开始找候选者的时候就已经保证了候选者不会认识后面的任何人，参见代码如下：
 
解法四：

bool knows(int a, int b);
class Solution {
public:
    int findCelebrity(int n) {
        int res = 0;
        for (int i = 0; i < n; ++i) {
            if (knows(res, i)) res = i;
        }
        for (int i = 0; i < res; ++i) {
            if (knows(res, i) || !knows(i, res)) return -1;
        }
        for (int i = res + 1; i < n; ++i) {
            if (!knows(i, res)) return -1;
        }
        return res;
    }
};

 
类似题目：
Find the Town Judge