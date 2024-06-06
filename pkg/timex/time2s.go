package timex

import (
	"fmt"
	"strconv"
)

func ConvertTimeToSeconds(timeStr string) (string, error) {
	lastChar := timeStr[len(timeStr)-1]
	valueStr := timeStr[:len(timeStr)-1]
	value, err := strconv.ParseInt(valueStr, 10, 64)
	if err != nil {
		return "", err
	}
	var secondData int64

	switch lastChar {
	case 'm':
		if value < 2 {
			return "", fmt.Errorf("时间间隔太短，请至少一次获取2分钟以上的数据", value)
		}
		secondData = value * 60
	case 'h':
		secondData = value * 3600
	case 'd':
		secondData = value * 86400
	case 'w':
		secondData = value * 604800
	}
	//整除100
	splitData := secondData / 100
	return fmt.Sprintf("%ds", splitData), nil
}
