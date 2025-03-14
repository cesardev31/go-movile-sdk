package main

import (
	"fmt"
	"log"
	"os"

	"github.com/cesardev31/go-movile-sdk/internal"
)

func main() {
	if len(os.Args) < 2 {
		printHelp()
		os.Exit(1)
	}

	cmd := os.Args[1]

	switch cmd {
	case "create":
		if len(os.Args) < 5 || os.Args[2] != "app" {
			fmt.Println("Uso: gosdk create app <app-name> <module-name>")
			os.Exit(1)
		}
		appName := os.Args[3]
		moduleName := os.Args[4]
		if err := internal.CreateNewProject(appName, moduleName); err != nil {
			log.Fatalf("Error creando el proyecto: %v", err)
		}
		fmt.Printf("Proyecto '%s' creado exitosamente.\n", appName)
	case "run":
		if len(os.Args) < 3 {
			fmt.Println("Falta el modo: debug | release")
			os.Exit(1)
		}
		mode := os.Args[2]
		runApp(mode)
	case "build":
		if len(os.Args) < 3 {
			fmt.Println("Falta el modo: debug | release")
			os.Exit(1)
		}
		mode := os.Args[2]
		buildApp(mode)
	case "test":
		internal.RunTests()
	default:
		fmt.Println("Comando no reconocido:", cmd)
		printHelp()
		os.Exit(1)
	}
}

func runApp(mode string) {
	// 1. Iniciar emulador si no está corriendo
	avdName, err := internal.GetFirstAVD()
	if err != nil {
		log.Fatalf("Error obteniendo AVD: %v", err)
	}
	err = internal.StartEmulatorIfNeeded(avdName)
	if err != nil {
		log.Fatalf("Error iniciando emulador: %v", err)
	}

	// 2. Compilar la app (modo debug o release)
	switch mode {
	case "debug":
		err = internal.FyneBuild("debug", "com.miempresa.miapp", "miapp")
	case "release":
		err = internal.FyneBuild("release", "com.miempresa.miapp", "miapp")
	default:
		log.Fatalf("Modo no soportado: %s", mode)
	}
	if err != nil {
		log.Fatalf("Error compilando APK: %v", err)
	}

	// 3. Instalar y lanzar la app
	err = internal.InstallAndLaunch("com.miempresa.miapp", "miapp.apk")
	if err != nil {
		log.Fatalf("Error instalando/lanzando la app: %v", err)
	}

	// 4. (Opcional) Ver logs
	internal.ShowLogs()
}

func buildApp(mode string) {
}

func printHelp() {
	fmt.Println("Uso: gosdk [create|run|build|test] <opciones>")
	fmt.Println("")
	fmt.Println("Comandos:")
	fmt.Println("  create app <app-name> <module-name>   Crea un nuevo proyecto con la estructura base")
	fmt.Println("  run debug     Inicia el emulador, compila en modo debug, instala y lanza la app")
	fmt.Println("  run release   Igual que run debug pero en modo release")
	fmt.Println("  build debug   Compila en modo debug y genera el APK (sin instalar)")
	fmt.Println("  build release Compila en modo release y genera el APK (sin instalar)")
	fmt.Println("  test          Ejecuta los tests")
}
