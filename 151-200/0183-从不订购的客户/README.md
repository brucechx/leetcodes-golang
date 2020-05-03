# [题目描述](https://leetcode-cn.com/problems/customers-who-never-order/)

```sql
select customers.name as 'Customers'
from customers
where customers.id not in
(
    select customerid from orders
);
```