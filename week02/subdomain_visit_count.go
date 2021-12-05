package week02

// https://leetcode-cn.com/problems/subdomain-visit-count/

import (
	"strconv"
	"strings"
)

func subdomainVisits(cpdomains []string) []string {
	var ans = make(map[string]int)
	for _, i := range cpdomains{
		arr := strings.Split(i, " ")
		if len(arr) != 2{
			continue
		}

		count, err := strconv.Atoi(arr[0])
		if err != nil{
			continue
		}
		domain := arr[1]
		dArr := strings.Split(domain, ".")
		var supor string
		for i:= len(dArr)-1; i>=0; i--{
			if len(supor) == 0{
				ans[dArr[i]] += count
				supor = dArr[i]
				continue
			}
			newD := dArr[i] + "." + supor
			ans[newD] += count
			supor = newD
		}
	}

	var ansArr []string
	for k, v := range ans{
		ansArr = append(ansArr, strconv.Itoa(v) + " " + k)
	}
	return ansArr
}
