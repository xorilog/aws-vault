AWS-VAULT credential server access from docker containers.

## Linux steps:

### Using the server configured role.

1. Terminal 1: aws-vault --debug exec --server --lazy my-role -- docker compose up --build aws-vault-proxy socat-unix-to-tcp socat-tcp-to-unix
2. Terminal 2: export AWS_CONTAINER_CREDENTIALS_RELATIVE_URI=/
2. Terminal 2: docker compose run testapp
2. Terminal 2: aws sts get-caller-identity

### Assuming a role from the server configured role.

1. Terminal 1: aws-vault --debug exec --server --lazy my-role -- docker compose up --build aws-vault-proxy socat-unix-to-tcp socat-tcp-to-unix
2. Terminal 2: export AWS_CONTAINER_CREDENTIALS_RELATIVE_URI=/role-arn/arn:aws:iam::xxxxxxxxxxxx:role/xxxxxxxxxxxxxxxxx
2. Terminal 2: docker compose run testapp
2. Terminal 2: aws sts get-caller-identity


## MacOS & Windows Docker desktop steps:

### Using the server configured role.

1. Terminal 1: aws-vault --debug exec --server --lazy my-role -- docker compose up --build aws-vault-proxy
2. Terminal 2: export AWS_CONTAINER_CREDENTIALS_RELATIVE_URI=/
2. Terminal 2: docker compose run testapp
2. Terminal 2: aws sts get-caller-identity

### Assuming a role from the server configured role.

1. Terminal 1: aws-vault --debug exec --server --lazy my-role -- docker compose up --build aws-vault-proxy
2. Terminal 2: export AWS_CONTAINER_CREDENTIALS_RELATIVE_URI=/role-arn/arn:aws:iam::xxxxxxxxxxxx:role/xxxxxxxxxxxxxxxxx
2. Terminal 2: docker compose run testapp
2. Terminal 2: aws sts get-caller-identity
