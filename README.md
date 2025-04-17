# Forefinger

Forefinger - это высокопроизводительная Go библиотека для взаимодействия с Ethereum JSON-RPC API. Библиотека предлагает
оптимизированный и удобный интерфейс для работы с Ethereum узлами.

## Особенности

- Эффективная работа с пулом соединений
- Пакетные вызовы для оптимизации производительности
- Поддержка всех стандартных методов Ethereum JSON-RPC
- Экономичное использование памяти благодаря внутренней структуре моделей данных
- Поддержка фильтрации логов и событий
- Интуитивно понятный API для Go разработчиков

## Установка

```bash
go get -u github.com/s4bb4t/forefinger
```

## Быстрый старт

```go
package main

import (
	"context"
	"fmt"
	"github.com/s4bb4t/forefinger/pkg/client"
	"math/big"
)

func main() {
	// Создание клиента с указанием URL Ethereum узла и размера пула соединений
	c, err := client.NewClient("https://mainnet.infura.io/v3/YOUR-API-KEY", 5)
	if err != nil {
		panic(err)
	}
	defer c.Close()

	// Получение последнего блока
	block, err := c.BlockByNumber(context.Background(), nil)
	if err != nil {
		panic(err)
	}

	// Вывод номера блока
	fmt.Printf("Последний блок: %s\n", block.Number())
}
```

## Основные возможности

### Инициализация клиента

```go
// Создание клиента с пулом из 5 соединений
client, err := client.NewClient("https://mainnet.infura.io/v3/YOUR-API-KEY", 5)
if err != nil {
// обработка ошибки
}
defer client.Close()
```

### Запросы блоков и транзакций

```go
// Получение блока по номеру
block, err := client.BlockByNumber(ctx, big.NewInt(14000000))

// Получение блока по хешу
blockHash := common.HexToHash("0x...")
block, err := client.BlockByHash(ctx, blockHash)

// Получение транзакции по хешу
txHash := common.HexToHash("0x...")
tx, err := client.TxByHash(ctx, txHash)

// Получение квитанции транзакции
receipt, err := client.TxReceipt(ctx, txHash)
```

### Работа с аккаунтами и смарт-контрактами

```go
// Получение баланса аккаунта
address := common.HexToAddress("0x...")
balance, err := client.Balance(ctx, address, nil)

// Чтение кода смарт-контракта
code, err := client.Code(ctx, address, nil)

// Вызов метода смарт-контракта
data := []byte{...} // ABI-кодированные данные вызова
result, err := client.CallContract(ctx, address, data, nil)

// Оценка газа для транзакции
gas, err := client.EstimateGas(ctx, data, nil)
```

### Работа с фильтрами и логами

```go
// Создание нового фильтра
filter := models.NewFilter().
FromBlock("latest"). // Начальный блок
ToBlock(big.NewInt(14000000)). // Конечный блок
Address("0x123...").           // Фильтрация по адресу контракта
Topic("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef") // Event signature для Transfer

// Применение фильтра для получения логов
logs, err := client.Logs(ctx, filter)

// Создание и использование фильтра на стороне узла
filterId, err := client.NewFilter(ctx, filter)
if err != nil {
// обработка ошибки
}

// Получение изменений для фильтра
changes, err := client.FilterChanges(ctx, filterId)

// Удаление фильтра после использования
client.UninstallFilter(ctx, filterId)
```

## Пакетные запросы

Forefinger поддерживает два типа пакетных запросов для оптимизации производительности:

### BatchCall

Позволяет выполнить несколько однотипных запросов в одном HTTP-вызове:

```go
// Создание списка результатов
results := make([]*big.Int, 10)

// Создание списка аргументов
args := make([][]any, 10)
for i := 0; i < 10; i++ {
args[i] = []any{common.HexToAddress(fmt.Sprintf("0x%x", i)), "latest"}
}

// Выполнение пакетного запроса балансов
err, errs := client.BatchCall(
ctx,
5, // размер пакета
methods.Balance,
args,
&results,
)

// Проверка ошибок и обработка результатов
```

### SequenceBatchCall

Позволяет выполнить последовательность разнотипных запросов:

```go
// Создание последовательности запросов
sequence := methods.Sequence{
{
Method: methods.BlockByNumber,
Args:   []any{nil, true},
Result: &block,
},
{
Method: methods.Balance,
Args:   []any{common.HexToAddress("0x..."), "latest"},
Result: &balance,
},
}

// Выполнение пакетного запроса
err, errs := client.SequenceBatchCall(ctx, 2, &sequence)
```

## Структура модели данных

Forefinger использует эффективную внутреннюю структуру для моделей данных, разделяя часто используемые и редко
используемые поля:

```go
// Получение информации о блоке
block, _ := client.BlockByNumber(ctx, nil)

// Доступ к часто используемым полям (хранятся непосредственно в структуре)
blockNumber := block.Number() // *big.Int
timestamp := block.Timestamp() // *big.Int
txs := block.Transactions() // Transactions

// Доступ к дополнительным полям (хранятся в сжатом формате)
hash, _ := block.Hash() // common.Hash
miner, _ := block.Miner() // common.Address
difficulty, _ := block.Difficulty() // *big.Int
```

## Конкурентная обработка

Библиотека обеспечивает безопасную конкурентную работу с использованием пула соединений:

```go
// Создание клиента с пулом из 10 соединений
client, _ := client.NewClient("https://ethereum.node.url", 10)

// Выполнение параллельных запросов
var wg sync.WaitGroup
for i := 0; i < 20; i++ {
wg.Add(1)
go func (blockNum int64) {
defer wg.Done()
block, err := client.BlockByNumber(context.Background(), big.NewInt(blockNum))
// обработка блока
}(int64(15000000 + i))
}
wg.Wait()
```

## Лицензия

[MIT License](LICENSE)