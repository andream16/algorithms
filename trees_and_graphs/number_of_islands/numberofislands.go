type island struct {
    lands []*land
}

type land struct {
    name string
    x int
    y int
    ground bool
}

func newLand(x, y int, g byte) *land {
    return &land{
        name: strconv.Itoa(x)+strconv.Itoa(y),
        ground: g == '1',
        x: x,
        y: y,
    }
}

type queue struct {
    lands []*land
}

func (q *queue) empty() bool {
    return len(q.lands) == 0
}

func (q *queue) enqueue(l *land) {
    if q.empty() {
        q.lands = []*land{}
    }
    q.lands = append(q.lands, l)
}

func (q *queue) dequeue() *land {
    l := q.lands[0]
    q.lands = q.lands[1:]
    return l
}

func numIslands(grid [][]byte) int {

    if len(grid) == 0 || len(grid[0]) == 0 {
        return 0
    }
    
    islands := []*island{}
    seen := map[string]struct{}{}
    saw := func(n string, h map[string]struct{}) bool {
        _, ok := h[n]
        return ok
    }
    q := &queue{}
    q.enqueue(newLand(0, 0, grid[0][0]))
    
    for !q.empty() {
        
        l := q.dequeue()
        if saw(l.name, seen) {
            continue
        }
        
        seen[l.name] = struct{}{}
        
        adj := getAdjacentLands(l.x, l.y, grid)
        
        if l.ground {
            islands = appendToIslands(l, islands)
        }
        
        for _, al := range adj {
            q.enqueue(al)
        }
        
    }
    
    return len(islands)
    
}

func appendToIslands(l *land, islands []*island) []*island {
    for i, isl := range islands {
        for _, lnd := range isl.lands {
            if l.x == lnd.x || l.y == lnd.y {
                isl.lands = append(isl.lands, l)
                islands[i] = isl
                return islands
            }
        }
    }
    return append(islands, &island{
        lands: []*land{
            l,
        },
    })
}

func getAdjacentLands(x, y int, grid [][]byte) []*land {
    lx, ly := len(grid)-1, len(grid[0])-1
    if lx <= 1 {
        return make([]*land, 0)
    }
    if x == 0 {
        if y == 0 {
            return []*land{
                newLand(x, y+1, grid[x][y+1]),
                newLand(x+1, y, grid[x+1][y]),
            }
        }
        if y == ly {
            return []*land{
                newLand(x, y-1, grid[x][y-1]),
                newLand(x+1, y, grid[x+1][y]),
            }
        }
        return []*land{
            newLand(x, y+1, grid[x][y+1]),
            newLand(x+1, y, grid[x+1][y]),
            newLand(x, y-1, grid[x][y-1]),
        }
    }
    if x == lx {
        if y == 0 {
            return []*land{
                newLand(x-1, y, grid[x-1][y]),
                newLand(x, y+1, grid[x][y+1]),
            }
        }
        if y == ly {
            return []*land{
                newLand(x-1, y, grid[x-1][y]),
                newLand(x, y-1, grid[x][y-1]),
            }
        }
        return []*land{
            newLand(x-1, y, grid[x-1][y]),
            newLand(x, y+1, grid[x][y+1]),
            newLand(x, y-1, grid[x][y-1]),
        }
    }
    if y == 0 {
        return []*land{
            newLand(x-1, y, grid[x-1][y]),
            newLand(x, y+1, grid[x][y+1]),
            newLand(x+1, y, grid[x+1][y]),
        }
    }
    if y == ly {
        return []*land{
            newLand(x-1, y, grid[x-1][y]),
            newLand(x, y-1, grid[x][y-1]),
            newLand(x+1, y, grid[x+1][y]),
        }
    }
    return []*land{
        newLand(x-1, y, grid[x-1][y]),
        newLand(x, y+1, grid[x][y+1]),
        newLand(x+1, y, grid[x+1][y]),
        newLand(x, y-1, grid[x][y-1]),
    }
}
