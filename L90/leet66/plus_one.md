LeetCode 66.加1
            
            
                
                    
                    原创                                                                                                                                            *晴儿*
                                        发布于2019-01-29 10:41:04                    
                    阅读数 53
                    
                        
                        收藏
                    
                
                
                
                
                                        展开
                
            
        
    
    
        
                
                
                                    
         
            
                                        
                
                                            给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一。

最高位数字存放在数组的首位， 数组中每个元素只存储一个数字。

你可以假设除了整数 0 之外，这个整数不会以零开头。

示例 1:

输入: [1,2,3]
输出: [1,2,4]
解释: 输入数组表示数字 123。


示例 2:

输入: [4,3,2,1]
输出: [4,3,2,2]
解释: 输入数组表示数字 4321。

 

class Solution:    def plusOne(self, digits):        """        :type digits: List[int]        :rtype: List[int]        """        i=len(digits)-1        carry=1        while i>=0:            digits[i]=digits[i]+carry            if(digits[i]>9):                digits[i]=digits[i]-10                carry=1            else:                carry=0                break            i=i-1        if carry==1:            digits.insert(0,1)        return digits            
————————————————
版权声明：本文为CSDN博主「*晴儿*」的原创文章，遵循 CC 4.0 BY-SA 版权协议，转载请附上原文出处链接及本声明。
原文链接：https://blog.csdn.net/qq_34919415/article/details/86686567