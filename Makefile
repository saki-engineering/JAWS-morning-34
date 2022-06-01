bootstrap:
	. ./config.sh && cd ./infra \
	&& cdk bootstrap \
		--qualifier $${PROJECT_NAME_ALIAS} \
		--toolkit-stack-name $${TOOLKIT_STACK_NAME} \
		--tags Project=$${PROJECT_NAME}

deploy:
	cd ./src && CGO_ENABLED=0 GOOS=linux go build -o main .
	. ./config.sh && cd ./infra \
	&& cdk deploy \
		-c projectName=$${PROJECT_NAME} \
		-c projectNameAlias=$${PROJECT_NAME_ALIAS} \
		-c resourceStackName=$${RESOURCE_STACK_NAME}

destroy:
	. ./config.sh && cd ./infra \
	&& cdk destroy \
		-c projectName=$${PROJECT_NAME} \
		-c projectNameAlias=$${PROJECT_NAME_ALIAS} \
		-c resourceStackName=$${RESOURCE_STACK_NAME}
	. ./config.sh && aws cloudformation delete-stack \
		--stack-name $${TOOLKIT_STACK_NAME}
