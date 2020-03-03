# Lavagna-go
Una semplice lavagna su cui poter scrivere.  
Frontend: https://gitlab.com/Leonardoalgeri/lavagnafrontend  
Backend: https://github.com/LeonardoAlgeri/Lavagna-go  
DockerHub: https://hub.docker.com/r/leonardoalgeri/lavagna-go

## Deploy
In fase di deploy è necessario fornire all'applicativo le seguenti varibili d'ambiente:
- Per il database mysql  
  - **SQL_URL** il link dell'host sql seguito da porta
  - **SQL_NAME** nome del database
  - **SQL_USER** l'utente del database sql 
  - **SQL_PASSWORD** la password del database sql


## Docker
È possibile eseguire il seguente programma su docker seguendo questi passi:
1) Settare un file env.list contentente il valore delle variabili d'ambiente. Trovate un esempio [qui](https://docs.docker.com/engine/reference/commandline/run/)  cercando **env**  
2) Eseguire le seguenti istruzioni

```
docker run --env-file env.list -dp 8080:8080 --name lavagna --restart always leonardoalgeri/lavagna-go
```
3) Attendere qualche minuto per l'avvio di spring
  
Per fermare il container è sufficiente la seguente operazione
```
docker stop lavagna
```

## Update dell'applicazione
L'update va fatto a mano.   

Per quanto riguarda docker eseguire le seguenti istruzioni  
```
docker stop lavagna
docker rm lavagna
docker pull leonardoalgeri/lavagna-go
```
Ed eseguire poi di nuovo il container
```
docker run --env-file env.list -dp 8080:8080 --name lavagna --restart always leonardoalgeri/lavagna-go
```

## Warning
In caso di problemi con la visualizzazione di emoji sulla lavagna controllare che la codifica del database sia **utf8mb4_0900_ai_ci**