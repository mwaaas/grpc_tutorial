deploy_base_image:
	docker build -t mwaaas/grpc_tutorial:base_image -f BaseDockerfile .
	docker push mwaaas/grpc_tutorial:base_image

probuf:
	protoc helloworld.proto --go_out=plugins=grpc:.