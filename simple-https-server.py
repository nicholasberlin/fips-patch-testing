import http.server
import ssl

PORT = 8443

httpd = http.server.HTTPServer(('0.0.0.0', PORT), http.server.SimpleHTTPRequestHandler)

httpd.socket = ssl.wrap_socket(httpd.socket,
                               keyfile="key.pem",
                               certfile="cert.pem",
                               server_side=True)

print(f"Serving on https://localhost:{PORT}")
httpd.serve_forever()
