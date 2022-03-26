## Generics

### Summary
1. [Introduction](#introduction)
2. [Examples](#examples)

### Introduction
- What is `T`, `K` and `V` ?
  - Like `i` and `j` used in _for loops_, those are just letters, `T` for example was adopted because it was the first letter of the work _Type_.
  - `K` and `V` are commonly used in functions that deals with maps, so `Key` and `Value`.
- What is `any`?
  - `any` basically is a more friendly name for `interface{}`
- What is `comparable`?
  - `comparable` is to only allow types that work with `==` or `!=` operators.

### Examples
1. [Simple](./01_simple_example/)
  - A example based on AWS SKD, that uses pointers of `string` and `int`
