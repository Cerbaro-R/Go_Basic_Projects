package main

import "fmt"

// 1. Definição das Structs (A "planta" do objeto)
type Atributos struct {
	HP    int
	Nivel int
}

type Personagem struct {
	Nome   string
	Classe string
	Status Atributos // Composição: Atributos dentro de Personagem
}

// 2. Método com Pointer Receiver (*)
// Precisamos do "*" porque queremos ALTERAR o HP do personagem original.
func (p *Personagem) ReceberDano(quantidade int) {
	fmt.Printf("💥 %s recebeu %d de dano!\n", p.Nome, quantidade)
	p.Status.HP -= quantidade

	if p.Status.HP <= 0 {
		p.Status.HP = 0
		fmt.Printf("💀 %s foi derrotado...\n", p.Nome)
	}
}

// 3. Método com Value Receiver (Sem o "*")
// Aqui o Go cria uma CÓPIA do personagem. O original não muda.
func (p Personagem) MostrarStatus() {
	fmt.Printf("--- STATUS ATUAL ---\n")
	fmt.Printf("Nome: %s | HP: %d | Nivel: %d\n", p.Nome, p.Status.HP, p.Status.Nivel)
	fmt.Printf("--------------------\n\n")
}

func main() {
	// Inicializando o personagem
	heroi := Personagem{
		Nome:   "Aragorn",
		Classe: "Ranger",
		Status: Atributos{
			HP:    100,
			Nivel: 1,
		},
	}

	// Mostrando status inicial
	heroi.MostrarStatus()

	// Aplicando dano (Isso altera o objeto original via ponteiro)
	heroi.ReceberDano(30)
	heroi.ReceberDano(20)

	// Mostrando status após o combate
	heroi.MostrarStatus()

	// Exemplo do que NÃO acontece sem ponteiro:
	// Se ReceberDano não tivesse o *, o HP ainda seria 100 aqui.
}
