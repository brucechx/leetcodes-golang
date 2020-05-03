# [题目描述](https://leetcode-cn.com/problems/duplicate-emails/)

```sql
select Email from
(
  select Email, count(Email) as num
  from Person
  group by Email
) as statistic
where num > 1
;

```