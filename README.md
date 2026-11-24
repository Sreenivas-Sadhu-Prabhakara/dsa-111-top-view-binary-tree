# Top View of a Binary Tree

> Striver SDE Sheet 111 · Day 17 Binary Tree — bottom view's one-flag twin: BFS per column, but the FIRST arrival wins instead of the last. Go, stdlib only, tested.

## Problem

Seen from directly above, one node is visible per vertical column — the highest. Columns are horizontal distances (root 0, left −1, right +1).

## Approach

Identical BFS-with-column-tags to the bottom view (#110); the only change is the write policy: **write-if-absent**. BFS's top-to-bottom order guarantees the first node recorded in a column is the shallowest, so later (deeper) nodes in the same column are ignored rather than overwriting. The pair of problems is a nice demonstration that the traversal often isn't the algorithm — the *reduction policy* over the traversal is.

**Complexity:** O(n) time, O(n) space.

## Run

    go test ./...   # verify
    go run .        # demo

---

One of 191 repos solving Striver's SDE Sheet in Go — stdlib only, one problem per repo, every solution tested. Browse the set: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
