  curl -X POST -k http://localhost:8080/example/echo -d '{"name": "gRPC-HTTP is working!"}'
  curl -X POST -k http://localhost:8080/example/streamecho -d '{"name": "gRPC-HTTP is working!"}'