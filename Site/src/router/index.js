import { createRouter, createWebHashHistory } from "vue-router";
import AuthenticationPage from "@/views/Authentication/Authentication.vue";
import HomePage           from "@/views/Home/home.vue"
import PolicyPage         from "@/views/Policy/Policy.vue"
import XOGamePage    from "@/views/Games/XOGame/XOGame.vue"
import XOGame3X3Page from "@/views/Games/XOGame/3X3/XOGame_3X3.vue"

const routes = [
    {
        path     : "/",
        redirect : "/home",
    },
    {
        name      : "authentication"  ,
        path      : "/authentication" ,
        component :AuthenticationPage ,
    },
    {
        path      : "/home"  ,
        name      : "home"   ,
        component : HomePage ,
    },
    {
        path      : "/policy"  ,
        name      : "policy"   ,
        component : PolicyPage ,
    },
    {
        path      : "/xogame"  ,
        name      : "xogame"   ,
        component : XOGamePage ,
    },
    {
        path      : "/xogame/3x3"  ,
        name      : "xogame_3x3"   ,
        component : XOGame3X3Page ,
    },


]

const router = createRouter({
  history: createWebHashHistory(),
  routes,
});

export default router;