SELECT MAX(Salary) as SecondHighestSalary
FROM Employee
WHERE Salary < (select Max(Salary) from Employee)