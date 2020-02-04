[LeetCode] 204. Count Primes 质数的个数 

 
Count the number of prime numbers less than a non-negative number, n.
Example:
Input: 10
Output: 4
Explanation: There are 4 prime numbers less than 10, they are 2, 3, 5, 7.
References: 
How Many Primes Are There?
Sieve of Eratosthenes
Credits:
Special thanks to @mithmatt for adding this problem and creating all test cases.
 
这道题给定一个非负数n，让我们求小于n的质数的个数，题目中给了充足的提示，解题方法就在第二个提示 埃拉托斯特尼筛法 Sieve of Eratosthenes 中，这个算法的过程如下图所示：
 

 
我们从2开始遍历到根号n，先找到第一个质数2，然后将其所有的倍数全部标记出来，然后到下一个质数3，标记其所有倍数，一次类推，直到根号n，此时数组中未被标记的数字就是质数。我们需要一个 n-1 长度的 bool 型数组来记录每个数字是否被标记，长度为 n-1 的原因是题目说是小于n的质数个数，并不包括n。 然后来实现埃拉托斯特尼筛法，难度并不是很大，代码如下所示：
  

class Solution {
public:
    int countPrimes(int n) {
        int res = 0;
        vector<bool> prime(n, true);
        for (int i = 2; i < n; ++i) {
            if (!prime[i]) continue;
            ++res;
            for (int j = 2; i * j < n; ++j) {
                prime[i * j] = false;
            }
        }
        return res;
    }
};
