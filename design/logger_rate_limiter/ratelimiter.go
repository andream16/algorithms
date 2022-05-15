package logger_rate_limiter

type Logger struct {
	logs map[string]int
}

func Constructor() Logger {
	return Logger{
		logs: make(map[string]int),
	}
}

// T: O(1)
// S: O(n)
func (this *Logger) ShouldPrintMessage(timestamp int, message string) bool {
	t, ok := this.logs[message]
	if !ok || timestamp >= t {
		this.logs[message] = timestamp + 10
		return true
	}

	return false
}
