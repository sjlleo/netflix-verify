package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

// Netflix - Tesing URL
const Netflix = "https://www.netflix.com/title/"

// VerifyAreaAvailable - Check whether the area supports the NetFlix
func VerifyAreaAvailable() {
	status := GetResponse(80018499)
	fmt.Println("** NetFlix 解锁检测小工具 By \033[1;36m@sjlleo\033[0m **")
	if status == -1{
		fmt.Println("\033[0;33mNetFlix不在您测试的出口IP所在的地区提供服务，无法正常使用NF网站\033[0m")
	} else if status == -2 {
		fmt.Println("\033[0;33m网络异常，连接NetFlix失败，请检查后重试\033[0m")
	} else {
		fmt.Println("\033[0;32mNetFlix在您测试的出口IP所在的地区提供服务，宽松版权的自制剧可以解锁\033[0m")
	}
}

// VerifySelfMade - Check Self-Made Movie
func VerifySelfMade() bool {
	fmt.Println("->> 正在检查是否完整支持自制剧 <<-")
	SelfMade := 80197526
	if GetResponse(SelfMade) == 1 {
		fmt.Println("\033[0;32m支持解锁全部的自制剧\033[0m")
		return true
	} else {
		fmt.Println("\033[0;31m不支持解锁带有强版权的自制剧\033[0m")
		return false
	}
}

// VerifyNonSelfMade - Check NonSelf-Made Movie
func VerifyNonSelfMade() {
	fmt.Println("->> 正在检查支持的NetFlix地区 <<-")
	MovieSpecialID := [32]int{70143836, 80057401, 70215460, 81324707, 81370688, 81269607, 81307998, 81327393, 81052553, 81299764, 81029396, 81240443, 81297517, 70241077, 81043527, 81232435, 81396420, 81351679, 60036229, 81290755, 81016002, 81382683, 80227207, 70295176, 81186099, 70271605, 81288438, 81323786, 81347799, 70108563, 81191900, 60020039}
	MovieCountry := [32]string{"通用地", "新加坡", "香港", "南非", "日本", "美国", "泰国", "俄罗斯", "德国", "韩国", "比利时", "冰岛", "印度", "英国", "土耳其", "罗马尼亚", "波兰", "瑞典", "立陶宛", "法国", "阿根廷", "澳大利亚", "加拿大", "台湾/马来西亚", "荷兰", "西班牙", "匈牙利", "以色列", "巴西", "捷克共和国", "斯洛伐克", "希腊"}
	MainLanguage := [32]string{"无","中文简体、繁体","中文简体、繁体","英语","日语","英语","泰语、英文","俄语","德语、英语","韩语","荷兰语、英语","英语","北印度语","英语","土耳其语","罗马尼亚语、西班牙语、英语","波兰语、英语","瑞典语、丹麦语","英语","法语、英语","西班牙语","英语","英语、法语","英语、粤语、中文简体、中文繁体","荷兰语","西班牙语","匈牙利语","阿拉伯语","葡萄牙语、西班牙语","捷克语","捷克语","希腊语"}

	if GetResponse(MovieSpecialID[0]) != 1 {
		fmt.Println("\033[0;31m不支持解锁非自制剧\033[0m")
		return
	}

	fmt.Println("\033[0;32m支持解锁全地区通用的非自制剧，正在检测地区...\033[0m")

	for i := 1; i < 32; i++ {
		if GetResponse(MovieSpecialID[i]) == 1 {
			fmt.Println("\033[0;36m原生IP地域解锁信息：\033[1;36m" + MovieCountry[i] + "区 NetFlix\033[0m")
			fmt.Println("\033[0;36m此地区NF的音轨/字幕主要是：\033[1;36m" + MainLanguage[i] + "\033[0m")
			return
		}
	}

	fmt.Println("\033[1;32m支持解锁全地区通用的非自制剧，但是不能解锁个性化地区的NetFlix\033[0m")

}

// GetResponse - Verify
func GetResponse(movieid int) int {
	resp, err := http.Get(Netflix + strconv.Itoa(movieid))
	if err != nil {
		fmt.Println(err)
		return -2
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	res := string(body)

	if strings.Contains(res, "Not Available") {
		return -1
	}

	if strings.Contains(res, "page-404") || strings.Contains(res, "NSEZ-403") {
		return 0
	}

	return 1
}

func main() {
	VerifyAreaAvailable()
	if VerifySelfMade() {
		VerifyNonSelfMade()
	}
}
