# Fluxograma Operacional do Sistema — Sistema Web de Cardápio para Pizzarias

![Arquitetura MVP](./images/arquitetura_mvp.png)

Este é o momento de dar "vida" à arquitetura. O fluxograma (ou fluxo de processos) documenta a jornada do usuário e como os componentes do sistema interagem para realizar uma tarefa.

O fluxo operacional da plataforma é dividido em três macroprocessos:

1. Jornada do Cliente (Do "Oi" à Entrega)
2. Fluxo de Gestão da Pizzaria (Configuração e Pedidos)
3. Orquestração Técnica (Backend e Automação)

---

# 1. Jornada do Cliente (Do "Oi" à Entrega)

Este fluxo descreve a experiência do consumidor final.

| Passo | Ator | Ação | Componente Técnico |
| :--- | :--- | :--- | :--- |
| **1.1** | Cliente | Envia mensagem "Oi" (ou qualquer texto) no WhatsApp da Pizzaria. | WhatsApp -> Evolution API |
| **1.2** | Sistema | Identifica Tenant, busca template de saudação e envia link do cardápio. | Go API -> Postgres -> Redis (Queue) |
| **1.3** | Cliente | Recebe a mensagem e clica no link: `meusaas.com/pizzaria-do-ze`. | WhatsApp -> Navegador |
| **1.4** | Cliente | Navega pelo cardápio, escolhe sabores (meio-a-meio), tamanho e adicionais. | PWA (React/Next.js) |
| **1.5** | PWA | Aplica lógica de cálculo baseada na pizza mais cara escolhida. | PWA (Frontend Logic) |
| **1.6** | Cliente | Preenche dados de entrega e forma de pagamento. Finaliza o pedido. | PWA -> Go API |
| **1.7** | Cliente | É redirecionado ao WhatsApp com o resumo do pedido formatado e envia a mensagem. | PWA -> WhatsApp (Link generator) |
| **1.8** | Cliente | Recebe mensagem automática confirmando o recebimento do pedido. | Evolution API -> Go API |
| **1.9** | Cliente | Recebe notificações sobre mudanças de status do pedido. | Go API -> Redis -> Evolution API |

---

# 2. Fluxo de Gestão da Pizzaria (Admin)

Este fluxo descreve como o dono da pizzaria interage com o SaaS.

| Passo | Ator | Ação | Componente Técnico |
| :--- | :--- | :--- | :--- |
| **2.1** | Dono | Faz login no painel administrativo. | Admin Dash -> Go API -> Postgres |
| **2.2** | Dono | Configura a loja: logo, cores e WhatsApp Business. | Admin Dash -> Go API -> Postgres/S3 |
| **2.3** | Dono | Configura estratégia de preço (`pricing_strategy = 'HIGHEST'`). | Admin Dash -> Go API -> Postgres (`config_json`) |
| **2.4** | Dono | Cadastra bairros, taxas, tamanhos, sabores e bebidas. | Admin Dash -> Go API -> Postgres |
| **2.5** | Dono | Mantém a tela de pedidos aberta durante horários de pico. | Admin Dash -> Go API (Concurrent Go Routines) |
| **2.6** | Dono | Recebe novo pedido em tempo real sem atualizar a página. | Go API -> WebSocket -> Admin Dash |
| **2.7** | Dono | Aceita pedido e altera status para "Em Preparo". | Admin Dash -> Go API -> Postgres |
| **2.8** | Dono | Atualiza status para "Pedido Pronto" e "Saiu para Entrega". | Admin Dash -> Go API -> Postgres |

---

# 3. Orquestração Técnica (Backend e Automação)

Este fluxo descreve os bastidores técnicos da aplicação, focando em performance e segurança.

| Passo | Ator | Ação | Componente Técnico |
| :--- | :--- | :--- | :--- |
| **3.1** | Go API | Recebe requisições da Web e da Evolution API. | Docker Container (Go) |
| **3.2** | Go API | Extrai o `tenant_id` da requisição e isola os dados da pizzaria. | Go Middleware -> Postgres (`tenant_id`) |
| **3.3** | Go API | Enfileira notificações ao invés de enviar imediatamente. | Go API -> Redis (Queue/Job) |
| **3.4** | Redis | Gerencia tarefas assíncronas de envio de mensagens. | Docker Container (Redis) |
| **3.5** | Go Worker | Processa jobs da fila e envia mensagens para a Evolution API. | Go Workers (Asynchronous process) |
| **3.6** | Sistema | Executa backup automático do banco de dados. | Go Cron Job -> `pg_dump` |
| **3.7** | Sistema | Envia o backup para armazenamento externo. | Go API -> S3/S4 Object Storage |

---

# Resumo do Fluxograma

## Entrada Business
```text
WhatsApp Bot (Triagem inicial)
```

---

## Entrada Venda
```text
PWA Web (Visual, rápido e lógico)
```

---

## Processamento Central
```text
API em Go rodando no Docker na VPS
```

---

## Saída de Notificações
```text
Redis gerenciando filas para evitar gargalos
```

---

## Segurança Passiva
```text
Backups automáticos para S3/S4
```