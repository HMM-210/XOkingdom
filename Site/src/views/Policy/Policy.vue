<template>
  <div class="body_title">
    <h1 class="title" :class="{ 'hidden-title': !show_title }">
      <router-link to="/">
        <span class="xletter">X</span><span class="oletter">O</span>kingdom
      </router-link>
    </h1>
    <div class="layout">
      <div class="tab-bar">
        <button :class="['tab-btn', { active: activeTab === 'terms' }]" @click="activeTab = 'terms'">Terms</button>
        <button :class="['tab-btn', { active: activeTab === 'privacy' }]" @click="activeTab = 'privacy'">Privacy</button>
      </div>
      <div class="content">
        <div class="policy-content" v-if="activeTab === 'terms'">
          <h2>Terms of Service</h2>
          <p class="subtitle">Last updated: 20 June 2026</p>
          <div class="consent-line">By using the Service, you agree to these Terms of Service.</div>

          <h3>1. Acceptance of Terms</h3>
          <p>By accessing or using the XOkingdom website, web application, and any related services (collectively, the "Service") operated by HMM210 ("Operator," "we," "us," or "our"), you agree to be bound by these Terms of Service ("Terms"). If you do not agree, you must not use the Service.</p>

          <h3>2. Eligibility</h3>
          <p>You must be at least 18 years of age to use the Service. By using the Service, you represent and warrant that you meet this age requirement and that you have the legal capacity to enter into these Terms.</p>

          <h3>3. Account Registration and Security</h3>
          <p>To access certain features, you may be required to create an account. You are responsible for maintaining the confidentiality of your login credentials and for all activities that occur under your account. You agree to notify us immediately of any unauthorized use of your account. We are not liable for any loss or damage arising from your failure to safeguard your account.</p>

          <h3>4. Acceptable Use</h3>
          <p>You agree not to:</p>
          <ul>
            <li>Use the Service for any unlawful purpose or in violation of any applicable law or regulation;</li>
            <li>Attempt to interfere with, disrupt, or degrade the integrity or performance of the Service;</li>
            <li>Reverse-engineer, decompile, or disassemble any aspect of the Service;</li>
            <li>Use automated means (bots, scrapers, etc.) to access or interact with the Service without our express permission;</li>
            <li>Submit or transmit any content that is harmful, abusive, harassing, or otherwise objectionable.</li>
          </ul>

          <h3>5. Intellectual Property</h3>
          <p>The Service, including its code, design, graphics, and underlying AI models, is the property of HMM210 and is protected by applicable intellectual property laws. You are granted a limited, non-exclusive, non-transferable, revocable license to use the Service for its intended purpose. No other rights are granted.</p>

          <h3>6. Feedback</h3>
          <p>If you provide us with any suggestions, ideas, or feedback regarding the Service, you grant HMM210 an irrevocable, royalty-free, perpetual, and worldwide license to use, modify, and incorporate such feedback without any obligation to you.</p>

          <h3>7. Privacy</h3>
          <p>Your use of the Service is governed by our Privacy Policy, which is incorporated into these Terms by reference.</p>

          <h3>8. Modifications to the Service or Terms</h3>
          <p>We reserve the right to modify or discontinue the Service at any time, with or without notice. We may also update these Terms from time to time. Material changes will be communicated via the Service; non-material changes take effect immediately. Your continued use of the Service after any modification constitutes your acceptance of the revised Terms.</p>

          <h3>9. Limitation of Liability and Disclaimer of Warranties</h3>
          <p>THE SERVICE IS PROVIDED "AS IS" AND "AS AVAILABLE" WITHOUT ANY WARRANTIES, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE, AND NON-INFRINGEMENT. WE DO NOT GUARANTEE THAT THE SERVICE WILL BE UNINTERRUPTED, ERROR-FREE, OR SECURE.</p>
          <p>TO THE MAXIMUM EXTENT PERMITTED BY LAW, HMM210 SHALL NOT BE LIABLE FOR ANY INDIRECT, INCIDENTAL, SPECIAL, CONSEQUENTIAL, OR PUNITIVE DAMAGES ARISING OUT OF OR RELATED TO YOUR USE OF THE SERVICE. OUR TOTAL LIABILITY TO YOU SHALL NOT EXCEED THE GREATER OF (A) THE AMOUNTS PAID BY YOU TO US IN THE TWELVE (12) MONTHS PRECEDING THE CLAIM, OR (B) USD 10.00.</p>

          <h3>10. Termination</h3>
          <p>We reserve the right to suspend or terminate your access to the Service at any time, for any reason or no reason, with or without notice. Upon termination, your right to use the Service immediately ceases. Sections 5, 6, 9, and 10 shall survive termination.</p>

          <h3>11. Governing Law and Dispute Resolution</h3>
          <p>These Terms are governed by the laws of the Republic of Iraq. Any disputes arising out of or relating to these Terms or the Service shall be resolved in the courts of Baghdad, Iraq. You agree to submit to the personal jurisdiction of such courts.</p>

          <h3>12. Limitation Period</h3>
          <p>Any claim or cause of action arising out of or related to these Terms or the Service must be filed within one (1) year after the claim accrues, or such claim is permanently barred.</p>

          <h3>13. Contact</h3>
          <p>For questions about these Terms, please contact us at: <a href="mailto:hassan210mmm@gmail.com">hassan210mmm@gmail.com</a></p>
        </div>

        <div class="policy-content" v-if="activeTab === 'privacy'">
          <h2>Privacy Policy</h2>
          <p class="subtitle">Last updated: 20 June 2026</p>
          <div class="consent-line">By using the Service, you consent to this Privacy Policy.</div>

          <h3>1. Introduction</h3>
          <p>HMM210 ("Operator," "we," "us," or "our") operates the XOkingdom Service. This Privacy Policy explains how we collect, use, disclose, and safeguard your information when you use our Service. By using the Service, you consent to the practices described in this policy.</p>

          <h3>2. Information We Collect</h3>
          <p>We collect the following categories of data:</p>
          <table>
            <thead>
              <tr><th>Category</th><th>Examples</th></tr>
            </thead>
            <tbody>
              <tr><td>Account Data</td><td>Nickname, email address, password (hashed), and other similar data reasonably collected for the same purposes.</td></tr>
              <tr><td>Profile Data</td><td>Gender, birthdate, first name, last name, and other similar data reasonably collected for the same purposes.</td></tr>
              <tr><td>Device Data</td><td>Browser type, operating system, screen resolution, CPU cores, language, timezone, color scheme preference, and other similar data reasonably collected for the same purposes.</td></tr>
              <tr><td>Network Data</td><td>IP address, country, city, internet service provider, and other similar data reasonably collected for the same purposes.</td></tr>
              <tr><td>Game Data</td><td>Game board state, moves, timer settings, difficulty level, win/loss/draw records, and other similar data reasonably collected for the same purposes.</td></tr>
              <tr><td>Communication Data</td><td>OTP verification codes, email correspondence, and other similar data reasonably collected for the same purposes.</td></tr>
              <tr><td>Analytics Data</td><td>Usage patterns, feature interactions, performance metrics, and other similar data reasonably collected for the same purposes.</td></tr>
              <tr><td>Any Other Data Necessary for Service Operation</td><td>Any other data that is reasonably necessary for the operation, improvement, or security of the Service, and other similar data reasonably collected for the same purposes.</td></tr>
            </tbody>
          </table>

          <h3>3. How We Collect Information</h3>
          <p>Information is collected when you: register an account, play a game, update your profile, communicate with us, or automatically through your device and browser settings.</p>

          <h3>4. Purpose of Processing</h3>
          <p>We process your information for the following purposes:</p>
          <ul>
            <li>To provide, maintain, and improve the Service;</li>
            <li>To authenticate your identity and secure your account;</li>
            <li>To enable gameplay, including AI-powered move suggestions;</li>
            <li>To communicate with you regarding your account or the Service;</li>
            <li>To detect and prevent fraudulent or malicious activity;</li>
            <li>To comply with legal obligations.</li>
          </ul>

          <h3>5. Legal Basis for Processing</h3>
          <p>We process your personal data based on: your consent (where required by applicable law), the necessity of processing for the performance of the Service, our legitimate interests in operating and improving the Service, and compliance with legal obligations.</p>

          <h3>6. Sharing of Information</h3>
          <p>We do not sell your personal data. We may share your information with:</p>
          <ul>
            <li><strong>Service providers:</strong> We engage trusted third-party service providers (including Brevo for email delivery and ipapi.co for IP geolocation) to assist in operating the Service. These providers are contractually bound to protect your data and use it only for the services we request.</li>
            <li><strong>Legal authorities:</strong> We may disclose information if required by law, legal process, or governmental request, or to protect the rights, property, or safety of HMM210, our users, or others.</li>
          </ul>

          <h3>7. Data Retention</h3>
          <p>We retain your personal data for as long as your account is active or as otherwise necessary to provide the Service, comply with legal obligations, resolve disputes, and enforce our agreements. The specific retention period is determined by the Operator based on operational needs.</p>

          <h3>8. Data Security</h3>
          <p>We implement appropriate technical and organizational measures to protect your data, including encryption of sensitive game data using AES-256-GCM. However, no method of transmission or storage is 100% secure, and we cannot guarantee absolute security.</p>

          <h3>9. Your Rights</h3>
          <p>Depending on your jurisdiction, you may have the right to access, correct, delete, or port your personal data, as well as the right to restrict or object to certain processing activities. These rights are subject to limitations under applicable law. To exercise your rights, please contact us at <a href="mailto:hassan210mmm@gmail.com">hassan210mmm@gmail.com</a>. We will respond to your request within the timeframe required by law.</p>

          <h3>10. Children's Privacy</h3>
          <p>The Service is not intended for individuals under the age of 18. We do not knowingly collect personal data from children under 18. If we become aware that a child under 18 has provided us with personal data, we will take steps to delete such information. If you believe we have collected data from a child under 18, please contact us immediately.</p>

          <h3>11. International Data Transfers</h3>
          <p>Your data may be processed on servers located in Iraq. By using the Service, you consent to the transfer of your data to these servers.</p>

          <h3>12. Changes to This Privacy Policy</h3>
          <p>We may update this Privacy Policy from time to time. Material changes will be communicated through the Service. Your continued use of the Service after changes take effect constitutes your acceptance of the revised policy.</p>

          <h3>13. Breach Notification</h3>
          <p>In the event of a data breach that poses a risk to your rights and freedoms, we will notify you and the relevant supervisory authority within 72 hours of becoming aware of the breach, to the extent required by applicable law.</p>

          <h3>14. Contact</h3>
          <p>For questions or concerns regarding this Privacy Policy or our data practices, please contact us at: <a href="mailto:hassan210mmm@gmail.com">hassan210mmm@gmail.com</a></p>
        </div>
      </div>
      <div class="footer">Reviewed by Bashir</div>
    </div>
  </div>
</template>

<style lang="scss" scoped>
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}
.body_title {
  width: 100%;
  height: 100vh;
  overflow: hidden;
  background: hsl(283, 86%, 14%);
  display: flex;
  justify-content: center;
  align-items: center;
}
.title {
  font-family: "Comic Sans MS", "Comic Sans", "Comic Neue", cursive;
  font-size: 150%;
  position: absolute;
  top: 1.5%;
  z-index: 2;
  background-color: rgb(8, 8, 29);
  padding: 5px 15px;
  border-radius: 25%;
  left: 50%;
  transform: translateX(-50%);
  .xletter { font-size: inherit; font-family: inherit; color: blue; }
  .oletter { font-size: inherit; font-family: inherit; color: purple; }
}
.hidden-title {
  opacity: 0 !important;
  pointer-events: none;
}
.layout {
  display: flex;
  flex-direction: column;
  width: 85%;
  height: 82%;
  max-width: 1000px;
  background: rgb(16, 16, 36);
  border-radius: 5%;
  overflow: hidden;
  position: absolute;
  bottom: 2%;
}
.tab-bar {
  display: grid;
  grid-template-columns: 1fr 1fr;
  background: rgb(8, 8, 29);
}
.tab-btn {
  background: rgb(8, 8, 29);
  border-style: solid;
  border-width: 0px 0px;
  border-color: hsl(283, 86%, 14%);
  color: #f7ebff;
  font-size: large;
  cursor: pointer;
  width: 100%;
  height: 100%;
  &:hover { background-color: rgb(4, 4, 34); }
  &.active { background-color: rgb(16, 16, 36); border-width: 0px 0px 5px 0px; border-color: rgb(1, 1, 110); }
}
.content {
  flex: 1;
  overflow-y: auto;
  padding: 30px 35px;
  scroll-behavior: smooth;
  &::-webkit-scrollbar { width: 8px; }
  &::-webkit-scrollbar-track { background: rgb(10, 10, 26); border-radius: 4px; }
  &::-webkit-scrollbar-thumb { background: #7c3aed; border-radius: 4px; }
}
.footer {
  text-align: center;
  padding: 12px;
  font-size: 13px;
  color: #6b6b8a;
  font-style: italic;
  border-top: 1px solid rgb(1, 1, 110);
  background: rgb(8, 8, 29);
}
.policy-content {
  font-size: 15px;
  line-height: 1.8;
  color: #e0d6f0;
  h2 {
    color: #ffffff;
    font-size: 24px;
    margin-bottom: 4px;
    font-weight: 700;
  }
  .subtitle {
    color: #9e9eb8;
    font-size: 13px;
    margin-bottom: 20px;
  }
  .consent-line {
    background: rgba(124, 58, 237, 0.15);
    padding: 10px 16px;
    border-radius: 6px;
    font-weight: 600;
    color: #ffffff;
    margin-bottom: 28px;
    font-size: 14px;
    border-left: 3px solid #7c3aed;
  }
  h3 {
    color: #ffffff;
    font-size: 17px;
    margin: 22px 0 8px;
    font-weight: 700;
  }
  p { margin-bottom: 10px; color: #c4b8db; }
  ul { margin: 6px 0 14px 20px; li { margin-bottom: 5px; color: #c4b8db; line-height: 1.7; } }
  a { color: #7c3aed; text-decoration: none; &:hover { text-decoration: underline; } }
  table { width: 100%; border-collapse: collapse; margin: 12px 0 16px; font-size: 14px;
    th { background: rgba(124, 58, 237, 0.2); color: #f7ebff; padding: 10px; text-align: left; border: 1px solid #2a1a5e; font-weight: 600; }
    td { padding: 10px; border: 1px solid #2a1a5e; color: #c4b8db; vertical-align: top; background: rgba(8, 8, 29, 0.3); }
  }
}
@media (max-width: 800px) {
  .layout { width: 92%; height: 88%; }
  .content { padding: 20px; }
  .tab-btn { font-size: 13px; padding: 10px 16px; }
}
.title a { text-decoration: none; color: inherit; }
</style>

<script>
export default {
  data() {
    return {
      show_title: true,
      activeTab: "terms",
    };
  },
  mounted() {
    this.show_title_func();
  },
  methods: {
    async show_title_func() {
      while (true) {
        const title = document.getElementsByClassName("title")[0];
        const layout = document.getElementsByClassName("layout")[0];
        if (!title || !layout) {
          await new Promise((resolve) => setTimeout(resolve, 500));
          continue;
        }
        const title_postion = title.getBoundingClientRect();
        const box_position = layout.getBoundingClientRect();
        const isColliding = !(
          title_postion.right < box_position.left ||
          title_postion.left > box_position.right ||
          title_postion.bottom < box_position.top ||
          title_postion.top > box_position.bottom
        );
        this.show_title = !isColliding;
        await new Promise((resolve) => setTimeout(resolve, 100));
      }
    },
  },
};
</script>
