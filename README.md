# Gitlab webhooks to AWTRIX 3 server

[AWTRIX 3](https://github.com/Blueforcer/awtrix3) is an open‑source custom firmware for the [Ulanzi Smart Pixel clock TC001](https://www.ulanzi.com/products/ulanzi-pixel-smart-clock-2882)

This app runs a server listening to gitlab webhooks.
When a webhook is received, it will send a message to a AWTRIX 3 server with an http request.

This can be used to display a notification when a pipeline begins and ends.

### Usage

```bash	
./gitlab-webhooks-awtrix3 --awtrix-host http://192.168.0.xxx --awtrix-username awtrix --awtrix-password awtrix --gitlab-secret xxxxxx --port 3000
```