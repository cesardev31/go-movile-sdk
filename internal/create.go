package internal

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func CreateNewProject(appName, moduleName string) error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	projectDir := filepath.Join(cwd, appName)

	dirs := []string{
		projectDir,
		filepath.Join(projectDir, "build"),
		filepath.Join(projectDir, "cmd"),
		filepath.Join(projectDir, "internal"),
	}
	for _, d := range dirs {
		if err := os.MkdirAll(d, 0755); err != nil {
			return fmt.Errorf("error creando directorio %s: %w", d, err)
		}
	}

	cmd := exec.Command("go", "mod", "init", moduleName)
	cmd.Dir = projectDir
	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("error creando go.mod: %v\nSalida: %s", err, string(output))
	}

	cmd = exec.Command("git", "init")
	cmd.Dir = projectDir
	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("error inicializando git: %v\nSalida: %s", err, string(output))
	}

	readmeContent := fmt.Sprintf("# %s\n\nInstrucciones para desarrollar la app usando el SDK.\n", appName)
	if err := os.WriteFile(filepath.Join(projectDir, "README.md"), []byte(readmeContent), 0644); err != nil {
		return fmt.Errorf("error creando README.md: %w", err)
	}

	configContent := `[App]
		Name = "` + appName + `"
		Version = "1.0.0"
		Debug = true
		
		[Server]
		Host = "api.miapp.com"
		Port = 443
		
		[Auth]
		TokenExpiry = 3600
		SecretKey = "supersecreto"
		
		[UI]
		Theme = "dark"
		`
	if err := os.WriteFile(filepath.Join(projectDir, "config.toml"), []byte(configContent), 0644); err != nil {
		return fmt.Errorf("error creando config.toml: %w", err)
	}

	// Crear .gitignore con contenido recomendado
	gitignoreContent := `# Binarios y artefactos
*.exe
*.dll
*.so
*.dylib
*.test
*.apk
*.aab
build/

# Dependencias vendorizadas
vendor/

# Archivos de sistema
.DS_Store
Thumbs.db

# Configuraciones de IDEs
.idea/
.vscode/

# Archivos temporales
*.log
tmp/
`
	if err := os.WriteFile(filepath.Join(projectDir, ".gitignore"), []byte(gitignoreContent), 0644); err != nil {
		return fmt.Errorf("error creando .gitignore: %w", err)
	}

	// Crear un archivo de ejemplo en cmd/main.go
	mainContent := `package main

import (
	"fmt"
)

func main() {
	fmt.Println("¡Hola, bienvenido a ` + appName + `!")
}
`
	if err := os.WriteFile(filepath.Join(projectDir, "cmd", "main.go"), []byte(mainContent), 0644); err != nil {
		return fmt.Errorf("error creando cmd/main.go: %w", err)
	}

	// Crear un archivo de ejemplo en internal/commands.go
	internalContent := `package internal

import "fmt"

// Ejemplo de función interna
func Hello() {
	fmt.Println("Hola desde internal!")
}
`
	if err := os.WriteFile(filepath.Join(projectDir, "internal", "commands.go"), []byte(internalContent), 0644); err != nil {
		return fmt.Errorf("error creando internal/commands.go: %w", err)
	}

	return nil
}
