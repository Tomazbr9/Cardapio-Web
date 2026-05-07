# Requisitos — Sistema Web de Cardápio para Pizzarias

## 1. Requisitos Funcionais (RF)

### 1.1 Para o Cliente (Cardápio Web / PWA)

#### RF-01 — Visualização de Cardápio
O sistema deve permitir a visualização dos produtos organizados por categorias, como:
- Pizzas
- Bebidas
- Esfihas
- Sobremesas
- Combos

---

#### RF-02 — Montagem de Pizza
O sistema deve permitir:
- Escolha de até 3 sabores por pizza
- Seleção do tamanho:
  - Pequena (P)
  - Média (M)
  - Grande (G)
  - Gigante (GG)
- Adição de bordas recheadas
- Inclusão de observações adicionais

---

#### RF-03 — Carrinho de Compras
O sistema deve possuir:
- Adição de produtos ao carrinho
- Remoção de produtos
- Edição de itens
- Atualização automática do valor total

---

#### RF-04 — Cálculo de Entrega
O sistema deve:
- Permitir seleção do bairro do cliente
- Calcular automaticamente a taxa de entrega
- Exibir o valor antes da finalização do pedido

---

#### RF-05 — Finalização via WhatsApp
Ao finalizar o pedido, o sistema deve:
- Gerar um texto formatado com:
  - Produtos
  - Quantidades
  - Valores
  - Endereço
  - Forma de pagamento
- Abrir automaticamente o WhatsApp da pizzaria
- Enviar os dados do pedido através de link pré-formatado

---

#### RF-06 — Controle de Funcionamento
O sistema deve:
- Permitir configuração de horário de funcionamento
- Bloquear novos pedidos quando a pizzaria estiver fechada
- Exibir mensagem informando indisponibilidade

---

## 1.2 Para o Dono da Pizzaria (Painel Administrativo)

### RF-07 — Gestão de Catálogo
O sistema deve possuir CRUD para:
- Categorias
- Produtos
- Sabores
- Tamanhos
- Bordas recheadas

---

### RF-08 — Configuração de Regras de Preço
O sistema deve permitir definir regras para pizzas com múltiplos sabores:
- Cobrança pelo sabor mais caro
- Cobrança pela média dos sabores

---

### RF-09 — Gestão de Taxas de Entrega
O sistema deve permitir:
- Cadastro de bairros
- Definição de taxa de entrega por bairro
- Edição e remoção de taxas

---

### RF-10 — Dashboard de Pedidos
O painel administrativo deve permitir:
- Visualizar pedidos em tempo real
- Alterar status do pedido:
  - Em preparo
  - Saiu para entrega
  - Finalizado
  - Cancelado

---

### RF-11 — Configurações da Loja
O sistema deve permitir:
- Upload de logo
- Configuração de cores do cardápio
- Cadastro de endereço
- Configuração do número/link do WhatsApp
- Configuração de horários de funcionamento

---

# 2. Requisitos Não Funcionais (RNF)

## RNF-01 — Multi-tenancy (Isolamento de Dados)
O sistema deve suportar múltiplas pizzarias na mesma infraestrutura, garantindo:
- Separação total dos dados
- Segurança entre tenants
- Isolamento lógico por cliente

---

## RNF-02 — Mobile First (PWA)
O cardápio deve:
- Ser otimizado para dispositivos móveis
- Funcionar como Progressive Web App (PWA)
- Possuir carregamento rápido em redes 3G/4G

---

## RNF-03 — Performance
O sistema deve:
- Possuir tempo inicial de carregamento inferior a 2 segundos
- Utilizar cache para otimização de recursos estáticos
- Minimizar requisições desnecessárias

---

## RNF-04 — Escalabilidade Horizontal
O backend deve:
- Permitir múltiplas instâncias simultâneas
- Escalar automaticamente em horários de pico
- Suportar aumento de acessos nas noites de sexta-feira e domingo

---

## RNF-05 — Persistência de Dados Relacional
O sistema deve utilizar:
- PostgreSQL como banco principal
- Estrutura relacional consistente
- Integridade transacional para pedidos e pagamentos

---

## RNF-06 — Disponibilidade
O sistema deve buscar:
- Disponibilidade mínima de 99,5%
- Alta estabilidade entre 18h e 00h
- Estratégias de recuperação em caso de falhas
