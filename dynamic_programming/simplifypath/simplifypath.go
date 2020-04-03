func simplifyPath(path string) string {
    const rootPath = "/"

    switch path {
    case "":
        return ""
    case "/../":
        return rootPath
    }
    
    lastIdx := len(path)-1
    
    if path[lastIdx] == '/' {
        path = path[:lastIdx]
    }

    path = strings.Replace(path, "//", "/", -1)
    
    fp := []string{}
    
    for _, s := range strings.Split(path, "/") {
        switch s {
            case "..":
            if len(fp) == 0 {
                continue
            }
            fp = fp[:len(fp)-1]
            continue
            case ".", "":
            continue
        }
        fp = append(fp, s)
    }
    
    return "/" + strings.Join(fp, "/")
    
}
