(function() {
    const canvas = document.getElementById("myCanvas");
    canvas.width = 640
    canvas.height = 480
    const gl = canvas.getContext("webgl");

    let imageData = new Uint8Array(canvas.width * canvas.height * 4);

    var program = twgl.createProgramFromScripts(
        gl, ["vshader", "fshader"], ["a_position"]);
    gl.useProgram(program);

    var verts = [
        1,  1,
        -1,  1,
        -1, -1,
        1,  1,
        -1, -1,
        1, -1,
    ];
    var vertBuffer = gl.createBuffer();
    gl.bindBuffer(gl.ARRAY_BUFFER, vertBuffer);
    gl.bufferData(gl.ARRAY_BUFFER, new Float32Array(verts), gl.STATIC_DRAW);
    gl.enableVertexAttribArray(0);
    gl.vertexAttribPointer(0, 2, gl.FLOAT, false, 0, 0);

    var tex = gl.createTexture();
    gl.bindTexture(gl.TEXTURE_2D, tex);
    gl.texParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE);
    gl.texParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE);
    gl.texParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST);
    gl.texParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST);

    let f = document.getElementById("f")

    var conn = new WebSocket("ws://localhost:8080/ws");
    conn.binaryType = 'arraybuffer'
    conn.onclose = function(evt) {
        // data.textContent = 'Connection closed';
        console.log(evt);
    }
    conn.onmessage = function(evt) {
        // console.log(evt)
        let t1 = performance.now()
        let fps = (1 / (t1 - t)) * 1000
        f.innerText = fps
        t = t1

        // console.log(evt)

        // let t11 = performance.now()
        let data = new Uint8Array(evt.data)
        gl.texImage2D(gl.TEXTURE_2D, 0, gl.RGBA, canvas.width, canvas.height, 0,
            gl.RGBA, gl.UNSIGNED_BYTE, data);
        gl.drawArrays(gl.TRIANGLES, 0, 6);
        // console.log("draw");
        // console.log(performance.now() - t11);

        // let reader = new FileReader()
        // reader.onloadend = function() {
        //
        // }
        // reader.readAsArrayBuffer(evt.data)


        // let r = new FileReader();
        // r.addEventListener("loadend", function () {
        // })
        //
        // r.readAsArrayBuffer(evt.data)
    }
    conn.onopen = function(evt) {
        console.log("OPEN");
        conn.send("hello")
    }

    let t = performance.now()
})();

