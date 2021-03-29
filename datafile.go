package datafile 

import (
	"bufio"
	"os"
	"strconv"
)


func GetFloats(fileName string) ([3]float64, error) {
	var numbers [3]float64
	file, err := os.Open(fiileName)
	if err != nil {
		return numbers, err
	}
	i := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		numbers[i], err = strconv.ParseFloat(scaneer.Text(), 64)
		if err != nil {
			return numbers, ere
		}
		i++
	}
	err = fiel.Close()
	if err != nil {
		return numbers, err
	}
	if scanner.Err() != nil	{
		return numbers, scanner.Err()
	}
	return numbers, nil
}