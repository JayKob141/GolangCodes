package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
    "strconv"
    "math/rand"
    "math"
    "time"
    "sync"
    "sort"
)

type JobShopSpecification struct {
    numberOfMachines int
    numberOfJobs int
    operationsMatrix [][]int
    timesMatrix [][]int
}

type GeneticAlgorithmParameters struct {
    numberOfGenerations int
    populationSize int
    percentOfCross float64
    percentOfMutation float64
}

func readFile(fullPath string) JobShopSpecification {

    jobShop := JobShopSpecification{}

    // Open the file.
    f, _ := os.Open(fullPath)

    // Create a new Scanner for the file.
    scanner := bufio.NewScanner(f)

    // Rading the first line and parsing # of jobs and machines
    scanner.Scan()
    line := scanner.Text()
    splitted := strings.Split(line, " ")

    fmt.Println("Readed line: ")
    fmt.Println(line)

    numberOfJobs, _ :=  strconv.ParseInt(splitted[0], 10, 0) 
    numberOfMachines, _ :=  strconv.ParseInt(splitted[1], 10, 0)
    jobShop.numberOfJobs = int(numberOfJobs)
    jobShop.numberOfMachines = int(numberOfMachines)

    // Initializing empty matrices 
    jobShop.operationsMatrix = make([][]int, jobShop.numberOfJobs)
    jobShop.timesMatrix = make([][]int, jobShop.numberOfJobs)
    for i := 0 ; i < jobShop.numberOfJobs; i++ {
        jobShop.operationsMatrix[i] = make([]int, numberOfMachines)
        jobShop.timesMatrix[i] = make([]int, numberOfMachines)
    }

    // Assigning the remaining values from the file
    var iOperation int
    for j := 0 ; j < jobShop.numberOfJobs ; j++ {
        scanner.Scan()
        line := scanner.Text()
        splitted := strings.Split(line, " ")

        iOperation = 0
        for o := 0 ; o < len(splitted) ; o = o + 2{
            x,_ := strconv.ParseInt( splitted[o], 10, 0)
            y,_ := strconv.ParseInt( splitted[o + 1], 10, 0)

            iMachine := int(x)
            kTime := int(y)

            jobShop.operationsMatrix[j][iOperation] = iMachine
            jobShop.timesMatrix[j][iMachine] = kTime

            iOperation += 1
        }
    }

    fmt.Printf("Number of jobs: %d \n", jobShop.numberOfJobs)
    fmt.Printf("Number of machines: %d \n", jobShop.numberOfMachines)
    fmt.Println("Operations matrix : ")
    fmt.Println(jobShop.operationsMatrix)
    fmt.Println("Times matrix : ")
    fmt.Println(jobShop.timesMatrix)
    
    return jobShop
}

// ###########################
//
// 		MAIN FUNCTION
//
// ###########################

func main() {
    argsWithoutProg := os.Args[1:]

    if len(argsWithoutProg) == 0 {
        fmt.Println("Expected 1 argument (instance file)")
        return
    }
    
    fmt.Println(argsWithoutProg)

    jobshop := readFile( argsWithoutProg[0] )

    ga := GeneticAlgorithmParameters { numberOfGenerations: 2000, populationSize: 256, percentOfCross: 0.3, percentOfMutation: 0.1 }
     
    geneticAlgorithm(ga, jobshop)
}


func fillWithRandomJobs(sequences []int, sequenceSize, numberOfSequences, numberOfJobs, numberOfMachines, idGoRoutine int, wg *sync.WaitGroup) {
    defer wg.Done()

    counters := make([]int, numberOfJobs)
    r := rand.New( rand.NewSource( time.Now().Unix() * int64(idGoRoutine) * 17) )

    for s := 0 ; s < numberOfSequences ; s++ {

        low := s * sequenceSize
        high := low + sequenceSize
        sequence := sequences[ low:high ]
        for j := 0 ; j < numberOfJobs; j++ {
            counters[j] = 0
        }

        for i := 0 ; i < sequenceSize; i++ {
            randomJob := r.Intn( numberOfJobs )
            for counters[randomJob] >= numberOfMachines {
                randomJob = r.Intn( numberOfJobs )
            }
            sequence[i] = randomJob
            counters[randomJob] += 1
        }
    }
}

func Max(x, y int) int{
    if x > y {
        return x
    }
    return y
}

func Min(x, y int) int{
    if x < y {
        return x
    }
    return y
}

func computeMakespanOfSequence(sequence []int, sequenceSize int, jobshop JobShopSpecification) int {

    ganttChart := make( []int, jobshop.numberOfMachines )
    iOperation := make( []int, jobshop.numberOfJobs )
    previousFinalTimes :=  make( []int, jobshop.numberOfJobs )
    makespan := 0

    for _, job := range sequence {
        o := iOperation[job]
        m := jobshop.operationsMatrix[job][o]
        d := jobshop.timesMatrix[job][m]

        if o == 0 {
            ganttChart[m] += d
        } else {
            ganttChart[m] = Max(previousFinalTimes[job], ganttChart[m]) + d
        }

        iOperation[job] += 1
        previousFinalTimes[job] = ganttChart[m] 
        makespan = Max( ganttChart[m], makespan )
    }

    return makespan
}

func computeSolutions(sequences, makespans []int, numberOfSequences, sequenceSize, idGoRoutine int, jobshop JobShopSpecification, wg *sync.WaitGroup) {
    defer wg.Done()

    for s := 0; s < numberOfSequences; s++{
        low := s * sequenceSize
        high := low + sequenceSize
        makespan := computeMakespanOfSequence( sequences[low:high], sequenceSize, jobshop)
        makespans[ numberOfSequences * idGoRoutine + s] = makespan
    }
}

func mutationOverSequences(sequences, makespans []int, numberOfSequences, sequenceSize, idGoRoutine int, jobshop JobShopSpecification, wg *sync.WaitGroup, rng *rand.Rand) {
    defer wg.Done()
}

func geneticAlgorithm(ga GeneticAlgorithmParameters, jobshop JobShopSpecification){

    numberOfGoRoutines := 4
    
    //  Init population
    sequenceSize := jobshop.numberOfJobs * jobshop.numberOfMachines
    sequences := make( []int, ga.populationSize * sequenceSize )
    makespans := make( []int, ga.populationSize )
    // ganttChart := make( []int, ga.populationSize * jobshop.numberOfMachines )
    // previousFinalTimes :=  make( []int, ga.populationSize * jobshop.numberOfJobs )
    // iOperation :=  make( []int, ga.populationSize * jobshop.numberOfJobs )
    indices := make( []int, ga.populationSize )

    // generate random init population
    var wg sync.WaitGroup
    wg.Add(numberOfGoRoutines)
    N := ga.populationSize / numberOfGoRoutines
    for r := 0 ; r < numberOfGoRoutines ; r++{
        low := r * N * sequenceSize
        high := low + N * sequenceSize
        go fillWithRandomJobs( sequences[low:high], sequenceSize,N, jobshop.numberOfJobs, jobshop.numberOfMachines, r , &wg)
    }
    wg.Wait()

    wg.Add(numberOfGoRoutines)
    N = ga.populationSize / numberOfGoRoutines
    for r := 0 ; r < numberOfGoRoutines ; r++{
        low := r * N * sequenceSize
        high := low + N * sequenceSize
        go computeSolutions( sequences[low:high], makespans, N, sequenceSize, r, jobshop, &wg)
    }
    wg.Wait()


    rngs := make( []*rand.Rand, numberOfGoRoutines)
    for i := 0 ; i < numberOfGoRoutines ; i++ {
        rngs[i] = rand.New( rand.NewSource( time.Now().Unix() * int64(i + 1) * 17) )
    }

    bestSolution := math.MaxInt32

    for k := 0 ; k < ga.populationSize ; k++ {
        indices[k] = k

        // calculate initial best solution
        bestSolution = Min( bestSolution, makespans[k] )
    }

    limitCross := int( float64(ga.populationSize) * ga.percentOfCross )
    if limitCross % 2 != 0 {
        limitCross--
    }
    fmt.Println("Number of individuals to be crossed ", limitCross)
    numberOfChilds := limitCross / 2
    childs := make([]int, numberOfChilds * sequenceSize)
    childsMakespans := make([]int, numberOfChilds)
    childsIndices := make([]int, numberOfChilds)

    for k := 0 ; k < numberOfChilds ; k++ {
        childsIndices[k] = k
    }

    r := rand.New( rand.NewSource( time.Now().Unix() * int64(1) * 17) )

    for g := 0 ; g < ga.numberOfGenerations ; g++ {
        sort.SliceStable(indices, func(i, j int) bool { 
            x := indices[i]
            y := indices[j]
            return makespans[x] < makespans[y] 
        })

        // Selection and cross operation
        for p := 0 ; p < limitCross ; p = p + 2 {

            index1 := indices[p]
            low1 := index1 * sequenceSize 
            high1 := low1 + sequenceSize 
            parent1 :=  sequences[ low1: high1 ]

            index2 := indices[p+1]
            low2 := index2 * sequenceSize 
            high2 := low2 + sequenceSize 
            parent2 :=  sequences[ low2: high2 ]

            low3 := (p/2) * sequenceSize
            high3 := low3 + sequenceSize
            child := childs[low3: high3]
            crossSequences( parent1, parent2, child, p % 2 == 0, jobshop, 0, r)

            mk := computeMakespanOfSequence(child, sequenceSize, jobshop) 
            bestSolution = Min( bestSolution, mk )
            childsMakespans[ p/2 ] = mk
        }

        sort.SliceStable(childsIndices, func(i, j int) bool { 
            x := childsIndices[i]
            y := childsIndices[j]
            return childsMakespans[x] < childsMakespans[y] 
        })

        // integrate children to population
        k := ga.populationSize
        lowestMakespanChildrens := childsMakespans[ childsIndices[0] ]
        for makespans[ indices[k - 1] ] > lowestMakespanChildrens && ( (ga.populationSize) - k < numberOfChilds - 1) {
            k--
        }
        for ; k < ga.populationSize; k++ {
            iChild := ga.populationSize - k
            makespans[ indices[k] ] = childsMakespans[ childsIndices[iChild] ]

            low := iChild * sequenceSize
            high := low + sequenceSize
            child := childs[low: high]

            low2 := indices[k] * sequenceSize
            high2 := low2 + sequenceSize
            sequence := sequences[low2: high2]
            for j := 0 ; j < sequenceSize ; j++ {
                sequence[j] = child[j]
            }
        }
        
        // mutation operation
        wg.Add(numberOfGoRoutines)
        for r := 0 ; r < numberOfGoRoutines ; r++{
            low := r * N * sequenceSize
            high := low + N * sequenceSize

            low2 := r * N
            high2 := low2 + N

            go mutationOverSequences(sequences[low:high], makespans[low2:high2], N, sequenceSize, r, jobshop, &wg, rngs[r])
        }
        wg.Wait()

    }

    for k := 0 ; k < ga.populationSize ; k++ {
        s := indices[k]

        low := s * sequenceSize
        high := low + sequenceSize
        printableSequence := sequences[low:high]
        makespan := makespans[s]
        for _, value := range printableSequence {
            fmt.Printf("%d ", value)
        }
        fmt.Printf(" | %d\n", makespan)
    }

    fmt.Println("Best solution found =", bestSolution)
    
}

func crossSequences( parent1, parent2, child []int, odd bool, jobshop JobShopSpecification, idGoRoutine int, r *rand.Rand ) {
    sequenceSize := len(parent1)
    // r := rand.New( rand.NewSource( time.Now().Unix() * int64(idGoRoutine) * 17) )

    // choose 2 points
    p1 := r.Intn( sequenceSize )
    p2 := r.Intn( sequenceSize )

    if p1 > p2 {
        p1, p2 = p2, p1
    }

    if odd {
        parent1, parent2 = parent2, parent1
    }

    // make cross
    for k := 0; k < p1; k++{
        child[k] = parent1[k]
    }
    for k := p1; k < p2; k++{
        child[k] = parent2[k]
    }
    for k := p2; k < sequenceSize; k++{
        child[k] = parent1[k]
    }

    histogram := make([]int, jobshop.numberOfJobs)
    for _, job := range child {
        histogram[job]++
    }

    // fix sequence
    o := 0
    for job, _ := range histogram {
        for histogram[job] < jobshop.numberOfMachines {
            // search and element on the child that exceeds the total allowed
            for histogram[ child[o] ] <= jobshop.numberOfMachines {
                o++
            }

            j := child[o]
            child[o] = job
            histogram[j]--
            histogram[job]++
            o++
        }
    }

    // assert sequence
    for _, job := range child {
        if histogram[job] != jobshop.numberOfMachines {
            fmt.Println("Fatal error, some sequence of machine has", histogram[job],"occurrences of job",job,"on his sequence")
            os.Exit(3)
        }
    }

}