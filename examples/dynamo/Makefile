localstack:
	SERVICES=dynamodb localstack start &

table:
	awslocal dynamodb create-table --table-name TestTable \
		--attribute-definitions AttributeName=Name,AttributeType=S \
		--key-schema AttributeName=Name,KeyType=HASH \
		--provisioned-throughput ReadCapacityUnits=5,WriteCapacityUnits=5
