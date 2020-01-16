
版权声明：本文为博主原创文章，未经博主允许不得转载。有事联系：coordinate@live.com https://blog.csdn.net/qq_17550379/article/details/83444673 
给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。
设计一个算法来计算你所能获取的最大利润。你可以尽可能地完成更多的交易（多次买卖一支股票）。
**注意：**你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
示例 1:
输入: [7,1,5,3,6,4]
输出: 7
解释: 在第 2 天（股票价格 = 1）的时候买入，在第 3 天（股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5-1 = 4 。随后，在第 4 天（股票价格 = 3）的时候买入，在第 5 天（股票价格 = 6）的时候卖出, 这笔交易所能获得利润 = 6-3 = 3 。
示例 2:
输入: [1,2,3,4,5]
输出: 4
解释: 在第 1 天（股票价格 = 1）的时候买入，在第 5 天 （股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5-1 = 4 。注意你不能在第 1 天和第 2 天接连购买股票，之后再将它们卖出。因为这样属于同时参与了多笔交易，你必须在再次购买前出售掉之前的股票。
示例 3:
输入: [7,6,4,3,1]
输出: 0
解释: 在这种情况下, 没有交易完成, 所以最大利润为 0。
解题思路
这是之前Leetcode 121：买卖股票的最佳时机（最详细的解法！！！）问题的扩展。这次我们学聪明了，我们首先试试贪心算法可不可以解决这个问题。这次我们不用再考虑交易次数的问题，那么这个问题就变得非常简单了。首先这不是一个在线问题，也就是我们知道明天的股价是不是比今天高，所以我们只要明天的股价高于今天的股价，那么我们就今天买明天卖即可。这样最终的收益一定是最优的。
class Solution:
    def maxProfit(self, prices):
        """
        :type prices: List[int]
        :rtype: int
        """
        result = 0
        for i in range(len(prices)-1):
            if prices[i] < prices[i+1]:
                result += prices[i+1] - prices[i]

        return result
当然我们也可以像之前一样使用动态规划的方法。但是我们这里由于买卖次数不限，所有我们的状态方程也有所变化
b u y [ i ] = m a x ( b u y [ i − 1 ] , s e l l [ i − 1 ] − p r i c e s [ i ] ) buy[i]=max(buy[i-1],sell[i-1]-prices[i]) buy[i]=max(buy[i−1],sell[i−1]−prices[i])
s e l l [ i ] = m a x ( s e l l [ i − 1 ] , b u y [ i − 1 ] + p r i c e s [ i ] ) sell[i]=max(sell[i-1],buy[i-1]+prices[i]) sell[i]=max(sell[i−1],buy[i−1]+prices[i])
我们这里主要的变化是buy，很容易理解，因为我们之前问题只有一次交易机会，所以buy之前不会出现sell这种状态，也就是相当于0-prices[i]。而现在我们的状态可以从buy->sell，也可以从sell->buy，自然就变成了sell[i-1]-prices[i]。
现在我们考虑边界情况，很显然buy[0]=-prices[0]，sell[0]=0（当天买然后再当天卖）。
class Solution:
    def maxProfit(self, prices):
        """
        :type prices: List[int]
        :rtype: int
        """
        if not prices:
            return 0

        len_prices = len(prices)
        buy, sell = [0]*len_prices, [0]*len_prices
        buy[0] = -prices[0]

        for i in range(1, len_prices):
            buy[i] = max(buy[i-1], sell[i-1] - prices[i])
            sell[i] = max(sell[i-1], buy[i-1] + prices[i])

        return sell[-1]