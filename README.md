# XOkingdom

Full-stack Tic-Tac-Toe with Q-learning AI, AES-GCM encryption, Go backend, Vue frontend.

لعبة XO كاملة الأطراف مع ذكاء اصطناعي Q-learning وتشفير AES-GCM.

## Structure

```
XOkingdom/
├── Server/          — Go backend (API, auth, game logic, AI)
│   ├── XOkingdom.go — Main server
│   ├── brain.json   — Q-table (72,712 states)
│   ├── go.mod / go.sum
│   └── .env         — SECRET_KEY_HEX, SMTP_PASS
│
├── Site/            — Vue 3 frontend
│   ├── src/         — Source code
│   ├── index.html
│   ├── package.json
│   └── vite.config.js
│
├── .gitignore
└── README.md
```

## Setup

**Server:**
```
cd Server
set SECRET_KEY_HEX=your_32_byte_hex_key
set SMTP_PASS=your_brevo_password
go run XOkingdom.go
```

**Site:**
```
cd Site
npm install
npm run dev
```

## Features

- Q-learning AI (5M self-play training rounds)
- AES-256-GCM encryption for all game data
- OTP email authentication
- Malicious move detection
- Rate limiting
- Responsive dark UI
- Privacy Policy & Terms of Service

## License

MIT
