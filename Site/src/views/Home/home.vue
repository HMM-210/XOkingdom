<template>
  <div class="body_title">
    <div class="signout" @click="signout">Sign out</div>
    <h1 class="title" :class="{ 'hidden-title': !show_title }">
      <router-link to="/">
        <span class="xletter">X</span><span class="oletter">O</span>kingdom
      </router-link>
    </h1>
    <br />
    <div class="mainbox">
      <p></p>
      <button @click="this.$router.push('/xogame')">
        XO Game
      </button>
    </div>
    <div class="animationbox"></div>
  </div>
</template>

<style scoped>
.signout {
  position: absolute;
  top : 2%;
  right: 2%;
  z-index: 5;
  &:hover {
    text-decoration: underline;
  }
}
.mainbox {
  overflow:auto;
  border-radius: 5%;
  display: grid;
  background-color: rgb(16, 16, 36);
  z-index: 2;
  align-items: center;
  grid-template-columns: 1fr;
  grid-template-rows: repeat(3,5%);
  overflow: hidden;
  font-size: large;
  color: #f7ebff;
  transition: all 0.3s ease;
  position: absolute;
  bottom: 5%;
  top: 10%;
  gap: 3%;

  width: 55%;
  height: 75%;
  top: 15%;

  @media (max-width: 1450px) {
    width: 55%;
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
  }
  button {
    width: auto;
    margin: 0% 5%;
    background-color: rgb(30, 23, 94);
    border-radius: 5% ;
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
    };
  },
  mounted() {
    this.check_access();
    this.show_title_func();
  },
  methods: {
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
    signout() {
      localStorage.removeItem("token");
      this.check_access()
    },
    async show_title_func() {
      while (true) {
        const title = document.getElementsByClassName("title")[0];
        const mainbox = document.getElementsByClassName("mainbox")[0];

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
  },
};
</script>
