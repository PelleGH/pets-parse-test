// package main

// import (
// 	"fmt"
// 	"os"
// 	"strings"
// )

// type Node struct {
// 	NodeName  string
// 	NodeEdges map[string][]string
// }

// func main() {
// 	// Read the file content
// 	data, err := os.ReadFile("OWL.txt")
// 	if err != nil {
// 		panic(err)
// 	}

// 	content := string(data)
// 	nodesMap := make(map[string]*Node)

// 	// Split data into node sections
// 	nodeDataSections := strings.Split(content, "\n\n")

// 	for _, nodeData := range nodeDataSections {
// 		lines := strings.Split(nodeData, "\n")
// 		if len(lines) == 0 {
// 			continue
// 		}

// 		// Extract the node label
// 		labelParts := strings.SplitN(lines[0], " ", 2)
// 		if len(labelParts) == 0 {
// 			continue
// 		}
// 		label := strings.TrimPrefix(labelParts[0], "minecraft:")

// 		if _, exists := nodesMap[label]; !exists {
// 			nodesMap[label] = &Node{NodeName: label, NodeEdges: make(map[string][]string)}
// 		}

// 		// Process attributes (edges)
// 		for _, attribute := range lines[1:] {
// 			if strings.HasPrefix(attribute, "    minecraft:") {
// 				edgeData := strings.TrimPrefix(attribute, "    minecraft:")
// 				edgeParts := strings.SplitN(edgeData, " ", 2)
// 				if len(edgeParts) < 2 {
// 					continue
// 				}
// 				edgeLabel := edgeParts[0]
// 				nodesList := strings.Split(edgeParts[1], ", ")

// 				for _, nodeName := range nodesList {
// 					nodeLabel := strings.TrimSuffix(strings.TrimPrefix(nodeName, "minecraft:"), ";")
// 					nodeLabel = strings.TrimRight(nodeLabel, " .")

// 					if _, exists := nodesMap[nodeLabel]; !exists {
// 						nodesMap[nodeLabel] = &Node{NodeName: nodeLabel, NodeEdges: make(map[string][]string)}
// 					}

// 					nodesMap[label].NodeEdges[edgeLabel] = append(nodesMap[label].NodeEdges[edgeLabel], nodeLabel)
// 				}
// 			}
// 		}
// 	}

// 	// Print nodes and their edges in a structured format
// 	for _, node := range nodesMap {
// 		fmt.Printf("Node: %s\n", node.NodeName)
// 		for edge, connections := range node.NodeEdges {
// 			fmt.Printf("  %s:\n", edge)
// 			for _, conn := range connections {
// 				fmt.Printf("    -> %s\n", conn)
// 			}
// 		}
// 		fmt.Println()
// 	}
// }
