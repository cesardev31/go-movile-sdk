package internal

import (
	"fmt"
	"os"
	"os/exec"
)

func FyneBuild(mode, appID, appName string) error {

	if err := os.Chdir("cmd"); err != nil {
		return err
	}
	defer os.Chdir("..")

	args := []string{"package", "-os", "android", "--appID", appID, "--name", appName}
	if mode == "release" {
		args = append(args, "--release")
	}
	args = append(args, "-icon", "../assets/icon.png")

	cmd := exec.Command("fyne", args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error compilando con fyne: %v\nSalida: %s", err, output)
	}

	if _, err := os.Stat(appName + ".apk"); err == nil {
		if err := os.Rename(appName+".apk", "../"+appName+".apk"); err != nil {
			return fmt.Errorf("error moviendo apk: %v", err)
		}
	} else if !os.IsNotExist(err) {
		return err
	}

	fmt.Println("Compilación completada con éxito en modo", mode)
	return nil
}
