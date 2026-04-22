func sortColors(nums []int) {
    c := [3]int{}
    for _, v := range nums {
        c[v]++
    }
    t := 0
    for range c[0] {
        nums[t] = 0
        t++
    }
    for range c[1] {
        nums[t] = 1
        t++
    }
    for range c[2] {
        nums[t] = 2
        t++
    }
}
