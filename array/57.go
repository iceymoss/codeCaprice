package array

func generateMatrix(n int) [][]int {
	//每一圈的起始位置
	startx, starty := 0, 0
	//表示圈数
	var loop int = n / 2
	//中间位置
	var center int = n / 2

	//从1开始填充值
	count := 1
	//每一个边界的拐角，给到下一个边界处理, 需要控制每一条边遍历的长度，每次循环右边界收缩一位
	offset := 1
	//先生成n*n的矩阵
	ans := make([][]int, n)
	for i := 0; i < n; i++ {
		ans[i] = make([]int, n)
	}

	for loop > 0 {
		i, j := startx, starty
		//由左向右，行不变
		for j = starty; j < n-offset; j++ {
			ans[i][j] = count
			count++
		}

		//由上到下，列不变
		for i = startx; i < n-offset; i++ {
			ans[i][j] = count
			count++
		}

		//由右向左, 行不变，j在上一个操作中已经移到了最右边
		for ; j > starty; j-- {
			ans[i][j] = count
			count++
		}

		for ; i > startx; i-- {
			ans[i][j] = count
			count++
		}
		//向内圈缩
		startx++
		starty++
		offset++
		loop--
	}

	//如果是奇数，中间的数需要自己填充
	if n%2 == 1 {
		ans[center][center] = n * n
	}
	return ans
}
