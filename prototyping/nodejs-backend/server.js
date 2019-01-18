const WebSocket = require('ws')

let wss = new WebSocket.Server({
    port: 8081,
})

wss.on('connection', function connection(ws) {
    ws.on('message', function incoming(message) {
        console.log('received: %s', message);
    });

    ws.send('something');

    const size = 320 * 200 * 4

    let arr = new Uint8Array(size)

    // for (let i = 0; i < size; i += 4) {
    //     arr[i] = 127
    //     arr[i + 1] = 127
    //     arr[i + 2] = 127
    //     arr[i + 3] = 255
    //     ws.send(arr)
    //     console.log("send framebuffer");
    // }
    let i = 0
    setTimeout(sendNext, 1)
    function sendNext() {
        // console.log(i);
        if (i < size) {
            arr[i++] = 127
            arr[i++] = 127
            arr[i++] = 127
            arr[i++] = 255
            ws.send(arr)
            setTimeout(sendNext, 1)
        }
    }
});


var webs = require('connect');
var serveStatic = require('serve-static');
webs().use(serveStatic(__dirname)).listen(8080, function(){
    console.log('Server running on 8080...');
});
