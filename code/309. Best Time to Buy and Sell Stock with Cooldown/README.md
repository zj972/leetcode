# 309. Best Time to Buy and Sell Stock with Cooldown
Say you have an array for which the `i<sup>th</sup>` element is the price of a given stock on day `i`.

Design an algorithm to find the maximum profit. You may complete as many transactions as you like (ie, buy one and sell one share of the stock multiple times) with the following restrictions:

* You may not engage in multiple transactions at the same time (ie, you must sell the stock before you buy again).

* After you sell your stock, you cannot buy stock on next day. (ie, cooldown 1 day)

##### Example:
```
prices = [1, 2, 3, 0, 2]
maxProfit = 3
transactions = [buy, sell, cooldown, buy, sell]
```
![image](https://raw.githubusercontent.com/zj972/leetcode/master/code/309.%20Best%20Time%20to%20Buy%20and%20Sell%20Stock%20with%20Cooldown/%E7%8A%B6%E6%80%81%E8%BD%AC%E7%A7%BB%E5%9B%BE.png)
