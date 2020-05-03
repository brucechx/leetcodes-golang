# [题目描述](https://leetcode-cn.com/problems/rising-temperature/)

```text
SELECT b.Id
FROM Weather as a,Weather as b
WHERE a.Temperature < b.Temperature and DATEDIFF(a.RecordDate,b.RecordDate) = -1;

```