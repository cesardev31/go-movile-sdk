package internal

import (
	"fmt"
	"log"
	"os/exec"
)

func BuildDebug() {
	fmt.Println("Compilando APK en modo debug...")
	cmd := exec.Command("fyne", "package", "-os", "android", "-icon", "assets/icon.png")
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Error compilando debug: %v\nSalida: %s", err, output)
	}
	fmt.Println("APK Debug generado con éxito.")
}

func BuildRelease() {
	fmt.Println("Compilando APK en modo release (producción)...")
	cmd := exec.Command("fyne", "package", "-os", "android", "-release",
		"-appID", "com.miempresa.miapp",
		"-icon", "assets/icon.png",
		"-keystore", "path/to/keystore.jks",
		"-alias", "mi_alias",
		"-storepass", "mi_storepass",
		"-keypass", "mi_keypass")
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Error compilando release: %v\nSalida: %s", err, output)
	}
	fmt.Println("APK Release generado con éxito.")
}

func RunTests() {
	fmt.Println("Ejecutando tests...")
	cmd := exec.Command("go", "test", "./...")
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Error en tests: %v\nSalida: %s", err, output)
	}
	fmt.Println("Tests ejecutados correctamente:\n", string(output))
}
