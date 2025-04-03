package license

import (
	"encoding/hex"
	"net"
	"os/exec"
	"runtime"
	"strings"

	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"time"
)

type License struct {
	MachineCode string `json:"machine_code"`
	ExpireAt    int64  `json:"expire_at"`
	Signature   string `json:"signature"`
}

var secretKey = "my-secret-key" // License 签名密钥

// 生成 License 签名
func generateSignature(machineCode string, expireAt int64) string {
	data := fmt.Sprintf("%s:%d", machineCode, expireAt)
	h := hmac.New(sha256.New, []byte(secretKey))
	h.Write([]byte(data))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

// 生成 License
func generateLicense(machineCode string, days int) string {
	expireAt := time.Now().Add(time.Duration(days) * 24 * time.Hour).Unix()
	signature := generateSignature(machineCode, expireAt)

	license := License{
		MachineCode: machineCode,
		ExpireAt:    expireAt,
		Signature:   signature,
	}

	licenseBytes, _ := json.Marshal(license)
	return base64.StdEncoding.EncodeToString(licenseBytes)
}

func saveLicenseToFile(license string) error {
	return ioutil.WriteFile("license.txt", []byte(license), 0644)
}

// 读取 License
func loadLicenseFromFile() (string, error) {
	data, err := ioutil.ReadFile("license.txt")
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func parseLicense(encodedLicense string) (*License, error) {
	licenseBytes, err := base64.StdEncoding.DecodeString(encodedLicense)
	if err != nil {
		return nil, err
	}

	var license License
	if err := json.Unmarshal(licenseBytes, &license); err != nil {
		return nil, err
	}

	// 验证签名
	expectedSignature := generateSignature(license.MachineCode, license.ExpireAt)
	if license.Signature != expectedSignature {
		return nil, errors.New("invalid signature")
	}

	// 检查是否过期
	if time.Now().Unix() > license.ExpireAt {
		return nil, errors.New("license expired")
	}

	return &license, nil
}

// 记录上次运行时间
type LastRun struct {
	Timestamp int64 `json:"timestamp"`
}

// 保存上次运行时间
func saveLastRunTime() error {
	data, _ := json.Marshal(LastRun{Timestamp: time.Now().Unix()})
	return ioutil.WriteFile("last_run.json", data, 0644)
}

// 读取上次运行时间
func loadLastRunTime() (int64, error) {
	data, err := ioutil.ReadFile("last_run.json")
	if err != nil {
		return 0, err
	}

	var lastRun LastRun
	json.Unmarshal(data, &lastRun)
	return lastRun.Timestamp, nil
}

// 检测系统时间是否被回退
func checkTimeValidity() bool {
	lastRun, err := loadLastRunTime()
	if err == nil && time.Now().Unix() < lastRun {
		return false // 系统时间被回退，拒绝启动
	}
	saveLastRunTime()
	return true
}

// 获取 MAC 地址
func getMacAddress() string {
	interfaces, err := net.Interfaces()
	if err != nil {
		return "unknown"
	}
	for _, iface := range interfaces {
		if len(iface.HardwareAddr) > 0 {
			return iface.HardwareAddr.String()
		}
	}
	return "unknown"
}

// 获取 CPU ID
func getCpuID() string {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("wmic", "cpu", "get", "ProcessorId")
	case "linux":
		cmd = exec.Command("cat", "/proc/cpuinfo")
	case "darwin":
		cmd = exec.Command("sysctl", "-n", "machdep.cpu.brand_string")
	default:
		return "unknown"
	}
	output, err := cmd.Output()
	if err != nil {
		return "unknown"
	}
	return strings.TrimSpace(string(output))
}

// 生成机器码
func generateMachineCode() string {
	data := getMacAddress() + getCpuID()
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}