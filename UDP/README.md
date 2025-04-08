# UDP (User Datagram Protocol)

UDP (User Datagram Protocol) is used in various software architectures for its simplicity, speed, and low overhead compared to other protocols like TCP. It’s a connectionless, lightweight transport layer protocol that does not guarantee delivery, order, or error checking, but it’s much faster and efficient.

---

## 🚀 Why Use UDP

- **Low Latency**: No need to establish a connection, reducing time before data is sent.
- **Lightweight**: Minimal overhead (no handshakes, acknowledgements, or retransmissions).
- **Broadcast & Multicast Support**: Can easily be sent to multiple devices.
- **Suitable for Real-Time Apps**: Gaming, VoIP, streaming.
- **Out-of-Order Tolerance**: Applications that can tolerate packet loss or reordering use UDP.

---

## 🔧 Common Use Cases

### 🎥 Live Video Streaming / 🎮 Online Gaming
Many VoIP and video conferencing applications use UDP due to its lower overhead and ability to tolerate packet loss. Real-time communications benefit from UDP’s reduced latency compared to TCP.

### 🌐 DNS (Domain Name Service)
DNS queries typically use UDP for their fast and lightweight nature. Although DNS can also use TCP for large responses or zone transfers, most queries are handled via UDP.

### 💹 Market Data Multicast
In low-latency trading, UDP is utilized for efficient market data delivery to multiple recipients simultaneously.

### 📡 IoT with Constrained Resources
UDP is often used in IoT devices for communication, sending small packets of data between devices.

![Screenshot 2025-04-07 at 7 36 58 PM](https://github.com/user-attachments/assets/200af913-537a-4d80-a8b3-fb0793abcad4)
