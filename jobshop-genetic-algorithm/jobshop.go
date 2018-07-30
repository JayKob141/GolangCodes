package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
    "strconv"
)

type JobShopSpecification struct {
    numberOfMachines int
    numberOfJobs int
    operationsMatrix [][]int
    timesMatrix [][]int
}

/*
struct GeneticAlgorithmParameters {
    numberOfGenerations int
    populationSize int
    percentOfCross float64
    percentOfMutation float64
}
*/

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
        fmt.Println("Expected 2 argument (instance file and parameters file)")
        return
    }
    
    fmt.Println(argsWithoutProg)

    readFile( argsWithoutProg[0] )
     
    // TODO: parse configuration and GA parameters file

    geneticAlgorithm()
}

func geneticAlgorithm(){
    //  TODO: Init population

    // TODO: repeat N generations
        // selection of individuals 
        // cross operation
        // mutation operation
        // update solution with minimal cost
    
}
