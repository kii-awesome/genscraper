# genscraper
 
Reading wallet address balances from large genesis files can also filter addresses

## Features

- display address and balance
- single address filter ( --address )

---

## Run

```
git clone https://github.com/kii-awesome/genscraper
cd genscraper
go run . --file /path/genesis.json
```

### Filter

```
go run . --file /path/genesis.json --address <ADDRESS>
```

