# XOkingdom

XO لعبة ضد ذكاء اصطناعي **ما رح تربح منه** (إلا إذا تحاكت).  
مشفرة، محمية، ومن حقك تخسر كأنك بطل.

```
XOkingdom/
├── Server/     🧠 الباك إند — Go + AES-256-GCM + Q-learning
├── Site/       🎨 الفرونت — Vue 3 + SCSS + Vite
└── README.md   🤷 وهذا
```

## وش يهمك؟

- الذكاء الاصطناعي تدرب **5 ملايين جولة** عشان يفضحك
- كل حركة مشفرة. كل لعبة مشفرة. حتى النتائج مشفرة.
- لو تحاول تخترق؟ السيرفر يطردك قبل ما تتندم
- OTP توثيق + rate limiting + كشف الحركات الخبيثة
- واجهة داكنة نظيفة ما تعمي عيونك

## شغله

**السيرفر:**
```
cd Server
set SECRET_KEY_HEX=your_32_byte_hex_key
set SMTP_PASS=your_brevo_password
go run XOkingdom.go
```

**الموقع:**
```
cd Site
npm install
npm run dev
```

## رخصة

MIT — كل شيء مسموح. اسرق، عدل، بع. ولا تنسى تذكرني.
