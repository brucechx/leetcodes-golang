package _180_连续出现的数字

/*
	select distinct l1.Num as ConsecutiveNums
from
Logs l1 join Logs l2 on (l1.Id + 1) = l2.Id and l1.Num = l2.Num
join Logs l3 on (l2.Id + 1) = l3.Id and l2.Num = l3.Num

 */
func main() {

}
