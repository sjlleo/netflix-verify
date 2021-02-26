package main

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

const Netflix = "https://www.netflix.com/title/"

func RequestIP(requrl string, ip string) string {
	if ip == "" {
		return "Error"
	}
	urlValue, err := url.Parse(requrl)
	if err != nil {
		return "Error"
	}
	host := urlValue.Host
	if ip == "" {
		ip = host
	}
	newrequrl := strings.Replace(requrl, host, ip, 1)
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{ServerName: host},
		},
		CheckRedirect: func(req *http.Request, via []*http.Request) error { return http.ErrUseLastResponse },
		Timeout:       5 * time.Second,
	}
	req, err := http.NewRequest("GET", newrequrl, nil)
	if err != nil {
		//return errors.New(strings.ReplaceAll(err.Error(), newrequrl, requrl))
		return "Error"
	}
	req.Host = host
	req.Header.Set("USER-AGENT", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.108 Safari/537.36")
	resp, err := client.Do(req)
	if err != nil {
		//return errors.New(strings.ReplaceAll(err.Error(), newrequrl, requrl))
		return "Error"
	}
	defer resp.Body.Close()

	Header := resp.Header

	if Header["Location"] == nil {
		return "Ban"
	} else {
		return strings.Split(Header["Location"][0], "/")[3]
	}
}

func ParseIP(s string) int {
	ip := net.ParseIP(s)
	if ip == nil {
		return 0
	}
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '.':
			return 4
		case ':':
			return 6
		}
	}
	return 0
}

func UnblockTest(MoiveID int, ip string) bool {
	testURL := Netflix + strconv.Itoa(MoiveID)
	reCode := RequestIP(testURL, ip)
	if strings.Contains(reCode, "Ban") {
		return false
	} else {
		return true
	}
}

func ShellPrinter(Num int) {
	switch Num {
	case 0:
		fmt.Println("** NetFlix 解锁检测小工具 By \033[1;36m@sjlleo\033[0m **")
		break
	case 1:
		fmt.Println("\033[0;33mNetFlix不为您测试的出口IP提供服务\033[0m")
		break
	case 2:
		fmt.Println("\033[0;32mNetFlix在您测试的出口IP所在的地区提供服务，宽松版权的自制剧可以解锁\033[0m")
		break
	case 3:
		fmt.Println("\033[0;36m[IPV4测试]\033[0m")
		break
	case 4:
		fmt.Println("\033[0;36m\n[IPV6测试]\033[0m")
		break
	case 5:
		fmt.Println("\033[0;32m支持解锁全部的自制剧\033[0m")
		break
	case 6:
		fmt.Println("\033[0;31m不支持解锁带有强版权的自制剧\033[0m")
		break
	case 7:
		fmt.Println("->> 正在检查是否完整支持自制剧 <<-")
		break
	case 8:
		fmt.Println("->> 正在检查支持的NetFlix地区 <<-")
		break
	case 9:
		fmt.Println("\033[0;32m支持解锁非自制剧\033[0m")
		break
	case 10:
		fmt.Println("\033[0;31m不支持解锁非自制剧\033[0m")
		break
	}

}

func FindCountry(Code string) string {
	countryName := []string{"阿富汗", "奥兰群岛", "阿尔巴尼亚", "阿尔及利亚", "美属萨摩亚", "安道尔", "安哥拉", "安圭拉", "南极洲", "安提瓜和巴布达", "阿根廷", "亚美尼亚", "阿鲁巴", "澳大利亚", "奥地利", "阿塞拜疆", "巴哈马", "巴林", "孟加拉国", "巴巴多斯", "白俄罗斯", "比利时", "伯利兹", "贝宁", "百慕大", "不丹", "玻利维亚", "波黑", "博茨瓦纳", "布维岛", "巴西", "英属印度洋领地", "文莱", "保加利亚", "布基纳法索", "布隆迪", "柬埔寨", "喀麦隆", "加拿大", "佛得角", "开曼群岛", "中非", "乍得", "智利", "中国", "圣诞岛", "科科斯（基林）群岛", "哥伦比亚", "科摩罗", "刚果（布）", "刚果（金）", "库克群岛", "哥斯达黎加", "科特迪瓦", "克罗地亚", "古巴", "塞浦路斯", "捷克", "丹麦", "吉布提", "多米尼克", "多米尼加", "厄瓜多尔", "埃及", "萨尔瓦多", "赤道几内亚", "厄立特里亚", "爱沙尼亚", "埃塞俄比亚", "福克兰群岛（马尔维纳斯）", "法罗群岛", "斐济", "芬兰", "法国", "法属圭亚那", "法属波利尼西亚", "法属南部领地", "加蓬", "冈比亚", "格鲁吉亚", "德国", "加纳", "直布罗陀", "希腊", "格陵兰", "格林纳达", "瓜德罗普", "关岛", "危地马拉", "格恩西岛", "几内亚", "几内亚比绍", "圭亚那", "海地", "赫德岛和麦克唐纳岛", "梵蒂冈", "洪都拉斯", "香港", "匈牙利", "冰岛", "印度", "印度尼西亚", "伊朗", "伊拉克", "爱尔兰", "英国属地曼岛", "以色列", "意大利", "牙买加", "日本", "泽西岛", "约旦", "哈萨克斯坦", "肯尼亚", "基里巴斯", "朝鲜", "韩国", "科威特", "吉尔吉斯斯坦", "老挝", "拉脱维亚", "黎巴嫩", "莱索托", "利比里亚", "利比亚", "列支敦士登", "立陶宛", "卢森堡", "澳门", "前南马其顿", "马达加斯加", "马拉维", "马来西亚", "马尔代夫", "马里", "马耳他", "马绍尔群岛", "马提尼克", "毛利塔尼亚", "毛里求斯", "马约特", "墨西哥", "密克罗尼西亚联邦", "摩尔多瓦", "摩纳哥", "蒙古", "黑山", "蒙特塞拉特", "摩洛哥", "莫桑比克", "缅甸", "纳米比亚", "瑙鲁", "尼泊尔", "荷兰", "荷属安的列斯", "新喀里多尼亚", "新西兰", "尼加拉瓜", "尼日尔", "尼日利亚", "纽埃", "诺福克岛", "北马里亚纳", "挪威", "阿曼", "巴基斯坦", "帕劳", "巴勒斯坦", "巴拿马", "巴布亚新几内亚", "巴拉圭", "秘鲁", "菲律宾", "皮特凯恩", "波兰", "葡萄牙", "波多黎各", "卡塔尔", "留尼汪", "罗马尼亚", "俄罗斯联邦", "卢旺达", "圣赫勒拿", "圣基茨和尼维斯", "圣卢西亚", "圣皮埃尔和密克隆", "圣文森特和格林纳丁斯", "萨摩亚", "圣马力诺", "圣多美和普林西比", "沙特阿拉伯", "塞内加尔", "塞尔维亚", "塞舌尔", "塞拉利昂", "新加坡", "斯洛伐克", "斯洛文尼亚", "所罗门群岛", "索马里", "南非", "南乔治亚岛和南桑德韦奇岛", "西班牙", "斯里兰卡", "苏丹", "苏里南", "斯瓦尔巴岛和扬马延岛", "斯威士兰", "瑞典", "瑞士", "叙利亚", "台湾", "塔吉克斯坦", "坦桑尼亚", "泰国", "东帝汶", "多哥", "托克劳", "汤加", "特立尼达和多巴哥", "突尼斯", "土耳其", "土库曼斯坦", "特克斯和凯科斯群岛", "图瓦卢", "乌干达", "乌克兰", "阿联酋", "英国", "美国", "美国本土外小岛屿", "乌拉圭", "乌兹别克斯坦", "瓦努阿图", "委内瑞拉", "越南", "英属维尔京群岛", "美属维尔京群岛", "瓦利斯和富图纳", "西撒哈拉", "也门", "赞比亚", "津巴布韦"}
	countryCode := []string{"af", "ax", "al", "dz", "as", "ad", "ao", "ai", "aq", "ag", "ar", "am", "aw", "au", "at", "az", "bs", "bh", "bd", "bb", "by", "be", "bz", "bj", "bm", "bt", "bo", "ba", "bw", "bv", "br", "io", "bn", "bg", "bf", "bi", "kh", "cm", "ca", "cv", "ky", "cf", "td", "cl", "cn", "cx", "cc", "co", "km", "cg", "cd", "ck", "cr", "ci", "hr", "cu", "cy", "cz", "dk", "dj", "dm", "do", "ec", "eg", "sv", "gq", "er", "ee", "et", "fk", "fo", "fj", "fi", "fr", "gf", "pf", "tf", "ga", "gm", "ge", "de", "gh", "gi", "gr", "gl", "gd", "gp", "gu", "gt", "gg", "gn", "gw", "gy", "ht", "hm", "va", "hn", "hk", "hu", "is", "in", "id", "ir", "iq", "ie", "im", "il", "it", "jm", "jp", "je", "jo", "kz", "ke", "ki", "kp", "kr", "kw", "kg", "la", "lv", "lb", "ls", "lr", "ly", "li", "lt", "lu", "mo", "mk", "mg", "mw", "my", "mv", "ml", "mt", "mh", "mq", "mr", "mu", "yt", "mx", "fm", "md", "mc", "mn", "me", "ms", "ma", "mz", "mm", "na", "nr", "np", "nl", "an", "nc", "nz", "ni", "ne", "ng", "nu", "nf", "mp", "no", "om", "pk", "pw", "ps", "pa", "pg", "py", "pe", "ph", "pn", "pl", "pt", "pr", "qa", "re", "ro", "ru", "rw", "sh", "kn", "lc", "pm", "vc", "ws", "sm", "st", "sa", "sn", "rs", "sc", "sl", "sg", "sk", "si", "sb", "so", "za", "gs", "es", "lk", "sd", "sr", "sj", "sz", "se", "ch", "sy", "tw", "tj", "tz", "th", "tl", "tg", "tk", "to", "tt", "tn", "tr", "tm", "tc", "tv", "ug", "ua", "ae", "gb", "index", "um", "uy", "uz", "vu", "ve", "vn", "vg", "vi", "wf", "eh", "ye", "zm", "zw"}
	for i, v := range countryCode {
		if strings.Contains(Code, v) {
			return countryName[i]
		}
	}
	return Code
}

func main() {
	var ipv4, ipv6 string

	var areaAvailableID = 80018499
	var SelfMadeAvailableID = 80197526
	var NonSelfMadeAvailableID = 70143836

	dns := "www.netflix.com"

	// 解析ip地址
	ns, err := net.LookupHost(dns)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Err: %s", err.Error())
		return
	}

	switch {
	case len(ns) != 0:
		for _, n := range ns {

			if ParseIP(n) == 4 {
				ipv4 = n
			}
			if ParseIP(n) == 6 {
				ipv6 = "[" + n + "]"
			}

		}

	}

	ShellPrinter(0)

	testURL := Netflix + strconv.Itoa(areaAvailableID)
	ipv4CountryCode := RequestIP(testURL, ipv4)
	ipv6CountryCode := RequestIP(testURL, ipv6)

	if !strings.Contains(ipv4CountryCode, "Error") {
		ShellPrinter(3)
		if strings.Contains(ipv4CountryCode, "Ban") {
			ShellPrinter(1)
		} else {
			ShellPrinter(2)
			ShellPrinter(7)
			if UnblockTest(SelfMadeAvailableID, ipv4) {
				ShellPrinter(5)
				ShellPrinter(8)
				if UnblockTest(NonSelfMadeAvailableID, ipv4) {
					ShellPrinter(9)
					fmt.Println("\033[0;36m原生IP地域解锁信息：\033[1;36m" + FindCountry(ipv4CountryCode) + "区 NetFlix\033[0m")
				} else {
					ShellPrinter(10)
				}
			} else {
				ShellPrinter(6)
			}
		}
	}

	if !strings.Contains(ipv6CountryCode, "Error") {
		ShellPrinter(4)
		if strings.Contains(ipv6CountryCode, "Ban") {
			ShellPrinter(1)
		} else {
			ShellPrinter(2)
			ShellPrinter(7)
			if UnblockTest(SelfMadeAvailableID, ipv6) {
				ShellPrinter(5)
				ShellPrinter(8)
				if UnblockTest(NonSelfMadeAvailableID, ipv6) {
					ShellPrinter(9)
					fmt.Println("\033[0;36m原生IP地域解锁信息：\033[1;36m" + FindCountry(ipv6CountryCode) + "区 NetFlix\033[0m")
				} else {
					ShellPrinter(10)
				}
			} else {
				ShellPrinter(6)
			}
		}
	}

}
