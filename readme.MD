# My Trees
Write your own trees.

## BST
The simplest, no self balancing.

Create and find with `New()/NewFromArray()`, and `Find()`. `Find()` returns error if a node is not found. Show single node info with `.Info()`, show entire tree with `PrintSubtreeLevelOrder()`
```
myTree := bst.NewFromArray([]int{4, 7, 1, 6, 8, 2, 3, -2})
myTree.PrintSubtreeLevelOrder()

// output
Root: 4 Size: 8 Level: 4
               4                              
       1               7              
  -2       2       6       8      
 _   _   _   3   _   _   _   _  
================================

found, _ := myTree.Find(7)
found.PrintSubtreeLevelOrder()
found.PrintNodeInfo()

// output
Root: 7 Size: 3 Level: 2
 7   
6 8 
====
  4  
  |
  7
 / \
6   8
```

Soft deleting node marks `deleted=true`, but does not change tree structure / delete the node.
```
myTree.SoftDelete(-2)
myTree.PrintSubtreeLevelOrder()

// output
Root: 4 Size: 8 Level: 4
       4               
   1       7       
 X   2   6   8   
_ _ _ 3 _ _ _ _ 
================
```

Insert will not balance the tree, inserting duplicate element returns an error.
```
myTree.Insert(11)
myTree.Insert(9)
_, err := myTree.Insert(9)
fmt.Println(err)
myTree.PrintSubtreeLevelOrder()

// output
failed to insert node, node already in sub-tree
Root: 4 Size: 9 Level: 4
       4               
   1       7       
 X   2   6   8   
_ _ _ 3 _ _ _ 9 
================
```

No self balancing also means the performance degrades to a linked list
```
worseTree := bst.NewFromArray([]int{1, 2, 3, 4, 5})
worseTree.PrintSubtreeLevelOrder()

Root: 1 Size: 5 Level: 5
               1                               
       _               2               
   _       _       _       3       
 _   _   _   _   _   _   _   4   
_ _ _ _ _ _ _ _ _ _ _ _ _ _ _ 5 
================================
```

`HardDelete(x)` will remove a node from the tree, change its structure and update the size metadata.
```
myTree.PrintSubtreeLevelOrder()

myTree, _ = myTree.HardDelete(7)
myTree.PrintSubtreeLevelOrder()

myTree, _ = myTree.HardDelete(4)
myTree.PrintSubtreeLevelOrder()

// output
Root: 4 Size: 9 Level: 4
       4               
   1       7       
 X   2   6   8   
_ _ _ 3 _ _ _ 9 
================
Root: 4 Size: 8 Level: 4
       4               
   1       8       
 X   2   6   9   
_ _ _ 3 _ _ _ _ 
================
Root: 6 Size: 7 Level: 4
       6               
   1       8       
 X   2   _   9   
_ _ _ 3 _ _ _ _ 
================
```