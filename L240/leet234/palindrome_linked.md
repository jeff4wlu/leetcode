[LeetCode] 234. Palindrome Linked List 回文链表 

 
Given a singly linked list, determine if it is a palindrome.
Example 1:
Input: 1->2
Output: false
Example 2:
Input: 1->2->2->1
Output: true
Follow up:
Could you do it in O(n) time and O(1) space?
 
这道题让我们判断一个链表是否为回文链表，LeetCode 中关于回文串的题共有六道，除了这道，其他的五道为 Palindrome Number，Validate Palindrome，Palindrome Partitioning，Palindrome Partitioning II 和 Longest Palindromic Substring。链表比字符串难的地方就在于不能通过坐标来直接访问，而只能从头开始遍历到某个位置。那么根据回文串的特点，我们需要比较对应位置的值是否相等，一个非常直接的思路就是先按顺序把所有的结点值都存入到一个栈 stack 里，然后利用栈的后入先出的特性，就可以按顺序从末尾取出结点值了，此时再从头遍历一遍链表，就可以比较回文的对应位置了，若不同直接返回 false 即可，参见代码如下：
 
解法一：

class Solution {
public:
    bool isPalindrome(ListNode* head) {
        ListNode *cur = head;
        stack<int> st;
        while (cur) {
            st.push(cur->val);
            cur = cur->next;
        }
        while (head) {
            int t = st.top(); st.pop();
            if (head->val != t) return false;
            head = head->next;
        }
        return true;
    }
};

 
我们也可以用迭代的形式来实现，此时需要使用一个全局变量结点 cur，先初始化为头结点，可以有两种写法，一种写在函数外面的全局变量，或者是在递归函数的参数中加上引用，也表示使用的是全局变量。然后对头结点调用递归函数，在递归函数中，首先判空，若为空则直接返回 true，否则就对下一个结点调用递归函数，若递归函数返回 true 且同时再当前结点值跟 cur 的结点值相同的话，就表明是回文串，否则就不是，注意每次 cur 需要指向下一个结点，参见代码如下：
 
解法二：

class Solution {
public:
    bool isPalindrome(ListNode* head) {
        ListNode *cur = head;
        return helper(head, cur);
    }
    bool helper(ListNode* node, ListNode*& cur) {
        if (!node) return true;
        bool res = helper(node->next, cur) && (cur->val == node->val);
        cur = cur->next;
        return res;
    }
};

 
其实上面的两种解法重复比较一些结点，因为只要前半个链表和后半个链表对应值相等，就是一个回文链表，而并不需要再比较一遍后半个链表，所以我们可以找到链表的中点，这个可以用快慢指针来实现，使用方法可以参见之前的两篇 Convert Sorted List to Binary Search Tree 和 Reorder List，使用快慢指针找中点的原理是 fast 和 slow 两个指针，每次快指针走两步，慢指针走一步，等快指针走完时，慢指针的位置就是中点。我们还需要用栈，每次慢指针走一步，都把值存入栈中，等到达中点时，链表的前半段都存入栈中了，由于栈的后进先出的性质，就可以和后半段链表按照回文对应的顺序比较了，参见代码如下：
 
解法三：

class Solution {
public:
    bool isPalindrome(ListNode* head) {
        if (!head || !head->next) return true;
        ListNode *slow = head, *fast = head;
        stack<int> st{{head->val}};
        while (fast->next && fast->next->next) {
            slow = slow->next;
            fast = fast->next->next;
            st.push(slow->val);
        }
        if (!fast->next) st.pop();
        while (slow->next) {
            slow = slow->next;
            int tmp = st.top(); st.pop();
            if (tmp != slow->val) return false;
        }
        return true;
    }
};

 
这道题的 Follow up 让我们用 O(1) 的空间，那就是说我们不能使用 stack 了，那么如何代替 stack 的作用呢，用 stack 的目的是为了利用其后进先出的特点，好倒着取出前半段的元素。那么现在不用 stack 了，如何倒着取元素呢。我们可以在找到中点后，将后半段的链表翻转一下，这样我们就可以按照回文的顺序比较了，参见代码如下：
 
解法四：

class Solution {
public:
    bool isPalindrome(ListNode* head) {
        if (!head || !head->next) return true;
        ListNode *slow = head, *fast = head;
        while (fast->next && fast->next->next) {
            slow = slow->next;
            fast = fast->next->next;
        }
        ListNode *last = slow->next, *pre = head;
        while (last->next) {
            ListNode *tmp = last->next;
            last->next = tmp->next;
            tmp->next = slow->next;
            slow->next = tmp;
        }
        while (slow->next) {
            slow = slow->next;
            if (pre->val != slow->val) return false;
            pre = pre->next;
        }
        return true;
    }
};

 
Githbu 同步地址：
https://github.com/grandyang/leetcode/issues/234
 
类似题目：
Palindrome Number
Validate Palindrome
Palindrome Partitioning
Palindrome Partitioning II
Longest Palindromic Substring
Reverse Linked List