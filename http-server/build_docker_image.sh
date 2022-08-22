# Copy the grpc service's protobuf definitions
# as we need them to build the package server
cp -r ../grpc-server/service .
docker build -t davarski/pkgserver --progress plain . 
# Remove it after building
rm -r service
