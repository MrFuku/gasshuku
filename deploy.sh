#!/bin/sh

eval $(aws ecr get-login --no-include-email --region ap-northeast-1)
docker push $AWS_ACCOUNT_ID.dkr.ecr.ap-northeast-1.amazonaws.com/gasshuku:latest
./scripts/ecs-deploy --cluster gasshuku-cluster --service-name gasshuku-service --image $AWS_ACCOUNT_ID.dkr.ecr.ap-northeast-1.amazonaws.com/gasshuku:latest --timeout 600
