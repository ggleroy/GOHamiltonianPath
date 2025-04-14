# GOHamiltonianPath
Hamiltonian Path Finder algorithm using Golang!  
A **Foundations of Algorithm Design and Analysis** assignment.  

## About  
This project implements a **Hamiltonian Path** algorithm for undirected graphs.  
A **Hamiltonian Path** visits each vertex exactly once.

### Algorithm breakdown

- **Validation**:  
  The input is checked to ensure it is a square adjacency matrix with only `0`s and `1`s.
  
- **Initialization**:  
  A `path` slice is created to store the traversal. It starts at vertex `0`.

- **Backtracking Search**:  
  At each step, the algorithm tries to add a valid next vertex connected to the last vertex.

- **Validity Check**:  
  A vertex can only be added if:
  - It is adjacent to the last vertex.
  - It has not been visited yet.

- **Result**:  
  If a Hamiltonian path is found, it returns the path, execution duration, number of vertices, and edges.

## Structure  
This repository consists of a single **Go** file: `main.go`.  

## How to Run  
Ensure **Go** is installed on your machine, then execute:  

```sh
go run main.go
```

- **Problem Complexity**:  
  The **Hamiltonian Path Problem** belongs to the class **NP-Complete**.

- **Justification**:
  - It is in **NP**: a given path can be verified in polynomial time.
  - It is **NP-Complete**: finding such a path is as hard as the Traveling Salesman Problem (TSP) without the need to return to the starting vertex.
  - Solving it in polynomial time would solve many other NP problems.

---

### Asymptotic Time Complexity

- **Backtracking Search Complexity**:  
  Each step tries all possible vertices not yet visited.

- **Worst-Case Complexity**:  
  \[
  O(n!)
  \]
  where \( n \) is the number of vertices.

- **Method Used**:  
  Manual operation counting based on the number of possible paths.

---

### Master Theorem Applicability

- **Applicability**:  
  The **Master Theorem** cannot be applied here.

- **Reason**:  
  The algorithm is not a simple divide-and-conquer recurrence of the form:
  \[
  T(n) = aT(n/b) + f(n)
  \]
  It is an exhaustive combinatorial search, not a recursively shrinking problem.

---

### Case Complexity Analysis

| Case         | Description                                                                 | Complexity |
|--------------|-----------------------------------------------------------------------------|------------|
| Best Case    | A Hamiltonian path is found early without backtracking much.                | \( O(n^2) \) |
| Average Case | Several vertices explored before finding a path.                            | Between \( O(n^2) \) and \( O(n!) \) |
| Worst Case   | All permutations explored with no path found or path found at the end.       | \( O(n!) \) |

 -***In sparse or complex graphs, the algorithm may take **exponential time**, making it impractical for large graphs.***
