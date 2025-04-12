# 🚀 Crypto Ticker Aggregator API

**Ticker** is a fast and extensible Go-based API that fetches cryptocurrency ticker data from multiple exchanges and stores it in a MongoDB database.  

---

## 🌐 Supported Exchanges

- Gemini  
- Kraken  
- Crypto.com  
- CEX.io  
- FMFW  
- Bitfinex  

---

## ✨ Features

- 📡 Real-time crypto ticker retrieval  
- 🗄️ MongoDB integration for persistent storage  
- 🧩 Easily extensible architecture for new exchanges  
- 🧾 RESTful API with OpenAPI specification  
- ⚙️ Configuration using TOML files  

---

## 🧰 Getting Started

### 1️⃣ Clone the repository

```bash
git clone https://github.com/CIDgravity/Ticker.git
cd ticker
```

### 2️⃣ Configure the application

Copy and customize the sample config:

```bash
cp config/config.toml.sample ./config.toml
```

### 3️⃣ Build the binary

```bash
make build
```

This will produce the `cidgravity-ticker` binary in the root of the project.

### 4️⃣ Run the application

```bash
./cidgravity-ticker --config path/to/config.toml
```

> If you run the binary from the same directory as your config, the `--config` flag is optional.

---

## 🧪 Development Tools

| Command         | Description                     |
|----------------|---------------------------------|
| `make test`     | Run unit tests                  |
| `make lint`     | Execute linters                 |
| `make openapi`  | Build OpenAPI documentation     |
| `make build`    | Compile the application binary  |

---

## 🔌 Adding a New Exchange

Adding support for another exchange involves three steps:

### 1. Create a new fetcher

Add a Go file in:

```
internal/exchange/new_exchange.go
```

### 2. Register the exchange

In `service/exchange_service.go` (around line 51), initialize the exchange:

```go
new_exchange.New()
```

### 3. Configure trading pair mappings

Edit:

```
config/exchange.go
```

And add the appropriate pair mappings for your new exchange.

---

## 📖 API Documentation

The API is documented with OpenAPI.

Generate it locally with:

```bash
make openapi
```

Or view the hosted version: [📘 API Docs](#)

---

## 📄 License

This project is licensed under the **MIT License**.  
See the [LICENSE](./LICENSE) file for more details.

---

## 🤝 Contributions

We welcome contributions of any kind!  
Feel free to open issues, suggest features, or submit pull requests to help improve **Ticker**.