package utils

const (
    Black   = "\033[30m"
    Red     = "\033[31m"
    Green   = "\033[32m"
    Yellow  = "\033[33m"
    Blue    = "\033[34m"
    Magenta = "\033[35m"
    Cyan    = "\033[36m"
    White   = "\033[37m"
    Purple  = "\033[35;1m"
    Pink    = "\033[38;5;200m"
    Orange  = "\033[38;5;208m"
    Reset   = "\033[0m"

    BgWhite = "\033[47m"
    BgBlack = "\033[40m"
    BgRed     = "\033[41m"
    BgGreen   = "\033[42m"
    BgYellow  = "\033[43m"
    BgBlue    = "\033[44m"
    BgMagenta = "\033[45m"
    BgCyan    = "\033[46m"
    BgGrey    = "\033[47m"
    BgReset   = "\033[0m"
)

func Color(text string, color string) string {
    return color + text + Reset
}

func ColorWithBackground(text string, textColor string, bgColor string) string {
    return textColor + bgColor + text + Reset
}