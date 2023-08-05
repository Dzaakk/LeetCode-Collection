# Write your MySQL query statement below
SELECT emp.name as Employee FROM Employee emp, Employee mgr 
WHERE emp.managerId=mgr.id AND emp.salary>mgr.salary