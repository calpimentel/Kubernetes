Add a gateway:

```
-> Criando o gateway, isto nos conect com o kubernetes ingressa este route path em particular
kubectl create -f hostname-gateway.yaml
```

-> O gateway vomos criar o serço
And then associate a virtual service and basic route to the gateway:

```
    > kubectl create -f hostname-virtualservice.yaml
```

->Vamos checar o serviço
Get "service" NodePort access:

```
-> A segunda forma de checar é pegando o ip e porta e ver a saida do serviço
    > curl -I http://$(minikube ip):31575
```

-> Podemos checar o stio syste
Get "gateway" NodePort access:
```

Talk to the port 80 nodeport mapping
```
-> Tambem vamos olhar o getway model na porta 31380 que é a porta padrão
    > curl -I -Hhost:hostname.example.com http://$(minikube ip):31380
```
-> A porta padrão por alguma razão nao estiver disponivel podemos checar o serviço stio do sistema
   para ver a porta que está exporta
    > kubectl get svc -n istio-system
```
