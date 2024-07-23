Linux Steps:

1. Terminal 1: aws-vault --debug exec --server --lazy my-role -- docker compose up --build aws-vault-proxy socat-unix-to-tcp socat-tcp-to-unix
2. Terminal 2: export AWS_CONTAINER_CREDENTIALS_RELATIVE_URI=/role-arn/arn:aws:iam::xxxxxxxxxxxx:role/xxxxxxxxxxxxxxxxx
2. Terminal 2: docker compose run testapp
2. Terminal 2: aws sts get-caller-identity
