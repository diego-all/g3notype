# run-from-gh

Local

    go run main.go start --db postgres my-ska

From Repo

    go run github.com/diego-all/run-from-gh@latest init --db postgres my-ska (Referencia)


    git remote add origin git@github.diego-all:diego-all/evergreen-con.git  (esta funciono para evergreen)
    git remote add origin git@github.diego-all:diego-all/run-from-gh.git


    go run main.go init --db postgres my-ska

    Generando proyecto 'my-ska' con base de datos 'postgres'


    go run github.com/diego-all/run-from-gh@latest init --db postgres my-ska

    root@pho3nix:/home/diegoall/MAESTRIA_ING/CLI/PRUEBACLI# go run github.com/diego-all/run-from-gh@latest init --db postgres my-ska
    go: downloading github.com/diego-all/run-from-gh v0.0.0-20240615221752-c6170d014454
    Generando proyecto 'my-ska' con base de datos 'postgres'


Al parecer tiene 2 comandos: my-cli-app e init

comando raiz


Usage:
  my-cli-app init [nombre del proyecto] [flags]


    go run github.com/diego-all/run-from-gh@latest init --db postgres my-ska


    go run github.com/diego-all/run-from-github@latest init --db postgres --config /ruta/al/archivo/config.json my-ska
