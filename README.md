# go-merkle
go-merkle implements a simple merkle tree in Golang.  It allows to obtain proofs and prove the validity of issued proofs.

## Introduction
Merkle trees are binary trees built from the given data entries.  The input entries i.e. leafs are first hashed.  Then each subsequent row is built by hashing two element pairs from the current row using the smaller of the two as the first element when hashing.  This is repeated until only 1 element remains.

## Design
There are many different approaches to handling the remainder i.e. uneven input list.  This implementation carries the odd remainder to the next level until consumed as shown in the diagram below:

```
     52ddb7875b4b  666eec13ab35  7e4729b49fce  31fcacdabf8a  8f5564afe39d
               \    /                      \    /                 |
                \  /                        \  /                  |
                 \/                          \/                   V
            40a65894031a                fed10b5a5178         8f5564afe39d
                             \    /                               |
                              \  /                                |
                               \/                                 V
                          21fd9dfe5c21                       8f5564afe39d
                                             \    /
                                              \  /
                                               \/
                                          8846d886ed99
```