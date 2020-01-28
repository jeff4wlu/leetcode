[LeetCode] 155. Min Stack 最小栈 

 
Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.
push(x) -- Push element x onto stack.
pop() -- Removes the element on top of the stack.
top() -- Get the top element.
getMin() -- Retrieve the minimum element in the stack.
 
Example:
MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.getMin();   --> Returns -3.
minStack.pop();
minStack.top();      --> Returns 0.
minStack.getMin();   --> Returns -2.
 
这道最小栈跟原来的栈相比就是多了一个功能，可以返回该栈的最小值。使用两个栈来实现，一个栈来按顺序存储 push 进来的数据，另一个用来存出现过的最小值。代码如下:
 
C++ 解法一： 

class MinStack {
public:
    MinStack() {}    
    void push(int x) {
        s1.push(x);
        if (s2.empty() || x <= s2.top()) s2.push(x);
    }    
    void pop() {
        if (s1.top() == s2.top()) s2.pop();
        s1.pop();
    }  
    int top() {
        return s1.top();
    }    
    int getMin() {
        return s2.top();
    }
    
private:
    stack<int> s1, s2;
};

 
Java 解法一：

public class MinStack {
    private Stack<Integer> s1 = new Stack<>();
    private Stack<Integer> s2 = new Stack<>();
    
    public MinStack() {}  
    public void push(int x) {
        s1.push(x);
        if (s2.isEmpty() || s2.peek() >= x) s2.push(x);
    }
    public void pop() {
        int x = s1.pop();
        if (s2.peek() == x) s2.pop();
    }   
    public int top() {
        return s1.peek();
    }  
    public int getMin() {
        return s2.peek();
    }
}

 
需要注意的是上面的 Java 解法中的 pop() 中，为什么不能用注释掉那两行的写法，博主之前也不太明白为啥不能对两个 stack 同时调用 peek() 函数来比较，如果是这种写法，那么不管 s1 和 s2 对栈顶元素是否相等，永远返回 false。这是为什么呢，这就要看 Java 对于peek的定义了，对于 peek() 函数的返回值并不是 int 类型，而是一个 Object 类型，这是一个基本的对象类型，如果直接用双等号 == 来比较的话，肯定不会返回 true，因为是两个不同的对象，所以一定要先将一个转为 int 型，然后再和另一个进行比较，这样才能得到想要的答案，这也是 Java 和 C++ 的一个重要的不同点吧。
那么下面再来看另一种解法，这种解法只用到了一个栈，还需要一个整型变量 min_val 来记录当前最小值，初始化为整型最大值，然后如果需要进栈的数字小于等于当前最小值 min_val，则将 min_val 压入栈，并且将 min_val 更新为当前数字。在出栈操作时，先将栈顶元素移出栈，再判断该元素是否和 min_val 相等，相等的话将 min_val 更新为新栈顶元素，再将新栈顶元素移出栈即可，参见代码如下：
 
C++ 解法二： 

class MinStack {
public:
    MinStack() {
        min_val = INT_MAX;
    }  
    void push(int x) {
        if (x <= min_val) {
            st.push(min_val);
            min_val = x;
        }
        st.push(x);
    }   
    void pop() {
        int t = st.top(); st.pop();
        if (t == min_val) {
            min_val = st.top(); st.pop();
        }
    }  
    int top() {
        return st.top();
    }    
    int getMin() {
        return min_val;
    }

private:
    int min_val;
    stack<int> st;
};

 
Java 解法二：

public class MinStack {
    private int min_val = Integer.MAX_VALUE;
    private Stack<Integer> s = new Stack<>();
    
    public MinStack() {}  
    public void push(int x) {
        if (x <= min_val) {
            s.push(min_val);
            min_val = x;
        }
        s.push(x);
    }    
    public void pop() {
        if (s.pop() == min_val) min_val = s.pop();
    }   
    public int top() {
        return s.peek();
    }    
    public int getMin() {
        return min_val;
    }
}

 
Github 同步地址：
https://github.com/grandyang/leetcode/issues/155
 
类似题目：
Sliding Window Maximum
Max Stack