let conn

(function() {
    const canvas = document.getElementById("mycanvas");
    canvas.width = 640
    canvas.height = 480
    const ctx = canvas.getContext("2d");

    let lockedMouse = false

    canvas.addEventListener("mousedown", function (e) {
        canvas.requestPointerLock()
    })

    document.addEventListener("pointerlockchange", function (e) {
        if (!lockedMouse) {
            console.log("locked")
            lockedMouse = true
            canvas.addEventListener("mousemove", updatemouse)
        } else {
            console.log("unlocked")
            lockedMouse = false
            canvas.removeEventListener("mousemove", updatemouse)
        }
    })




    const f = document.getElementById("f")

    conn = new WebSocket("ws://localhost:8080/ws");
    conn.binaryType = 'arraybuffer'
    conn.onclose = function(evt) {
        data.textContent = 'Connection closed';
    }
    let imagedata = ctx.createImageData(canvas.width,canvas.height)
    let fps = 0
    conn.onmessage = function(evt) {
        let t1 = performance.now()
        fps = (1 / (t1 - t)) * 1000
        // console.log("fps: ", fps)
        f.innerHTML = fps

        let data = new Uint8Array(evt.data)
        // console.log(data)
        imagedata.data.set(data)
        ctx.putImageData(imagedata, 0, 0)

        t = t1
    }
    conn.onopen = function(evt) {
        console.log("OPEN");
        conn.send("hello")
    }

    let t = performance.now()
})();

function updatemouse(move) {
    console.log(move.movementX, move.movementY)
    conn.send({x: move.movementX, y: move.movementY})
}