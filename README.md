# XOkingdom — Tic-Tac-Toe Game

A full-stack Tic-Tac-Toe (XO) game with a Go backend, Vue 3 frontend, and Q-learning AI trained over 5 million self-play rounds.

---

## Project Structure

```
XOkingdom/
├── Server/          # Go backend
│   ├── XOkingdom.go # Main server file
│   ├── brain.json   # Q-table (72,712 game states)
│   ├── go.mod       # Module definition
│   └── go.sum       # Dependency checksums
│
├── Site/            # Vue 3 frontend
│   ├── src/         # Source code
│   ├── index.html   # Entry HTML
│   ├── package.json # Dependencies
│   └── vite.config.js # Build config
│
├── .gitignore
└── README.md
---

## Server (Go)

### Tech Stack
- Go with Gin framework for routing and HTTP
- GORM for PostgreSQL database ORM
- PostgreSQL on port 5432

### Encryption
- AES-256-GCM for all game data
- 32-byte key from `SECRET_KEY_HEX` environment variable
- Random 12-byte nonce per encryption operation
- Double encryption for MatchID

### AI
- Q-learning reinforcement learning
- `brain.json`: 72,712 game states
- Loaded once at server startup via `init()`
- `bestMove()` selects the optimal move

### Authentication
- OTP via Brevo SMTP email
- SMTP password from `SMTP_PASS` environment variable
- Device fingerprinting: browser, OS, CPU cores, screen, timezone, language, battery, geolocation (via ipapi.co)

### API Routes

| Route | Method | Description |
|-------|--------|-------------|
| `/api/auth/send-otp` | POST | Send verification code |
| `/api/auth/verify-otp` | POST | Verify the code |
| `/api/auth/auto-login` | POST | Auto-login via saved token |
| `/api/game/xo/3x3/move` | POST | Submit a move |
| `/api/game/xo/3x3/give-up` | POST | Forfeit the game |
| `/api/game/xo/3x3/resume-playing` | POST | Resume a saved game |
| `/api/game/xo/3x3/time-out` | POST | Time ran out |

### Rate Limiting
- OTP: 1 request per 60 seconds
- Game: 5 requests per second

### Anti-Cheat
- `is_malicious_move()`: Validates that only one cell changed from empty to the player's symbol
- `is_malicious_game()`: Validates game settings (difficulty, symbol, click, timer, board)

### Game Logic
- `roundnum` (0-9) calculated from filled cells count
- `is_win()`: Checks 8 win paths with non-empty cell validation
- Rounds: 0 (AI starts first), 1 (Player starts), 2-7 (alternating), 8 (AI final move), 9 (Player final move)

### Statistics
- Wins, Losses, Draws, Total — updated on every game end
- Responses: WIN, LOSE, DRAW — includes final board state

---

## Site (Vue 3)

### Tech Stack
- Vue 3 (Options API)
- Vite build tool
- SCSS styling

### Pages

**Authentication.vue** — Login & registration:
- Two tabs: Log in and Sign in
- OTP verification
- Collects device info: browser, OS, CPU, screen, timezone, language, battery, network, IP, country, city
- Validation: nickname (3-15 chars), password (8+ chars), age (18+)

**XOGame_3X3.vue** — Game page:
- 14-minute countdown timer
- 3×3 game board
- Settings: difficulty, symbol (X/O), starter
- `startup()`: Initialize the game
- `make_move()`: Send move, receive result

**Policy.vue** — Privacy Policy & Terms of Service:
- Two tabs for navigation
- "Reviewed by Bashir" footer
- Dark theme

**home.vue** — Home page

**XOGame.vue** — Game selection page

### Networking
- `config.js`: Uses `window.location.hostname` to determine API URL
- 11 fetch calls across the project

### Fonts
- "Comic Sans MS" on Windows, "Comic Neue" elsewhere (via Google Fonts)

---

## Python — AI Training

`Tic-Tac-Toe_3X3.py`:
- Q-learning training over 5,000,000 self-play rounds
- Output: `brain.json` with 72,712 states

---

## Design

- Background: `hsl(283, 86%, 14%)` (dark purple)
- Containers: `rgb(16, 16, 36)` (dark blue)
- Active buttons: 5px bottom border in `rgb(1, 1, 110)`
- Hover: `rgb(4, 4, 34)`
- Logo: X in blue, O in purple
- Floating title auto-hides on collision with the container

---

## Running Locally

### Requirements
- Go 1.20+
- Node.js
- PostgreSQL on port 5432

### Server
```
cd Server
# Create .env file with your keys
go run XOkingdom.go
```

### Site
```
cd Site
npm install
npm run dev
```

Open `http://localhost:5173` in your browser.

---

## License

MIT

---

# XOkingdom — لعبة XO (Tic-Tac-Toe)

مشروع متكامل للعبة إكس-أو يشمل خادماً خلفياً بلغة Go مع تشفير AES-256-GCM وذكاء اصطناعي قائم على التعلم المعزز (Q-learning)، وواجهة أمامية بلغة Vue 3.

---

## هيكل المشروع

```
XOkingdom/
├── Server/          # الخادم الخلفي (Go)
│   ├── XOkingdom.go # ملف السيرفر الرئيسي
│   ├── brain.json   # جدول Q (72,712 حالة لعبة)
│   ├── go.mod       # تعريف الحزمة والتبعيات
│   └── go.sum       # التواقيع الرقمية للحزم
│
├── Site/            # الواجهة الأمامية (Vue 3)
│   ├── src/         # مجلد المصدر
│   ├── index.html   # الصفحة الرئيسية
│   ├── package.json # تعريف التبعيات
│   └── vite.config.js # إعدادات البناء
│
├── .gitignore       # الملفات المستثناة من git
└── README.md        # هذا الملف
```

---

## السيرفر (Go)

### التقنيات
- Go مع إطار Gin لإدارة المسارات وخدمة HTTP
- GORM للربط مع قاعدة البيانات PostgreSQL
- PostgreSQL على المنفذ 5432

### التشفير
- AES-256-GCM لتشفير جميع بيانات اللعبة
- مفتاح 32 بايت يُقرأ من متغير البيئة `SECRET_KEY_HEX`
- nonce عشوائي 12 بايت لكل عملية تشفير
- MatchID مشفر مرتين (تشفير مزدوج)

### الذكاء الاصطناعي
- خوارزمية Q-learning
- ملف `brain.json` يحتوي 72,712 حالة لعبة
- يُحمل في الذاكرة مرة واحدة عند إقلاع السيرفر
- البحث عن أفضل حركة متاحة

### المصادقة
- OTP عبر البريد الإلكتروني باستخدام Brevo SMTP
- كلمة مرور SMTP من متغير البيئة `SMTP_PASS`
- معلومات الجهاز: المتصفح، نظام التشغيل، المعالج، الشاشة، الموقع الجغرافي

### المسارات

| المسار | الطريقة | الوظيفة |
|--------|---------|---------|
| `/api/auth/send-otp` | POST | إرسال رمز التحقق للبريد |
| `/api/auth/verify-otp` | POST | التحقق من الرمز |
| `/api/auth/auto-login` | POST | دخول تلقائي بالتوكن |
| `/api/game/xo/3x3/move` | POST | إرسال حركة اللاعب |
| `/api/game/xo/3x3/give-up` | POST | استسلام |
| `/api/game/xo/3x3/resume-playing` | POST | استئناف اللعبة |
| `/api/game/xo/3x3/time-out` | POST | انتهاء الوقت |

### تحديد المعدل
- OTP: طلب واحد لكل 60 ثانية
- اللعبة: 5 طلبات في الثانية

### كشف التلاعب
- التحقق من أن اللاعب غير خلية واحدة فقط من فارغة إلى رمزه
- التحقق من صحة إعدادات اللعبة

### منطق اللعبة
- يُحتسب رقم الجولة (0-9) بعدد الخلايا المملوءة
- فحص 8 مسارات فوز مع التأكد من أن الخلايا غير فارغة
- الجولات: 0 (الكمبيوتر يسبق)، 1 (اللاعب يبدأ)، 2-7 (تبادل)، 8 (آخر حركة للكمبيوتر)، 9 (آخر حركة للاعب)

### الإحصائيات
- فوز، خسارة، تعادل، المجموع — تُحدث عند نهاية كل لعبة

---

## الموقع (Vue 3)

### التقنيات
- Vue 3 (Options API)
- Vite للبناء
- SCSS للتصميم

### الصفحات

**التسجيل** — تسجيل الدخول والاشتراك:
- تبويبان: دخول واشتراك
- التحقق برمز OTP
- معلومات الجهاز: المتصفح، نظام التشغيل، المعالج، الشاشة، المنطقة الزمنية، اللغة
- التحقق: الاسم (3-15 حرف)، كلمة المرور (8 أحرف فأكثر)، العمر (18 سنة فأكثر)

**اللعبة** — صفحة اللعبة:
- مؤقت 14 دقيقة
- رقعة 3×3
- إعدادات: الصعوبة، الرمز، البادئ

**السياسة** — سياسة الخصوصية وشروط الخدمة
- تبويبان
- توقيع "Reviewed by Bashir"
- ألوان داكنة

### الخطوط
- Comic Sans MS على ويندوز، Comic Neue على باقي الأجهزة

---

## Python — التدريب

تدريب Q-learning عبر 5,000,000 جولة لعب ذاتي
المخرجات: brain.json (72,712 حالة)

---

## التشغيل

### المتطلبات
- Go الإصدار 1.20 فأحدث
- Node.js
- PostgreSQL

### تشغيل السيرفر
```
cd Server
# إنشاء ملف .env
go run XOkingdom.go
```

### تشغيل الموقع
```
cd Site
npm install
npm run dev
```

ثم افتح `http://localhost:5173`

---

## الترخيص

MIT
