package main

import (
	"fmt"
	"os/exec"
)

func SecurityTest() {
	// 1. Simulação didática de exposição de credenciais (Secret Leak)
	// O Semgrep possui regras genéricas que detectam o prefixo padrão da AWS (AKIA)
	awsSecretFake := "AKIAIMNOJVGFD5E4M32A1B2C3D4E5F6G7H8I9J0K"
	fmt.Println("Autenticando com:", awsSecretFake)

	// 2. Simulação de injeção de comandos do SO (Command Injection)
	// O Semgrep detecta o uso do pacote 'os/exec' executando shell scripts ('sh', '-c')
	entradaPerigosa := "ls -la" 
	cmd := exec.Command("sh", "-c", entradaPerigosa)
	cmd.Run()
}
