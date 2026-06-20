<template>
  <div class="body_title" >
    <h1 class="title" :class="{ 'hidden-title': !show_title }">
      <router-link to="/">
        <span class="xletter">X</span><span class="oletter">O</span>kingdom
      </router-link>
    </h1>
    <br>
    <div
      class="mainbox"
      :class="choosen_tab"
      v-if="choosen_tab == 'login' || choosen_tab == 'signin'"
    >
      <button
        class="log_in_button"
        @click="choosen_tab = 'login'"
        :class="{
          choosen_tab: choosen_tab === 'login',
          is_choose: is_choose === true,
        }"
      >
        Log in
      </button>
      <button
        class="Sign_in_button"
        @click="choosen_tab = 'signin'"
        :class="{
          choosen_tab: choosen_tab === 'signin',
          is_choose: is_choose === true,
        }"
      >
        Sign in
      </button>
      <div class="loginsection" v-if="choosen_tab == 'login'">
        <p class="welcomeword"><b>Good to see you again, buddy</b></p>
        <p class="Demandforinformation">
          <span class="note">Note:</span>Please follow the instrection to log
          in.
        </p>
        <label id="label" for="username">Enter your nick name:</label>
        <input
          id="nickname"
          type="nickname"
          autocomplete="username"
          placeholder="Nick name"
        />
        <label id="label" for="email">Enter your email:</label>
        <input
          id="email"
          type="text"
          autocomplete="email"
          placeholder="Email"
        />
        <label id="label" for="password">Enter your password:</label>
        <input
          id="password"
          type="password"
          autocomplete="password"
          placeholder="Password"
        />
        <div class="divloginbutton">
          <button id="loginbutton" @click="printerrors">log in</button>
        </div>
      </div>
      <small class="Policy">By clicking "Log in" or "Sign in", you agree to our <router-link to="/policy">Terms & Privacy Policy</router-link>.</small>
      <div class="signinsection" v-if="choosen_tab == 'signin'">
        <p class="welcomeword2"><b>New here? Let's build your legend!</b></p>
        <p class="Demandforinformation2">
          <span class="note">Note:</span>Please follow the instrection to sign
          in.
        </p>
        <label id="fullnamelabel" for="fullname">Enter your full name:</label>
        <div id="fullname">
          <input
            id="firstname"
            type="firstname"
            autocomplete="firstname"
            placeholder="First name"
          />
          <input
            id="lastname"
            type="lastname"
            autocomplete="lastname"
            placeholder="Last name"
          />
        </div>
        <label id="nickname2label" for="nickname2">Enter your nick name:</label>
        <input
          id="nickname2"
          type="nickname"
          autocomplete="username"
          placeholder="Nick name"
        />
        <label id="email2label" for="email2">Enter your email:</label>
        <input
          id="email2"
          type="text"
          autocomplete="email"
          placeholder="Email"
        />
        <label id="password2label" for="password2">Enter your password:</label>
        <div id="fullpassword">
          <input
            id="password2"
            type="password"
            autocomplete="password"
            placeholder="Password"
          />
          <input
            id="conformpassword2"
            type="password"
            autocomplete="password"
            placeholder="Conform your password"
          />
        </div>
        <label id="birthdatelabel" for="birthdate">Enter your birthdate:</label>
        <input id="birthdate" type="date" autocomplete="birthdate" />
        <label id="gender" for="gender">Enter your gender:</label>
        <div class="genderboxmale">
          <input
            id="gendermale"
            type="radio"
            autocomplete="gender"
            value="Male"
            name="user_gender"
          />
          <label for="gendermale">Male</label>
        </div>
        <div class="genderboxfemale">
          <input
            id="genderfemale"
            type="radio"
            autocomplete="gender"
            value="Female"
            name="user_gender"
          />
          <label for="genderfemale">Female</label>
        </div>
        <div class="divsigninbutton">
          <button id="signinbutton" @click="printerrors">sign in</button>
        </div>
      </div>
    </div>

    <div
      class="mainbox_verification"
      :class="choosen_tab"
      v-if="
        choosen_tab != 'login' &&
        choosen_tab != 'signin'
      "
    >
      <h2>
        We have sent the verification code to the email address you provided.
      </h2>

      <div class="verification_code_box">
        <label for="verification_code" id="verification_code_label"
          >Enter your verification code:</label
        >
        <br />
        <input
          type="number"
          id="verification_code_input"
          placeholder="Verification code"
        />
      </div>
      <button class="verificationbutton" @click="check_verification_code">
        check verification code
      </button>
    </div>


    <div class="animationbox"></div>
  </div>
</template>

<style lang="scss">
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
  font-family: "Comic Sans MS", "Comic Sans", "Comic Neue", cursive;
}
button {
  &:hover {
      background-color: rgb(4, 4, 34);
    }
}
.title {
  font-family: "Comic Sans MS", "Comic Sans", "Comic Neue", cursive;
  font-size: 150%;
  position: absolute;
  top: 1.5%;
  z-index: 2;
  background-color: rgb(8, 8, 29);
  padding: 5px;
  border-radius: 25%;
  .xletter {
    font-size: inherit;
    font-family: inherit;
    color: blue;
  }
  .oletter {
    font-size: inherit;
    font-family: inherit;
    color: purple;
  }
}
.body_title {
  font-size: large;
  color: #f7ebff;
  display: flex;
  position: absolute;
  top: 0;
  left: 0;
  height: 100vh;
  width: 100vw;
  background-color: hsl(283, 86%, 14%);
  justify-content: center;
  align-items: center;
  overflow: hidden;
}
.hidden-title {
  opacity: 0 !important;
  pointer-events: none;
}

.choosen_tab {
  border-style: solid;
  border-color: rgb(1, 1, 110) !important;
  border-width: 0px 0px 5px 0px !important;
  background-color: rgb(16, 16, 36) !important;
}
.mainbox_verification {
  display: grid;
  background-color: rgb(16, 16, 36);
  z-index: 2;
  border-radius: 5%;
  align-items: center;
  grid-template-columns: 1fr;
  grid-template-rows: 1fr 1fr 1fr;
  overflow: hidden;
  font-size: large;
  color: #f7ebff;
  transition: all 0.3s ease;
  position: absolute;
  bottom: 5%;
  top: 10%;
  @media (max-width: 1450px) {
    width: 75%;
    height: 70%;
  }

  @media (max-width: 1100px) {
    width: 85%;
    height: 75%;
  }
  @media (max-width: 950px) {
    width: 95%;
    height: 90%;
    top: 9%;
  }
  h2 {
    width: 93%;
    margin-left: 5%;
    margin-right: 7%;
    margin-bottom: 12%;
    margin-top: 3%;
  }

  .verification_code_box {
    width: 100%;
    #verification_code_input {
      width: 93%;
      margin-left: 3%;
      margin-right: 5%;
    }
    #verification_code_label {
      margin: 3%;
    }
  }
  button {
    margin: auto;
    border-style: solid;
    border-width: 0px 0px;
    border-color: hsl(283, 86%, 14%);
    font-size: large;
    color: #f7ebff;
    &:hover {
      background-color: rgb(4, 4, 34);
    }
    position: absolute;
    left: 50%;
    transform: translateX(-50%);
    bottom: 20%;
    height: 5%;
    width: 55%;
    background-color: rgb(17, 17, 75);
  }
}
  .Policy {
    font-size: small;
    position: absolute;
    justify-self: center;
    bottom: 0%;
    height: 5%;
    margin: 3%;
    margin-top: 0%;
  }

a {
  color: hsl(283, 93%, 41%);
}
.mainbox {
  border-radius: 5%;
  display: grid;
  background-color: rgb(16, 16, 36);
  z-index: 2;
  align-items: center;
  grid-template-columns: 1fr 1fr;
  grid-template-rows: 5%;
  overflow: hidden;
  font-size: large;
  color: #f7ebff;
  transition: all 0.3s ease;
  position: absolute;
  bottom: 5%;
  top: 10%;

  &.login {
    width: 75%;
    height: 75%;
    top: 15%;

    @media (max-width: 1450px) {
      width: 75%;
      height: 70%;
    }

    @media (max-width: 1100px) {
      width: 85%;
      height: 75%;
    }
    @media (max-width: 950px) {
      width: 85%;
      height: 90%;
      top: 9%;
      .divloginbutton {
        position: absolute;
        bottom: 15%;
      }
      .mainbox {
        position: absolute;
        bottom: 10%;
      }
    }
  }

  &.signin {
    width: 75%;
    height: 85%;
    position: absolute;
    top: 10%;
    .divsigninbutton {
      position: absolute;
      bottom: 10%;
    }

    @media (max-width: 1350px) {
      width: 85%;
      height: 90%;
      position: absolute;
      top: 9%;
    }

    @media (max-width: 1005px) {
      .signinsection {
        height: 75%;
        margin: 10%;
        margin-top: 0;
        align-self: start;
        display: grid;
        grid-template-columns: repeat(2, 1fr);
        grid-template-rows: auto;
        grid-column: span 2;
        align-items: center;
      }
      #fullnamelabel,
      #fullpasswordlabel {
        grid-column: span 3;
      }
      #fullname,
      #nickname2label,
      #email2label,
      #password2label,
      #birthdatelabel,
      #gender {
        grid-column: span 3;
        width: 100%;
      }
      #firstname {
        width: 49%;
        margin-right: 1%;
      }
    }
    #lastname {
      width: 49%;
      margin-left: 1%;
    }
    #fullpassword {
      grid-column: span 2;
      width: 100%;
      #password2 {
        width: 49%;
        margin-right: 1%;
      }
      #conformpassword2 {
        width: 49%;
        margin-left: 1%;
      }
    }
  }

  .log_in_button {
    background-color: rgb(8, 8, 29);
    width: 100%;
  }
  .Sign_in_button {
    background-color: rgb(8, 8, 29);
    width: 100%;
  }
  button {
    margin: auto;
    margin-top: 0%;
    width: 100%;
    height: 100%;
    border-style: solid;
    border-width: 0px 0px;
    border-color: hsl(283, 86%, 14%);
    font-size: large;
    color: #f7ebff;
    &:hover {
      background-color: rgb(4, 4, 34);
    }
  }
}

.animationbox {
  background: linear-gradient(
    hsl(283, 86%, 14%) 1%,
    rgb(13, 13, 58) 45% 55%,
    hsl(283, 86%, 14%) 99%
  );
  position: absolute;
  top: 100%;
  width: 100%;
  height: 25%;
  animation: floatUp 15s ease-out infinite;
  animation-delay: 1s;
  z-index: 1;
  filter: blur(0 20px);
}

@keyframes floatUp {
  0% {
    transform: translateY(0);
  }
  100% {
    transform: translateY(-130vh);
  }
}
.loginsection {
  height: 75%;
  margin: 10%;
  margin-top: 0;
  align-self: start;
  display: grid;
  grid-template-columns: 1fr;
  grid-template-rows: auto;
  grid-column: span 2;
  align-items: center;
  .welcomeword {
    justify-self: center;
    font-size: xx-large;
  }
  .Demandforinformation {
    .note {
      color: red;
    }
  }
  #loginbutton {
    background-color: rgb(17, 17, 75);
    grid-row: 9;
  }
  input {
    color: black;
    font-weight: bold;
  }
}
.divloginbutton {
  position: absolute;
  justify-self: center;
  height: 5%;
  width: 40%;
  left: 50%;
  transform: translateX(-50%);

  @media (max-height: 899px) {
    bottom: 10%;
  }

  @media (min-height: 900px) {
    bottom: 10%;
  }
}
#password {
  width: 100%;
}
@media (min-width: 999px) {
  .signinsection {
    height: 75%;
    margin: 10%;
    margin-top: 0;
    align-self: start;
    display: grid;
    grid-template-columns: repeat(3, 1fr);
    grid-template-rows: auto;
    grid-column: span 2;
    align-items: center;
    gap: 5px;
    #fullname {
      grid-column: span 2;
      #firstname {
        width: 49%;
        margin-right: 1%;
      }
      #lastname {
        width: 49%;
        margin-left: 1%;
      }
      #fullpassword {
        grid-row: span 2;
        #password2 {
          width: 49%;
          margin-right: 1%;
        }
        #conformpassword2 {
          width: 49%;
          margin-left: 1%;
        }
      }
    }
  }
}
.welcomeword2 {
  justify-self: center;
  font-size: xx-large;
  grid-column: span 3;
  text-align: center;
}
input {
  color: black;
  font-weight: bold;
}

.Demandforinformation2 {
  grid-column: span 3;
  justify-self: start;
  .note {
    color: red;
  }
}
#signinbutton {
  background-color: rgb(17, 17, 75);
}
#nickname2 {
  justify-self: start;
  grid-column: span 2;
  width: 100%;
}
#email2 {
  justify-self: start;
  grid-column: span 2;
  width: 100%;
}
#gender {
  justify-self: start;
  margin-top: 2%;
}
#birthdate {
  justify-self: start;
  grid-column: span 2;
  width: 100%;
}
#gendermale {
  margin-left: 5%;
}
#genderfemale {
  margin-left: 5%;
}
.divsigninbutton {
  position: absolute;
  justify-self: center;
  bottom: 20%;
  height: 5%;
  width: 40%;
  background-color: rgb(17, 17, 75);
  left: 50%;
  transform: translateX(-50%);
}
.title a { text-decoration: none; color: inherit; }
</style>

<script>
import { API_BASE } from '@/config.js'

export default {
  data() {
    return {
      show_title: true,
      is_otp: true,
      times: 1,
      is_ready: false,
      is_ready2: false,
      process_type: null,
      user_nickname: "",
      user_first_name: "",
      user_last_name: "",
      user_Gender: null,
      user_birthdate: null,
      user_password: "",
      user_confirm_password: "",
      user_email: "",
      choosen_tab: "login",
      errors: [],
      errors2: [],
      is_choose: false,
      server_token: null,
      user_lang: "",
      user_timezone: "",
      screen_w: null,
      screen_h: null,
      user_browser: null,
      user_os: null,
      cpu_cores: null,
      prefers_dark: null,
      is_online: false,
      battery_level: null,
      is_charging: null,
      user_ip: null,
      user_country: null,
      user_city: null,
      user_Internet_service_provider: null,
      show_title: true,
      is_otp: true,
      times: 1,

    };
  },
  mounted() {
    this.is_log_in();
    this.show_title_func();
    this.refresh_error2();
  },
  methods: {
    async is_log_in() {
      let local_token = localStorage.getItem("token");
      if (local_token) {
        try {
          let response = await fetch(API_BASE + "/api/auth/auto-login", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({ token: local_token }),
          });
          let result = await response.json();
          if (result.status === "OK") {
              this.$router.push('/home')
          } else {
            localStorage.removeItem("token");
          }
        } catch (error) {
        }
      }
    },

    async check_verification_code() {
      let codeInput = document.getElementById("verification_code_input");
      let codeValue = codeInput ? codeInput.value : "";

      let oldErrors = document.querySelectorAll(".floating-errors");
      oldErrors.forEach((msg) => msg.remove());

      let times_local = this.times;
      if (times_local > 5) {
        this.print_single_error (
          "verification_code_input",
          "Attempts limit reached; please reload the page to try again.",
        );
        return;
      }
      if (codeValue === "") {
        this.print_single_error(
          "verification_code_input",
          "Please enter the verification code",
        );
        return;
      }
      if (codeValue.length !== 6) {
        this.print_single_error(
          "verification_code_input",
          "Incorrect code, please try again",
        );
        this.times += 1;
        return;
      }

      try {
        let emailValue = this.user_email;
        let nicknameValue = this.user_nickname;
        let response = await fetch(API_BASE + "/api/auth/verify-otp", {
          method: "POST",
          headers: { "Content-Type": "application/json" },
          body: JSON.stringify({
            email: emailValue,
            otp: codeValue,
            nickname: nicknameValue
          }),
        });

        let result = await response.json();

        if (result.status === "OK") {
          if (result.token) {
            localStorage.setItem("token", result.token);
          }
          this.$router.push('/home')
        }else if (result.status === "TOOMANY"){
          this.print_single_error (
          "verification_code_input",
          "Attempts limit reached; please reload the page to try again.",
        );
        } else {
          this.print_single_error(
            "verification_code_input",
            "Incorrect code, please try again",
          );
          this.times += 1;
        }
      } catch (error) {
        this.print_single_error(
          "verification_code_input",
          "Connection error, try again later",
        );
        this.times += 1;
      }
    },

    print_single_error(elementid, message) {
      let element = document.getElementById(elementid);
      if (element) {
        let small = document.createElement("small");
        small.className = "floating-errors";
        small.style.color = "red";
        small.style.fontSize = "small";
        small.style.position = "fixed";
        small.style.zIndex = "998";
        small.innerText = message;

        let elementposition = element.getBoundingClientRect();
        let wariningpositoiny =
          window.innerHeight * 0.005 + elementposition.bottom;
        let wariningpositoinx = elementposition.left;

        small.style.left = wariningpositoinx + "px";
        small.style.top = wariningpositoiny + "px";
        document.body.appendChild(small);
      }
    },

    finderrors() {
      let error_list = [];

      let nickname = document.getElementById("nickname")
        ? document.getElementById("nickname").value
        : "";
      let email = document.getElementById("email")
        ? document.getElementById("email").value
        : "";
      let password = document.getElementById("password")
        ? document.getElementById("password").value
        : "";

      let firstname = document.getElementById("firstname")
        ? document.getElementById("firstname").value
        : "";
      let lastname = document.getElementById("lastname")
        ? document.getElementById("lastname").value
        : "";
      let nickname2 = document.getElementById("nickname2")
        ? document.getElementById("nickname2").value
        : "";
      let email2 = document.getElementById("email2")
        ? document.getElementById("email2").value
        : "";
      let password2 = document.getElementById("password2")
        ? document.getElementById("password2").value
        : "";
      let conformpassword2 = document.getElementById("conformpassword2")
        ? document.getElementById("conformpassword2").value
        : "";
      let birthdate = document.getElementById("birthdate")
        ? document.getElementById("birthdate").value
        : "";
      let birth = new Date(birthdate)
      let selectedGender = document.querySelector(
        'input[name="user_gender"]:checked',
      );
      let cani = true;
      const nicknamePattern = /^[a-zA-Z0-9_]+$/;
      let at = false;
      let point = false;

      if (this.choosen_tab === "login") {
        if (nickname === "") {
          error_list.push(["nickname", "Please fill this field"]);
        } else if (nickname.length > 15 || nickname.length < 3) {
          error_list.push([
            "nickname",
            "Nickname should be between 3 and 15 characters",
          ]);
        } else if (!nicknamePattern.test(nickname)) {
          error_list.push([
            "nickname",
            "Only letters, numbers, and '_' are allowed",
          ]);
        }
        if (email === "") {
          error_list.push(["email", "Please fill this field"]);
        } else {
          for (let letter of email) {
            if (letter === "@") at = true;
            if (letter === ".") point = true;
          }
          if (at === false || point === false) error_list.push(["email", "Invalid email"]);
        }
        if (password === "") {
          error_list.push(["password", "Please fill this field"]);
          cani = false;
        } else if (password.length < 8) {
          error_list.push([
            "password",
            "Password should be 8 letters at least",
          ]);
          cani = false;
        }
      }

      if (this.choosen_tab === "signin") {
        if (firstname === "") {
          error_list.push(["firstname", "Please fill this field"]);
        }
        if (lastname === "") {
          error_list.push(["lastname", "Please fill this field"]);
        }
        if (nickname2 === "") {
          error_list.push(["nickname2", "Please fill this field"]);
        } else if (nickname2.length > 15 || nickname2.length < 3) {
          error_list.push([
            "nickname2",
            "nickname2 should be between 3 and 15 characters",
          ]);
        } else if (!nicknamePattern.test(nickname2)) {
          error_list.push([
            "nickname2",
            "Only letters, numbers, and '_' are allowed",
          ]);
        }

        if (email2 === "") {
          error_list.push(["email2", "Please fill this field"]);
        } else {
          for (let letter of email2) {
            if (letter === "@") at = true;
            if (letter === ".") point = true;
          }
          if (at === false || point === false) error_list.push(["email2", "Invalid email"]);
        }
        if (birthdate === "") {
          error_list.push(["birthdate", "Please fill this field"]);
          cani = false;
        }else if (isNaN(birth.getTime())) {
          error_list.push(["birthdate", "Invalid date format. Please enter a full date (YYYY-MM-DD)."]);
          cani = false;
        }else{
          let minDate = new Date();
          let maxDate = new Date();
          minDate.setFullYear(minDate.getFullYear() - 18);
          maxDate.setFullYear(maxDate.getFullYear() - 150);
          if (birth > new Date()) {
            error_list.push(["birthdate", "You're not born yet. Please come back later."]);
            cani = false;
          }else if (birth > minDate) {
            error_list.push(["birthdate", "You must be at least 18 years old to sign in."]);
            cani = false;
          }else if (birth < maxDate && cani) {
            error_list.push(["birthdate", "Don't forget your heart medication, grandpa!"]);
            cani = false;
          }
        }
        if (password2 === "") {
          error_list.push(["password2", "Please fill this field"]);
          cani = false;
        } else if (password2.length < 8) {
          error_list.push([
            "password2",
            "Password should be 8 letters at least",
          ]);
          cani = false;
        }

        if (conformpassword2 === "") {
          error_list.push(["conformpassword2", "Please fill this field"]);
          cani = false;
        } else if (conformpassword2.length < 8) {
          error_list.push([
            "conformpassword2",
            "Password should be 8 letters at least",
          ]);
          cani = false;
        }

        if (
          password2 !== "" &&
          conformpassword2 !== "" &&
          password2 !== conformpassword2 &&
          cani
        ) {
          error_list.push([
            "password2",
            "Passwords do not match. Please try again",
          ]);
        }
        if (!selectedGender) {
          error_list.push(["gendermale", "Please select your gender"]);
        }
      }

      this.errors = error_list;
      return error_list;
    },

    async printerrors() {
      this.errors2 = [];
      this.finderrors();
      if (this.errors.length === 0 && this.errors2.length === 0) {
          await this.gotinfofromuser();
          await this.send_info_and_response();
          if (this.errors.length > 0) {
              for (let error of this.errors) {
                  this.print_single_error(error[0], error[1]);
              }
          }
          if (this.errors2.length > 0) {
              for (let error of this.errors2) {
                  this.print_single_error(error[0], error[1]);
              }
          }
      }
      this.is_choose = true;
      while (this.is_choose) {
        this.finderrors();
        let all_errors = [...this.errors, ...this.errors2];

        let oldErrors = document.querySelectorAll(".floating-errors");
        oldErrors.forEach((msg) => msg.remove());
        if (this.errors.length > 0) {
          this.errors2 = []
        }
        if (all_errors.length > 0) {
          
          for (let error of all_errors) {
            let elementid = error[0];
            let element = document.getElementById(elementid);

            if (element) {
              let small = document.createElement("small");
              small.className = "floating-errors";

              small.style.color = "red";
              small.style.fontSize = "small";
              small.style.position = "fixed";
              small.style.zIndex = "998";
              small.innerText = error[1];

              let elementposition = element.getBoundingClientRect();
              let wariningpositoiny =
                window.innerHeight * 0.005 + elementposition.bottom;
              let wariningpositoinx = elementposition.left;

              small.style.left = wariningpositoinx + "px";
              small.style.top = wariningpositoiny + "px";

              document.body.appendChild(small);
              this.is_ready = false
            }
          }
        } else {
          this.is_ready = true
          break;
        }

        await new Promise((resolve) => setTimeout(resolve, 20));
      }
    },
    async refresh_error2() {
      while (true) {
        if (this.errors2.length > 0) {
          for (let error of this.errors2) {
            let el = document.getElementById(error[0]);
            if (!el) continue;
            let small = document.createElement("small");
            small.className = "floating-errors";
            small.style.color = "red";
            small.style.fontSize = "small";
            small.style.position = "fixed";
            small.style.zIndex = "998";
            small.innerText = error[1];
            let rect = el.getBoundingClientRect();
            small.style.left = rect.left + "px";
            small.style.top = (window.innerHeight * 0.005 + rect.bottom) + "px";
            document.body.appendChild(small);
          }
        }
        await new Promise((resolve) => setTimeout(resolve, 20));
      }
    },
    async gotinfofromuser() {
      this.user_lang = navigator.language || navigator.userLanguage;
      this.user_timezone = Intl.DateTimeFormat().resolvedOptions().timeZone;
      this.screen_w = window.screen.width;
      this.screen_h = window.screen.height;
      const sys_info = navigator.userAgent;
      this.user_birthdate = document.getElementById("birthdate")?.value || "";
      this.user_first_name = document.getElementById("firstname")?.value || "";
      this.user_last_name = document.getElementById("lastname")?.value || "";

      if (sys_info.indexOf("Win") !== -1) this.user_os = "Windows";
      else if (sys_info.indexOf("Mac") !== -1) this.user_os = "macOS";
      else if (sys_info.indexOf("Android") !== -1) this.user_os = "Android";
      else if (
        sys_info.indexOf("iPhone") !== -1 ||
        sys_info.indexOf("iPad") !== -1
      )
        this.user_os = "iOS";
      else if (sys_info.indexOf("Linux") !== -1) this.user_os = "Linux";

      if (sys_info.indexOf("Chrome") !== -1)
        this.user_browser = "Google Chrome";
      else if (sys_info.indexOf("Safari") !== -1) this.user_browser = "Safari";
      else if (sys_info.indexOf("Firefox") !== -1)
        this.user_browser = "Firefox";
      else if (sys_info.indexOf("Edge") !== -1)
        this.user_browser = "Microsoft Edge";
      this.cpu_cores = String(navigator.hardwareConcurrency || "error");
      this.prefers_dark = window.matchMedia(
        "(prefers-color-scheme: dark)",
      ).matches;
      this.is_online = navigator.onLine;
      window.addEventListener("offline", () => {
        this.is_online = false;
      });

      window.addEventListener("online", () => {
        this.is_online = true;
      });
      if (navigator.getBattery) {
        try {
          const battery = await navigator.getBattery();

          this.battery_level = `${Math.round(battery.level * 100)}%`;
          this.is_charging = battery.charging;

          battery.addEventListener("levelchange", () => {
            this.battery_level = `${Math.round(battery.level * 100)}%`;
          });

          battery.addEventListener("chargingchange", () => {
            this.is_charging = battery.charging;
          });
        } catch (error) {
          this.battery_level = "error";
        }
      } else {
        this.battery_level = "error";
      }
      try {
        let response = await fetch("https://ipapi.co/json/");

        if (response.ok) {
          let ipData = await response.json();

          this.user_ip = ipData.ip;
          this.user_country = ipData.country_name;
          this.user_city = ipData.city;
          this.user_Internet_service_provider = ipData.org;
        }
      } catch (error) {
        this.user_ip = "error";
        this.user_country = "error";
        this.user_city = "error";
        this.user_Internet_service_provider = "error";
      }
    },
    async send_info_and_response() {
  const isLogin = this.choosen_tab === "login";
  const btn = document.getElementById(isLogin ? "loginbutton" : "signinbutton");
  const originalText = btn.innerText;
  let dots = 1;
  btn.innerText = ".";
  const dotInterval = setInterval(() => {
    dots = (dots % 4) + 1;
    btn.innerText = ".".repeat(dots);
  }, 1000);
  const restore = () => { clearInterval(dotInterval); btn.innerText = originalText; };

    this.user_nickname = document.getElementById(isLogin ? "nickname" : "nickname2")?.value || "";
    this.user_email = document.getElementById(isLogin ? "email" : "email2")?.value || "";
    this.user_password = document.getElementById(isLogin ? "password" : "password2")?.value || "";
    this.user_Gender = document.querySelector('input[name="user_gender"]:checked')?.value || "";
    this.user_birthdate = document.getElementById("birthdate")?.value || "";
    this.user_first_name = document.getElementById("firstname")?.value || "";
    this.user_last_name = document.getElementById("lastname")?.value || "";
  
  const payload = {
    action: this.choosen_tab,
    nickname: this.user_nickname,
    email: this.user_email,
    password: this.user_password,
    gender: this.user_Gender,
    birthdate: this.user_birthdate,
    firstname: this.user_first_name,
    lastname: this.user_last_name,
    device: {
      lang: this.user_lang,
      timezone: this.user_timezone,
      screen_w: this.screen_w,
      screen_h: this.screen_h,
      os: this.user_os,
      browser: this.user_browser,
      cpu_cores: this.cpu_cores,
      prefers_dark: String(this.prefers_dark)
    },
    network: {
      ip: this.user_ip,
      country: this.user_country,
      city: this.user_city,
      isp: this.user_Internet_service_provider
    }
  };

  let response = await fetch(API_BASE + "/api/auth/send-otp", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(payload), 
  });

  let result = await response.json();

  if (result.status === "OK") {
    restore();
    this.choosen_tab = "verification";
    this.is_choose = false;
    this.is_otp = true;
  } else if (result.status === "WRONG") {
    restore();
    this.is_choose = false;
    if (isLogin) {
        this.errors2.push(["nickname", "Credentials incorrect"]);
        this.errors2.push(["password", "Credentials incorrect"]);
        this.errors2.push(["email", "Credentials incorrect"]);
    } else {
        this.errors2.push(["nickname2", "Username is taken"]);
    }

  } else if (result.status === "MALICIOUS") {
    restore();
    localStorage.removeItem("token");
    this.police();
  }else {
    restore();
        this.errors2.push(["nickname2", "Error"]);
        this.errors2.push(["nickname", "Error"]);
        this.errors2.push(["password", "Error"]);
  }
},
    police() {
      const alertScreen = document.createElement("div");
      const alertText = document.createElement("h1");

      Object.assign(alertScreen.style, {
        position: "fixed",
        top: "0",
        left: "0",
        width: "100vw",
        height: "100vh",
        zIndex: "999999",
        display: "flex",
        justifyContent: "center",
        alignItems: "center",
      });

      alertText.innerText = "YOU ARE MALICIOUS";
      alertText.style.fontSize = "5rem";
      alertText.style.fontWeight = "bold";

      alertScreen.appendChild(alertText);
      document.body.appendChild(alertScreen);

      let isRed = true;
      const police = setInterval(() => {
        alertScreen.style.backgroundColor = isRed ? "red" : "blue";
        alertText.style.color = isRed ? "blue" : "red";
        isRed = !isRed;
      }, 300);
      setTimeout(() => {
        clearInterval(police);
        window.location.replace("https://www.google.com");
      }, 5000);
    },

    async show_title_func() {
      while (true) {
        const title = document.getElementsByClassName("title")[0];
        const mainbox = document.getElementsByClassName("mainbox")[0];
        const mainbox_verification = document.getElementsByClassName(
          "mainbox_verification",
        )[0];

        const activeBox = mainbox || mainbox_verification;

        if (!title || !activeBox) {
          await new Promise((resolve) => setTimeout(resolve, 500));
          continue;
        }

        const title_postion = title.getBoundingClientRect();
        const box_position = activeBox.getBoundingClientRect();

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