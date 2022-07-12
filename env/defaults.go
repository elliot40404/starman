package env

import (
	"fmt"
	"os"
)

const TIME_UNIT string = "Seconds"
const DELAY int = 10
const COMMENT_CHAR string = "@REM"

func HOMEDIR() string {
	home, _ := os.UserHomeDir()
	return fmt.Sprintf("%s\\.sm", home)
}

func LOG_FILE() string {
	return fmt.Sprintf("%s\\sm_startup.log", HOMEDIR())
}

func STARTUP_DIR() string {
	home, _ := os.UserHomeDir()
	return fmt.Sprintf("%s\\AppData\\Roaming\\Microsoft\\Windows\\Start Menu\\Programs\\Startup", home)
}

func STARTUP_FILE() string {
	home, _ := os.UserHomeDir()
	return fmt.Sprintf("%s\\AppData\\Roaming\\Microsoft\\Windows\\Start Menu\\Programs\\Startup\\sm_startup.vbs", home)
}

func STARTUP_FILE_DEFAULT() string {
	return fmt.Sprintf("Set WshShell = CreateObject(\"WScript.Shell\")\nWshShell.Run chr(34) & \"%s\" & Chr(34), 0\nSet WshShell = Nothing", CONFIG_FILE())
}

// func STARTUP_FILE() string {
// 	home, _ := os.UserHomeDir()
// 	return fmt.Sprintf("%s\\AppData\\Roaming\\Microsoft\\Windows\\Start Menu\\Programs\\Startup\\sm_startup.bat", home)
// }

func CONFIG_FILE() string {
	return fmt.Sprintf("%s\\sm_startup.bat", HOMEDIR())
}

func CONFIG_FILE_DEFAULT() string {
	return fmt.Sprintf("@ECHO OFF\n\n@REM DELAY\nTIMEOUT /T %d", DELAY)
}
