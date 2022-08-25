import grpc
import ping_pb2
import ping_pb2_grpc

class StrategyPingPong:
    def execute(self):
           with grpc.insecure_channel('host.docker.internal:3001') as channel:
                stub = ping_pb2_grpc.HealthCheckStub(channel)
                ping_request = ping_pb2.Ping(text = "Ping")
                pong_reply = stub.PingPong(ping_request)
                print(pong_reply)

    def supports(self, x):
        return x == 1

class StrategyPingsPong:
    def execute(self):
        print("Not implemented yet")

    def supports(self, x):
        return False

class StrategyPingPongs:
    def execute(self):
       with grpc.insecure_channel('host.docker.internal:3001') as channel:
            stub = ping_pb2_grpc.HealthCheckStub(channel)
            ping_request = ping_pb2.Ping(text = "Ping")

            response_iterator = stub.PingPongs(ping_request)

            result = ""
            for response in response_iterator:
                result += response.text
            print(result)

    def supports(self, x):
        return x == 3

class StrategyPingsPongs:
    def execute(self):
        with grpc.insecure_channel('host.docker.internal:3001') as channel:
            stub = ping_pb2_grpc.HealthCheckStub(channel)
            ping_request_iterator = self.generateRequest(5)
            response = stub.PingsPong(ping_request_iterator)
            print(response)

    def generateRequest(self, count):
        for _ in range(0, count):
            yield ping_pb2.Ping(text = "Ping")

    def supports(self, x):
        return x == 2

def inputMode():
    print('PingPong = 1')
    print('PingsPong = 2')
    print('PingPongs = 3')
    print('PingsPongs = 4')
    return int(input("Input wanted mode: "))

if __name__ == "__main__":
    strategies = [
        StrategyPingPong(),
        StrategyPingsPong(),
        StrategyPingPongs(),
        StrategyPingsPongs()
    ]

    mode = inputMode()

    for strategy in strategies:
        if(strategy.supports(mode)):
            strategy.execute()

