import { createMemoryHistory, createRouter } from 'vue-router'
import { loadRemote } from './loadRemote'
import App from './App.vue'

// no need to put defineAsyncComponent, becouse component is async by default
const routes = [
    {
        path: "/",
        component: App
    },
    {
        path: "/app",
        component: () => loadRemote('remote', 'http://localhost:8080', './AppView')
    },
    {
        path: "/about",
        component: () => loadRemote('remote', 'http://localhost:8080', './AboutView')
    },
    {
        path: "/profile",
        component: () => loadRemote('remote', 'http://localhost:8080', './ProfileView')
    }
]

const router = createRouter({
    history: createMemoryHistory(),
    routes,
})

export default router