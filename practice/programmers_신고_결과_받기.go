import (
	"strings"
)

func solution(ids []string, reports []string, k int) []int {
    reportSet := map[string]struct{}{}
	for _, report := range reports {
		reportSet[report] = struct{}{}
	}
	reportCnt := map[string]int{}
	reportedBy := map[string][]string{}
	for report := range reportSet {
		names := strings.Split(report, " ")
		reportCnt[names[1]]++
		reportedBy[names[1]] = append(reportedBy[names[1]], names[0])
	}
	var blockedNames []string
	for name, cnt := range reportCnt {
		if cnt >= k {
			blockedNames = append(blockedNames, name)
		}
	}
	result := make([]int, len(ids))
	for _, blockedName := range blockedNames {
		reporters := reportedBy[blockedName]
		for _, reporter := range reporters {
			for i := range ids {
				if ids[i] == reporter {
					result[i]++
				}
			}
		}
	}
	return result
}
