// internal/emulator.go
package internal

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
	"time"
)

func GetFirstAVD() (string, error) {
	cmd := exec.Command("emulator", "-list-avds")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" && !strings.HasPrefix(line, "INFO") {
			return line, nil
		}
	}
	return "", fmt.Errorf("no se encontró ningún AVD")
}

func StartEmulatorIfNeeded(avdName string) error {
	cmd := exec.Command("adb", "shell", "getprop", "sys.boot_completed")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	if strings.Contains(string(output), "1") {
		fmt.Println("El emulador ya está corriendo.")
		return nil
	}
	fmt.Printf("Iniciando emulador con AVD: %s...\n", avdName)
	emuCmd := exec.Command("emulator", "-avd", avdName)
	if err := emuCmd.Start(); err != nil {
		return fmt.Errorf("no se pudo iniciar el emulador: %w", err)
	}
	time.Sleep(30 * time.Second)
	return nil
}

func InstallAndLaunch(appID, apkName string) error {
	installCmd := exec.Command("adb", "install", "-r", apkName)
	output, err := installCmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error instalando apk: %v\nSalida: %s", err, output)
	}
	// Lanzar
	launchCmd := exec.Command("adb", "shell", "am", "start", "-n", fmt.Sprintf("%s/org.golang.app.GoNativeActivity", appID))
	output, err = launchCmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error lanzando app: %v\nSalida: %s", err, output)
	}
	return nil
}

// ShowLogs muestra los logs con adb logcat
func ShowLogs() {
	fmt.Println("Mostrando logs. Presiona Ctrl+C para salir...")
	cmd := exec.Command("adb", "logcat")
	cmd.Stdout = bytes.NewBuffer(nil)
	cmd.Stderr = bytes.NewBuffer(nil)
	cmd.Stdout = NewPassthroughWriter()
	cmd.Stderr = NewPassthroughWriter()
	_ = cmd.Run()
}

func NewPassthroughWriter() *passthroughWriter {
	return &passthroughWriter{}
}

type passthroughWriter struct{}

func (p *passthroughWriter) Write(b []byte) (int, error) {
	return fmt.Print(string(b))
}
