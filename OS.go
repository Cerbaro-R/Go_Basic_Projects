package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type cliente struct {
	id       int
	nome     string
	telefone string
}

type Servico struct {
	Descricao string
	Valor     float64
}

type os_struct struct {
	id         int
	descricao  string
	servicos   []Servico
	valorTotal float64
	status     string
	dono       cliente
}

// criarCliente solicita os dados do cliente e retorna um novo cliente
func criarCliente(id int) cliente {
	reader := bufio.NewReader(os.Stdin)
	var c cliente
	c.id = id
	fmt.Println("\n--- Cadastro de Cliente ---")
	fmt.Print("Nome: ")
	nome, _ := reader.ReadString('\n')
	c.nome = strings.TrimSpace(nome)

	fmt.Print("Telefone: ")
	telefone, _ := reader.ReadString('\n')
	c.telefone = strings.TrimSpace(telefone)
	return c
}

// listarClientes exibe todos os clientes cadastrados
func listarClientes(clientes []cliente) {
	fmt.Println("\n--- Lista de Clientes ---")
	if len(clientes) == 0 {
		fmt.Println("Nenhum cliente cadastrado.")
		return
	}
	for _, c := range clientes {
		fmt.Printf("ID: %d, Nome: %s, Telefone: %s\n", c.id, c.nome, c.telefone)
	}
	fmt.Println("-------------------------")
}

// criarOS solicita os dados da Ordem de Serviço
func criarOS(id int, c cliente) os_struct {
	reader := bufio.NewReader(os.Stdin)
	var novaOS os_struct
	novaOS.id = id
	novaOS.dono = c
	novaOS.status = "Aberto"
	novaOS.valorTotal = 0.0

	fmt.Println("\n--- Cadastro de Ordem de Serviço ---")
	fmt.Print("Descrição geral da OS: ")
	descricao, _ := reader.ReadString('\n')
	novaOS.descricao = strings.TrimSpace(descricao)

	fmt.Println("Ordem de Serviço criada com sucesso! Adicione serviços a ela.")
	return novaOS
}

// recalcularValorTotal atualiza o valor total da OS somando os serviços
func (o *os_struct) recalcularValorTotal() {
	var total float64
	for _, s := range o.servicos {
		total += s.Valor
	}
	o.valorTotal = total
}

// adicionarServicoOS adiciona um novo serviço a uma OS existente
func adicionarServicoOS(o *os_struct) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("\n--- Adicionar Serviço ---")
	fmt.Printf("Adicionando serviço à OS %d: %s\n", o.id, o.descricao)

	fmt.Print("Descrição do serviço: ")
	desc, _ := reader.ReadString('\n')

	fmt.Print("Valor do serviço: ")
	valorStr, _ := reader.ReadString('\n')
	valor, err := strconv.ParseFloat(strings.TrimSpace(valorStr), 64)
	if err != nil {
		fmt.Println("Valor inválido. O serviço não foi adicionado.")
		return
	}

	novoServico := Servico{
		Descricao: strings.TrimSpace(desc),
		Valor:     valor,
	}

	o.servicos = append(o.servicos, novoServico)
	o.recalcularValorTotal()
	fmt.Println("Serviço adicionado com sucesso!")
}

// finalizarOS altera o status de uma OS para "Finalizada" com verificação
func finalizarOS(o *os_struct) {
	reader := bufio.NewReader(os.Stdin)
	var confirmacao string

	if len(o.servicos) == 0 {
		fmt.Printf("Atenção: A OS %d não possui serviços (Valor Total: R$ 0.00). Deseja realmente finalizar? (s/n): ", o.id)
	} else {
		fmt.Printf("Você realmente deseja finalizar a OS %d (%s)? (s/n): ", o.id, o.descricao)
	}

	confirmacao, _ = reader.ReadString('\n')
	if strings.TrimSpace(strings.ToLower(confirmacao)) == "s" {
		o.status = "Finalizada"
		fmt.Println("OS finalizada com sucesso!")
	} else {
		fmt.Println("Finalização da OS cancelada.")
	}
}

// listarOSs exibe todas as ordens de serviço cadastradas
func listarOSs(ordens []os_struct) {
	fmt.Println("\n--- Lista de Ordens de Serviço ---")
	if len(ordens) == 0 {
		fmt.Println("Nenhuma OS cadastrada.")
		return
	}
	for _, o := range ordens {
		fmt.Println("==============================")
		fmt.Printf("ID: %d\n", o.id)
		fmt.Printf("Cliente: %s (ID: %d)\n", o.dono.nome, o.dono.id)
		fmt.Printf("Descrição Geral: %s\n", o.descricao)
		fmt.Printf("Status: %s\n", o.status)
		fmt.Println("--- Serviços ---")
		if len(o.servicos) == 0 {
			fmt.Println("Nenhum serviço adicionado.")
		} else {
			for i, s := range o.servicos {
				fmt.Printf("%d. %s - R$ %.2f\n", i+1, s.Descricao, s.Valor)
			}
		}
		fmt.Println("----------------")
		fmt.Printf("Valor Total: R$ %.2f\n", o.valorTotal)
		fmt.Println("==============================")
	}
}

func main() {
	var clientes []cliente
	var ordensDeServico []os_struct
	var escolha string
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n--- Sistema de OS ---")
		fmt.Println("1. Cadastrar Cliente")
		fmt.Println("2. Listar Clientes")
		fmt.Println("3. Criar Ordem de Serviço")
		fmt.Println("4. Listar Ordens de Serviço")
		fmt.Println("5. Adicionar Serviço a OS")
		fmt.Println("6. Finalizar Ordem de Serviço")
		fmt.Println("7. Sair")
		fmt.Print("Escolha uma opção: ")
		fmt.Scanln(&escolha)

		switch escolha {
		case "1":
			novoCliente := criarCliente(len(clientes) + 1)
			clientes = append(clientes, novoCliente)
			fmt.Println("Cliente cadastrado com sucesso!")
		case "2":
			listarClientes(clientes)
		case "3":
			var clienteSelecionado cliente
			var clienteEncontrado bool

			if len(clientes) == 0 {
				fmt.Println("Nenhum cliente cadastrado. Vamos cadastrar um novo.")
				clienteSelecionado = criarCliente(len(clientes) + 1)
				clientes = append(clientes, clienteSelecionado)
				clienteEncontrado = true
			} else {
				listarClientes(clientes)
				fmt.Print("Digite o ID do cliente ou '0' para cadastrar um novo: ")
				input, _ := reader.ReadString('\n')
				id, _ := strconv.Atoi(strings.TrimSpace(input))

				if id == 0 {
					clienteSelecionado = criarCliente(len(clientes) + 1)
					clientes = append(clientes, clienteSelecionado)
					clienteEncontrado = true
				} else {
					for _, c := range clientes {
						if c.id == id {
							clienteSelecionado = c
							clienteEncontrado = true
							break
						}
					}
				}
			}

			if clienteEncontrado {
				novaOS := criarOS(len(ordensDeServico)+1, clienteSelecionado)
				ordensDeServico = append(ordensDeServico, novaOS)
			} else {
				fmt.Println("Cliente não encontrado.")
			}
		case "4":
			listarOSs(ordensDeServico)
		case "5":
			listarOSs(ordensDeServico)
			if len(ordensDeServico) == 0 {
				continue
			}
			fmt.Print("Digite o ID da OS para adicionar um serviço: ")
			idStr, _ := reader.ReadString('\n')
			id, err := strconv.Atoi(strings.TrimSpace(idStr))
			if err != nil {
				fmt.Println("ID inválido.")
				continue
			}

			var osEncontrada *os_struct
			for i := range ordensDeServico {
				if ordensDeServico[i].id == id {
					osEncontrada = &ordensDeServico[i]
					break
				}
			}

			if osEncontrada != nil {
				if osEncontrada.status == "Finalizada" {
					fmt.Println("Não é possível adicionar serviços a uma OS já finalizada.")
				} else {
					adicionarServicoOS(osEncontrada)
				}
			} else {
				fmt.Println("OS não encontrada.")
			}
		case "6":
			listarOSs(ordensDeServico)
			if len(ordensDeServico) == 0 {
				continue
			}
			fmt.Print("Digite o ID da OS que deseja finalizar: ")
			idStr, _ := reader.ReadString('\n')
			id, err := strconv.Atoi(strings.TrimSpace(idStr))
			if err != nil {
				fmt.Println("ID inválido.")
				continue
			}

			var osEncontrada *os_struct
			for i := range ordensDeServico {
				if ordensDeServico[i].id == id {
					osEncontrada = &ordensDeServico[i]
					break
				}
			}

			if osEncontrada != nil {
				if osEncontrada.status == "Finalizada" {
					fmt.Println("Esta OS já foi finalizada.")
				} else {
					finalizarOS(osEncontrada)
				}
			} else {
				fmt.Println("OS não encontrada.")
			}
		case "7":
			fmt.Println("Saindo do sistema...")
			return
		default:
			fmt.Println("Opção inválida. Tente novamente.")
		}
	}
}
