package tools

import "os"

func PathExists(path string) (bool, error) {
	//os.Stat 的err的几种情况, 第一种没有错误,目录存在;第二种 错误为os.isNotExist 目录不存在; 第三种 其他错误
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
