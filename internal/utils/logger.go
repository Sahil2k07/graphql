package utils

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/charmbracelet/log"
)

// ANSI color codes with bold for console output
const (
	reset         = "\033[0m"
	bold          = "\033[1m"
	brightRed     = "\033[91m"
	violet        = "\033[35m"
	green         = "\033[32m"
	cyan          = "\033[36m"
	blue          = "\033[34m"
	brightWhite   = "\033[97m"
	uniqueTeal    = "\033[38;5;37m"
	brightYellow  = "\033[93m"
	brightMagenta = "\033[95m"
	brightBlue    = "\033[94m"
	brightCyan    = "\033[96m"
	maroon        = "\033[38;5;124m"
	gray          = "\033[90m"
)

// Color helpers
func colorizeKey(text, colorCode string) string {
	return bold + colorCode + text + reset
}

func colorizeValue(text, colorCode string) string {
	return colorCode + text + reset
}

// Extracts readable stack trace info
func getCleanStackTrace(skip, maxFrames int) (string, string) {
	pc := make([]uintptr, maxFrames+skip)
	n := runtime.Callers(skip, pc)
	frames := runtime.CallersFrames(pc[:n])

	var sb strings.Builder
	count := 0
	var firstFrameStr string

	for {
		frame, more := frames.Next()
		line := fmt.Sprintf("  %s (%s:%d)\n", frame.Function, filepath.Base(frame.File), frame.Line)

		if count == 0 {
			firstFrameStr = fmt.Sprintf("%s (%s:%d)", frame.Function, filepath.Base(frame.File), frame.Line)
		}

		sb.WriteString(line)
		count++
		if count >= maxFrames || !more {
			break
		}
	}

	return firstFrameStr, sb.String()
}

// Prints formatted, colored log output for debugging GraphQL errors
func LogGraphQLError(err error) {
	logger := log.Default()

	errKey := colorizeKey("Error", brightRed)
	errMsg := colorizeValue(err.Error(), brightRed)

	locationKey := colorizeKey("Location", violet)
	locationVal, stackTrace := getCleanStackTrace(4, 10)
	locationValColored := colorizeValue(locationVal, green)

	stackHeader := colorizeKey("Stack Trace:", brightCyan)

	logger.Errorf("\n🐞 GraphQL Error\n"+
		"%s : %s\n"+
		"%s : %s\n"+
		"%s\n%s\n",
		errKey, errMsg,
		locationKey, locationValColored,
		stackHeader,
		stackTrace,
	)
}
