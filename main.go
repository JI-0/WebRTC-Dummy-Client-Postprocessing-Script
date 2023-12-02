package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

const original_file_name = "timestamps"

type ClientList map[int]int

var clients ClientList

const maxPeers = 100

func main() {
	original_file, original_file_err := os.Open(original_file_name + ".csv")
	if original_file_err != nil {
		log.Fatal(original_file_err)
	}
	defer original_file.Close()
	new_name := "Processed_" + original_file_name + ".csv"
	new_name_sum := "Sum_" + original_file_name + ".csv"
	new_file, new_file_err := os.OpenFile(new_name, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if new_file_err != nil {
		log.Fatal(new_file_err)
	}
	new_file_sum, new_file_sum_err := os.OpenFile(new_name_sum, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if new_file_sum_err != nil {
		log.Fatal(new_file_err)
	}

	scanner := bufio.NewScanner(original_file)
	writer := bufio.NewWriter(new_file)
	writer_sum := bufio.NewWriter(new_file_sum)

	clients = make(ClientList)
	min := 16976116344130
	max := 0
	maxK := 0
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "New client:") {
			newLine := ""
			// maxK := 0
			// for k := range clients {
			// 	if k > maxK {
			// 		maxK = k
			// 	}
			// }
			avg := 0
			minP := 16976116344130
			maxP := 0
			for i := 0; i < maxK && i < maxPeers; i++ {
				// cli, ok := clients[i]
				// if !ok {
				// 	newLine += "0,"
				// 	continue
				// }
				sumOfC := clients[i]
				// println(sumOfC)
				pacPerS := 1000 * sumOfC / (max - min)
				newLine += strconv.Itoa(pacPerS) + ","
				avg += pacPerS
				if pacPerS < minP {
					minP = pacPerS
				}
				if pacPerS > maxP {
					maxP = pacPerS
				}
				// if clients[cli] != 0 {
				// 	newLine += strconv.Itoa(1000*clients[cli]/(max-min)) + ","
				// 	// writer.WriteString(strconv.Itoa(cli) + "," + strconv.Itoa(1000*sum/(max-min)) + "\n")
				// } else {
				// 	// newLine += strconv.Itoa(1000*clients[cli]/(max-min)) + ","
				// 	newLine += "0,"
				// 	// writer.WriteString(strconv.Itoa(cli) + "," + strconv.Itoa(0) + "\n")
				// }
			}
			newLine += "\n"
			writer.WriteString(newLine)
			if avg != 0 {
				avg = avg / maxK
			}
			writer_sum.WriteString(strconv.Itoa(avg) + "," + strconv.Itoa(minP) + "," + strconv.Itoa(maxP) + "\n")

			maxK += 1
			min = 16976116344130
			max = 0
			clients = make(ClientList)
			continue
		}
		parts := strings.Split(scanner.Text(), ",")
		cliI, _ := strconv.Atoi(parts[0])
		newT, _ := strconv.Atoi(parts[1])
		if newT < min {
			min = newT
		}
		if newT > max {
			max = newT
		}

		_, ok := clients[cliI]
		if !ok {
			clients[cliI] = 1
			continue
		}
		clients[cliI] += 1
	}

	writer.Flush()
	writer_sum.Flush()
	new_file.Close()
	new_file_sum.Close()
}
