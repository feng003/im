package license

import (
	"os"
	"fmt"
	"testing"
)

func TestLicense(t *testing.T) {
	machineCode := generateMachineCode()
	fmt.Println("Machine Code:", machineCode)

	// 读取 License
	licenseStr, err := loadLicenseFromFile()
	if err != nil {
		fmt.Println("No valid license found. Generating a new one...")
		licenseStr = generateLicense(machineCode, 30) // 生成 30 天 License
		saveLicenseToFile(licenseStr)
	}

	// 验证 License
	license, err := parseLicense(licenseStr)
	if err != nil {
		fmt.Println("License invalid:", err)
		fmt.Println("System exiting...")
		os.Exit(1)
	}

	fmt.Printf("license is %+v \n", license)
	// 检测是否篡改时间
	if !checkTimeValidity() {
		fmt.Println("System time tampered. Exiting...")
		os.Exit(1)
	}

	fmt.Println("License is valid. System running...")	
}