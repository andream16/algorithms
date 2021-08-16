package island

func dfs(grid [][]byte, i, j int) {
    if grid[i][j] != '1' {
        return
    }
    
    grid[i][j] = '0'
    
    for _, dir := range [][]int{{1,0},{0,1},{-1,0},{0,-1}} {
        x, y := i+dir[0], j+dir[1]
        if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[i]) {
            continue
        }
        dfs(grid, x, y)
    }
}

// We only need to call dfs for every land that we find while looping the matrix.
// When we find a land in the main loop, it means that is a new land.
// The DFS function will take care of recursively call itself to check the 4 adjacent lands.
// We change the value from 0 to 1 when we find lands so we don't take extra space and don't visit the same location
// more than once.
//
// T: O(r x c) where r is the number of rows and c is the number of columns.
// S: O(r x c) as in the worst case the queue can contain all the elements.
func numIslands(grid [][]byte) int {
    counter := 0
    for i := 0; i<len(grid); i++ {
        for j:=0; j<len(grid[i]); j++ {
            if grid[i][j] == '1' {
                counter++
                dfs(grid, i, j)
            }
        }
    }
    return counter
}
