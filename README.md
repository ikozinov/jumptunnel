## Jumptunnel

Jumptunnel is used to set up tunnels to internal services and databases using SSH tunnels via jumphost. This is especially convenient for pulling staging services into a local deployment.

This is sometimes necessary when using services (or databases) behind a corporate VPN or firewall.

### Docker-Compose

Add your credentials to .env file (use example.env as template):
- username to connect to jumphost (TUNNEL)
- jumphost FQDN or IP-address (TUNNEL)
- your ssh private key (PRIVATE_KEY)
- passphrase for private key, if needed (PASSPHRASE)

### Kubernetes

Create a secret containing some ssh keys:

kubectl create secret generic jumptunnel --from-file=ssh-privatekey=/path/to/.ssh/id_rsa --from-literal=username=testuser--from-literal=passphrase=topsecret


Deploy manifest like that:

```
apiVersion: apps/v1
kind: Deployment
metadata:
  name: pgpassport-tunnel
  labels:
    app: pgpassport-tunnel
spec:
  replicas: 1
  selector:
    matchLabels:
      app: pgpassport-tunnel
  template:
    metadata:
      labels:
        app: pgpassport-tunnel
    spec:
      containers:
      - name: jumptunnel
        image: ikozinov/jumptunnel
        ports:
        - containerPort: 5432
        env:
        - name: LISTEN_PORT
          value: "5432"
        - name: DESTINATION
          value: "postgres.internal.network:5432"
        - name: PRIVATE_KEY
          valueFrom:
            secretKeyRef:
              name: jumptunnel
              key: privatekey
        - name: TUNNEL
          valueFrom:
            secretKeyRef:
              name: jumptunnel
              key: tunnel
        - name: PASSPHRASE
          valueFrom:
            secretKeyRef:
              name: jumptunnel
              key: passphrase
---
apiVersion: v1
kind: Service
metadata:
  name: pgpassport
spec:
  selector:
    app: pgpassport-tunnel
  ports:
    - protocol: TCP
      port: 5432
      targetPort: 5432
---
```

Change detination (postgres.internal.network) and ports (5432) to what you need