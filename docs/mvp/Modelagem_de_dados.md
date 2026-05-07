# Estrutura de Banco de Dados — Sistema Web de Cardápio para Pizzarias

# 1. Núcleo do Multi-tenant

A tabela `tenants` é a raiz principal da aplicação.  
Cada pizzaria possuirá um identificador único (`UUID`) para garantir isolamento de dados entre clientes.

## Tabela: tenants

```sql
CREATE TABLE tenants (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    slug VARCHAR(255) UNIQUE NOT NULL, -- Ex: 'pizzaria-do-ze'
    whatsapp_number VARCHAR(20),
    config_json JSONB, -- Cores, logo, horário de funcionamento
    is_active BOOLEAN
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

## Responsabilidades da tabela
- Cadastro das pizzarias
- Configurações visuais do cardápio
- Controle multi-tenant
- Integração principal com WhatsApp
- Configurações gerais da loja

---

# 2. Categorias e Produtos Gerais

Essa estrutura é utilizada para produtos simples e categorias do cardápio, como:
- Refrigerantes
- Sobremesas
- Esfihas
- Combos

---

## Tabela: categories

```sql
CREATE TABLE categories (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tenant_id UUID REFERENCES tenants(id),
    name VARCHAR(100) NOT NULL,
    position INT DEFAULT 0, -- Para ordenar no menu
    is_active BOOLEAN DEFAULT true
);
```

## Responsabilidades
- Organização visual do cardápio
- Ordenação de categorias
- Ativação/desativação de categorias

---

## Tabela: products

```sql
CREATE TABLE products (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tenant_id UUID REFERENCES tenants(id),
    category_id UUID REFERENCES categories(id),
    name VARCHAR(255) NOT NULL,
    description TEXT,
    base_price DECIMAL(10,2) NOT NULL,
    image_url VARCHAR(255),
    is_pizza BOOLEAN DEFAULT false, -- Define se segue a lógica de sabores
    is_active BOOLEAN DEFAULT true
);
```

## Responsabilidades
- Cadastro geral de produtos
- Produtos simples e pizzas
- Controle de disponibilidade
- Associação com categorias

---

# 3. A Lógica Complexa: Pizzas

As pizzas possuem regras específicas:
- Variação de preço por tamanho
- Escolha de múltiplos sabores
- Controle de sabores premium
- Regras de meio-a-meio

---

# 3.1 Tamanhos de Pizza

## Tabela: pizza_sizes

```sql
CREATE TABLE pizza_sizes (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tenant_id UUID REFERENCES tenants(id),
    name VARCHAR(50) NOT NULL, -- Ex: 'Grande'
    slices INT, -- Ex: 8
    max_flavors INT DEFAULT 2 -- Limite de sabores para esse tamanho
);
```

## Responsabilidades
- Definição de tamanhos
- Controle de quantidade de sabores
- Quantidade de fatias

---

# 3.2 Sabores de Pizza

## Tabela: pizza_flavors

```sql
CREATE TABLE pizza_flavors (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tenant_id UUID REFERENCES tenants(id),
    name VARCHAR(100) NOT NULL,
    description TEXT,
    price_modifier DECIMAL(10,2) DEFAULT 0 -- Valor extra se o sabor for premium
);
```

## Responsabilidades
- Cadastro de sabores
- Definição de sabores premium
- Controle de descrição dos sabores

---

# 3.3 Relacionamento Preço x Tamanho x Sabor

Essa tabela é essencial para calcular corretamente:
- Pizza meio-a-meio
- Pizza com 3 sabores
- Diferenças de preço entre tamanhos

---

## Tabela: pizza_flavor_prices

```sql
CREATE TABLE pizza_flavor_prices (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    flavor_id UUID REFERENCES pizza_flavors(id),
    size_id UUID REFERENCES pizza_sizes(id),
    price DECIMAL(10,2) NOT NULL
);
```

## Responsabilidades
- Precificação dinâmica
- Controle de preço por tamanho
- Regras de cálculo para múltiplos sabores

---

# 4. Pedidos e Itens

Estrutura responsável pelo registro completo dos pedidos realizados pelos clientes.

---

# 4.1 Pedidos

## Tabela: orders

```sql
CREATE TABLE orders (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tenant_id UUID REFERENCES tenants(id),
    customer_name VARCHAR(255) NOT NULL,
    customer_whatsapp VARCHAR(20) NOT NULL,
    total_amount DECIMAL(10,2) NOT NULL,
    delivery_fee DECIMAL(10,2) DEFAULT 0,
    status VARCHAR(50) DEFAULT 'PENDING', -- PENDING, PREPARING, DISPATCHED, DELIVERED
    payment_method VARCHAR(50),
    address_json JSONB, -- Rua, número, bairro
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

## Responsabilidades
- Registro principal do pedido
- Controle de status
- Armazenamento de endereço
- Dados do cliente
- Controle financeiro

---

# 4.2 Itens do Pedido

## Tabela: order_items

```sql
CREATE TABLE order_items (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    order_id UUID REFERENCES orders(id),
    product_id UUID REFERENCES products(id),
    size_id UUID REFERENCES pizza_sizes(id), -- NULL se não for pizza
    quantity INT NOT NULL,
    subtotal DECIMAL(10,2) NOT NULL,
    notes TEXT
);
```

## Responsabilidades
- Produtos pertencentes ao pedido
- Quantidade de itens
- Subtotal individual
- Observações do cliente

---

# 4.3 Sabores Escolhidos do Pedido

Essa tabela registra os sabores selecionados em pizzas meio-a-meio ou com múltiplos sabores.

---

## Tabela: order_item_flavors

```sql
CREATE TABLE order_item_flavors (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    order_item_id UUID REFERENCES order_items(id),
    flavor_id UUID REFERENCES pizza_flavors(id)
);
```

## Responsabilidades
- Relacionar sabores ao item do pedido
- Permitir pizzas com múltiplos sabores
- Base para cálculo final da pizza

---

## Tabela: bot_messages

```sql
CREATE TABLE bot_messages (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tenant_id UUID REFERENCES tenants(id),
    trigger_type VARCHAR(50), -- 'WELCOME', 'ORDER_CONFIRMED', 'ORDER_DISPATCHED'
    message_text TEXT NOT NULL -- Ex: "Olá! Veja nosso cardápio aqui: {link_cardapio}"
);
```

# Considerações Arquiteturais

## Banco de Dados
- PostgreSQL
- Uso de UUID para segurança e escalabilidade
- JSONB para configurações flexíveis

---

## Estratégia Multi-tenant

Todas as tabelas principais utilizam:
```sql
tenant_id UUID REFERENCES tenants(id)
```

Isso garante:
- Isolamento de dados
- Segurança entre pizzarias
- Escalabilidade horizontal
- Facilidade de manutenção

---

# Benefícios da Estrutura

## Escalabilidade
- Suporte para múltiplas pizzarias
- Fácil expansão horizontal

## Flexibilidade
- Configurações independentes por tenant
- Regras de preço customizáveis

## Performance
- Estrutura otimizada para consultas rápidas
- Separação lógica eficiente

## Manutenção
- Fácil evolução do sistema
- Modularização da lógica de negócio

