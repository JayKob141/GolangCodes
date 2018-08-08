# Job-Shop Scheduling Problem

An implementation of the clasical job-shop scheduling problem using Genetic algorithm.


# Explanation

The job-shop consists on the following:

We have N number of jobs that must be processed on M machines.
Each job is composed of M number of operations, each with a specific processing time and a number of machine in which have to be processed. 

The objective of the problem is to find a sequence of operations for all jobs, such that total processing time of the scheduled operations is the minimal posible.

The operations of the problem have the following conditions:

- The operations have a predefined order of execution.
- Each job is processed on each machine exactly once.
- Only 1 operation per machine is allowed to be processed at any given time.
- Given that the operations of any job has a predefined order, 2 operations of the same job cannot be processed concurrently

The problem has NP-Complete complexity.
