package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var hName string
var edges []string

type Node struct {
	NodeName string
	NodeID   int
}
type Tuple struct { // FOR EDGES
	key   string // ex. obtainedBy
	value string // ex. recipe
}

func main() {
	nodeLst := map[Node][]Tuple{} // NODE HASHMAP WITH A TUPLE LIST (EDGES) AS VALUE

	file, err := os.Open("Example Data_A.ttl")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	var lines []string
	var newNode Node
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)

		if strings.HasPrefix(line, "minecraft:") {
			temp := strings.TrimPrefix(line, "minecraft:")

			wrd := getWrd(temp) // FIRST WORD

			newNode.NodeName = wrd // ASSIGN TO NODENAME
			fmt.Printf(wrd + "\n\n")

		} else if strings.HasPrefix(line, "	nodeOntology:hasID ") {
			temp := strings.TrimPrefix(line, "	nodeOntology:hasID ")

			wrd := getWrd(temp) // FIRST WORD

			i, err := strconv.Atoi(wrd) // STRING TO INT
			if err != nil {
				fmt.Println("Error reading file:", err)
			}
			newNode.NodeID = i // ASSIGN TO NODEID
			fmt.Println(i)
		} else if strings.HasPrefix(line, "    minecraft:obtainedBy") || (strings.HasPrefix(line, "    minecraft:hasInput")) || (strings.HasPrefix(line, "    minecraft:hasOutput") || (strings.HasPrefix(line, "    minecraft:usedInStation"))) {
			var tempTuple Tuple
			if strings.HasPrefix(line, "    minecraft:obtainedBy") {
				temp := strings.TrimPrefix(line, "    minecraft:obtainedBy minecraft:")
				wrd := getWrd(temp)
				tempTuple.key = "obtainedBy"
				tempTuple.value = wrd
			} else if strings.HasPrefix(line, "    minecraft:hasInput") {
				temp := strings.TrimPrefix(line, "    minecraft:hasInput minecraft:")
				wrd := getWrd(temp)
				tempTuple.key = "hasInput"
				tempTuple.value = wrd
			} else if strings.HasPrefix(line, "    minecraft:hasOutput") {
				temp := strings.TrimPrefix(line, "    minecraft:hasOutput minecraft:")
				wrd := getWrd(temp)
				tempTuple.key = "hasOutput"
				tempTuple.value = wrd
			} else if strings.HasPrefix(line, "    minecraft:usedInStation") {
				temp := strings.TrimPrefix(line, "    minecraft:usedInStation minecraft:")
				wrd := getWrd(temp)
				tempTuple.key = "usedInStation"
				tempTuple.value = wrd
			}
			nodeLst[newNode] = append(nodeLst[newNode], tempTuple)
			//fmt.Println(nodeLst)
		}
		if strings.HasSuffix(line, ";") {
			continue
		} else if strings.HasSuffix(line, ".") { // END OF NODE
			var tempTuple Tuple
			nodeLst[newNode] = append(nodeLst[newNode], tempTuple)
			//nodeLst = append(nodeLst, newNode)	// APPEND TO LIST OF NODES

		} else {
			continue // NEWLINE/EMPTY SPACE
		}
		//fmt.Println(line)

	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
	fmt.Println(nodeLst)
	fmt.Printf("bomba")
}
func getWrd(w string) string {
	wrd := ""
	for i := range w {
		if w[i] == ' ' {
			wrd = w[0:i]
			break
		}
	}
	return wrd
}
