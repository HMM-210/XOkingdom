<template>
  <div class="body_title">
    <h1 class="title" :class="{ 'hidden-title': !show_title }">
      <router-link to="/">
        <span class="xletter">X</span><span class="oletter">O</span>kingdom
      </router-link>
    </h1>
    <div class="mainbox">
      <div class="timer"></div>
      <div class="gamebox">
        <div
          class="cell"
          @click="
            current_click = 0;
            make_move();
          "
        >
          <span
            id="cell1"
            :class="{
              'blue-x': current_board[0] === 'X',
              'purple-o': current_board[0] === 'O',
            }"
          >
            {{ current_board[0] }}
          </span>
        </div>
        <div
          class="cell"
          @click="
            current_click = 1;
            make_move();
          "
        >
          <span
            id="cell2"
            :class="{
              'blue-x': current_board[1] === 'X',
              'purple-o': current_board[1] === 'O',
            }"
          >
            {{ current_board[1] }}
          </span>
        </div>
        <div
          class="cell"
          @click="
            current_click = 2;
            make_move();
          "
        >
          <span
            id="cell3"
            :class="{
              'blue-x': current_board[2] === 'X',
              'purple-o': current_board[2] === 'O',
            }"
          >
            {{ current_board[2] }}
          </span>
        </div>
        <div
          class="cell"
          @click="
            current_click = 3;
            make_move();
          "
        >
          <span
            id="cell4"
            :class="{
              'blue-x': current_board[3] === 'X',
              'purple-o': current_board[3] === 'O',
            }"
          >
            {{ current_board[3] }}
          </span>
        </div>
        <div
          class="cell"
          @click="
            current_click = 4;
            make_move();
          "
        >
          <span
            id="cell5"
            :class="{
              'blue-x': current_board[4] === 'X',
              'purple-o': current_board[4] === 'O',
            }"
          >
            {{ current_board[4] }}
          </span>
        </div>
        <div
          class="cell"
          @click="
            current_click = 5;
            make_move();
          "
        >
          <span
            id="cell6"
            :class="{
              'blue-x': current_board[5] === 'X',
              'purple-o': current_board[5] === 'O',
            }"
          >
            {{ current_board[5] }}
          </span>
        </div>
        <div
          class="cell"
          @click="
            current_click = 6;
            make_move();
          "
        >
          <span
            id="cell7"
            :class="{
              'blue-x': current_board[6] === 'X',
              'purple-o': current_board[6] === 'O',
            }"
          >
            {{ current_board[6] }}
          </span>
        </div>
        <div
          class="cell"
          @click="
            current_click = 7;
            make_move();
          "
        >
          <span
            id="cell8"
            :class="{
              'blue-x': current_board[7] === 'X',
              'purple-o': current_board[7] === 'O',
            }"
          >
            {{ current_board[7] }}
          </span>
        </div>
        <div
          class="cell"
          @click="
            current_click = 8;
            make_move();
          "
        >
          <span
            id="cell9"
            :class="{
              'blue-x': current_board[8] === 'X',
              'purple-o': current_board[8] === 'O',
            }"
          >
            {{ current_board[8] }}
          </span>
        </div>
      </div>

      <div class="settingsbox">
        <label for="difficultyselect">Difficulty:</label>
        <select
          id="difficultyselect"
          v-model="selected_difficulty"
          :disabled="is_start"
        >
          <option value="hard">Hard</option>
          <option value="medium">Medium</option>
          <option value="easy">Easy</option>
        </select>

        <label for="whostartsselect">Who starts:</label>
        <select
          id="whostartsselect"
          v-model="selected_whostarts"
          :disabled="is_start"
        >
          <option value="player">Player</option>
          <option value="computer">Computer</option>
        </select>

        <label for="symbolselect">Symbols:</label>
        <select
          id="symbolselect"
          v-model="selected_symbol"
          :disabled="is_start"
        >
          <option value="x">X</option>
          <option value="o">O</option>
        </select>

        <label for="timerselect">Timer:</label>
        <select id="timerselect" v-model="selected_timer" :disabled="is_start">
          <option value="10seconds">10 seconds</option>
          <option value="30seconds">30 seconds</option>
          <option value="1minute">1 minute</option>
        </select>
        <button class="startandstopbutton" @click="start_and_end_button">
          {{ is_start ? "Give up" : "Start Game" }}
        </button>
      </div>
    </div>
    <div class="animationbox"></div>
  </div>
</template>

<style lang="scss" scoped>
.timer {
  position: absolute;
  top: -30px;
  left: 50%;
  transform: translateX(-50%);
  z-index: 99;
  text-align: center;
}
button {
  &:hover {
    background-color: rgb(4, 4, 34);
  }
}
.mainbox {
  @media (max-width: 1450px) {
    width: 70%;
    height: 55%;
  }

  @media (max-width: 1100px) {
    width: 85%;
    height: 65%;
  }

  @media (max-width: 950px) {
    width: 90%;
    height: 65%;
  }
  @media (max-width: 850px) {
    width: 95%;
    height: 90%;
    top: 9%;
  }

  button {
    background-color: rgb(30, 23, 94);
    border-radius: 5%;
  }
}
@media (min-width: 850px) {
  .mainbox {
    padding: 5%;
    overflow: visible;
    border-radius: 5%;
    display: grid;
    background-color: rgb(16, 16, 36);
    z-index: 2;
    align-items: center;
    grid-template-columns: repeat(2, 1fr);
    grid-template-rows: repeat(1, 1fr);
    font-size: large;
    color: #f7ebff;
    transition: all 0.3s ease;
    position: absolute;
    bottom: 5%;
    top: 25%;
    width: 75%;
    height: 55%;
    gap: 5%;

    .gamebox {
      display: grid;
      z-index: 3;
      align-items: center;
      grid-template-columns: repeat(3, 1fr);
      grid-template-rows: repeat(3, 1fr);
      width: 100%;
      height: 100%;
      gap: 2px;
      .cell {
        width: 100%;
        height: 100%;
        border: 1px solid rgb(210, 161, 255);

        container-type: inline-size;

        display: flex;
        justify-content: center;
        align-items: center;

        span {
          font-size: 60cqw;
          font-weight: bold;
          line-height: 1;
          display: block;
          text-align: center;
          &.blue-x {
            color: blue;
          }

          &.purple-o {
            color: purple;
          }
        }
      }

      .cell:nth-child(3n + 1) {
        border-left: none;
      }
      .cell:nth-child(3n) {
        border-right: none;
      }
      .cell:nth-child(-n + 3) {
        border-top: none;
      }
      .cell:nth-child(n + 7) {
        border-bottom: none;
      }
    }
    .settingsbox {
      display: grid;
      z-index: 3;
      align-items: center;
      grid-template-columns: repeat(2, 1fr);
      grid-template-rows: repeat(5, 1fr);
      width: 100%;
      height: 100%;
      gap: 2%;
      .startandstopbutton {
        grid-column: span 2;
        height: 60%;
      }
    }
  }
}
.hidden-title {
  opacity: 0 !important;
  pointer-events: none;
}

@media (max-width: 850px) {
  .mainbox {
    padding: 5%;
    overflow: visible;
    border-radius: 5%;
    display: grid;
    background-color: rgb(16, 16, 36);
    z-index: 2;
    align-items: center;
    grid-template-columns: repeat(1, 1fr);
    grid-template-rows: repeat(2, 1fr);
    overflow: visible;
    font-size: medium;
    color: #f7ebff;
    transition: all 0.3s ease;
    position: absolute;
    bottom: 5%;
    width: 70%;
    height: 70%;
    top: 20%;
    gap: 2%;
    .gamebox {
      font-weight: bold;
      display: grid;
      z-index: 3;
      align-items: center;
      grid-template-columns: repeat(3, 1fr);
      grid-template-rows: repeat(3, 1fr);
      width: 100%;
      height: 100%;
      gap: 2px;
      .cell {
        width: 100%;
        height: 100%;
        border: 1px solid rgb(210, 161, 255);

        container-type: inline-size;

        display: flex;
        justify-content: center;
        align-items: center;

        span {
          font-size: 60cqw;
          font-weight: bold;
          line-height: 1;
          display: block;
          text-align: center;
          &.blue-x {
            color: blue;
          }

          &.purple-o {
            color: purple;
          }
        }
      }

      .cell:nth-child(3n + 1) {
        border-left: none;
      }
      .cell:nth-child(3n) {
        border-right: none;
      }
      .cell:nth-child(-n + 3) {
        border-top: none;
      }
      .cell:nth-child(n + 7) {
        border-bottom: none;
      }
    }
    .settingsbox {
      padding: 5%;
      display: grid;
      z-index: 3;
      align-items: center;
      grid-template-columns: repeat(2, 1fr);
      grid-template-rows: repeat(5, 1fr);
      width: 100%;
      height: 100%;
      gap: 2%;
      #difficultyselect {
        max-width: 100%;
      }
      .startandstopbutton {
        grid-column: span 2;
      }
    }
    @media (max-height: 830px) {
      width: 70%;
      height: 80%;
      top: 10%;
    }
  }
}
.hidden-title {
  opacity: 0 !important;
  pointer-events: none;
}
.title a { text-decoration: none; color: inherit; }
</style>

<script>
import { API_BASE } from '@/config.js'

export default {
  data() {
    return {
      have_access: false,
      show_title: true,
      is_start: false,
      current_board: ["", "", "", "", "", "", "", "", ""],
      current_click: null,
      is_player_turn: false,
      timer_on: false,
      _timer_interval: null,
      _seconds: 0,
      _minutes: 0,

      selected_difficulty: "hard",
      selected_whostarts: "player",
      selected_symbol: "x",
      selected_timer: "10seconds",
    };
  },
  mounted() {
    this.startup();
    this.check_access();
    this.show_title_func();
    this.timer();
  },
  methods: {
    async timer() {
      if (this._timer_interval) {
        clearInterval(this._timer_interval);
      }
      if (this.selected_timer === "10seconds") {
        this._seconds = 10;
        this._minutes = 0;
      } else if (this.selected_timer === "30seconds") {
        this._seconds = 30;
        this._minutes = 0;
      } else if (this.selected_timer === "1minute") {
        this._seconds = 0;
        this._minutes = 1;
      }
      let local_token = localStorage.getItem("token");
      this._timer_interval = setInterval( async () => {
        let timer = document.getElementsByClassName("timer")[0];
        if (!this.timer_on) {
          clearInterval(this._timer_interval);
          return;
        }
        timer.innerText =
          String(this._minutes) + ":" + String(this._seconds).padStart(2, "0");
        if (this._minutes > 0 && this._seconds === 0) {
          this._minutes--;
          this._seconds = 59;
        } else if (this._seconds > 0) {
          this._seconds--;
        } else if (this._seconds == 0){
          this.current_board = ["", "", "", "", "", "", "", "", ""]
          this.is_start = false;
          this.timer_on = false;
          document.getElementsByClassName("timer")[0].innerText = "00:00";
          try {
          let response = await fetch(
            API_BASE + "/api/game/xo/3x3/time-out",
            {
              method: "POST",
              headers: { "Content-Type": "application/json" },
              body: JSON.stringify({ token: local_token}),
            },
          );
          let result = await response.json();
          timer.innerText = "Time out :(";
          } catch {

          }
        }
      }, 1000);
    },
    async startup() {
      let local_token = localStorage.getItem("token");
      if (local_token) {
        try {
          let response = await fetch(
            API_BASE + "/api/game/xo/3x3/resume-playing",
            {
              method: "POST",
              headers: { "Content-Type": "application/json" },
              body: JSON.stringify({ token: local_token }),
            },
          );
          let result = await response.json();
          if (result.status === "OK") {
            this.current_board = result.board;
            this.is_start = true;
            this.is_player_turn = true;
            this.timer_on = true;
            this.timer();
          } else if (result.status === "WRONG") {
            this.is_start = false;
          } else {
            this.is_start = false;
            localStorage.removeItem("token");
          }
        } catch (error) {
          this.is_start = false;
          this.timer_on = false;
        }
      } else {
        this.is_start = false;
        this.timer_on = false;
      }
    },
    async check_access() {
      let local_token = localStorage.getItem("token");
      if (local_token) {
        try {
          let response = await fetch(
            API_BASE + "/api/auth/auto-login",
            {
              method: "POST",
              headers: { "Content-Type": "application/json" },
              body: JSON.stringify({ token: local_token }),
            },
          );
          let result = await response.json();
          if (result.status === "OK") {
            this.have_access = true;
          } else {
            localStorage.removeItem("token");
            this.have_access = false;
            this.$router.push("/authentication");
          }
        } catch (error) {
          this.have_access = false;
          this.$router.push("/authentication");
        }
      } else {
        this.have_access = false;
        this.$router.push("/authentication");
      }
    },
    async show_title_func() {
      while (true) {
        const title = document.getElementsByClassName("title")[0];
        const mainbox = document.getElementsByClassName("timer")[0];

        const activeBox = mainbox;

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
    make_move() {
      let timer = document.getElementsByClassName("timer")[0];

      let selected_symbol = this.selected_symbol.toUpperCase();
      let current_click = this.current_click;
      if (
        this.current_board[current_click] !== "" ||
        !this.is_start ||
        !this.is_player_turn
      ) {
        return;
      } else {
        this.is_player_turn = false;
        this.current_board[current_click] = selected_symbol;
      }
      let local_token = localStorage.getItem("token");
          fetch(API_BASE + "/api/game/xo/3x3/move", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({
          token: local_token,
          board: this.current_board,
          timer: this.selected_timer,
          click: this.current_click,
          difficulty: this.selected_difficulty,
          who_starts: this.selected_whostarts,
          symbol: this.selected_symbol,
        }),
      })
        .then((res) => res.json())
        .then((result) => {
          if (result.status === "OK") {
            this.current_board = result.board;
            this.is_player_turn = true;
            if (this.selected_timer === "10seconds") {
              this._seconds = 10;
              this._minutes = 0;
            } else if (this.selected_timer === "30seconds") {
              this._seconds = 30;
              this._minutes = 0;
            } else if (this.selected_timer === "1minute") {
              this._seconds = 0;
              this._minutes = 1;
            }
          } else if (result.status === "MALICIOUS") {
            this.is_start = false;
            this.timer_on = false;
            localStorage.removeItem("token");
            this.police();
          } else if (result.status === "WIN") {
            this.current_board = result.board;
            this.is_start = false;
            this.timer_on = false;
            timer.innerText = "You win ;)";
          } else if (result.status === "DRAW") {
            this.current_board = result.board;
            this.is_start = false;
            this.timer_on = false;
            timer.innerText = "Draw";
          } else if (result.status === "LOSE") {
            this.current_board = result.board;
            this.is_start = false;
            this.timer_on = false;

            timer.innerText = "You lose :(";
          }
        });
    },
    start_and_end_button() {
      this.current_board = ["", "", "", "", "", "", "", "", ""];
      if (this.is_start === false) {
        if (this.selected_whostarts == "player") {
          document.getElementsByClassName("timer")[0].innerText = "00:00";
          this.is_start = true;
          this.is_player_turn = true;
          this.timer_on = true;
          this.timer();
        } else {
          document.getElementsByClassName("timer")[0].innerText = "00:00";
          this.is_start = true;
          this.is_player_turn = false;
          this.timer_on = true;
          this.timer();
          let local_token = localStorage.getItem("token");
      fetch(API_BASE + "/api/game/xo/3x3/move", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({
              token: local_token,
              board: this.current_board,
              timer: this.selected_timer,
              click: this.current_click,
              difficulty: this.selected_difficulty,
              who_starts: this.selected_whostarts,
              symbol: this.selected_symbol,
            }),
          })
            .then((res) => res.json())
            .then((result) => {
              if (result.status === "OK") {
                this.current_board = result.board;
                this.is_player_turn = true;
              } else if (result.status === "MALICIOUS") {
                this.is_start = false;
                this.timer_on = false;
                localStorage.removeItem("token");
                this.police();
              }
            })
            .catch((error) => {
              this.is_start = false;
              this.timer_on = false;
            });
        }
      } else {
        document.getElementsByClassName("timer")[0].innerText = "00:00";
        this.board = ["", "", "", "", "", "", "", "", ""]
        let local_token = localStorage.getItem("token");
        fetch(API_BASE + "/api/game/xo/3x3/give-up", {
          method: "POST",
          headers: { "Content-Type": "application/json" },
          body: JSON.stringify({
            token: local_token,
          }),
        })
          .then((res) => res.json())
          .then((result) => {
            if (result.status === "OK") {
              this.is_start = false;
              this.timer_on = false;
              this.current_board = ["", "", "", "", "", "", "", "", ""];
              this.is_player_turn = false;
            }
          });
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
  },
};
</script>
