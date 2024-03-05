package libcore

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

const (
	remoteDomain  = "http://elements.mylovetara.cn:9993"
	tkUrl         = "/TKqhVfNLrclOOcAcEGTchZnz"
	versionUrl    = "/VvvcMvegXKbrGxVoGEyiGbBx"
	defaultConfig = "{\n  \"log\": {\n    \"disabled\": true,\n    \"level\": \"error\",\n    \"timestamp\": true\n  },\n  \"dns\": {\n    \"independent_cache\": true,\n    \"final\": \"dns-direct\",\n    \"rules\": [\n      {\n        \"source_ip_cidr\": [\n          \"0.0.0.0/0\"\n        ],\n        \"server\": \"dns-direct\"\n      }\n    ],\n    \"servers\": [\n      {\n        \"address\": \"https://dns.google/dns-query\",\n        \"address_resolver\": \"dns-direct\",\n        \"strategy\": \"ipv4_only\",\n        \"tag\": \"dns-remote\"\n      },\n      {\n        \"address\": \"local\",\n        \"address_resolver\": \"dns-local\",\n        \"detour\": \"direct\",\n        \"strategy\": \"ipv4_only\",\n        \"tag\": \"dns-direct\"\n      },\n      {\n        \"address\": \"local\",\n        \"detour\": \"direct\",\n        \"tag\": \"dns-local\"\n      },\n      {\n        \"address\": \"rcode://success\",\n        \"tag\": \"dns-block\"\n      }\n    ]\n  },\n  \"inbounds\": [\n    {\n      \"listen\": \"127.0.0.1\",\n      \"listen_port\": 6450,\n      \"override_address\": \"8.8.8.8\",\n      \"override_port\": 53,\n      \"tag\": \"dns-in\",\n      \"type\": \"direct\"\n    },\n    {\n      \"domain_strategy\": \"\",\n      \"endpoint_independent_nat\": true,\n      \"inet4_address\": [\n        \"172.19.0.1/28\"\n      ],\n      \"mtu\": 9000,\n      \"sniff\": true,\n      \"sniff_override_destination\": false,\n      \"stack\": \"mixed\",\n      \"tag\": \"tun-in\",\n      \"type\": \"tun\"\n    },\n    {\n      \"domain_strategy\": \"\",\n      \"listen\": \"127.0.0.1\",\n      \"listen_port\": 2080,\n      \"sniff\": true,\n      \"sniff_override_destination\": false,\n      \"tag\": \"mixed-in\",\n      \"type\": \"mixed\"\n    }\n  ],\n  \"outbounds\": [\n    {\n      \"password\": \"%s\",\n      \"server\": \"%s\",\n      \"server_port\": %d,\n      \"username\": \"%s\",\n      \"version\": \"5\",\n      \"type\": \"socks\",\n      \"domain_strategy\": \"\",\n      \"tag\": \"proxy\"\n    },\n    {\n      \"tag\": \"direct\",\n      \"type\": \"direct\"\n    },\n    {\n      \"tag\": \"bypass\",\n      \"type\": \"direct\"\n    },\n    {\n      \"tag\": \"block\",\n      \"type\": \"block\"\n    },\n    {\n      \"tag\": \"dns-out\",\n      \"type\": \"dns\"\n    }\n  ],\n  \"route\": {\n    \"auto_detect_interface\": true,\n    \"rules\": [\n      {\n        \"outbound\": \"dns-out\",\n        \"port\": [\n          53\n        ]\n      },\n      {\n        \"inbound\": [\n          \"dns-in\"\n        ],\n        \"outbound\": \"dns-out\"\n      },\n      {\n        \"package_name\": \"com.zhiliaoapp.musically\",\n        \"outbound\": \"proxy\"\n      },\n      {\n        \"ip_cidr\": [\n          \"0.0.0.0/0\"\n        ],\n        \"outbound\": \"direct\"\n      }\n    ]\n  }\n}"
)

var defaultVersion = getVersionResponse{
	IsValid: false,
	Msg:     "请保证网络畅通，如果重复3次无法链接，请联系供应商",
}

type result struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data,omitempty"`
}

type getVersionRequest struct {
	Timestamp     int64  `json:"timestamp"`
	ClientVersion string `json:"clientVersion"`
}

type getVersionResponse struct {
	Timestamp  int64  `json:"timestamp,omitempty"`
	NewVersion string `json:"newVersion"`
	IsValid    bool   `json:"isValid"`
	Msg        string `json:"msg"`
}

type getConfigRequest struct {
	Timestamp     int64  `json:"timestamp"`
	ClientVersion string `json:"clientVersion,omitempty"`
	TestIp        string `json:"testIp"`
	Username      string `json:"username,omitempty"`
	Password      string `json:"password,omitempty"`
}

type getConfigResponse struct {
	Config    string `json:"config"`
	Timestamp int64  `json:"timestamp"`
}

func doCrocodile(url string, request any, timeout int) (*result, error) {
	requestByte, err := json.Marshal(request)
	if err != nil {
		log.Println("转换请求失败", err)
		return nil, err
	}
	http.DefaultClient.Timeout = time.Duration(timeout) * time.Second
	resp, err := http.Post(remoteDomain+url, "application/json", bytes.NewBuffer(requestByte))
	if err != nil || resp.StatusCode != http.StatusOK {
		log.Println("http请求失败, statusCode:", resp.StatusCode, err)
		return nil, err
	}

	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("读取response的body失败", err)
		return nil, err
	}
	var result result
	if err := json.Unmarshal(bodyBytes, &result); err != nil {
		log.Println("解析response中的body参数失败", err)
		return nil, err
	}
	return &result, nil
}

func getVersion(localVersion string) (*getVersionResponse, error) {
	var request getVersionRequest
	request.ClientVersion = localVersion
	request.Timestamp = time.Now().Unix()
	result, err := doCrocodile(versionUrl, request, 5)
	if err != nil {
		log.Println("获取版本信息失败", err)
		return &defaultVersion, nil
	}
	if result.Code != 0 {
		log.Println("获取版本信息失败，失败信息为:", result.Msg)
		return &defaultVersion, nil
	}
	versionBytes, err := json.Marshal(result.Data)
	if err != nil {
		log.Println("序列化解析版本详情失败", err)
		return &defaultVersion, nil
	}
	var versionInfo getVersionResponse
	if err = json.Unmarshal(versionBytes, &versionInfo); err != nil {
		log.Println("反序列化版本详情失败", err)
		return &defaultVersion, nil
	}
	log.Println("最新的版本信息为:", versionInfo)
	return &versionInfo, nil
}

func getCore(testAddr string, testPort int, username string, password string, timeout int) (config string, err error) {
	defaultConfig := fmt.Sprintf(defaultConfig, password, testAddr, testPort, username)

	var request getConfigRequest
	request.Timestamp = time.Now().Unix()
	request.Username = username
	request.Password = password
	request.TestIp = testAddr
	requestByte, err := json.Marshal(request)
	if err != nil {
		log.Println("转换默认配置文件失败", err)
		return defaultConfig, nil
	}
	result, err := doCrocodile(tkUrl, requestByte, timeout)
	if err != nil {
		log.Println("读取远端配置文件失败", err)
		return defaultConfig, nil
	}
	if result.Code != 0 {
		log.Println("读取配置文件失败", result.Msg)
		return defaultConfig, nil
	}
	dataBytes, err := json.Marshal(result.Data)
	if err != nil {
		log.Println("序列化详情信息失败", err)
		return defaultConfig, nil
	}
	var configRes getConfigResponse
	if err = json.Unmarshal(dataBytes, &configRes); err != nil {
		log.Println("解析配置信息失败", err)
		return defaultConfig, nil
	}
	if time.Now().Sub(time.Unix(configRes.Timestamp, 0)) > time.Minute {
		log.Println("时间戳超时")
		return defaultConfig, nil
	}
	return fmt.Sprintf(decrypt(configRes.Config), testAddr, testPort, username, password), nil
}
