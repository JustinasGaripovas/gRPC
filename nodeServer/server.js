import PingPongService from './pingPongService.js'
import * as grpc from '@grpc/grpc-js';
import * as protoLoader from '@grpc/proto-loader';

const PROTO_PATH = './ping.proto';
const port = '3001';
const ip = '0.0.0.0:';

const packageDefinition = protoLoader.loadSync(
    PROTO_PATH,
    {
            keepCase: true,
            longs: String,
            enums: String,
            defaults: true,
            oneofs: true
    });

let healthCheck = grpc.loadPackageDefinition(packageDefinition).healthCheck;

function main() {
        let server = new grpc.Server();
        let service = new PingPongService();

        server.addService(healthCheck.HealthCheck.service,{
                    "pingPong": service.pingPong,
                    "pingPongs": service.pingPongs,
                    "pingsPong": service.pingsPong,
                    "pingsPongs": service.pingsPongs,
            });
        server.bindAsync(ip + port, grpc.ServerCredentials.createInsecure(), () => {
                server.start();
        });
}

main();