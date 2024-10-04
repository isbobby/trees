# My Trees
Write your own trees.

## BST
The simplest, no self balancing.

```
myTree := bst.NewFromArray([]int{4, 7, 1, 6, 8, 2, 3})
myTree.PrintSubtreeLevelOrder()

// output
Root: 4 Size: 7 Level: 4
       4 
      1 7 
    _ 2 6 8 
_ _ _ 3 _ _ _ _ 
================

myTree.SoftDelete(6)
myTree.PrintSubtreeLevelOrder()

// output
Root: 4 Size: 7 Level: 4
       4 
      1 7 
    _ 2 X 8 
_ _ _ 3 _ _ _ _ 
================

myTree.Insert(11)
myTree.Insert(-2)
myTree.Insert(9)
myTree.PrintSubtreeLevelOrder()

// output
Root: 4 Size: 10 Level: 5
               4 
              1 7 
           -2 2 X 8 
       _ _ _ 3 _ _ _ 11 
_ _ _ _ _ _ _ _ _ _ _ _ _ _ 9 _ 
================================
```