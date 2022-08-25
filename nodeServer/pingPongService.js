export default class PingPongService{
    pingPong(call, callback) {
        callback(null, {text: call.request.text + 'Pong'});
    }

    pingPongs(call) {
        call.write({text: 'Ping'})

        for (let i = 0; i < 3; i++) {
            call.write({text: 'Pong'})
        }

        call.end();
    }

    pingsPong(call, callback) {
        let pings = ""

        call.on('data', function(ping) {
            pings += ping.text
        });

        call.on('end', function() {
            callback(null, {text: pings + "Pong"});
        });
    }

    pingsPongs(call, callback) {
        callback(null, {text: call.request.text + 'Pong'});
    }
}