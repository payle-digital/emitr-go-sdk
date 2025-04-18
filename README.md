# Emitr Go SDK

[![Go Reference](https://pkg.go.dev/badge/github.com/payle-digital/emitr-go.svg)](https://pkg.go.dev/github.com/payle-digital/emitr-go)

Cliente oficial em Go para o **[Emitr](https://payle.digital)** — um sistema leve e de alta performance para event streaming e event sourcing criado pelo **[Payle](https://payle.digital)**.

---

## 🚀 Instalação

```bash
  go get github.com/payle-digital/emitr-go@latest
```

## ⚙️ Como usar
### 1. Criação do client

```go
import "github.com/payle-digital/emitr-go"

client := emitr.NewClient("<broker-url>", "<api-key>")
```

### 2. Publicação de eventos

```go
err := client.Produce("customer.created", `message`)
if err != nil {
    panic(err)
}
```

### 3. Consumo de eventos

```go
go client.Consume("customer.created", "worker-1", func(msg emitr.IncomingMessage) error {
    fmt.Printf("Recebido offset %d: %s\n", msg.Offset, msg.Payload)
    return nil
})
```

### 4. Formato da mensagem recebida
```go
type IncomingMessage struct {
    Offset    int64
    Timestamp time.Time
    Key       string
    Headers   map[string]string
    Payload   string
}
```

## 💡 O que o Emitr faz por você
- ✉️ Armazena mensagens por tópico de forma durável
- 🧠 Cada listener consome com offset individual
- ⚡️ Alta performance mesmo em disco local
- ☁️ Ideal para event sourcing, webhooks, integrações
- 🔐 Cada cliente se conecta via API Key que identifica seu cluster

## 🔐 Segurança
- Cada Client se conecta com uma API Key única
- A API Key determina:
  - O cluster acessado
  - As permissões disponíveis (futuro: leitura/gravação por tópico)

## 📝 Licença

Este projeto está sob a licença MIT.