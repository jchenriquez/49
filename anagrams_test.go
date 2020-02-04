package anagrams

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sort"
	"testing"
)

type Test struct {
	Input  []string   `json: "input"`
	Output [][]string `json: "Output"`
}

func toComparableString(strs [][]string) string {

	ret := make([]string, 0, len(strs))

	for _, st := range strs {

		t := make([]string, len(st))
		copy(t, st)

		sort.Slice(t, func(i, j int) bool {
			return t[i] < t[j]
		})

		var conct string

		for _, tstr := range t {
			conct = fmt.Sprintf("%s%s", conct, tstr)
		}

		ret = append(ret, conct)
	}

	sort.Slice(ret, func(i, j int) bool {
		return ret[i] < ret[j]
	})

	var retStr string

	fmt.Println("46")
	for index, str := range ret {
		retStr = fmt.Sprintf("%s%s", retStr, str)
		if index < len(ret)-1 {
			retStr = fmt.Sprintf("%s|", retStr)
		}
	}

	return retStr
}

func TestGroupAnagrams(tst *testing.T) {

	f, err := os.Open("./test.json")

	if err != nil {
		tst.Error(err)
	}

	defer f.Close()

	reader := bufio.NewReader(f)
	decoder := json.NewDecoder(reader)

	for {
		tests := make(map[string]Test)
		err := decoder.Decode(&tests)
		fmt.Println("line 73")

		if err == nil {

			for name, caseTest := range tests {
				tst.Run(name, func(st *testing.T) {
					testResult := GroupAnagrams(caseTest.Input)

					if toComparableString(testResult) != toComparableString(caseTest.Output) {
						tst.Errorf("error with test %v\n testResult %v\n", caseTest, testResult)
					}
				})
			}

		} else if err == io.EOF {
			break
		} else {
			tst.Error(err)
		}
	}

}
