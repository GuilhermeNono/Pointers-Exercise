## 1) Básico: disparar goroutines
**Objetivo:** entender execução concorrente.

- Crie uma função `say(msg string)` que imprime a mensagem 5 vezes com `time.Sleep`.
- Rode duas goroutines com mensagens diferentes.
- No `main`, use `time.Sleep` para não encerrar cedo (depois troque por `WaitGroup`).

---

## 2) Sincronização com `sync.WaitGroup`
**Objetivo:** esperar várias goroutines terminarem corretamente.

- Lance 3 goroutines simulando tarefas (`task 1`, `task 2`, `task 3`).
- Cada tarefa dorme um tempo aleatório e imprime início/fim.
- Use `wg.Add`, `defer wg.Done`, `wg.Wait`.

---

## 3) Comunicação com channels
**Objetivo:** trocar dados entre goroutines.

- Faça uma goroutine calcular o quadrado de números 1..5 e enviar para um `chan int`.
- No `main`, leia e imprima os resultados.
- Feche o canal quando terminar (`close(ch)`).

---

## 4) Worker Pool
**Objetivo:** padrão muito usado em backend.

- Crie:
    - `jobs := make(chan int, 10)`
    - `results := make(chan int, 10)`
- Suba 3 workers (goroutines) processando jobs.
- Envie 10 jobs e colete resultados.
- Cada worker pode “processar” com `Sleep` + multiplicação.

---

## 5) Fan-out / Fan-in
**Objetivo:** paralelizar processamento e juntar saída.

- Entrada: lista de números.
- Fan-out: 4 goroutines processam (ex: dobro do número).
- Fan-in: juntar tudo em um único canal de saída.
- Dica: use `WaitGroup` para fechar o canal final.

---

## 6) Timeout e cancelamento (`context`)
**Objetivo:** controlar goroutines longas.

- Crie função `fetchData(ctx context.Context)` que leva ~3s.
- Rode com `context.WithTimeout(..., 1*time.Second)`.
- Faça a goroutine respeitar `ctx.Done()` e retornar cancelamento.

---

## 7) Evitar race condition
**Objetivo:** segurança de dados compartilhados.

- Faça contador global incrementado por 100 goroutines.
- Primeiro sem proteção (resultado inconsistente).
- Depois corrija com:
    - `sync.Mutex`, e
    - alternativa com channel (estilo actor).
- Rode com `go run -race`.
