# 排序

## 排序的几个基本概念
+ 内排序：在对待排序数据存放在内存中进行的排序过程。
+ 外排序：待排数据量太大，无法一次性将所有待排序数据放入内存中，在排序过程中需要对磁盘等外部储存器进行访问。
+ 比较排序：排序过程中需要对数据关键字进行比较。
+ 非比较排序：排序过程中不需要对数据关键字进行比较。
+ 排序算法的稳定性：在每一次单趟排序后，相同关键字的相对顺序不变。

分析排序通常从以下几个点着手：
+ 算法基本原理及实现
+ 时间复杂度
+ 稳定性

## 排序的分类
![](https://github.com/blockchain-develop/gopro/blob/master/img/sort/kind%20of%20sort.png)

## 排序的分析
![](https://github.com/blockchain-develop/gopro/blob/master/img/sort/analysis%20of%20sort.png)

## 排序参考
[十大排序算法介绍](https://www.runoob.com/w3cnote/ten-sorting-algorithm.html)

## 比较排序

包括快速排序、插入排序、简单选择排序、冒泡排序、归并排序、堆排序。

比较排序可以通过决策树分析算法的性能，比较排序算法的时间复杂度下限为O(nlgn)。

+ [MIT算法导论-比较排序性能分析](https://www.bilibili.com/video/BV1Tb411M7FA?p=5)

### 快速排序
1. 快速排序的分治
2. 快速排序算法
3. 快速排序的尾递归优化
4. 快速排序算法时间复杂度分析(最坏情况、最优情况)
5. 随机化随机排序

+ [MIT算法导论-快速排序](https://www.bilibili.com/video/BV1Tb411M7FA?p=4)
+ [指示器随机变量](https://www.cnblogs.com/yuanquanxi/p/10241303.html)
+ [go语言实现的快速排序](./memsort/comsort/swapsort/quicksort)

### 插入排序

+ [go语言实现的插入排序](./memsort/comsort/insertsort/insertsort)

### 简单选择排序

+ [go语言实现的选择排序](./memsort/comsort/selectsort/simpleselectsort)

### 冒泡排序

+ [go语言实现的冒泡排序](./memsort/comsort/swapsort/bubblingsort)

### 归并排序

+ [go语言的归并排序](./memsort/comsort/insertsort/mergesort)

### 堆排序

+ [go语言的堆排序](./memsort/comsort/selectsort/heapsort)

### 随机化二叉搜索树排序

+ [MIT算法导论-二叉搜索树](https://www.bilibili.com/video/BV1Tb411M7FA?p=9)

## 非比较排序

包括计数排序、基数排序、桶排序

### 计数排序

可以在线性时间内处理小范围数据。

+ [go语言的计数排序](./memsort/uncomsort/countingsort)
+ [MIT算法导论-计数排序](https://www.bilibili.com/video/BV1Tb411M7FA?p=5)

### 基数排序

可以在线性时间内处理大范围数据。

+ [go语言的基数排序](./memsort/uncomsort/radixsort)
+ [MIT算法导论-基数排序](https://www.bilibili.com/video/BV1Tb411M7FA?p=5)

### 桶排序

+ [go语言的桶排序](./memsort/uncomsort/bucketsort)