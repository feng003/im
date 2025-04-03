package tool

import (
	"sort"
	"strconv"
	"strings"
)

// FindLatestVersion 获得最新版本号
func FindLatestVersion(versions []string) string {
	if len(versions) == 0 {
		return ""
	}

	// 使用 sort.Slice 对版本号进行排序
	sort.Slice(versions, func(i, j int) bool {
		return isNewerVersion(versions[i], versions[j])
	})

	return versions[0]
}

func isNewerVersion(v1, v2 string) bool {
	// 版本号按 "." 进行分割
	ver1 := splitVersion(v1)
	ver2 := splitVersion(v2)

	// 逐级比较版本号
	for i := 0; i < len(ver1) && i < len(ver2); i++ {
		if ver1[i] > ver2[i] {
			return true
		} else if ver1[i] < ver2[i] {
			return false
		}
	}

	// 版本号前几部分相同，比较长度
	return len(ver1) > len(ver2)
}

func splitVersion(version string) []int {
	var ver []int
	// 按 "." 分割版本号，并转换为整数
	for _, v := range strings.Split(version, ".") {
		num, _ := strconv.Atoi(v)
		ver = append(ver, num)
	}
	return ver
}
