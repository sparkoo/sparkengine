console.log("hello")

const WIDHT = 640
const HEIGHT = 480

const canvas = document.getElementById("myCanvas")
canvas.width = WIDHT
canvas.height = HEIGHT

const fpsField = document.getElementById("fps")

const ctx = canvas.getContext("2d")
const FRAMEBUFFER_SIZE = WIDHT * HEIGHT * 4
let FRAMEBUFFER = new Uint8ClampedArray(FRAMEBUFFER_SIZE)

let FB = new Array(FRAMEBUFFER_SIZE)

const imageData = ctx.createImageData(WIDHT, HEIGHT)

console.log("initialized")


function noise() {
    // console.log("generating noise")
    let c = (Math.random() * 1000) % 255
    // FRAMEBUFFER = Array.
    let start = new Date().getTime()
    for (let i = 0; i < FRAMEBUFFER_SIZE; i += 4) {
        // FRAMEBUFFER.set([c, c, c, c], i)

        FB[i] = c
        FB[i + 1] = c
        FB[i + 2] = c
        FB[i + 3] = c
    }
    console.log(new Date().getTime() - start, "ms")
}

let frameStart = new Date().getTime()

ctx.fillStyle = "orange"

function loop() {
    let t = new Date().getTime()
    let diff = t - frameStart
    let fps = (1 / diff) * 1000
    fpsField.innerHTML = fps
    frameStart = t

    // console.log("loop", fps)
    noise()
    // ball()
    createImageBitmap(new Blob([1,2,3])).then(function(blob) {
        ctx.drawImage(blob, 0, 0, WIDHT, HEIGHT)
    })

    requestAnimationFrame(loop)
}

let x = 0
let y = 0
const speed = 2
let vx = speed
let vy = speed
const size = 10

function ball() {
    if (x + vx < 0 || (x + vx) + size > WIDHT) {
        vx *= -1
    }
    if (y + vy < 0 || (y + vy) + size > HEIGHT) {
        vy *= -1
    }
    x += vx
    y += vy
    ctx.clearRect(0, 0, WIDHT, HEIGHT)
    ctx.fillRect(x, y, 10, 10)
}

loop()