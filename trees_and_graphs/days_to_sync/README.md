Given a matrix of `0s` and `1s`, where `0` means that a server is not synced and `1` means that it is synced.
Return the number of days that it would take to sync all the servers not synced.
A synced server can sync any non synced server if it's adjacent (`up, down, left, right`).

Example:

- Input:

```
[[1, 0, 0, 0],
[0, 0, 0, 0],
[0, 0, 0, 0],,
[0, 0, 0, 0]]
```

- Output: `6`

Companies [`Amazon`]

Source `Amazon Interview`
