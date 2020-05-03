# [题目描述](https://leetcode-cn.com/problems/employees-earning-more-than-their-managers/)

```sql
SELECT
     a.NAME AS Employee
FROM Employee AS a JOIN Employee AS b
     ON a.ManagerId = b.Id
     AND a.Salary > b.Salary
;
```