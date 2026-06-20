# XOkingdom — Tic-Tac-Toe Game

A full-stack Tic-Tac-Toe (XO) game with a Go backend, Vue 3 frontend, and Q-learning AI trained over 5 million self-play rounds.

---

## Project Structure

Two folders only:

**Server/** — Go backend
- XOkingdom.go (server logic, routing, encryption, AI)
- brain.json (Q-table, 72,712 states)
- go.mod, go.sum

**Site/** — Vue 3 frontend
- src/ (all source components)
- index.html, package.json, vite.config.js

---

## Server (Go)

**Stack:** Go + Gin + GORM + PostgreSQL
**Encryption:** AES-256-GCM, 32-byte key from `SECRET_KEY_HEX`, random 12-byte nonce, double encryption for MatchID
**AI:** Q-learning with `brain.json` (72,712 states), loaded once at startup, `bestMove()` for optimal move selection
**Auth:** OTP via Brevo SMTP (`SMTP_PASS`), device fingerprinting (browser, OS, CPU, screen, location)

**API Routes:**
- `POST /api/auth/send-otp` — send verification code
- `POST /api/auth/verify-otp` — verify the code
- `POST /api/auth/auto-login` — auto-login via saved token
- `POST /api/game/xo/3x3/move` — submit a move
- `POST /api/game/xo/3x3/give-up` — forfeit the game
- `POST /api/game/xo/3x3/resume-playing` — resume a saved game
- `POST /api/game/xo/3x3/time-out` — time ran out

**Rate Limiting:** OTP 1/60s, Game 5/s
**Anti-Cheat:** `is_malicious_move()` validates single-cell changes, `is_malicious_game()` validates settings
**Game Logic:** roundnum 0-9, `is_win()` checks 8 paths, rounds alternate AI/Player
**Stats:** Wins, Losses, Draws, Total — updated on game end

---

## Site (Vue 3)

**Stack:** Vue 3 (Options API) + Vite + SCSS

**Pages:**
- **Authentication.vue** — login/register with OTP, device info collection, validation (nickname 3-15 chars, password 8+ chars, age 18+)
- **XOGame_3X3.vue** — game board, 14-min timer, difficulty/symbol/starter settings, `startup()` and `make_move()`
- **Policy.vue** — Privacy Policy & Terms of Service, two tabs, "Reviewed by Bashir" footer, dark theme
- **home.vue** — home page
- **XOGame.vue** — game selection page

**Networking:** `config.js` uses `window.location.hostname` for API URL, 11 fetch calls
**Fonts:** Comic Sans MS (Windows), Comic Neue (elsewhere, via Google Fonts)

---

## Python — AI Training

- Q-learning over 5,000,000 self-play rounds
- Output: `brain.json` (72,712 states)

---

## Design

Background `hsl(283, 86%, 14%)`, containers `rgb(16, 16, 36)`, active buttons `rgb(1, 1, 110)` with 5px bottom border, hover `rgb(4, 4, 34)`. Logo: X in blue, O in purple. Floating title auto-hides on collision.

---

## Running Locally

**Requirements:** Go 1.20+, Node.js, PostgreSQL on port 5432

```
cd Server && go run XOkingdom.go    # start backend
cd Site   && npm install && npm run dev  # start frontend
```

Open `http://localhost:5173`

---

## License

MIT

---

# XOkingdom — لعبة XO (Tic-Tac-Toe)

مشروع متكامل للعبة إكس-أو يشمل خادماً خلفياً بلغة Go مع تشفير AES-256-GCM وذكاء اصطناعي Q-learning، وواجهة أمامية بلغة Vue 3.

---

## هيكل المشروع

مجلدان رئيسيان:

**Server/** — الخادم الخلفي (Go)
- XOkingdom.go (منطق السيرفر، التشفير، الذكاء الاصطناعي)
- brain.json (جدول Q، 72,712 حالة)
- go.mod, go.sum

**Site/** — الواجهة الأمامية (Vue 3)
- src/ (جميع مكونات المصدر)
- index.html, package.json, vite.config.js

---

## السيرفر (Go)

**التقنيات:** Go + Gin + GORM + PostgreSQL
**التشفير:** AES-256-GCM، مفتاح 32 بايت من `SECRET_KEY_HEX`، nonce عشوائي 12 بايت، تشفير مزدوج للـ MatchID
**الذكاء الاصطناعي:** Q-learning مع `brain.json` (72,712 حالة)، يُحمل عند الإقلاع
**المصادقة:** OTP عبر Brevo SMTP (`SMTP_PASS`)، معلومات الجهاز (المتصفح، النظام، المعالج، الشاشة، الموقع)

**المسارات:**
- `POST /api/auth/send-otp` — إرسال رمز التحقق
- `POST /api/auth/verify-otp` — التحقق من الرمز
- `POST /api/auth/auto-login` — دخول تلقائي بالتوكن
- `POST /api/game/xo/3x3/move` — إرسال حركة
- `POST /api/game/xo/3x3/give-up` — استسلام
- `POST /api/game/xo/3x3/resume-playing` — استئناف اللعبة
- `POST /api/game/xo/3x3/time-out` — انتهاء الوقت

**تحديد المعدل:** OTP طلب/60ث، اللعبة 5 طلبات/ث
**كشف التلاعب:** التحقق من تغيير خلية واحدة فقط، التحقق من صحة الإعدادات
**منطق اللعبة:** الجولات 0-9، فحص 8 مسارات فوز، تبادل الأدوار
**الإحصائيات:** فوز، خسارة، تعادل، المجموع

---

## الموقع (Vue 3)

**التقنيات:** Vue 3 (Options API) + Vite + SCSS

**الصفحات:**
- **التسجيل** — دخول/اشتراك مع OTP، معلومات الجهاز، التحقق (الاسم 3-15 حرف، كلمة المرور 8+ أحرف، العمر 18+)
- **اللعبة** — رقعة 3×3، مؤقت 14 دقيقة، إعدادات الصعوبة والرمز والبادئ
- **السياسة** — سياسة الخصوصية وشروط الخدمة، توقيع "Reviewed by Bashir"
- **الرئيسية** و **اختيار اللعبة**

**الربط:** `config.js` يستخدم `window.location.hostname`، 11 طلب fetch
**الخطوط:** Comic Sans MS (ويندوز)، Comic Neue (بقية الأجهزة)

---

## Python — التدريب

تدريب Q-learning عبر 5,000,000 جولة. المخرجات: brain.json (72,712 حالة)

---

## التصميم

خلفية `hsl(283, 86%, 14%)`، صناديق `rgb(16, 16, 36)`، أزرار نشطة `rgb(1, 1, 110)` مع حد سفلي 5px، hover `rgb(4, 4, 34)`. الشعار: X أزرق، O بنفسجي. العنوان يختفي تلقائيا عند التصادم.

---

## التشغيل

**المتطلبات:** Go 1.20+، Node.js، PostgreSQL

```
cd Server && go run XOkingdom.go
cd Site   && npm install && npm run dev
```

ثم افتح `http://localhost:5173`

---

## الترخيص

MIT
