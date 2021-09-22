package utils

import "fmt"

func PrintBanner(){
	fmt.Println(
		"   _____        _    _             \n" +
			"  / ____|      | |  | |            \n" +
			" | |      ___  | |__| |__   __ ___ \n" +
			" | |     / _ \\ |  __  |\\ \\ / // __|\n" +
			" | |____| (_) || |  | | \\ V / \\__ \\\n" +
			"  \\_____|\\___/ |_|  |_|  \\_/  |___/  v0.3.0  \n" +
			"                                  --By HeyJavaBean")

}

var printLog = false

func UseLog(){
	printLog = true
}

func UnuseLog(){
	printLog = false
}


func PrintLog(a ...interface{}){
	if printLog{
		fmt.Println(a)
	}
}
