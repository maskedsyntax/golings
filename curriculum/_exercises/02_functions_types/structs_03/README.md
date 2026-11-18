# structs_03: Modifying a copy

Because structs are values, `for _, s := range students` gives you a **copy** of each
student. Changing that copy does nothing to the slice.

## Your task

`Curve(students []Student, bonus int)` should add `bonus` to every student's `Score` **in
place**, capping each score at 100. Right now it has no effect. Fix it.

Reference: https://gobyexample.com/structs
