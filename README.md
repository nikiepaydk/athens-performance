# athens-performance

These tests were run a Ubuntu with a Ryzen 9 3950X3D with 64 GB memory.

## How to Run

### Run test directly with default proxy: proxy.golang.org
`docker compose --profile default up --build --remove-orphans`

```
go-app-default  | real  0m 2.43s
go-app-default  | user  0m 0.89s
go-app-default  | sys   0m 0.46s
```

### Run test with athens
`docker compose --profile athens up --build --remove-orphans`

```
go-app-athens  | real   4m 33.56s
go-app-athens  | user   0m 0.79s
go-app-athens  | sys    0m 0.50s
```

### Performance degradation: x113 📉