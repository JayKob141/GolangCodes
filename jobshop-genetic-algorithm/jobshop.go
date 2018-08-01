package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
    "strconv"
    "math/rand"
    "time"
	"sync"
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

    ga := GeneticAlgorithmParameters { numberOfGenerations: 16, populationSize: 16, percentOfCross: 0.5, percentOfMutation: 0.1 }
     
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

func geneticAlgorithm(ga GeneticAlgorithmParameters, jobshop JobShopSpecification){

    numberOfGoRoutines := 4
    
    //  Init population
    sequenceSize := jobshop.numberOfJobs * jobshop.numberOfMachines
    sequences := make( []int, ga.populationSize * sequenceSize )
    makespans := make( []int, ga.populationSize )
    // ganttChart := make( []int, ga.populationSize * jobshop.numberOfMachines )
    // previousFinalTimes :=  make( []int, ga.populationSize * jobshop.numberOfJobs )
    // iOperation :=  make( []int, ga.populationSize * jobshop.numberOfJobs )
    // indices := make( []int, ga.populationSize )

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

    for s := 0 ; s < ga.populationSize ; s++ {
        low := s * sequenceSize
        high := low + sequenceSize
        printableSequence := sequences[low:high]
        makespan := computeMakespanOfSequence(printableSequence, sequenceSize, jobshop)
        makespans[s] = makespan
        for _, value := range printableSequence {
            fmt.Printf("%d ", value)
        }
        fmt.Printf(" | %d\n", makespan)
    }

    for g := 0 ; g < ga.numberOfGenerations ; g++ {
        // TODO:
        // sort sequences by makespan

        // selection of individuals
        
        // cross operation
        // mutation operation

        // evaluate population (compute makespans)

        // update solution with minimal cost
    }
    
}
