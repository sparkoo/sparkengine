const {app, BrowserWindow} = require('electron') // http://electronjs.org/docs/api
const path = require('path') // https://nodejs.org/api/path.html
const url = require('url') // https://nodejs.org/api/url.html

let window = null

// Wait until the app is ready
app.once('ready', () => {
    // Create a new window
    window = new BrowserWindow({
        width: 640,
        height: 480,
        // Don't show the window until it ready, this prevents any white flickering
        show: false,
        resizable: false,
        autoHideMenuBar: true,
        title: "Bouncing ball demo",
        fullscreenable: true
    })

    // Load a URL in the window to the local index.html path
    window.loadURL(url.format({
        pathname: path.join(__dirname, '2d.html'),
        protocol: 'file:',
        slashes: true
    }))

    // Show window when page is ready
    window.once('ready-to-show', () => {
        window.show()
    })
})